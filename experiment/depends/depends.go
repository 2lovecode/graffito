package depends

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type DataContainer struct {
	mu sync.Mutex
	dataMap map[string]interface{}
}

func NewDataContainer() *DataContainer {
	return &DataContainer{
		mu:      sync.Mutex{},
		dataMap: make(map[string]interface{}),
	}
}

func (dc *DataContainer) Set(key string, value interface{}) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	dc.dataMap[key] = value
}

func (dc *DataContainer) Get(key string) (value interface{}) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	value, _ = dc.dataMap[key]

	return
}

type IService interface {
	Name() string
	Run(ctx context.Context, dc *DataContainer) error
	Decode(receiver interface{}) error
}

type Depends struct {
	serviceMap      map[string]IService // 服务map
	serviceCount    int32
	depends         map[string][]string // 依赖关系
	starts          []string            //第一波调用
	inversedDepends map[string][]string // 反依赖关系
	executeFlags    []int32
	executeMap      map[string]*int32 // 执行列表 每一个服务对应一条记录，记录中：0 - 未执行 1 - 执行完毕
	executeStatus   int32             // 总体进度表 0 - 未结束 1 - 已结束
	inChanels       map[string]chan bool
	outChanel       chan string
	timeout         time.Duration
	ctxCtrl			context.Context
	ctxCtrlCancel	context.CancelFunc
	dc 				*DataContainer
}

func NewDepends(timeout time.Duration) *Depends {
	if timeout == 0 {
		timeout = 1000 * time.Millisecond
	}
	return &Depends{
		serviceMap:      make(map[string]IService, 0),
		serviceCount:    0,
		depends:         make(map[string][]string, 0),
		inversedDepends: make(map[string][]string, 0),
		executeFlags:    make([]int32, 500),
		executeMap:      make(map[string]*int32),
		executeStatus:  0,
		timeout:         timeout,
		dc:				 NewDataContainer(),
	}
}

// 注册服务
func (me *Depends) Register(service IService) error {
	me.serviceMap[service.Name()] = service
	me.executeMap[service.Name()] = &me.executeFlags[int(me.serviceCount)]
	me.serviceCount++
	// todo 重复注册等情况的处理
	return nil
}

// 添加依赖关系
func (me *Depends) AddDepend(service IService, serviceDepends []IService) error {
	var dependNames []string
	for _, item := range serviceDepends {
		dependNames = append(dependNames, item.Name())
	}
	if _, ok := me.depends[service.Name()]; ok {
		me.depends[service.Name()] = append(me.depends[service.Name()], dependNames...)
	} else {
		me.depends[service.Name()] = dependNames
	}
	//todo 依赖检查等情况的处理
	return nil
}

// 执行
func (me *Depends) Execute(ctx context.Context) {
	me.ctxCtrl, me.ctxCtrlCancel = context.WithTimeout(context.Background(), me.timeout)
	// 输入输出通道
	me.inChanels = make(map[string]chan bool)
	me.outChanel = make(chan string)
	defer func() {
		// 错误处理
		if p := recover(); p != nil {
			if atomic.AddInt32(&me.executeStatus, 1) == 1 {
				for _, in := range me.inChanels {
					close(in)
				}
				close(me.outChanel)
			}
			fmt.Println(string(debug.Stack()))
		}
	}()

	// 初始化 解析出反依赖关系和第一波调用队列
	me.bootstrap()

	// 点火
	me.fire()

	// 调度
	me.dispatch()

	// 运行
	me.operate(ctx)

	<-me.ctxCtrl.Done()
	if atomic.AddInt32(&me.executeStatus, 1) == 1 {
		for _, in := range me.inChanels {
			close(in)
		}
		close(me.outChanel)
	}
}

// 初始化
func (me *Depends) bootstrap() error {

	for _, eachService := range me.serviceMap {
		me.inChanels[eachService.Name()] = make(chan bool)

		if depends, ok := me.depends[eachService.Name()]; !ok {
			me.starts = append(me.starts, eachService.Name())
		} else {
			for _, dp := range depends {
				if _, o := me.inversedDepends[dp]; o {
					me.inversedDepends[dp] = append(me.inversedDepends[dp], eachService.Name())
				} else {
					me.inversedDepends[dp] = []string{eachService.Name()}
				}
			}
		}
	}
	return nil
}

// 启动第一波调用
func (me *Depends) fire() error {
	tGo(func() error {
		for _, startName := range me.starts {
			me.inChanels[startName] <- true
		}
		return nil
	})

	return nil
}

// 调度
func (me *Depends) dispatch() error {
	tGo(func() error {
		for serviceName := range me.outChanel {
			if serviceName != "" {
				executeQueue := make([]string, 0)
				if atomic.LoadInt32(&me.executeStatus) == 0 {
					if inversedDepends, ok := me.inversedDepends[serviceName]; ok {
						for _, inversedDepend := range inversedDepends {
							// 检查依赖项
							flag := true
							if depends, o := me.depends[inversedDepend]; o {
								for _, depend := range depends {
									if atomic.LoadInt32(me.executeMap[depend]) != 1 {
										flag = false
										break
									}
								}
							}
							if flag {
								executeQueue = append(executeQueue, inversedDepend)
							}
						}
					}
				}

				if len(executeQueue) > 0 {
					for _, eachName := range executeQueue {
						if inChanel, ok := me.inChanels[eachName]; ok {
							inChanel <- true
						}
					}
				}
			}
		}
		return nil
	})
	return nil
}

func (me *Depends) operate(ctx context.Context) error {
	for _, service := range me.serviceMap {

		go func(in chan bool, s IService) {
			startC := make(chan bool)
			var addFlag int32 = 0

			defer func() {
				// 错误处理
				if p := recover(); p != nil {
					atomic.StoreInt32(me.executeMap[s.Name()], 1)
					fmt.Println(string(debug.Stack()))
				}
			}()
			if flag, open := <-in; flag && open {
				// 执行
				go func(startC chan bool, addFlag int32) {
					defer func() {
						if p := recover(); p != nil {
							if atomic.AddInt32(&addFlag, 1) == 1 {
								startC <- true
							}
							fmt.Println(string(debug.Stack()))
						}
					}()
					tNow := time.Now()
					s.Run(ctx, me.dc)
					if atomic.AddInt32(&addFlag, 1) == 1 {
						startC <- true
						close(startC)
					}
					eNow := time.Now()
					fmt.Println(s.Name(), "执行", eNow.Sub(tNow).Milliseconds())
				}(startC, addFlag)

				select {
				case <-me.ctxCtrl.Done():
					atomic.AddInt32(&addFlag, 2)
					if atomic.AddInt32(&me.executeStatus, 1) == 1 {
						for _, in := range me.inChanels {
							close(in)
						}
						close(me.outChanel)
					}
				case <-startC:
					atomic.StoreInt32(me.executeMap[s.Name()], 1)
					finishExecuteCount := 0
					finishFlag := false
					for _, flag := range me.executeMap {
						if *flag == 1 {
							finishExecuteCount++
						}
					}
					if finishExecuteCount >= int(me.serviceCount) {
						finishFlag = true
					}
					if !finishFlag {
						if atomic.LoadInt32(&me.executeStatus) == 0 {
							me.outChanel <- s.Name()
						}
					} else {
						(me.ctxCtrlCancel)()
					}
				}
			}

		}(me.inChanels[service.Name()], service)
	}
	return nil
}

func tGo(fn func() error) {
	go func() {
		defer func() {
			// 错误处理
			if p := recover(); p != nil {
				fmt.Println(debug.Stack())
			}
		}()
		fn()
	}()
}