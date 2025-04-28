package handler

import (
	"HTMX/components"
	"HTMX/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func H1(c *gin.Context) {
	view := components.CountryNames(models.Countries)
	view.Render(c.Request.Context(), c.Writer)
	c.String(http.StatusOK, "Hello, world")
}
