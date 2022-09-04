package routers

import (
	"net/http"

	res "gin-gorm/internals"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func InitRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/api/v1")
	r.RegisterStudentRoutes(v1)
	r.RegisterClassRoutes(v1)

	v1.GET("/ping", ping)

	return r

}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, res.SetOk("pong"))
}
