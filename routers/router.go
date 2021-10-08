package routers

import (
	"github.com/labstack/echo/v4"
)

func InitRouters() *echo.Echo {
	e := echo.New()
	// routers
	//e.Static("/static", "static/test.html")
	//e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	//	Level: 5,
	//}))
	//e.Static("/obj", "obj")
	api := e.Group("/api")
	{
		api.POST("/dw", Dw)
	}
	return e
}
