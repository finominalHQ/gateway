package middlewares

import (
	"github.com/gobuffalo/buffalo"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
)

func Logger() buffalo.MiddlewareFunc {
	return paramlogger.ParameterLogger
}
