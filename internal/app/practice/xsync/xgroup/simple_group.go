package xgroup

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/practice/base"
	"sync"
)

type SimpleGroup struct {
}

func NewSimpleGroup() *SimpleGroup {
	return &SimpleGroup{}
}

func (s *SimpleGroup) Name() base.Name {
	return base.SimpleGroup
}

func (s *SimpleGroup) Description() string {
	return ""
}

func (s *SimpleGroup) Run(ctx context.Context) (string, error) {

	wg := &sync.WaitGroup{}

	total := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ii := i
		go func() {
			defer wg.Done()
			total += ii
		}()
	}

	wg2 := &sync.WaitGroup{}
	wg2.Add(2)

	go func() {
		defer wg2.Done()
		fmt.Println("wait 1 start")
		wg.Wait()
		fmt.Println("total 1 ", total)
		fmt.Println("wait 1 end")
	}()

	go func() {
		defer wg2.Done()
		fmt.Println("wait 2 start")
		wg.Wait()
		fmt.Println("total 2 ", total)
		fmt.Println("wait 2 end")
	}()

	fmt.Println("wait 3 start")
	wg.Wait()
	fmt.Println("total 3 ", total)
	fmt.Println("wait 3 end")

	wg2.Wait()

	return "", nil
}
