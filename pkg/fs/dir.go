package fs

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取项目根目录路径
func GetProjectRootPath() (string, error) {
	// 获取当前工作目录路径
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 逐级向上检查文件系统，直到找到包含 go.mod 文件的目录
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, nil
		}

		// 到达文件系统的根目录时退出循环
		if wd == filepath.Dir(wd) {
			break
		}

		// 向上一级目录
		wd = filepath.Dir(wd)
	}

	return "", fmt.Errorf("unable to determine project root directory")
}

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
