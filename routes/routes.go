package routes

import (
	"HTMX/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", handler.H1)
}
