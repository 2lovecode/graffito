package tools

func init() {

}
type ToolsModule struct {
}

func NewToolsModule() ToolsModule {
	p := ToolsModule{}
	return p
}

func (mine ToolsModule)GetModuleName() string {
	return "tools"
}

func (mine ToolsModule)Register() map[string]interface{}{
	return map[string]interface{}{
		"hello" : Hello,
		"ccnum" : CalcCharNum,
	}
}