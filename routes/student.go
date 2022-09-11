package routers

import (
	"net/http"
	"strconv"

	"gin-gorm/controllers"
	dto "gin-gorm/dto"
	res "gin-gorm/helpers"

	"github.com/gin-gonic/gin"
)

func (r routes) RegisterStudentRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/students")

	route.GET("/", GetStudentsRoute)
	route.GET("/:id", GetStudentByIdRoute)
	route.POST("/", CreateStudentRoute)
	route.DELETE("/:id", DeleteStudentByIdRoute)
}

func GetStudentsRoute(c *gin.Context) {
	data, err := controllers.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res.SetOk(data))
}

func GetStudentByIdRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))
	}

	data, err := controllers.GetStudentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res.SetOk(data))
}

func CreateStudentRoute(c *gin.Context) {
	var req dto.CreateStudentRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))
	}

	data, err := controllers.CreateStudent(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusCreated, res.SetOk(data))
}

func DeleteStudentByIdRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.SetErr(err.Error()))
	}

	err = controllers.DeleteStudentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, res.SetErr(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res.SetOk(gin.H{}))
}
