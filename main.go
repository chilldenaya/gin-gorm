package main

import (
	"fmt"
	m "gin-gorm/models"
	r "gin-gorm/routes"
	"os"
)

func main() {
	m.InitDBConnection()

	r.InitRoutes().Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
