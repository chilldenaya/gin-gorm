package routers

import (
	"net/http"
	"strconv"

	"gin-gorm/controllers"
	dto "gin-gorm/dto"
	res "gin-gorm/helpers"

	"github.com/gin-gonic/gin"
)

func (r routes) RegisterClassRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/classes")

	route.GET("/", GetClassesRoute)
	route.GET("/:id", GetClassByIdRoute)

	route.POST("/", CreateClassRoute)
}

func GetClassesRoute(c *gin.Context) {
	data, err := controllers.GetClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, data)

		return
	}

	c.JSON(http.StatusOK, data)
}

func GetClassByIdRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))
	}

	data, err := controllers.GetClassById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, data)
}

func CreateClassRoute(c *gin.Context) {
	var req dto.CreateClassRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))
	}

	data, err := controllers.CreateClass(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, data)
}
