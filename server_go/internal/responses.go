package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InvalidRequestResponse() gin.H {
	return gin.H{"code": http.StatusBadRequest, "message": "Request invalid"}

}

func UploadCompleteResponse() gin.H {
	return gin.H{"code": http.StatusOK, "message": "Resources upload complete"}

}

func InvalidResIDResponse() gin.H {
	return gin.H{"code": http.StatusBadRequest, "message": "Invalid resource ID"}

}

func InternalErrDResponse() gin.H {
	return gin.H{"code": http.StatusInternalServerError, "message": "Internal error"}

}
