package api

import (
	v1 "github.com/alisilver78/go-uplod-download-service/api/v1"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	v1.Routes(e)
}
