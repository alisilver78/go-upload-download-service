package file

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {
	fileGroupe := e.Group("/file")
	fileGroupe.GET("/", fileBaseRoute)
}
