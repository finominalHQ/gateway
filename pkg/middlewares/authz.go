package middlewares

import (
	"gateway/pkg/identity"
	"gateway/pkg/util"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func AuthZ() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			status := identity.AuthZ(c.Request())
			if status == false {
				return util.Error(c, "gateway.identity.authz.error.title", "gateway.idenity.authz.error.message", http.StatusUnauthorized, nil)
			}

			return next(c)
		}
	}
}
