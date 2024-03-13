package chapter1_map

import (
	"context"
	"fmt"
	"github.com/2lovecode/graffito/internal/app/practice/base"
	"github.com/2lovecode/graffito/internal/app/practice/chapter1_map/help"
	"github.com/2lovecode/graffito/internal/app/practice/chapter1_map/qa1"
	"math/rand"
	"sort"

	"github.com/speps/go-hashids/v2"
	"github.com/spf13/cobra"
)

// map的实现原理
// 什么类型可以作为key，什么类型可以作为value
// 添加-扩容
// 删除-缩容
// 遍历无序
// 判等

type CMap struct {
	numbers map[int]base.QA
}

func New() *CMap {
	cm := &CMap{}
	cm.init()
	return cm
}

func (cm *CMap) GetQA(ctx context.Context, number int) base.QA {
	if _, ok := cm.numbers[number]; ok {
		return cm.numbers[number]
	}
	return help.New()
}

func (cm *CMap) init() {
	cm.numbers = map[int]base.QA{
		1: qa1.New(),
	}
}

func Run1() {
	//map赋值是指针赋值
	map1 := map[string]string{
		"a": "b",
	}
	var map2 map[string]string

	map2 = map1

	map2["c"] = "d"
	fmt.Println(map1, map2)
	//map是无序的
	gc := map[string]string{"a": "aaa", "b": "bbbbbbb", "c": "cccc"}
	for index, item := range gc {
		fmt.Println(index, item)
	}

	keys := make([]string, len(gc))

	i := 0
	for index, _ := range gc {
		keys[i] = index
		i++
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("key: %s value: %s\n", key, gc[key])
	}

	//map键值对调
	kv := map[string]int{"a": 1, "b": 2, "c": 3}
	vk := make(map[int]string, len(kv))

	for key, value := range kv {
		vk[value] = key
	}

	for k, v := range vk {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}

func Run2() {
	type SSS struct {
		D int
	}

	s1 := SSS{
		D: 1,
	}
	s2 := SSS{
		D: 1,
	}

	var tt = map[*SSS]int{}
	var ttt = map[SSS]int{}

	tt[&s1] = 33
	fmt.Println(tt[&s1], tt[&s2])

	ttt[s1] = 22
	fmt.Println(ttt[s1], ttt[s2])

	fmt.Println(s1 == s2)
}

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "map",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				switch args[0] {
				case "1":
					Run1()
				case "2":
					Run2()
				}
			}
		},
	}
}

func DefaultHash() *hashids.HashID {
	hash := hashids.NewData()
	hash.Salt = "APpAUptXsssdfffww"
	hash.MinLength = 6
	hash.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890+-*=@#$%"
	h, _ := hashids.NewWithData(hash)
	return h
}

func EncodeHash(num int64) string {
	h := DefaultHash()

	e, _ := h.EncodeInt64([]int64{num})
	return e
}

func DecodeHash(str string) int64 {
	h := DefaultHash()
	d, _ := h.DecodeInt64WithError(str)
	return d[0]
}

func consistentShuffleInPlace(alphabet []rune, salt []rune) {
	if len(salt) == 0 {
		return
	}

	for i, v, p := len(alphabet)-1, 0, 0; i > 0; i-- {
		p += int(salt[v])
		j := (int(salt[v]) + v + p) % i
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
		v = (v + 1) % len(salt)
	}
	return
}

func fisherYatesShuffleWithConsistency(arr []int, seed int64) {
	rng := rand.New(rand.NewSource(seed))

	for i := len(arr) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
