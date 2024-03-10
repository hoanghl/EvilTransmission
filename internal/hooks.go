package internal

import (
	"database/sql"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type REQ_TYPE string

const PATH_DIR_ROOT = "/Users/macos/Projects/xButler/EvilTransmission/db"

const (
	IMG_JPG  REQ_TYPE = "image/JPG"
	IMG_JPEG REQ_TYPE = "image/JPEG"
	IMG_PNG  REQ_TYPE = "image/png"
	VID_MP4  REQ_TYPE = "video/mp4"
)

func GetMediaInfo(ctx *gin.Context) {
	list_info, err := Conf.DB.GetMediaInfo()
	if err != nil {
		logger.Error("Error while loading")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, list_info)

}

func GetMedia(ctx *gin.Context) {
	_res_id := ctx.Param("res_id")
	res_id, err := strconv.Atoi(_res_id)
	if err != nil {
		logger.Error("Invalid res_id: ", res_id)
		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	// Get filename
	filename, res_type, err := Conf.DB.GetMedia(res_id)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, "Resource not found")
		return
	} else if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	logger.Debug(filename, " - ", res_type)

	// Read file
	filepath := path.Join(PATH_DIR_ROOT, res_type, filename)
	if _, err = os.Stat(filepath); err == nil {
		file, err := os.ReadFile(filepath)
		if err != nil {
			logger.Error("Read file error: ", err)
			ctx.JSON(http.StatusInternalServerError, "")
			return
		}
		var content_type string
		switch res_type {
		case "image":
			content_type = "image/png"
		case "thumbnail":
			content_type = "image/png"
		case "video":
			content_type = "application/mp4"
		default:
			panic(content_type)
		}

		ctx.Data(http.StatusOK, content_type, file)
		return
	}
}

// func GetRes(ctx *gin.Context) {

// 	// Get query args and do some sanity
// 	rid := ctx.Query("rid")
// 	thumbnail_ := ctx.Query("thumbnail")

// 	if rid == "" && thumbnail_ == "" {
// 		// Get all entries' id
// 		entries, err := Conf.DB.GetAllEntries()
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, ErrReqResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, entries)

// 		logger.Info("Sent list of ids")

// 		return
// 	}

// 	if rid == "" {
// 		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Request not contain field 'rid'"))
// 		logger.Error("Request not contain field 'rid'")
// 		return
// 	}

// 	if thumbnail_ == "" {
// 		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Request not contain field 'thumbnail'"))
// 		logger.Error("Request not contain field 'thumbnail'")
// 		return
// 	}
// 	thumbnail, err := strconv.ParseBool(thumbnail_)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Invalid value/type for field 'thumbnail'"))
// 		logger.Error("Invalid value/type for field 'thumbnail'")
// 		return
// 	}

// 	// Check existence of id in db and in storage
// 	entry, err := Conf.DB.GetEntry(DBEntry{ID: rid})
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, InternalErrResponse(err.Error()))
// 		logger.Error(err)
// 		return
// 	}

// 	logger.Infof("Found entry: %s", entry.ID)

// 	if _, existed := os.Stat(entry.Path); existed != nil {
// 		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
// 		return
// 	}

// 	// Start processing
// 	if thumbnail {
// 		logger.Info("Query: thumbnail")

// 		ext := filepath.Ext(entry.Path)
// 		if !strings.HasSuffix(ext, ".mp4") {
// 			ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Image cannot exist thumbnail"))
// 			logger.Error(err)
// 			return

// 		}
// 		pathWithoutExt := strings.TrimSuffix(entry.Path, filepath.Ext(entry.Path))
// 		pathThumb := fmt.Sprintf("%s_thumb.png", pathWithoutExt)
// 		data, err := os.ReadFile(pathThumb)
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
// 			logger.Error(err)
// 			return
// 		}

// 		ctx.Data(http.StatusOK, "image", data)

// 	} else {
// 		logger.Info("Query: entry")

// 		if _, existed := os.Stat(entry.Path); existed == nil {
// 			data, err := os.ReadFile(entry.Path)
// 			if err != nil {
// 				ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
// 				logger.Error(err)
// 				return
// 			}

// 			ext := filepath.Ext(entry.Path)
// 			if strings.HasSuffix(ext, "mp4") {
// 				ctx.Data(http.StatusOK, "video", data)
// 				logger.Info("Sent video")
// 			} else {
// 				ctx.Data(http.StatusOK, "image", data)
// 				logger.Info("Sent image")
// 			}
// 		} else {
// 			ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot read resource"))
// 			logger.Errorf("Cannot read resource: %s", entry.Path)
// 			return
// 		}

// 	}

// }

// func PostRes(ctx *gin.Context) {
// 	form, err := ctx.MultipartForm()
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Extracting MultipartForm error"))
// 		logger.Error(err)
// 		return
// 	}

// 	if len(form.Value["type"]) == 0 || len(form.File["file"]) == 0 {
// 		ctx.JSON(http.StatusBadRequest, InvalidRequestResponse("Field not found: 'type' or 'file'"))
// 		logger.Error(err)
// 		return
// 	}
// 	fileType := form.Value["type"][0]
// 	fileHeader := form.File["file"][0]

// 	// Get hash and check
// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot opean loaded file"))
// 		logger.Error(err)
// 		return
// 	}
// 	defer file.Close()

// 	// _, err = Conf.DB.GetEntry(DBEntry{Hashval: hash})
// 	// if err != nil && err != sql.ErrNoRows {
// 	// 	ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot query resource"))
// 	// 	logger.Error(err)
// 	// 	return
// 	// } else if err == nil {
// 	// 	ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Resource existed"))
// 	// 	return
// 	// }

// 	// Save uploaded file
// 	err = os.MkdirAll("uploads", os.ModePerm)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot create directory to store things"))
// 		logger.Error(err)
// 		return
// 	}

// 	var pathRes string
// 	fileID := uuid.NewString()
// 	switch fileType {
// 	case string(IMG_JPG):
// 	case string(IMG_JPEG):
// 		pathRes = fmt.Sprintf("uploads/%s.jpg", fileID)
// 	case string(VID_MP4):
// 		pathRes = fmt.Sprintf("uploads/%s.mp4", fileID)
// 	default:
// 		ctx.JSON(http.StatusInternalServerError, InvalidRequestResponse("Incorrect value for field 'fileType'"))
// 		logger.Error(err)
// 		return
// 	}

// 	err = ctx.SaveUploadedFile(fileHeader, pathRes)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot save resource file in server"))
// 		logger.Error(err)
// 		return
// 	}

// 	// Create thumbnail
// 	if fileType == string(VID_MP4) {
// 		ExtractThumbnail(pathRes)
// 	}

// 	// Send info to server
// 	err = Conf.DB.InsertEntry(DBEntry{
// 		ID:         fileID,
// 		Path:       pathRes,
// 		LastUpdate: time.Now(),
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, InternalErrResponse("Cannot insert resource to database"))
// 		logger.Error(err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, UploadCompleteResponse(fileID))

// }
