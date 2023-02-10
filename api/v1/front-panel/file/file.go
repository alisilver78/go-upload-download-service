package file

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func fileBaseRoute(c echo.Context) error {
	html := `
		<p style="font-weight:bold">File Base Route</p>
	`

	return c.HTML(http.StatusOK, html)
}

func fileDownload(c echo.Context) error {

	file := c.Param("file")

	return c.File(file)
}

func fileAttachment(c echo.Context) error {

	file := c.Param("file")

	return c.Attachment(file, file)
}

func fileInline(c echo.Context) error {

	file := c.Param("file")

	return c.Inline(file, file)
}
