package handler

import (
	"HTMX/components"
	"HTMX/models"
	"context"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

var count = 0

func H1(c *gin.Context) {

	nh := models.NewNowHandler(time.Now)

	// components.Button_eg("click me").Render(c.Request.Context(), c.Writer)

	// view := components.CountryNames(models.Countries)
	// view.Render(c.Request.Context(), c.Writer)

	Mouse_entered(c)

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

// func CounterExample(c *gin.Context) {
// 	components.Layout(count).Render(c.Request.Context(), c.Writer)
// }

func StreamComponent(c *gin.Context) {
	countries := models.Countries
	templ.Handler(components.StreamPage(countries), templ.WithStreaming()).ServeHTTP(c.Writer, c.Request)
}

func Mouse_entered(c *gin.Context) {
	components.MouseEntered(models.Countries).Render(c.Request.Context(), c.Writer)
}

func ShowMousePage(c *gin.Context) {
	components.MouseHere().Render(c.Request.Context(), c.Writer)
}

func IncrementCount(c *gin.Context) {
	count++
	components.Counter(count).Render(c.Request.Context(), c.Writer)
}

func RenderPage(c *gin.Context) {
	c.Status(200)
	c.Header("Content-Type", "text/html")
	components.Layout(components.Counter(count)).Render(c.Request.Context(), c.Writer)
}
