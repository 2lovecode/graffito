package main

import (
	"fmt"
)

func Echo(s  string) string{
	return s + "suffix"
}

func main() {
	//env := os.Environ()
	//proc := &os.ProcAttr{
	//	Env:   env,
	//	Files: []*os.File{
	//		os.Stdin,
	//		os.Stdout,
	//		os.Stderr,
	//	},
	//}
	//pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, proc)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}
	//
	//fmt.Printf("pid is : %v", pid)
	fmt.Println(Echo("hello"))
}
