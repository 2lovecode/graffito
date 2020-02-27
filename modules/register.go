package modules

import "graffito/modules/tools"

func RegisterModule() map[string]map[string]interface{}{
	return map[string]map[string]interface{}{
		"tools" : tools.Register(),
	}
}
