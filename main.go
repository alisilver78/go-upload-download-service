package main

import (
	"github.com/alisilver78/go-upload-download-service/server"
	"github.com/labstack/echo/v4"
)

func main() {
	// create echo object
	e := echo.New()

	// run server
	server.Run(e)
}
