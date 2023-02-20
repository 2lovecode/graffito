package algorithm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	algCommand := &cobra.Command{Use: "alg", Short: "数据结构"}

	arrCmd := &cobra.Command{Use: "arr", Run: func(cmd *cobra.Command, args []string) {
		myA := NewMyArray(2)

		testData := []int{3, 4, 1, 6}

		for k, v := range testData {
			fmt.Println("插入数组：", k, v)
			myA.Insert(v, k)
		}

		fmt.Println("----------------")

		fmt.Println("当前数组长度：", myA.Len)
		fmt.Println("下标为2的数组值：", myA.Get(2))

		fmt.Println("删除下标为2的值")
		myA.Del(2)

		fmt.Println("当前数组长度：", myA.Len)
		fmt.Println("下标为2的数组值：", myA.Get(2))

	}, Short: "数组实现", Example: "graffito alg arr"}

	algCommand.AddCommand(arrCmd)

	fiboCmd := &cobra.Command{Use: "fibo", Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("第%d个数的值是%d\n", FibonacciCnt, Fibonacci(FibonacciCnt))
		fmt.Printf("第%d个数的值是%d(缓存实现)\n", FibonacciCnt, FibonacciUseMem(FibonacciCnt))
	}, Short: "斐波那契数列", Example: "graffito alg fibo"}

	algCommand.AddCommand(fiboCmd)

	hashCmd := &cobra.Command{Use: "hash", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开放寻址哈希表")
		openAddrTestData := map[string]string{
			"k-1": "v-1",
			"k-2": "v-2",
			"k-3": "v-3",
			"k-4": "v-4",
			"k-5": "v-5",
		}
		openAddrHash := NewOpenAddrHash(5)
		BatchPut(openAddrHash, openAddrTestData)
		fmt.Printf("key：%s, value: %s, expValue: %s\n", "k-3", openAddrHash.Get("k-3"), "v-3")

		fmt.Println("分离链表哈希表")
		linkTestData := map[string]string{
			"k-1": "v-1",
			"k-2": "v-2",
			"k-3": "v-3",
			"k-4": "v-4",
			"k-5": "v-5",
			"k-6": "v-6",
			"k-7": "v-7",
		}
		linkHash := NewLinkHash(5)
		BatchPut(linkHash, linkTestData)
		fmt.Printf("key：%s, value: %s, expValue: %s\n", "k-6", linkHash.Get("k-6"), "v-6")

		fmt.Println("分离链表重散列哈希表")
		linkRehashTestData := map[string]string{
			"k-1":  "v-1",
			"k-2":  "v-2",
			"k-3":  "v-3",
			"k-4":  "v-4",
			"k-5":  "v-5",
			"k-6":  "v-6",
			"k-7":  "v-7",
			"k-8":  "v-8",
			"k-9":  "v-9",
			"k-10": "v-10",
			"k-11": "v-11",
			"k-12": "v-12",
			"k-13": "v-13",
		}
		linkReHashHash := NewLinkReHash(5)
		BatchPut(linkReHashHash, linkRehashTestData)
		fmt.Printf("key：%s, value: %s, expValue: %s\n", "k-10", linkReHashHash.Get("k-10"), "v-10")

	}, Short: "哈希表", Example: "graffito alg hash"}
	algCommand.AddCommand(hashCmd)

	heapCmd := &cobra.Command{Use: "heap", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("最大堆实现")
		heap := NewIntHeap(5)
		testData := []int{1, 9, 4, 10, 22, 14, 8, 7}
		for _, v := range testData {
			fmt.Println("push: ", v)
			heap.Push(v)
		}

		fmt.Println("------------")
		for _, _ = range testData {
			fmt.Println("pop:", heap.Pop())
		}
	}, Short: "堆"}
	algCommand.AddCommand(heapCmd)

	quickCmd := &cobra.Command{Use: "qs", Run: func(cmd *cobra.Command, args []string) {
		testData := []int{4, 5, 6, 3, 7, 9}
		fmt.Println("beforeSort: ", testData)
		QuickSort(testData, 0, len(testData)-1)
		fmt.Println("afterSort: ", testData)
	}, Short: "快排"}
	algCommand.AddCommand(quickCmd)

	return algCommand
}
