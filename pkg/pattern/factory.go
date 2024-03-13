package pattern

// pc接口
type Pc interface {
	RunPrint() string
}

// pc接口的具体实现
type lenovo struct{}

func (mine lenovo) RunPrint() string {
	return "联想"
}

func newLenovo() lenovo {
	return lenovo{}
}

type dell struct{}

func (mine dell) RunPrint() string {
	return "戴尔"
}

func newDell() dell {
	return dell{}
}

type hp struct{}

func newHp() hp {
	return hp{}
}

func (mine hp) RunPrint() string {
	return "惠普"
}

// 工厂方法实现
type PcType string

const (
	TypeLenovo PcType = "lenovo"
	TypeDell          = "dell"
	TypeHp            = "hp"
)

func NewPc(t PcType) Pc {
	var p Pc
	switch t {
	case "lenovo":
		p = newLenovo()
	case "dell":
		p = newDell()
	case "hp":
		p = newHp()
	}

	return p
}
