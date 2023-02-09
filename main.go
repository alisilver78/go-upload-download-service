package main

import (
	"github.com/alisilver78/go-uplod-download-service/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	server.Run(e)
}