package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func H1(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world")
}
