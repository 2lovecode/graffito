package practice


func init() {
}

type PracticeModule struct {
}

func NewPracticeModule() PracticeModule {
	p := PracticeModule{}
	return p
}

func (mine PracticeModule)GetModuleName() string {
	return "practice"
}

func (mine PracticeModule)Register() map[string]interface{}{
	return map[string]interface{}{
		"zktest": ZkTest,
		"cache" : CacheRun,
		"list" : ListRun,
		"map" : MapRun,
		"string" : StringRun,
		"json" : JsonRun,
		"struct": StructRun,
		"interface": InterfaceRun,
		"func": FunctionRun,
		"arr": ArrayRun,
		"fcopy": FileCopyRun,
		"channel": ChannelRun,
		"routine": RoutineRun,
		"gin" : GinRun,
		"http": HttpRun,
		"pattern": PatternRun,
		"sync": SyncRun,
		"time": TimeRun,
	}
}
