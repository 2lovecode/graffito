package practice

import (
	"graffito/practice/channel_x"
	"graffito/practice/dataloader_t"
	"graffito/practice/generics_x"
	"graffito/practice/interface_x"
	"graffito/practice/map_x"
	"graffito/practice/plan9"
	"graffito/practice/regexp_x"
	"graffito/practice/schedule_x"
	"graffito/practice/slice_x"
	"graffito/practice/ts"

	"github.com/spf13/cobra"
)

func NewPracticeCommand() *cobra.Command {

	pracCmd := &cobra.Command{Use: "prac", Short: "练习代码"}

	cmds := []*cobra.Command{
		slice_x.NewCommand(),      // 切片
		map_x.NewCommand(),        // map
		channel_x.NewCommand(),    // channel
		interface_x.NewCommand(),  // interface
		ts.NewCommand(),           // 测试
		generics_x.NewCommand(),   // 泛型
		regexp_x.NewCommand(),     // 正则表达式
		dataloader_t.NewCommand(), // graphql-dataloader
		schedule_x.NewCommand(),   // go调度器测试
		plan9.NewCommand(),        // plan9
	}

	pracCmd.AddCommand(cmds...)
	return pracCmd
}
