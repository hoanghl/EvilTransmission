package main

import (
	"github.com/io/server_go/internal"
	"github.com/sirupsen/logrus"
)

var logger = internal.GetLog(logrus.DebugLevel)

func main() {
	logger.Info("Initialize components")
	internal.Initialize()

	logger.Info("Start server")
	internal.StartServer()

}
