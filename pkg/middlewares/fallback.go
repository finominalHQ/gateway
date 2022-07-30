package middlewares

import (
	"gateway/pkg/util"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func Fallback() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			return util.Error(c, "gateway.local.fallback.error.title", "gateway.local.fallback.error.message", http.StatusNotFound, nil)
		}
	}
}
