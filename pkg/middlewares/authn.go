package middlewares

import (
	"gateway/pkg/identity"
	"gateway/pkg/util"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func AuthN() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {

			status := identity.AuthN(c.Request())
			if status == false {
				return util.Error(c, "gateway.identity.authn.error.title", "gateway.identity.authn.error.message", http.StatusUnauthorized, nil)
			}

			return next(c)
		}
	}
}
