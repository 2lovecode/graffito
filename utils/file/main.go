package file

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCurrentAbsPath() string {
	dir := GetCurrentAbsPathByExecutable()
	if !strings.Contains(GetProjectAbsDir(), dir) {
		return GetCurrentAbsPathByCaller(2)
	}
	return dir
}

func GetProjectAbsDir() string {
	dir, err := filepath.Abs("")
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// 获取当前执行文件绝对路径
func GetCurrentAbsPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径
func GetCurrentAbsPathByCaller(depth int) string {
	var absPath string
	_, filename, _, ok := runtime.Caller(depth)
	if ok {
		absPath = path.Dir(filename)
	}
	return absPath
}
