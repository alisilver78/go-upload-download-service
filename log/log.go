package log

import (
	"github.com/alisilver78/go-upload-download-service/configs"
	"github.com/labstack/gommon/log"
)

func Errorf(massage string, args ...string) {
	if configs.CFG.ErrorLog {
		argumants := ""
		for _, arg := range args {
			argumants += arg
		}
		log.Errorf(massage, argumants)
	}
}
