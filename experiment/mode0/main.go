package main

import (
	"fmt"
	"graffito/experiment/mode0"
	"graffito/experiment/mode0/card"
	"graffito/experiment/mode0/handler"
	"graffito/experiment/mode0/source"
)

// 一种批量多样数据的处理模式
/**
数据源产生的数据，形如
[
	{
		type: "big",
		name: "type_big",
	},
	{
		type: "small",
		name: "type_small",
	},
]

最终产生的数据，形如
[
	{
		name: "type_big_Big_Big_Hooooooo",
	},
	{
		name: "Hoooooo_type_small",
	},
]

*/
func main() {
	//数据聚合层 - 可以替换数据来源
	s0 			:= source.NewS0()
	aggLayer 	:= mode0.NewDataAggregationLayer(mode0.Source(s0))
	aggData 	:= aggLayer.Output()


	//数据处理层 - 可注册多个处理器
	dealerLayer := mode0.NewDealerLayer(aggData)

	//生成处理器 - 每个处理器可注册不同的处理逻辑
	bigCard 		:= card.NewBigCard()
	bigNameHandler 	:= handler.NewNameBig()
	//...
	bigCard.Register(bigNameHandler)

	smallCard 			:= card.NewSmallCard()
	smallNameHandler 	:= handler.NewNameSmall()
	//...
	smallCard.Register(smallNameHandler)

	dealerLayer.RegisterCard(bigCard)
	dealerLayer.RegisterCard(smallCard)

	fmt.Println(dealerLayer.Output())
}
