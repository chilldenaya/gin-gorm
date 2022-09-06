package main

import (
	m "gin-gorm/models"
	r "gin-gorm/routes"
)

func main() {
	m.InitDBConnection()

	r.InitRoutes().Run(":8080")
}
