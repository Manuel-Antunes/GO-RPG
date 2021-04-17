package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func dashboard(c echo.Context) error {
	return c.String(http.StatusOK, "Dashboard")
}

func main() {
	e := echo.New()
	e.GET("/", dashboard)
	e.Logger.Print("Listening on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
