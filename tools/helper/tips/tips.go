package tips

import (
	"fmt"
	"strconv"
	"strings"
)

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

type PrintItem struct {
	Text string
	Children []*PrintItem
}

func (tgs *TipGroups) CliPrintB() {
	if tgs != nil {
		maxLength := 0

		mainItem := &PrintItem{
			Text:     "\n"+CombineText(32, tgs.Name, tgs.Description, tgs.Comment)+"\n",
			Children: nil,
		}


		for _, eachGroup := range tgs.Tips {
			level1Title := CombineText(33, eachGroup.Group, eachGroup.Description, eachGroup.Comment)
			maxLength = Max(len(level1Title), maxLength)

			level1Item := &PrintItem{
				Text:     level1Title,
				Children: nil,
			}
			for _, eachTip := range eachGroup.Items {
				level2Title := CombineText(37, eachTip.Name, "", eachTip.Comment)
				maxLength = Max(len(level2Title), maxLength)
				level2Item := &PrintItem{
					Text:     level2Title,
					Children: nil,
				}
				for _, usageItem := range eachTip.Usage {
					level3Title := CombineText(36, usageItem.Name, usageItem.Description, usageItem.Comment)
					maxLength = Max(len(level3Title), maxLength)
					level2Item.Children = append(level2Item.Children, &PrintItem{
						Text:     level3Title,
						Children: nil,
					})
				}
				level1Item.Children = append(level1Item.Children, level2Item)
			}

			mainItem.Children = append(mainItem.Children, level1Item)
		}

		maxLength += 4*5

		RPrint(mainItem, maxLength, 0)
	}
}

func RPrint(item *PrintItem, maxLength int, level int) {
	if item != nil {
		text := strings.Repeat(" ", 4 * level) + item.Text
		l := maxLength - len(text)
		text += strings.Repeat(" ", l - 1)
		if level == 1 {
			fmt.Println(strings.Repeat("* ", maxLength/2))
		}

		fText := fmt.Sprintf("%-"+strconv.Itoa(maxLength)+"s", text)

		fmt.Println(fText)

		if item.Children != nil && len(item.Children) > 0 {
			for _, eachItem := range item.Children {
				RPrint(eachItem, maxLength, level+1)
			}
		}
		if level == 1 {
			fmt.Println(strings.Repeat("* ", maxLength/2))
		}
	}
	return
}
func CombineText(color int, name string, desc string, comment string) (title string) {
	title = "\033["+strconv.Itoa(color)+";48m"+name+"\033[0m"
	if desc != "" {
		title += " : " + desc
	}
	if comment != "" {
		title += "(" + comment + ")"
	}
	return
}

func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
