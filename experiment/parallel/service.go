package parallel

import "sync"

type IServer interface {
	 GoRequest()
	 ParseResponse() (data interface{}, err error)
	 GetName() (name string)
	 GetDepend() (list []string)
}

type PServer struct {
}

func (mine *PServer) GoRequest() {

}

func (mine *PServer) ParseResponse() (data interface{}, err error){
	return nil, nil
}

func (mine *PServer) GetName() (name string){
	return "p"
}

func (mine *PServer) GetDepend() (list []string){
	return []string{"p"}
}

type WaitQueue struct {
	Lock sync.Mutex
	Queue map[string]IServer
}

func (mine *WaitQueue) Has(name string) bool {
	mine.Lock.Lock()
	_, ok := mine.Queue[name]
	mine.Lock.Unlock()
	return ok
}

func (mine *WaitQueue) Del(name string) {
	if mine.Has(name) {
		mine.Lock.Lock()
		delete(mine.Queue, name)
		mine.Lock.Unlock()
	}
}

func (mine *WaitQueue) IsEmpty() bool{
	var emtpy bool
	mine.Lock.Lock()
	emtpy = len(mine.Queue) == 0
	mine.Lock.Unlock()
	return emtpy
}

type DoneQueue struct {
	Lock sync.Mutex
	Queue map[string]IServer
}

func (mine *DoneQueue) Has(name string) bool {
	mine.Lock.Lock()
	_, ok := mine.Queue[name]
	mine.Lock.Unlock()
	return ok
}

func (mine *DoneQueue) Add(name string, server IServer) {
	mine.Lock.Lock()
	mine.Queue[name] = server
	mine.Lock.Unlock()
}

func Run() {
	//var waitQueue WaitQueue
	//var doneQueue DoneQueue
	//
	//ch := make(chan bool)
	//
	//go func() {
	//	for {
	//		if waitQueue.IsEmpty() {
	//			close(ch)
	//			break
	//		}
	//		<-ch
	//	}
	//}()
	//
	//go func() {
	//
	//}
}


