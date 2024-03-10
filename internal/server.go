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
	r.GET("/res/", GetMediaInfo)
	r.GET("/res/:res_id", GetMedia)
	// r.POST("/res/", PostRes)
}

func StartServer() {
	r.Run(fmt.Sprintf(":%d", Conf.PORT))
}
