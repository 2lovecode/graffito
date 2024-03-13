package algorithm

import "github.com/spf13/cobra"

type LRUCache struct {
}

func (lc *LRUCache) Put() {

}

func (lc *LRUCache) Get() {

}

func NewLRUCache() *cobra.Command {
	return &cobra.Command{
		Use: "lru-cache",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
