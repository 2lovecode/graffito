package modules

import (
	"github.com/codegangsta/inject"
	//tools "graffito/modules/tools"
)

var moduleList map[string]map[string]interface{}

type InputParams struct {
	Module string
	Task string
	Params map[string]interface{}
}

func init() {
	moduleList = RegisterModule()
}

func Run(input InputParams) {
	inj := inject.New()
	//tools.Hello()
	s := moduleList["tools"]["hello"]
	inj.Invoke(s)
}

func getInvoke(module string, task string) interface{} {
	if m, ok := moduleList[module]; !ok {

	}
	if t, ok := moduleList[module][task]; !ok {

	}
}
