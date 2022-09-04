package internals

import "github.com/gin-gonic/gin"

func SetOk(data interface{}) gin.H {
	return gin.H{
		"data":    data,
		"message": "success",
	}
}
