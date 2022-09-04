package main

import (
	r "gin-gorm/routers"
)

func main() {
	router := r.InitRoutes()
	router.Run(":8080")
}
