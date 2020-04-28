package pattern

func init() {

}
type PatternsModule struct {
}

func NewPatternsModule() PatternsModule {
	p := PatternsModule{}
	return p
}

func (mine PatternsModule)GetModuleName() string {
	return "pattern"
}

func (mine PatternsModule)Register() map[string]interface{}{
	return map[string]interface{}{
		"singleton":   SingletonRun,
		"factory":   FactoryRun,
		"builder":   BuilderRun,
	}
}

