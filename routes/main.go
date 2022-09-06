package routers

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	Router *gin.Engine
}

func InitRoutes() routes {
	r := routes{
		Router: gin.Default(),
	}

	v1 := r.Router.Group("/api/v1")
	r.RegisterStudentRoutes(v1)
	r.RegisterClassRoutes(v1)
	r.RegisterHealthcheckRoutes(v1)

	return r

}

func (r routes) Run(addr ...string) error {
	return r.Router.Run()
}
