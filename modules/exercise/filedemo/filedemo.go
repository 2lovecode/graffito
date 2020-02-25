package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(from string, to string) {
	fromFile, err := os.Open(from)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fromFile.Close()

	toFile, err := os.Create(to)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer toFile.Close()

	io.Copy(toFile, fromFile)
}

func cat(f *os.File) {
	var buf [512]byte
	switch ns, err := f.Read(buf[:]); true {
	case ns < 0:
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	case ns == 0:
		return
	case ns > 0:
		if nw, ew := os.Stdout.Write(buf[0:ns]); nw != ns {
			fmt.Fprintf(os.Stderr, "error: %s\n", ew.Error())
		}
	}
}

func main() {
	//var msg, hh string
	//var err error
	//
	//fmt.Scanln(&hh)
	//
	//fmt.Println(hh, "morning")
	//
	//mReader := bufio.NewReader(os.Stdin)
	//
	//msg, err = mReader.ReadString('\n')
	//
	//if err == nil {
	//	fmt.Println(msg)
	//} else {
	//	fmt.Println("Error")
	//}
	//
	//fileHandle, err := os.Open("")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer fileHandle.Close()

	runtimePath, _ := os.Getwd()

	from := runtimePath + "/exercise/filedemo/from.dat"
	to := runtimePath + "/exercise/filedemo/to.dat"

	fmt.Println(from)
	fmt.Println("123456")
	copyFile(from, to)


}
