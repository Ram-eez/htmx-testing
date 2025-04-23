package main

import (
	"HTMX/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//router.GET("")

	routes.RegisterRoutes(router)
	router.Run(":8080")
}
