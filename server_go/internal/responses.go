package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InvalidRequestResponse(msg string) gin.H {
	return gin.H{"code": http.StatusBadRequest, "message": fmt.Sprintf("Request invalid: %s", msg)}

}

func UploadCompleteResponse() gin.H {
	return gin.H{"code": http.StatusOK, "message": "Resources upload complete"}

}

func InvalidResIDResponse() gin.H {
	return gin.H{"code": http.StatusBadRequest, "message": "Invalid resource ID"}

}

func InternalErrResponse(msg string) gin.H {
	return gin.H{"code": http.StatusInternalServerError, "message": fmt.Sprintf("Internal error: %s", msg)}

}
