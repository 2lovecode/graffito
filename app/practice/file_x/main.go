package file_x

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

func Run1() {
	runtimePath, _ := os.Getwd()

	from := runtimePath + "/data/file/from.dat"
	to := runtimePath + "/data/file/to.dat"

	fmt.Println(from)
	fmt.Println("123456")
	copyFile(from, to)
}
