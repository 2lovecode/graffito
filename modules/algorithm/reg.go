package algorithm

type AlgModule struct {

}

func NewAlgModule() AlgModule {
	p := AlgModule{}
	return p
}

func (mine AlgModule)GetModuleName() string {
	return "alg"
}

func (mine AlgModule)Register() map[string]interface{}{
	return map[string]interface{}{
		"array": ArrayRun,
		"link": LinkRun,
		"fibo" : FibonacciRun,
		"qsort" : QuickSortRun,
		"heap" : HeapRun,
		"hash" : HashRun,
	}
}
