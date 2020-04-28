package modules

import (
	"graffito/modules/algorithm"
	"graffito/modules/experiment"
	"graffito/modules/pattern"
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
	expModule := experiment.NewExpModule()
	patternModule := pattern.NewPatternsModule()

	return map[string]map[string]interface{}{
		toolsModule.GetModuleName() : toolsModule.Register(),
		practiceModule.GetModuleName() : practiceModule.Register(),
		algModule.GetModuleName() : algModule.Register(),
		expModule.GetModuleName() : expModule.Register(),
		patternModule.GetModuleName() : patternModule.Register(),
	}
}
