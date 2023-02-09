package server

import (
	"github.com/alisilver78/go-uplod-download-service/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
