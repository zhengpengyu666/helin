package routers

import (
	"github.com/labstack/echo/v4"
	"helin/routers/test"
)

func InitRouters() *echo.Echo {
	e := echo.New()
	// routers
	e.Static("/static", "static/test.html")
	api := e.Group("/api")
	{
		api.POST("/", test.Hello)
	}
	return e
}
