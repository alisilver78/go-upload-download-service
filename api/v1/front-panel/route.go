package frontpanel

import (
	"github.com/alisilver78/go-uplod-download-service/api/v1/front-panel/file"
	"github.com/labstack/echo/v4"
)

func Routes(v1Group *echo.Group) {
	fpG := v1Group.Group("/fp")
	file.Routes(fpG)

}
