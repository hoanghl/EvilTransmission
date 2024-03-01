package internal

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type REQ_TYPE string

const (
	IMG_JPG  REQ_TYPE = "image/JPG"
	IMG_JPEG REQ_TYPE = "image/JPEG"
	IMG_PNG  REQ_TYPE = "image/png"
	VID_MP4  REQ_TYPE = "video/mp4"
)

func GetRes(ctx *gin.Context) {

	// Get query args and do some sanity
	rid := ctx.Query("rid")
	thumbnail_ := ctx.Query("thumbnail")

	if rid == "" && thumbnail_ == "" {
		// Get all entries' resid
		entries, err := Conf.DB.GetAllEntry()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ErrReqResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, entries)

		logger.Info("Sent list of resID")

		return
	}

	if rid == "" {
		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Request not contain field 'rid'"))
		logger.Error("Request not contain field 'rid'")
		return
	}

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

	// Check existence of resid in db and in storage
	entry, err := Conf.DB.GetEntry(DBEntry{ResID: rid})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, InternalErrResponse(err.Error()))
		logger.Error(err)
		return
	}

	logger.Infof("Found entry: %s", entry.ResID)

	if _, existed := os.Stat(entry.Path); existed != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
		return
	}

	// Start processing
	if thumbnail {
		logger.Info("Query: thumbnail")

		ext := filepath.Ext(entry.Path)
		if !strings.HasSuffix(ext, ".mp4") {
			ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Image cannot exist thumbnail"))
			logger.Error(err)
			return

		}
		pathWithoutExt := strings.TrimSuffix(entry.Path, filepath.Ext(entry.Path))
		pathThumb := fmt.Sprintf("%s_thumb.png", pathWithoutExt)
		data, err := os.ReadFile(pathThumb)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
			logger.Error(err)
			return
		}

		ctx.Data(http.StatusOK, "image", data)

	} else {
		logger.Info("Query: entry")

		if _, existed := os.Stat(entry.Path); existed == nil {
			data, err := os.ReadFile(entry.Path)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
				logger.Error(err)
				return
			}

			ext := filepath.Ext(entry.Path)
			if strings.HasSuffix(ext, "mp4") {
				ctx.Data(http.StatusOK, "video", data)
				logger.Info("Sent video")
			} else {
				ctx.Data(http.StatusOK, "image", data)
				logger.Info("Sent image")
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
			logger.Errorf("Cannot read resource: %s", entry.Path)
			return
		}

	}

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
	fileHeader := form.File["file"][0]

	// Get hash and check
	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot opean loaded file"))
		logger.Error(err)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	io.Copy(&buf, file)
	hash := GetFileHash(buf.Bytes())

	_, err = Conf.DB.GetEntry(DBEntry{Hashval: hash})
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot query resource"))
		logger.Error(err)
		return
	} else if err == nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Resource existed"))
		return
	}

	// Save uploaded file
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot create directory to store things"))
		logger.Error(err)
		return
	}

	var pathRes string
	fileID := uuid.NewString()
	switch fileType {
	case string(IMG_JPG):
	case string(IMG_JPEG):
		pathRes = fmt.Sprintf("uploads/%s.jpg", fileID)
	case string(VID_MP4):
		pathRes = fmt.Sprintf("uploads/%s.mp4", fileID)
	default:
		ctx.JSON(http.StatusInternalServerError, InvalidRequestResponse("Incorrect value for field 'fileType'"))
		logger.Error(err)
		return
	}

	err = ctx.SaveUploadedFile(fileHeader, pathRes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot save resource file in server"))
		logger.Error(err)
		return
	}

	// Create thumbnail
	if fileType == string(VID_MP4) {
		ExtractThumbnail(pathRes)
	}

	// Send info to server
	err = Conf.DB.InsertEntry(DBEntry{
		ResID:      fileID,
		Path:       pathRes,
		LastUpdate: time.Now(),
		Hashval:    hash,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot insert resource to database"))
		logger.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, UploadCompleteResponse(fileID))

}
