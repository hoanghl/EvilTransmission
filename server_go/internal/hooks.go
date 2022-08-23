package internal

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var logger = GetLog()

func GetRes(ctx *gin.Context) {
	val, isExisted := ctx.Params.Get("rid")
	if !isExisted {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
	}

	_, err := strconv.Atoi(val)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
	}

	if path, ok := db[val]; ok {

		if _, existed := os.Stat(path.(string)); existed == nil {
			data, err := os.ReadFile(path.(string))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, nil)
			}

			ctx.Data(http.StatusOK, "image", data)
		}
	} else {
		ctx.JSON(http.StatusOK, UploadCompleteResponse())
	}

}
