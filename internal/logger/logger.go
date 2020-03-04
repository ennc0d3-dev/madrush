package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
}

func Log(msg interface{}) {
	log.Debug(msg)
}
func Fatal(msg interface{}) {
	log.Fatal(msg)
}
