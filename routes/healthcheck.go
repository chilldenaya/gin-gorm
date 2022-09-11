package routers

import (
	"net/http"

	"gin-gorm/controllers"
	res "gin-gorm/helpers"

	"github.com/gin-gonic/gin"
)

func (r routes) RegisterHealthcheckRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/healthchecks")

	route.GET("/app", GetAppHeartbeat)
	route.GET("/db", GetDBHeartbeat)
}

func GetAppHeartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, res.SetOk("App is up!"))
}

func GetDBHeartbeat(c *gin.Context) {
	err := controllers.GetDBHeartbeatController()
	if err != nil {
		c.JSON(http.StatusNotFound, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res.SetOk("DB is up!"))
}
