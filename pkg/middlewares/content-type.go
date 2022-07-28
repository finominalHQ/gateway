package middlewares

import (
	"github.com/gobuffalo/buffalo"
	contenttype "github.com/gobuffalo/mw-contenttype"
)

func ForceContentType() buffalo.MiddlewareFunc {
	return contenttype.Set("application/json")
}
