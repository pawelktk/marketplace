package front

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	host := c.Request().Host
	URI := c.Request().RequestURI
	qs := c.QueryString()

	c.Logger().Error(err, fmt.Sprintf(" on: %s%s%s error code: %d", host, URI, qs, code))
	c.String(code, fmt.Sprintf("error code: %d", code))
}

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	e.HTTPErrorHandler = customHTTPErrorHandler

	//e.Static("/", "css")

	e.Use(middleware.Static("static"))

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", MainHandler)
	e.GET("/cancel_add/", CancelAdd)
	e.GET("/cancel_edit/:id", CancelEdit)

	e.POST("/products/", HandleProductAdd)

	e.GET("/products/add", AddProduct)
	e.GET("/products/edit/:id", EditProduct)

	e.PUT("/products/:id", HandleProductUpdate)

}
