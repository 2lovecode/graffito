package ginx

import (
	"net/http"
	"sync"

	"github.com/2lovecode/graffito/internal/app/experiment/ginx/render"
)

type Engine struct {
	pool sync.Pool // 使用了一个pool，把context放入其中，可以复用context

	trees methodTrees // 方法树，不用map而使用array。array是内存中的连续空间，方便寻址
}

func NewEngine() *Engine {
	engine := &Engine{
		trees: make(methodTrees, 0, 9),
	}
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (engine *Engine) allocateContext() *XContext {
	return &XContext{engine: engine}
}

/*************  实现http包的ServeHTTP接口    **************/

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*XContext)

	c.Request = req
	c.ResponseWriter = w

	engine.handler(c)

	engine.pool.Put(c)
}

/******************* handler是一个重要的方法，框架中路由，中间件等特性均在此实现 *************************/
func (engine *Engine) handler(xc *XContext) {
	jsonResp := render.JSON{Data: "aaaaaa"}

	xc.Render(200, jsonResp)
}

// 添加路由 把路由放到 --基数树-- 中
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	// todo 检查下路由

	root := engine.trees.get(method)

	if root == nil {

	}
}
