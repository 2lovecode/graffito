package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:05", "2020-03-10 11:11:11")
	if err == nil {
		fmt.Println(t.After(time.Now()))
		fmt.Println(t.Format("2006-01-02"))
	}
}
