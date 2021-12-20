package excel

import (
	"fmt"

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
