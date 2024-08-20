package depends

import "fmt"

type Api interface {
	Do()
}

type ApiA struct {
}

func (a *ApiA) Do() {
	fmt.Println("a")
}

type ApiB struct {
}

func (a *ApiB) Do() {
	fmt.Println("b")
}

type ApiC struct {
}

func (a *ApiC) Do() {
	fmt.Println("c")
}
