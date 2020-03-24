package experiment

func init() {

}
type ExpModule struct {
}

func NewExpModule() ExpModule {
	p := ExpModule{}
	return p
}

func (mine ExpModule)GetModuleName() string {
	return "exp"
}

func (mine ExpModule)Register() map[string]interface{}{
	return map[string]interface{}{
		"parallel": ParallelRun,
		"event": EventRun,
		"permutation": PermutationRun,
	}
}
