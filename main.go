package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var tpl = &tplWrapper{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = tpl
	e.GET("/", indexHandler)
	e.Static("/static", "./public")
	e.Logger.Fatal(e.Start(":8081"))
}

func indexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

type tplWrapper struct {
	templates *template.Template
}

func (t *tplWrapper) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
