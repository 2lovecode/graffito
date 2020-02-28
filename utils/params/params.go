package params

type CommandParams struct {
	Module string
	Task string
}

type InputParamsInterface interface {
	GetInputNum() int
	GetFileList() []string
	GetInputPrefix() string
	GetString(key string) string
}

type InputParams struct {
	InputNum int
	FileList []string
	ParamsMap map[string]interface{}
}

func NewInputParams() InputParams{
	p := InputParams{}
	return p
}

func (mine InputParams) GetInputNum() int {
	return mine.InputNum
}

func(mine InputParams) GetFileList() []string {
	return mine.FileList
}
func (mine InputParams) GetInputPrefix() string {
	return "input"
}

func (mine InputParams) GetString(key string) string {
	v, ok := mine.ParamsMap[key]
	if ok {
		value, okk := v.(string)
		if okk {
			return value
		}
	}
	return ""
}
