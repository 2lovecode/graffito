package excel

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const TransTypeInsertSql = ""

type ReadExcelParam struct {
	FilePath  string
	SheetName string
	HeaderRow int
	StartRow  int
	EndRow    int
	StartCol  int
	EndCol    int
}

func ReadExcel(p *ReadExcelParam) {
	if p.SheetName == "" {
		p.SheetName = "sheet1"
	}

	fileHandler, err := excelize.OpenFile(p.FilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := fileHandler.GetRows(p.SheetName)
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
}

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "excel", Run: func(cmd *cobra.Command, args []string) {
		argn := len(args)

		var param *ReadExcelParam

		switch argn {
		case 1:
			param.FilePath = args[0]
		case 2:
			param.FilePath = args[0]
			param.SheetName = args[1]
		case 3:
			param.FilePath = args[0]
			param.SheetName = args[1]
			param.HeaderRow, _ = strconv.Atoi(args[2])
		case 4:
			param.FilePath = args[0]
			param.SheetName = args[1]
			param.HeaderRow, _ = strconv.Atoi(args[2])
			param.StartRow, _ = strconv.Atoi(args[3])
		case 5:
			param.FilePath = args[0]
			param.SheetName = args[1]
			param.HeaderRow, _ = strconv.Atoi(args[2])
			param.StartRow, _ = strconv.Atoi(args[3])
			param.EndRow, _ = strconv.Atoi(args[4])
		case 6:
			param.FilePath = args[0]
			param.SheetName = args[1]
			param.HeaderRow, _ = strconv.Atoi(args[2])
			param.StartRow, _ = strconv.Atoi(args[3])
			param.EndRow, _ = strconv.Atoi(args[4])
			param.StartCol, _ = strconv.Atoi(args[5])
		case 7:
			param.FilePath = args[0]
			param.SheetName = args[1]
			param.HeaderRow, _ = strconv.Atoi(args[2])
			param.StartRow, _ = strconv.Atoi(args[3])
			param.EndRow, _ = strconv.Atoi(args[4])
			param.StartCol, _ = strconv.Atoi(args[5])
			param.StartCol, _ = strconv.Atoi(args[6])
		}

		if param == nil {
			fmt.Println("参数错误")
		} else {
			ReadExcel(param)
		}
	}, Short: "Excel读取", Example: "graffito tools excel"}
}
