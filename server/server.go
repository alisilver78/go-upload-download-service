package server

import (
	"fmt"

	"github.com/alisilver78/go-upload-download-service/api"
	"github.com/alisilver78/go-upload-download-service/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(e *echo.Echo) {
	e.Use(middleware.Recover())

	//load configs
	configs.LoadConfig()

	if configs.CFG.ErrorAccessLog {
		e.Use(middleware.Logger())
	}

	//load apis
	api.Routes(e)

	//start server
	fmt.Printf("server started on http://%s:%s", configs.CFG.Domain, configs.CFG.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", configs.CFG.Domain, configs.CFG.Port)))
}
