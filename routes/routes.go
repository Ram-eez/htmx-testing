package routes

import (
	"HTMX/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Static("/static", "./static")
	// router.GET("/", handler.H1)

	router.POST("/increment", handler.IncrementCount)
	router.GET("/", handler.RenderPage)

	router.GET("/stream", handler.StreamComponent)

	// router.POST("/mouse", handler.Mouse_entered)
	// router.GET("/mouseexample", handler.ShowMousePage)
}
