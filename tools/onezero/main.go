package onezero

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func NewCommand() *cobra.Command {
	oneZeroCmd := &cobra.Command{Use: "one-zero", Run: func(cmd *cobra.Command, args []string) {
		str := "1 1 0 0 0 0 0 1 0 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 1 1 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 1 1 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 1 0 0 0 0 0 0 0 0 0 0 0 1 0 1 0 1 0 0"

		strS := strings.Split(str, " ")

		highStr := ""
		lowStr := ""
		split := " "
		cnt := 200
		for _, v := range strS {
			if v == "1" {
				highStr += fmt.Sprintf("%s%s", v, split)
				lowStr += fmt.Sprintf("%s%s", split, split)
			} else {
				highStr += fmt.Sprintf("%s%s", split, split)
				lowStr += fmt.Sprintf("%s%s", v, split)
			}
		}

		if len(highStr) > cnt {
			highS := make([]string, 0)
			lowS := make([]string, 0)
			for i := 0; i < len(highStr); i++ {
				if i%cnt == 0 {
					if i+cnt >= len(highStr) {
						highS = append(highS, highStr[i:])
						lowS = append(lowS, lowStr[i:])
						break
					} else {
						highS = append(highS, highStr[i:i+cnt])
						lowS = append(lowS, lowStr[i:i+cnt])
					}
				}
			}

			for k, _ := range highS {
				fmt.Printf("%s\n%s\n\n\n", highS[k], lowS[k])
			}

		} else {
			fmt.Printf("%s\n%s\n", highStr, lowStr)
		}

	}, Short: "one-zero", Example: "one-zero"}

	return oneZeroCmd
}
