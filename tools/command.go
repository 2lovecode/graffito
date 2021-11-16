package tools

import (
	"fmt"
	"graffito/tools/color"
	"graffito/tools/helper"
	"graffito/tools/redis"
	"graffito/tools/string_op"
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

	return toolsCmd
}
