package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRes(ctx *gin.Context) {

	// Get query args and do some sanity
	val := ctx.Query("rid")
	if val == "" {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
		return
	}

	thumbnail_ := ctx.Query("thumbnail")
	if thumbnail_ == "" {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
	}
	thumbnail, err := strconv.ParseBool(thumbnail_)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
	}

	if thumbnail {
		// TODO: HoangLe [Sep-20]: implement later
		// println("Continue this implementation")
	} else {
		if path, ok := db[val]; ok {
			if _, existed := os.Stat(path.(string)); existed == nil {
				data, err := os.ReadFile(path.(string))
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, nil)
				}

				ctx.Data(http.StatusOK, "image", data)
			}
		} else {
			ctx.JSON(http.StatusOK, InvalidResIDResponse())
		}
	}

}

func PostRes(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse())
	}
	defer file.Close()

	err = os.MkdirAll(".uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrDResponse())
	}

	dst, err := os.Create(fmt.Sprintf(".upload/%d.mp4", time.Now().UnixNano()))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrDResponse())
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrDResponse())
	}

	ctx.JSON(http.StatusOK, UploadCompleteResponse())

}
