package routers

import (
	"net/http"

	s "gin-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func (r routes) RegisterStudentRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/students")

	route.GET("/", GetStudentsRoute)
	route.GET("/:id", GetStudentByIdRoute)
}

func GetStudentsRoute(c *gin.Context) {
	c.JSON(http.StatusOK, s.GetStudentsController())
}

func GetStudentByIdRoute(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, s.GetStudentByIdController(id))
}
