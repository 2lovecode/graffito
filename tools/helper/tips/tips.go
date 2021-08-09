package tips

import "fmt"

type TipGroups struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Comment string `json:"comment"`
	Tips []*TipGroup `json:"tips"`
}

type TipGroup struct {
	Group string `json:"group"`
	Description string `json:"description"`
	Comment string `json:"comment"`
	Items []*TipItem `json:"items"`
}

type TipItem struct {
	Name string `json:"name"`
	Comment string `json:"comment"`
	Usage []*TipUsage `json:"usage"`
}

type TipUsage struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Comment string `json:"comment"`
}

func (tgs *TipGroups) CliPrintA() {
	if tgs != nil {
		fmt.Println(tgs.Name, tgs.Description, tgs.Comment)
		for _, eachGroup := range tgs.Tips {
			fmt.Println("  ", eachGroup.Group, eachGroup.Description, eachGroup.Comment)
			for _, eachTip := range eachGroup.Items {
				fmt.Println("    ", "\033[43;35;5m"+eachTip.Name+"\033[0m", eachTip.Comment)
				for _, usageItem := range eachTip.Usage {
					fmt.Println("      ",usageItem.Name, usageItem.Description, usageItem.Comment)
				}
			}
		}
	}
}
