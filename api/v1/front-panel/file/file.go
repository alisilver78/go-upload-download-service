package file

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func fileBaseRoute(c echo.Context) error {
	html := `<p style="font-weight:bold">File Base Route</p>`

	return c.HTML(http.StatusOK, html)
}