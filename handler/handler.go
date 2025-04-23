package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func h1(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world")
}
