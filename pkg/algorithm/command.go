package algorithm

import (
	"fmt"
	"github.com/spf13/cobra"
	"graffito/algorithm/hashmap"
	"graffito/algorithm/link"
	sort2 "graffito/algorithm/sort"
	"graffito/algorithm/tree"
)

func NewCommand() *cobra.Command {
	algCommand := &cobra.Command{Use: "alg", Short: "数据结构"}
	cmds := []*cobra.Command{
		linkCmd(),
		sortCmd(),
		hashmapCmd(),
		treeCmd(),
		otherCmd(),
	}
	algCommand.AddCommand(cmds...)

	return algCommand
}

func linkCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "link",
		Short: "链表",
	}
	cmds := []*cobra.Command{
		{
			Use: "array",
			Run: func(cmd *cobra.Command, args []string) {
				myA := link.NewMyArray(2)

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
			},
			Short: "数组实现",
		},
	}
	rootCmd.AddCommand(cmds...)
	return rootCmd
}

func sortCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "sort",
		Short: "排序",
	}
	cmds := []*cobra.Command{
		{
			Use: "quick",
			Run: func(cmd *cobra.Command, args []string) {
				testData := []int{4, 5, 6, 3, 7, 9}
				fmt.Println("beforeSort: ", testData)
				sort2.Quick(testData, 0, len(testData)-1)
				fmt.Println("afterSort: ", testData)
			},
			Short: "快排",
		},
	}

	rootCmd.AddCommand(cmds...)
	return rootCmd
}

func hashmapCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hashmap",
		Short: "哈希",
	}
	cmds := []*cobra.Command{
		{
			Use: "open-addressing",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("开放寻址哈希表")
				openAddrTestData := map[string]string{
					"k-1": "v-1",
					"k-2": "v-2",
					"k-3": "v-3",
					"k-4": "v-4",
					"k-5": "v-5",
				}
				openAddrHash := hashmap.NewOpenAddrHash(5)
				hashmap.BatchPut(openAddrHash, openAddrTestData)
				fmt.Printf("key：%s, value: %s, expValue: %s\n", "k-3", openAddrHash.Get("k-3"), "v-3")
			},
			Short: "开放寻址法",
		},
		{
			Use: "separate-chaining",
			Run: func(cmd *cobra.Command, args []string) {
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
				linkHash := hashmap.NewLinkHash(5)
				hashmap.BatchPut(linkHash, linkTestData)
				fmt.Printf("key：%s, value: %s, expValue: %s\n", "k-6", linkHash.Get("k-6"), "v-6")
			},
			Short: "分离链表法",
		},
		{
			Use: "separate-chaining-rehash",
			Run: func(cmd *cobra.Command, args []string) {
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
				linkReHashHash := hashmap.NewLinkReHash(5)
				hashmap.BatchPut(linkReHashHash, linkRehashTestData)
				fmt.Printf("key：%s, value: %s, expValue: %s\n", "k-10", linkReHashHash.Get("k-10"), "v-10")
			},
			Short: "分离链表重散列",
		},
	}

	rootCmd.AddCommand(cmds...)
	return rootCmd
}

func treeCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tree",
		Short: "树",
	}

	cmds := []*cobra.Command{
		{
			Use: "max-heap",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("最大堆实现")
				heap := tree.NewIntHeap(5)
				testData := []int{1, 9, 4, 10, 22, 14, 8, 7}
				for _, v := range testData {
					fmt.Println("push: ", v)
					heap.Push(v)
				}

				fmt.Println("------------")
				for _, _ = range testData {
					fmt.Println("pop:", heap.Pop())
				}
			},
			Short: "最大堆实现",
		},
	}

	rootCmd.AddCommand(cmds...)
	return rootCmd
}

func otherCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "other",
		Short: "其他",
	}

	cmds := []*cobra.Command{
		{
			Use: "fibonacci",
			Run: func(cmd *cobra.Command, args []string) {

			},
			Short: "斐波那契数列",
		},
	}

	rootCmd.AddCommand(cmds...)
	return rootCmd
}
