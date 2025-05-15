package routes

import (
	"HTMX/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Static("/static", "./static")
	router.GET("/", handler.H1)
}
