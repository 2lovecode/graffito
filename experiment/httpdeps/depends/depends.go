package depends

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync/atomic"
	"time"
)

type IService interface {
	Name() string
	Run(ctx context.Context) error
	Decode(receiver interface{}) error
}

type HttpDeps struct {
	serviceMap      map[string]IService // 服务map
	serviceCount    int32
	depends         map[string][]string // 依赖关系
	starts          []string            //第一波调用
	inversedDepends map[string][]string // 反依赖关系
	executeFlags    []int32
	executeMap      map[string]*int32 // 执行列表 每一个服务对应一条记录，记录中：0 - 未执行 1 - 执行完毕
	executecStatus  int32             // 总体进度表 0 - 未结束 1 - 已结束
	inChanels       map[string]chan bool
	outChanel       chan string
	ctrlCtx         context.Context
	ctrlCancel      context.CancelFunc
	timeout         time.Duration
}

func NewHttpDeps(timeout time.Duration) *HttpDeps {
	if timeout == 0 {
		timeout = 1000 * time.Millisecond
	}
	return &HttpDeps{
		serviceMap:      make(map[string]IService, 0),
		serviceCount:    0,
		depends:         make(map[string][]string, 0),
		inversedDepends: make(map[string][]string, 0),
		executeFlags:    make([]int32, 500),
		executeMap:      make(map[string]*int32),
		executecStatus:  0,
		timeout:         timeout,
	}
}

// 注册服务
func (httpD *HttpDeps) Register(service IService) error {
	httpD.serviceMap[service.Name()] = service
	httpD.executeMap[service.Name()] = &httpD.executeFlags[int(httpD.serviceCount)]
	httpD.serviceCount++
	// todo 重复注册等情况的处理
	return nil
}

// 添加依赖关系
func (httpD *HttpDeps) AddDepend(service IService, serviceDepends []IService) error {
	var dependNames []string
	for _, item := range serviceDepends {
		dependNames = append(dependNames, item.Name())
	}
	if _, ok := httpD.depends[service.Name()]; ok {
		httpD.depends[service.Name()] = append(httpD.depends[service.Name()], dependNames...)
	} else {
		httpD.depends[service.Name()] = dependNames
	}
	//todo 依赖检查等情况的处理
	return nil
}

// 执行
func (httpD *HttpDeps) Execute(ctx context.Context) {
	timeOutCtrl := time.After(httpD.timeout)
	// 输入输出通道
	httpD.inChanels = make(map[string]chan bool)
	httpD.outChanel = make(chan string)
	defer func() {
		atomic.StoreInt32(&(httpD.executecStatus), 1)
		close(httpD.outChanel)

		for _, in := range httpD.inChanels {
			close(in)
		}
	}()
	// 初始化 解析出反依赖关系和第一波调用队列
	httpD.bootstrap()

	// 点火
	httpD.fire()

	// 调度
	httpD.dispatch()

	// 运行
	httpD.operate(ctx)

	tick := time.Tick(1 * time.Millisecond)
	for {
		breakFlag := false
		select {
		case <-timeOutCtrl:
			atomic.StoreInt32(&httpD.executecStatus, 1)
			breakFlag = true
		case <-tick:
			finishExecuteCount := 0
			for _, flag := range httpD.executeMap {
				if atomic.LoadInt32(flag) == 1 {
					finishExecuteCount++
				}
			}
			if finishExecuteCount >= int(httpD.serviceCount) {
				atomic.StoreInt32(&httpD.executecStatus, 1)
				breakFlag = true
			}
		}

		if breakFlag {
			break
		}
	}
}

// 初始化
func (httpD *HttpDeps) bootstrap() error {
	httpD.ctrlCtx, httpD.ctrlCancel = context.WithTimeout(context.Background(), 2000*time.Millisecond)

	for _, eachService := range httpD.serviceMap {
		httpD.inChanels[eachService.Name()] = make(chan bool)

		if depends, ok := httpD.depends[eachService.Name()]; !ok {
			httpD.starts = append(httpD.starts, eachService.Name())
		} else {
			for _, dp := range depends {
				if _, o := httpD.inversedDepends[dp]; o {
					httpD.inversedDepends[dp] = append(httpD.inversedDepends[dp], eachService.Name())
				} else {
					httpD.inversedDepends[dp] = []string{eachService.Name()}
				}
			}
		}
	}
	return nil
}

// 启动第一波调用
func (httpD *HttpDeps) fire() error {
	tGo(func() error {
		for _, startName := range httpD.starts {
			httpD.inChanels[startName] <- true
		}
		return nil
	})

	return nil
}

// 调度
func (httpD *HttpDeps) dispatch() error {
	tGo(func() error {
		for serviceName := range httpD.outChanel {
			if serviceName != "" {
				executeQueue := make([]string, 0)
				if atomic.LoadInt32(&httpD.executecStatus) != 1 {
					if inversedDepends, ok := httpD.inversedDepends[serviceName]; ok {
						for _, inversedDepend := range inversedDepends {
							// 检查依赖项
							flag := true
							if depends, o := httpD.depends[inversedDepend]; o {
								for _, depend := range depends {
									if atomic.LoadInt32(httpD.executeMap[depend]) != 1 {
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
						if inChanel, ok := httpD.inChanels[eachName]; ok {
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

func (httpD *HttpDeps) operate(ctx context.Context) error {
	for _, service := range httpD.serviceMap {

		go func(in chan bool, s IService) {
			defer func() {
				// 错误处理
				if p := recover(); p != nil {
					fmt.Println(debug.Stack())
				}
			}()
			if flag, open := <-in; flag && open {
				tNow := time.Now()
				// 执行
				s.Run(ctx)
				atomic.StoreInt32(httpD.executeMap[s.Name()], 1)
				eNow := time.Now()
				fmt.Println(s.Name(), "执行", eNow.Sub(tNow).Milliseconds())
				if atomic.LoadInt32(&httpD.executecStatus) != 1 {
					httpD.outChanel <- s.Name()
				}
			}

		}(httpD.inChanels[service.Name()], service)
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
