package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

type pessoa struct {
	Nome  string
	Idade int
}

func dashboard(c echo.Context) error {
	jorge := pessoa{
		Nome:  "jorge",
		Idade: 10,
	}
	return c.Render(http.StatusOK, "index", jorge)
}

func main() {
	e := echo.New()
	e.GET("/", dashboard)
	//this object is used to store the Template engine object reference
	t := &Template{
		templates: template.Must(template.ParseGlob("src/views/*.html")),
	}
	//here we setup the static files system
	e.Static("/public", "src/public")
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
