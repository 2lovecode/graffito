package handler

type NameSmall struct {
	BaseHandler
}

func (h *NameSmall) GetName() string {
	return "name"
}

func (h *NameSmall) Output() interface{} {
	data := h.GetInputData()
	if data != nil {
		return "Hoooooo_" + data.Data.Name
	}
	return ""
}

func NewNameSmall() *NameSmall {
	return &NameSmall{}
}