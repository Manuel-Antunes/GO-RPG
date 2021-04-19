package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

func create_user(c echo.Context) error {
	form, error := c.MultipartForm()
	if error == nil {
		fmt.Println(form)
	}
	return c.Render(http.StatusOK, "dashboard", form)
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
func sign_up(c echo.Context) error {
	return c.Render(http.StatusOK, "sign-up", nil)
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.POST("/users", create_user)
	e.GET("/sign-up", sign_up)
	//here we setup the static files system
	e.Static("/public", "src/public")
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("src/views/*.html")),
	}
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
