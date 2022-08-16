package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routing ////////////////////////
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{"data": "hello"})
	})

	r.Run()
}
