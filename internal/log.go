package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = GetLog(logrus.DebugLevel)

func GetLog(logLevel logrus.Level) *logrus.Logger {

	return &logrus.Logger{
		Out:   os.Stderr,
		Level: logLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			PadLevelText:    true,
		},
	}
}
