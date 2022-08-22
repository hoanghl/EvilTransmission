package main

import (
	"github.com/io/server_go/internal"
)

var logger = internal.GetLog()

func main() {
	logger.Info("Initialize components")
	internal.Initialize()

	logger.Info("Start server")
	internal.StartServer()

}
