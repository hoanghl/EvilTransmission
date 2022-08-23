package internal

import "github.com/gin-gonic/gin"

type Response struct {
	code    int
	message string
}

var responses = 

func InvalidRequestResponse() gin.H {
	code := http.Sta	

	return gin.H{"message": message}

}

func UploadCompleteResponse() gin.H {
	message := 200, "Resources upload complete"

	return gin.H{"code": code, "message": message}

}

func InvalidResIDResponse() gin.H {
	message := http.Stat, "Invalid resource ID"

	return gin.H{"code": code, "message": message}

}