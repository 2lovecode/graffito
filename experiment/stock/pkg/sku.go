package pkg

import "fmt"

type Sku struct {
	ID string
}


func NewSku(id string) *Sku {
	return &Sku{
		ID: id,
	}
}


func (s *Sku) Print() {
	fmt.Println("Good Morning!")
}