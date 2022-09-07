package helpers

import "github.com/gin-gonic/gin"

func SetOk(data interface{}) gin.H {
	return gin.H{
		"data":    data,
		"message": "success",
	}
}

func SetErr(message string) gin.H {
	return gin.H{
		"data":    nil,
		"message": message,
	}
}
