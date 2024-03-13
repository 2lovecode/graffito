package chapter2_slice

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/practice/base"
	"github.com/2lovecode/graffito/internal/app/practice/chapter2_slice/help"
	"github.com/2lovecode/graffito/internal/app/practice/chapter2_slice/qa1"
	"github.com/spf13/cobra"
	"net/url"
	"sort"
)

type CSlice struct {
	numbers map[int]base.QA
}

func New() *CSlice {
	cm := &CSlice{}
	cm.init()
	return cm
}

func (cm *CSlice) GetQA(ctx context.Context, number int) base.QA {
	if _, ok := cm.numbers[number]; ok {
		return cm.numbers[number]
	}
	return help.New()
}

func (cm *CSlice) init() {
	cm.numbers = map[int]base.QA{
		1: qa1.New(),
	}
}

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "slice", Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "expansion-one-by-one":
				ExpansionOneByOne()
			case "expansion-multiple":
				ExpansionMultiple()
			case "func-params":
				FuncParams()
			case "4":
				Run_4()
			case "sorting":
				Sorting()
			case "substr":
				Substr()
			}
		}

	}}
}

func Run_4() {
	s1 := []int{1, 2, 3, 4}

	s2 := []*int{}

	for _, v := range s1 {
		s2 = append(s2, &v)
	}

	for _, v := range s2 {
		fmt.Println(*v)
	}
	originalUrl := "https://pic.ziroom.com/steward_images/60025745.png"
	finalUrl := originalUrl
	originalUrlObj, err := url.Parse(originalUrl)

	if err == nil {
		if originalUrlObj.RawQuery == "" {
			finalUrl = finalUrl + "?imageMogr2/thumbnail/384x/format/jpg"
		} else {
			finalUrl = finalUrl + "&imageMogr2/thumbnail/384x/format/jpg"
		}
	}
	// "imageMogr2/thumbnail/384x/format/jpg"
	fmt.Println(finalUrl)
}

func Sorting() {

	sortStrings := []string{"2", "7", "1", "5"}

	sort.Slice(sortStrings, func(i, j int) bool {
		return sortStrings[i] < sortStrings[j]
	})

	fmt.Println(sortStrings)
}
