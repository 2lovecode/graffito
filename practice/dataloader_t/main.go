package dataloader_t

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/graph-gophers/dataloader"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "dataloader",
		Run: func(cmd *cobra.Command, args []string) {
			Run(context.Background())
		},
		Short: "dataloader测试",
	}
}
func Run(ctx context.Context) error {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		for _, key := range keys {
			switch key.String() {
			case "key1":
				fmt.Println("key1-run")
				results = append(results, &dataloader.Result{Data: "hello"})
			case "key2":
				fmt.Println("key2-run")
				results = append(results, &dataloader.Result{Data: "world"})
			}

		}
		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)
	// loader := dataloader.NewBatchedLoader(batchFn, dataloader.WithCache(&dataloader.NoCache{}))

	thunk2 := loader.Load(context.TODO(), dataloader.StringKey("key2"))
	fmt.Println("thunk2-load")
	result2, _ := thunk2()

	thunk := loader.Load(context.TODO(), dataloader.StringKey("key1"))
	fmt.Println("thunk-load")

	thunk3 := loader.Load(context.TODO(), dataloader.StringKey("key1"))
	fmt.Println("thunk3-load")

	result, _ := thunk()
	result3, _ := thunk3()

	log.Printf("value: %#v", result)
	log.Printf("value: %#v", result2)
	log.Printf("value: %#v", result3)
	return nil
}
