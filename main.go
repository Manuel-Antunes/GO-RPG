package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

func dashboard(c echo.Context) error {
	return c.String(http.StatusOK, "Dashboard")
}

func main() {
	e := echo.New()
	e.GET("/", dashboard)
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.Logger.Print("Listening on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}

// this type is an existing template to serve html/templates into echo
// https://echo.labstack.com/guide/templates/
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
