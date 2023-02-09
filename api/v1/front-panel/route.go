package frontpanel

import (
	"github.com/alisilver78/go-uplod-download-service/api/v1/front-panel/file"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	file.Routes(e)

}
