package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouters() *echo.Echo {
	e := echo.New()
	// routers
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Static("/obj", "./obj")
	e.Static("/dist", "./static/dist")
	e.Static("/index", "./static/dist/index.html")
	e.Static("/", "./static/dist")
	//e.Static("/obj", "obj")
	api := e.Group("/api")
	{
		api.POST("/dw", Dw)
		api.POST("/uwb", Uwb)
		api.POST("/modelData", ModelData)
		api.POST("/nameList", NameList)
	}
	return e
}
