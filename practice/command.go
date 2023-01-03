package practice

import (
	"context"
	"graffito/practice/dataloader_t"
	"graffito/practice/generics_x"
	"graffito/practice/plan9"
	"graffito/practice/regexp_x"
	"graffito/practice/schedule_x"
	"graffito/practice/slice_x"
	"graffito/practice/ts"

	"github.com/spf13/cobra"
)

func NewPracticeCommand() *cobra.Command {

	pracCmd := &cobra.Command{Use: "prac", Short: "练习代码"}

	// 切片
	slice1OpCmd := &cobra.Command{Use: "slice", Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "expansion-one-by-one":
				slice_x.ExpansionOneByOne()
			case "expansion-multiple":
				slice_x.ExpansionMultiple()
			case "3":
				slice_x.Run_3()
			case "4":
				slice_x.Run_4()
			case "sorting":
				slice_x.Sorting()
			case "substr":
				slice_x.Substr()
			}
		}

	}}

	pracCmd.AddCommand(slice1OpCmd)

	// graphql-dataloader
	dataloaderCmd := &cobra.Command{
		Use: "dataloader",
		Run: func(cmd *cobra.Command, args []string) {
			dataloader_t.Run(context.Background())
		},
		Short: "dataloader测试",
	}

	pracCmd.AddCommand(dataloaderCmd)

	// go调度器测试
	goScheduleCmd := &cobra.Command{Use: "schedule", Run: func(cmd *cobra.Command, args []string) {
		schedule_x.Run()
	}}
	pracCmd.AddCommand(goScheduleCmd)

	// 测试
	plan9Cmd := &cobra.Command{Use: "plan9", Run: func(cmd *cobra.Command, args []string) {
		plan9.GoWithPlan9()
		plan9.Plan9WithGo()
	}}
	pracCmd.AddCommand(plan9Cmd)

	// 正则表达式
	regCmd := &cobra.Command{Use: "reg", Run: func(cmd *cobra.Command, args []string) {
		regexp_x.Run()
	}}
	pracCmd.AddCommand(regCmd)

	// 泛型
	genericCmd := &cobra.Command{Use: "generic", Run: func(cmd *cobra.Command, args []string) {
		generics_x.Run()
	}}
	pracCmd.AddCommand(genericCmd)

	// 测试
	tsCmd := &cobra.Command{Use: "ts", Run: func(cmd *cobra.Command, args []string) {
		ts.Run()
	}}
	pracCmd.AddCommand(tsCmd)

	return pracCmd
}
