package routers

import (
	"net/http"

	res "gin-gorm/internals"

	"github.com/gin-gonic/gin"
)

func (r routes) RegisterHealthcheckRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/healthchecks")

	route.GET("/app", GetAppHeartbeat)
	route.GET("/db", GetDBHeartbeat)
}

func GetAppHeartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, res.SetOk(nil))
}

func GetDBHeartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, res.SetOk(nil))
}
