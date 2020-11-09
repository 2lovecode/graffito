package main

import (
	"fmt"
	"graffito/algorithm/lib"
)

func main() {
	/*
		1.数组 -- 可通过下标随机访问 -- O(1)
		2.hash函数计算出的值必须随机且均匀分布 -- 信息论【信息的压缩】
		3.hash函数的时间和空间复杂度
		4.装载因子 -- 衡量冲突的概率 -- 通过扩容重散列降低装载因子
	*/
	testData := map[string]string{
		"k-1" : "v-1",
		"k-2" : "v-2",
		"k-3" : "v-3",
		"k-4" : "v-4",
		"k-5" : "v-5",
		//"k-6" : "v-6",
		//"k-7" : "v-7",
	}
	/**
	开放寻址法解决冲突
	开放寻址，有不同实现
		线性探测 - 实现
		二次探测
		双重散列
	*/
	fmt.Println("===开放寻址处理冲突===")
	openAddrHash := lib.NewOpenAddrHash(5)
	lib.BatchPut(openAddrHash, testData)
	lib.BatchGet(openAddrHash, testData)

	//分离链表处理冲突
	fmt.Println("===分离链表处理冲突===")
	linkHash := lib.NewLinkHash(2)
	lib.BatchPut(linkHash, testData)
	lib.BatchGet(linkHash, testData)


	//装载因子-重散列
	fmt.Println("===分离链表-重散列===")
	linkReHash := lib.NewLinkReHash(2)
	lib.BatchPut(linkReHash, testData)
	lib.BatchGet(linkReHash, testData)

	//
}
