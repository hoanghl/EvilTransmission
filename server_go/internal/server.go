package internal

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitServer() {
	// Set up GIN
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.Default()

	// Routing
	r.GET("/res/", GetRes)
}

func StartServer() {
	port, _ := os.LookupEnv("PORT")

	r.Run(fmt.Sprintf(":%s", port))
}
