package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRes(ctx *gin.Context) {
	println(ctx.Params)

	ctx.JSON(http.StatusOK, gin.H{"data": "nothin'"})

}

func PrintThingsd() {
	fmt.Println("hellow worlds")
}
