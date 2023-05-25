package universe

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"os"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "universe",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("universe")
			pwd, _ := os.Getwd()
			filePath := fmt.Sprintf("%s/other/universe/W020211027623974108131.xls", pwd)
			fileHandler, err := excelize.OpenFile(filePath)
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
			rows, err := fileHandler.GetRows("sheet1")
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, row := range rows {
				for _, colCell := range row {
					fmt.Print(colCell, "\t")
				}
				fmt.Println()
			}
		},
	}
}
