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
	route.DELETE("/:id", DeleteClassByIdRoute)
}

func GetClassesRoute(c *gin.Context) {
	data, err := controllers.GetClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res.SetOk(data))
}

func GetClassByIdRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))
	}

	data, err := controllers.GetClassById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res.SetOk(data))
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

	c.JSON(http.StatusCreated, res.SetOk(data))
}

func DeleteClassByIdRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))
	}

	err = controllers.DeleteClassById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res.SetOk(gin.H{}))
}
