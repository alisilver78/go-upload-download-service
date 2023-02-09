package api

import (
	v1 "github.com/alisilver78/go-uplod-download-service/api/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {
	ap1Group := e.Group("/api", middleware.RemoveTrailingSlash())
	v1.Routes(ap1Group)
}
