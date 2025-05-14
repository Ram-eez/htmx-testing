package handler

import (
	"HTMX/components"
	"HTMX/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func H1(c *gin.Context) {

	nh := models.NewNowHandler(time.Now)

	// components.Button_eg("click me").Render(c.Request.Context(), c.Writer)

	// view := components.CountryNames(models.Countries)
	// view.Render(c.Request.Context(), c.Writer)

	countries, err := models.GetCountries("country")

	vm := models.NewCountryListViewModel(countries, err)

	components.CountryListComponent(vm).Render(c.Request.Context(), c.Writer)

	components.Message(models.Countries[1].Name, models.Countries[1].States, models.Countries[1].Population).Render(c.Request.Context(), c.Writer)

	ctx := context.WithValue(context.Background(), models.ThemeContextKey, "dark")

	components.ThemeName().Render(ctx, c.Writer)
	prop_dril := components.Top("rameez")

	prop_dril.Render(c.Request.Context(), c.Writer)

	components.Page(nh.Now(), models.Countries).Render(c.Request.Context(), c.Writer)
	c.String(http.StatusOK, "Hello, world")
}
