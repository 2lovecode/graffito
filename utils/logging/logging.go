package logging

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	PanicLevel int = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)


type LogFile struct {
	level    int
	logTime  int64
	fileName string
	fileFd   *os.File
}

var logFile LogFile

func Config(logFolder string, level int) {
	logFile.fileName = logFolder
	logFile.level = level

	log.SetOutput(logFile)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func SetLevel(level int) {
	logFile.level = level
}

func DebugF(format string, args ...interface{}) {
	if logFile.level >= DebugLevel {
		log.SetPrefix("debug ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func InfoF(format string, args ...interface{}) {
	if logFile.level >= InfoLevel {
		log.SetPrefix("info ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func WarnF(format string, args ...interface{}) {
	if logFile.level >= WarnLevel {
		log.SetPrefix("warn ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func ErrorF(format string, args ...interface{}) {
	if logFile.level >= ErrorLevel {
		log.SetPrefix("error ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func FatalF(format string, args ...interface{}) {
	if logFile.level >= FatalLevel {
		log.SetPrefix("fatal ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (me LogFile) Write(buf []byte) (n int, err error) {
	if me.fileName == "" {
		fmt.Printf("consol: %s", buf)
		return len(buf), nil
	}

	if logFile.logTime+3600 < time.Now().Unix() {
		logFile.createLogFile()
		logFile.logTime = time.Now().Unix()
	}

	if logFile.fileFd == nil {
		return len(buf), nil
	}

	return logFile.fileFd.Write(buf)
}

func (me *LogFile) createLogFile() {
	logdir := "./"
	if index := strings.LastIndex(me.fileName, "/"); index != -1 {
		logdir = me.fileName[0:index] + "/"
		os.MkdirAll(me.fileName[0:index], os.ModePerm)
	}

	now := time.Now()
	filename := fmt.Sprintf("%s_%04d%02d%02d_%02d%02d", me.fileName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	if err := os.Rename(me.fileName, filename); err == nil {
		go func() {
			tarCmd := exec.Command("tar", "-zcf", filename+".tar.gz", filename, "--remove-files")
			tarCmd.Run()

			rmCmd := exec.Command("/bin/sh", "-c", "find "+logdir+` -type f -mtime +2 -exec rm {} \;`)
			rmCmd.Run()
		}()
	}

	for index := 0; index < 10; index++ {
		if fd, err := os.OpenFile(me.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeExclusive); nil == err {
			me.fileFd.Sync()
			me.fileFd.Close()
			me.fileFd = fd
			break
		}

		me.fileFd = nil
	}
}