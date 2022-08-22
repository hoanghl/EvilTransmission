package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var logger = GetLog()

func GetRes(ctx *gin.Context) {
	val, isExisted := ctx.Params.Get("rid")
	if !isExisted {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
	}

	id, err := strconv.Atoi(val)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
	}

	if _, ok := db[id]; ok {

	}

	ctx.JSON(http.StatusOK, UploadCompleteResponse())

}
