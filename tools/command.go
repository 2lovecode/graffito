package tools

import (
	"fmt"
	"graffito/tools/color"
	"graffito/tools/excel"
	"graffito/tools/helper"
	"graffito/tools/redis"
	"graffito/tools/sql2go"
	"graffito/tools/string_op"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func NewToolsCommand() *cobra.Command {

	toolsCmd := &cobra.Command{Use: "tools", Short: "工具箱"}

	stringOpCmd := &cobra.Command{Use: "count", Run: func(cmd *cobra.Command, args []string) {
		str := ""
		if len(args) > 0 {
			str = args[0]
		}
		strOp := string_op.NewStringOp(str)

		fmt.Println(strOp.Count())
	}, Short: "字符个数", Example: "graffito tools count abc"}

	toolsCmd.AddCommand(stringOpCmd)

	helperCmd := &cobra.Command{Use: "helper", Run: func(cmd *cobra.Command, args []string) {
		str := ""
		if len(args) > 0 {
			str = args[0]
		}
		h := helper.NewHelper(str)
		h.Run()
	}, Short: "命令列表", Example: "graffito tools helper"}
	toolsCmd.AddCommand(helperCmd)

	redisCmd := &cobra.Command{Use: "redis", Run: func(cmd *cobra.Command, args []string) {
		host := ""
		port := ""

		if len(args) > 0 {
			host = args[0]
			if len(args) > 1 {
				port = args[1]
			}
		}
		r := redis.NewRedisInstance(host, port)
		r.Run()
	}, Short: "命令列表", Example: "graffito tools helper"}
	toolsCmd.AddCommand(redisCmd)

	colorCmd := &cobra.Command{Use: "color", Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 3 {
			res, err := color.ConvertColorFormat(args[0], color.ColorFormat(strings.ToUpper(args[1])), color.ColorFormat(strings.ToUpper(args[2])))
			if err == nil {
				fmt.Println(res)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println("缺少参数！")
		}

		// fmt.Println(color.ConvertColorFormat("#FFC7A68D", color.ColorFormatAHEX, color.ColorFormatRGBA))
		// fmt.Println(color.ConvertColorFormat("rgba(199, 166, 141, 1)", color.ColorFormatRGBA, color.ColorFormatAHEX))
		// fmt.Println(color.ConvertColorFormat("#FFC7A68D", color.ColorFormatAHEX, color.ColorFormatRGBA))
		// fmt.Println(color.ConvertColorFormat("#FFC7A68D", color.ColorFormatAHEX, color.ColorFormatRGBA))
		// fmt.Println(color.ConvertColorFormat("#FFC7A68D", color.ColorFormatAHEX, color.ColorFormatRGBA))
	}, Short: "颜色转换", Example: "graffito tools color"}
	toolsCmd.AddCommand(colorCmd)

	excelCmd := &cobra.Command{Use: "excel", Run: func(cmd *cobra.Command, args []string) {
		argn := len(args)

		var param *excel.ReadExcelParam

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
			excel.ReadExcel(param)
		}
	}, Short: "Excel读取", Example: "graffito tools excel"}
	toolsCmd.AddCommand(excelCmd)

	sql2GoCmd := &cobra.Command{Use: "sql2go", Run: func(cmd *cobra.Command, args []string) {
		sql2go.Run()
	}, Short: "sql2go", Example: "graffito tools sql2go"}
	toolsCmd.AddCommand(sql2GoCmd)

	return toolsCmd
}
