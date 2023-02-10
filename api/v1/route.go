package v1

import (
	backpanel "github.com/alisilver78/go-upload-download-service/api/v1/back-panel"
	frontpanel "github.com/alisilver78/go-upload-download-service/api/v1/front-panel"
	"github.com/labstack/echo/v4"
)

func Routes(apiGroup *echo.Group) {
	v1Group := apiGroup.Group("/v1")
	frontpanel.Routes(v1Group)
	backpanel.Routes(v1Group)
}
