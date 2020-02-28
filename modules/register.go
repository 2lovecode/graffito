package modules

import (
	"graffito/modules/algorithm"
	"graffito/modules/practice"
	"graffito/modules/tools"
)

type ModuleInterface interface {
	Register() map[string]interface{}
	GetModuleName() string
}

func RegisterModule() map[string]map[string]interface{}{
	toolsModule := tools.NewToolsModule()
	practiceModule := practice.NewPracticeModule()
	algModule := algorithm.NewAlgModule()

	return map[string]map[string]interface{}{
		toolsModule.GetModuleName() : toolsModule.Register(),
		practiceModule.GetModuleName() : practiceModule.Register(),
		algModule.GetModuleName() : algModule.Register(),

	}
}
