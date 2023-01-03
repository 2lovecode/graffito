package slice_x

import (
	"fmt"
	"net/url"
	"sort"
)

func myAppend(s []int) []int {
	s = append(s, 10)
	return s
}

func changeSlice(s []int) {
	s[0] = 10
}
func Run_3() {
	// 实验作为函数参数传递
	s := []int{1, 2, 3}

	newS := myAppend(s)

	fmt.Println(s, newS)

	base := []int{0, 1, 2, 3, 4, 5, 6}

	o := base[2:4:5]
	newO := myAppend(o)
	//newO = myAppend(newO)
	fmt.Println(base, o, newO)

	base1 := []int{0, 1, 2, 3, 4, 5, 6}
	o1 := base1[2:4]
	newO1 := myAppend(o1)
	//newO1 = myAppend(newO1)
	fmt.Println(base1, o1, newO1)

	base3 := []int{0, 1, 2}
	changeSlice(base3)
	fmt.Println(base3)
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
