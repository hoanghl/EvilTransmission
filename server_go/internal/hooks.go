package internal

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func GetRes(ctx *gin.Context) {

	// Get query args and do some sanity
	rid := ctx.Query("rid")
	if rid == "" {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Request not contain field 'rid'"))
		logger.Error("Request not contain field 'rid'")
		return
	}

	thumbnail_ := ctx.Query("thumbnail")
	if thumbnail_ == "" {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Request not contain field 'thumbnail'"))
		logger.Error("Request not contain field 'thumbnail'")
		return
	}
	thumbnail, err := strconv.ParseBool(thumbnail_)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Invalid value/type for field 'thumbnail'"))
		logger.Error("Invalid value/type for field 'thumbnail'")
		return
	}
	fmt.Print(thumbnail)

	// if path, ok := db[rid]; ok {
	// 	if thumbnail {
	// 		ext := filepath.Ext(path.(string))

	// 		logger.Infof("Extension: %s", ext)

	// 		if !strings.HasSuffix(ext, ".mp4") {
	// 			ctx.JSON(http.StatusInternalServerError, InternalErrDResponse("Image cannot exist thumbnail"))
	// 			logger.Error(err)
	// 			return

	// 		}
	// 		pathWithoutExt := strings.TrimSuffix(path.(string), filepath.Ext(path.(string)))
	// 		pathThumb := fmt.Sprintf("%s_thumb.png", pathWithoutExt)
	// 		data, err := os.ReadFile(pathThumb)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusInternalServerError, InternalErrDResponse("Cannot read resource"))
	// 			logger.Error(err)
	// 			return
	// 		}

	// 		ctx.Data(http.StatusOK, "image", data)

	// 	} else {
	// 		if _, existed := os.Stat(path.(string)); existed == nil {
	// 			data, err := os.ReadFile(path.(string))
	// 			if err != nil {
	// 				ctx.JSON(http.StatusInternalServerError, InternalErrDResponse("Cannot read resource"))
	// 				logger.Error(err)
	// 				return
	// 			}

	// 			logger.Infof("len data: %d", len(data))

	// 			ext := filepath.Ext(path.(string))
	// 			if strings.HasSuffix(ext, "mp4") {
	// 				ctx.Data(http.StatusOK, "video", data)
	// 				logger.Info("Sent video")
	// 			} else {
	// 				ctx.Data(http.StatusOK, "image", data)
	// 				logger.Info("Sent image")
	// 			}
	// 		} else {
	// 			ctx.JSON(http.StatusInternalServerError, InternalErrDResponse("Cannot read resource"))
	// 			logger.Errorf("Cannot read resource: %s", path.(string))
	// 			return
	// 		}

	// 	}

	// } else {
	// 	ctx.JSON(http.StatusOK, InvalidResIDResponse())
	// }
}

func PostRes(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Extracting MultipartForm error"))
		logger.Error(err)
		return
	}

	if len(form.Value["type"]) == 0 || len(form.File["file"]) == 0 {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Field not found: 'type' or 'file'"))
		logger.Error(err)
		return
	}
	fileType := form.Value["type"][0]
	file := form.File["file"][0]

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrDResponse("Cannot create directory to store things"))
		logger.Error(err)
		return
	}

	var pathRes string
	fileID := uuid.NewString()
	if fileType == "image/jpg" || fileType == "image/jpeg" {
		pathRes = fmt.Sprintf("uploads/%s.jpg", fileID)
	} else if fileType == "image/png" {
		pathRes = fmt.Sprintf("uploads/%s.", fileID)
	} else if fileType == "image/mp4" {
		pathRes = fmt.Sprintf("uploads/%s.mp4", fileID)
	} else {
		ctx.JSON(http.StatusInternalServerError, InvalidRequestResponse("Incorrect value for field 'fileType'"))
		logger.Error(err)
		return
	}

	err = ctx.SaveUploadedFile(file, pathRes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrDResponse("Cannot save resource file in server"))
		logger.Error(err)
		return
	}

	// Create thumbnail
	if fileType == "image/mp4" {
		ExtractThumbnail(pathRes)
	}

	ctx.JSON(http.StatusOK, UploadCompleteResponse())

}
