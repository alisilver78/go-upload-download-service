package file

import "github.com/labstack/echo/v4"

func Routes(fpG *echo.Group) {
	fileGp := fpG.Group("/file")
	fileGp.GET("", fileBaseRoute)
	fileGp.GET("/download/:file", fileDownload)
	fileGp.GET("/download/attachment/:file", fileAttachment)
	fileGp.GET("/download/inline/:file", fileInline)
}
