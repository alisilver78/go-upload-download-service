package api

import (
	"net/http"

	v1 "github.com/alisilver78/go-upload-download-service/api/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {
	//route page
	e.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h3>route page</h3>`)
	})

	ap1Group := e.Group("/api", middleware.RemoveTrailingSlash())
	v1.Routes(ap1Group)
}
