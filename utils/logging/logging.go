package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

var TLog = logrus.New()

func init () {
	TLog.SetOutput(os.Stderr)
}