package depends

type Depends struct {
	apiMap  map[string]Api
	depends map[string][]string
}

func NewDepends() *Depends {
	return &Depends{
		apiMap: map[string]Api{
			"a": &ApiA{},
			"b": &ApiB{},
			"c": &ApiC{},
		},
		depends: map[string][]string{
			"c": []string{"a", "b"},
		},
	}
}

func (dp *Depends) Bind(a string, b []string) {

}

func (dp *Depends) Do() {
	list := []string{"a", "b"}
	for _, v := range list {
		aa := dp.apiMap[v]
		go func() {
			aa.Do()
		}()
	}
}
