package v1

import (
	backpanel "github.com/alisilver78/go-uplod-download-service/api/v1/back-panel"
	frontpanel "github.com/alisilver78/go-uplod-download-service/api/v1/front-panel"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	frontpanel.Routes(e)
	backpanel.Routes(e)
}
