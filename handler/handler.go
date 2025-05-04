package handler

import (
	"HTMX/components"
	"HTMX/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func H1(c *gin.Context) {

	nh := models.NewNowHandler(time.Now)

	// components.Button_eg("click me").Render(c.Request.Context(), c.Writer)

	// view := components.CountryNames(models.Countries)
	// view.Render(c.Request.Context(), c.Writer)

	components.Page(nh.Now(), models.Countries).Render(c.Request.Context(), c.Writer)
	c.String(http.StatusOK, "Hello, world")
}
