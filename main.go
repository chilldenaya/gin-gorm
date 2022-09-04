package main

import (
	r "gin-gorm/routes"
)

func main() {
	router := r.InitRoutes()
	router.Run(":8080")
}
