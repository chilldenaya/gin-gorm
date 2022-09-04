package routers

import (
	"net/http"

	s "gin-gorm/services"

	"github.com/gin-gonic/gin"
)

func (r routes) RegisterClassRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/classes")

	route.GET("/", GetClassesRoute)
	route.GET("/:id", GetClassByIdRoute)
}

func GetClassesRoute(c *gin.Context) {
	c.JSON(http.StatusOK, s.GetClassesService())
}

func GetClassByIdRoute(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, s.GetClassByIdService(id))
}
