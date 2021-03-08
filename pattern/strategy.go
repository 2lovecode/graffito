package pattern

type Operator interface {
	Apply(a int, b int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(a int, b int) int {
	return o.Operator.Apply(a, b)
}

type Addition struct{}

func (Addition) Apply(a int, b int) int {
	return a + b
}

type Multiplication struct{}

func (Multiplication) Apply(a int, b int) int {
	return a * b
}
