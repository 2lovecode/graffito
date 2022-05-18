package experiment

import (
	"context"
	"fmt"
	"graffito/experiment/cache"
	"graffito/experiment/depends"
	"graffito/experiment/event"
	"graffito/experiment/mode0"
	"graffito/experiment/mode0/card"
	"graffito/experiment/mode0/handler"
	"graffito/experiment/mode0/source"
	"graffito/experiment/search"
	"graffito/utils/params"
	"time"

	"github.com/spf13/cobra"
)

func NewExperimentCommand() *cobra.Command {

	expCmd := &cobra.Command{Use: "exp", Short: "试验代码"}

	// stringOpCmd := &cobra.Command{Use: "g3n", Run: func(cmd *cobra.Command, args []string) {
	// 	g3n.Run()
	// }, Short: "g3n游戏包示例代码"}

	// expCmd.AddCommand(stringOpCmd)
	cacheCmd := &cobra.Command{Use: "cross-cache", Run: func(cmd *cobra.Command, args []string) {
		cache.CrossRun()
	}, Short: "缓存穿透"}

	expCmd.AddCommand(cacheCmd)

	eventCmd := &cobra.Command{Use: "event", Run: func(cmd *cobra.Command, args []string) {
		mEvent := event.NewMEvent()

		myC := &event.MyCar{}
		myP := &event.MyPhone{}

		mEvent.On("car", myC)
		mEvent.On("car", myP)

		p := params.NewPayload()
		p.Set("name", "vivo")

		mEvent.Trigger("car", p)
	}, Short: "事件"}
	expCmd.AddCommand(eventCmd)

	mode0Cmd := &cobra.Command{Use: "mode0", Run: func(cmd *cobra.Command, args []string) {

		//数据聚合层 - 可以替换数据来源
		s0 := source.NewS0()
		aggLayer := mode0.NewDataAggregationLayer(mode0.Source(s0))
		aggData := aggLayer.Output()

		//数据处理层 - 可注册多个处理器
		dealerLayer := mode0.NewDealerLayer(aggData)

		//生成处理器 - 每个处理器可注册不同的处理逻辑
		bigCard := card.NewBigCard()
		bigNameHandler := handler.NewNameBig()
		//...
		bigCard.Register(bigNameHandler)

		smallCard := card.NewSmallCard()
		smallNameHandler := handler.NewNameSmall()
		//...
		smallCard.Register(smallNameHandler)

		dealerLayer.RegisterCard(bigCard)
		dealerLayer.RegisterCard(smallCard)

		ret := dealerLayer.Output()

		fmt.Println(ret[0], "{\"name\":\"type_big_Big_Big_Hooooooo\"}")
		fmt.Println(ret[1], "{\"name\":\"Hoooooo_type_small\"}")
	}, Short: "mode0"}
	expCmd.AddCommand(mode0Cmd)

	dependsCmd := &cobra.Command{Use: "depends", Run: func(cmd *cobra.Command, args []string) {
		ctx := context.WithValue(context.TODO(), "q", "test")

		sA := depends.NewServiceA()
		sB := depends.NewServiceB()
		sC := depends.NewServiceC()

		hd := depends.NewDepends(100 * time.Millisecond)

		hd.Register(sA)
		hd.Register(sB)
		hd.Register(sC)

		hd.AddDepend(sC, []depends.IService{sB})
		hd.AddDepend(sB, []depends.IService{sA})

		hd.Execute(ctx)

		sAData := depends.ServiceAData{}
		sBData := depends.ServiceBData{}
		sCData := depends.ServiceCData{}

		sA.Decode(&sAData)
		sB.Decode(&sBData)
		sC.Decode(&sCData)

		fmt.Println(sAData, sBData, sCData)
	}, Short: "depends"}
	expCmd.AddCommand(dependsCmd)

	searchCmd := &cobra.Command{Use: "search", Run: func(cmd *cobra.Command, args []string) {
		search.Run()
	}, Short: "search"}
	expCmd.AddCommand(searchCmd)

	return expCmd
}
