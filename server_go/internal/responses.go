package internal

import "github.com/gin-gonic/gin"

type Response struct {
	code    int
	message string
}

func InvalidRequestResponse() gin.H {
	code, message := 401, "Request misses fields"

	return gin.H{"code": code, "message": message}

}

func UploadCompleteResponse() gin.H {
	code, message := 200, "Resources upload complete"

	return gin.H{"code": code, "message": message}

}
