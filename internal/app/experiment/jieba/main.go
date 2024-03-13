package jieba

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yanyiwu/gojieba"
)

func Run() {
	fmt.Println(GetDataPath("123"))

	dpath := GetDataPath("dict.txt")
	hpath := GetDataPath("hmm_model.utf8")
	ipath := GetDataPath("idf.utf8")
	upath := GetDataPath("user.dict.utf8")
	spath := GetDataPath("stop_words.utf8")

	jb := gojieba.NewJieba(dpath, hpath, upath, ipath, spath)

	// jb := gojieba.NewJieba()
	defer jb.Free()

	list := jb.Tag("大学")
	fmt.Println(list)
}

func GetDataPath(filename string) string {

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	appDataPath := filepath.Join(workPath, "experiment/jieba/data", filename)

	return appDataPath
}
