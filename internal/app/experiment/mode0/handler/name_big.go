package handler


type NameBig struct {
	BaseHandler
}

func (h *NameBig) GetName() string {
	return "name"
}

func (h *NameBig) Output() interface{} {
	data := h.GetInputData()
	if data != nil {
		return data.Data.Name + "_Big_Big_Hooooooo"
	}
	return ""
}

func NewNameBig() *NameBig {
	return &NameBig{}
}