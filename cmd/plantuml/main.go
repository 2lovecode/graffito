package main

import (
	"fmt"
	"graffito/utils/file"
	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {

	dataDirName := "data"
	imagesDirName := "images"
	currentPath := file.GetCurrentAbsPath()

	dataDirPath := fmt.Sprintf("%s/%s", currentPath, dataDirName)
	imagesDirPath := fmt.Sprintf("%s/%s", currentPath, imagesDirName)
	plantumlPath := fmt.Sprintf("%s/plantuml-1.2022.7.jar", currentPath)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				dataFileName := ""
				imageFilePath := ""

				dataFileName = path.Base(event.Name)
				dataFileExt := path.Ext(dataFileName)
				imageFilePath = fmt.Sprintf("%s/%s.png", imagesDirPath, strings.TrimSuffix(dataFileName, dataFileExt))
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					generateCmd := exec.Command("java", "-jar", plantumlPath, event.Name, "-o", imageFilePath)
					err := generateCmd.Run()
					if err != nil {
						log.Printf("Err: %v", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dataDirPath)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}
