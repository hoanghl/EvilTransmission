package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLog() *logrus.Logger {
	return &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			PadLevelText:    true,
		},
	}
}
