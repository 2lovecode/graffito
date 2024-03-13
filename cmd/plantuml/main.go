package main

import (
	"fmt"
	"github.com/2lovecode/graffito/pkg/fs"
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

type runtime struct {
	projectRootDir string
	appDir         string
	dataDir        string
	imageDir       string
	plantumlLib    string
}

func main() {
	r := buildRuntime()
	if checkErr := checkEnv(r); checkErr != nil {
		log.Fatal(checkErr)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		closeErr := watcher.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					generateCmd := exec.Command("java", "-jar", r.plantumlLib, event.Name, "-o", r.imageDir)
					runErr := generateCmd.Run()
					if runErr != nil {
						log.Printf("Err: %v", err)
						return
					}
				}
			case watchErr, ok := <-watcher.Errors:
				log.Println("error:", watchErr)
				if !ok {
					return
				}
			}
		}
	}()

	err = watcher.Add(r.dataDir)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func buildRuntime() *runtime {
	projectRoot, _ := fs.GetProjectRootPath()

	plantumlPath := fmt.Sprintf("%s/third_party/plantuml/plantuml-1.2022.7.jar", projectRoot)

	appPath := fs.GetCurrentAbsPath()

	dataDirName := "data"
	imagesDirName := "images"

	dataDirPath := fmt.Sprintf("%s/%s", appPath, dataDirName)
	imagesDirPath := fmt.Sprintf("%s/%s", appPath, imagesDirName)

	return &runtime{
		projectRootDir: projectRoot,
		appDir:         appPath,
		dataDir:        dataDirPath,
		imageDir:       imagesDirPath,
		plantumlLib:    plantumlPath,
	}
}

func checkEnv(r *runtime) error {

	if _, plantumlLibErr := os.Stat(r.plantumlLib); plantumlLibErr != nil {
		return plantumlLibErr
	}

	// 执行 java -version 命令
	cmd := exec.Command("java", "-version")

	// 执行命令并捕获输出和错误
	output, err := cmd.CombinedOutput()

	// 检查命令执行结果
	if err != nil {
		return fmt.Errorf("java环境未安装")
	}

	// 输出命令执行结果
	fmt.Println(string(output))

	return nil
}
