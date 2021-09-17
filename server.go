package main

import (
	"github.com/labstack/echo/v4"
	"helin/router"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/aa", func(context echo.Context) error {
		hello := router.Hello()
		return context.String(http.StatusOK, hello)
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/static", "static/test.html")
	e.Logger.Fatal(e.Start(":8080"))
}
