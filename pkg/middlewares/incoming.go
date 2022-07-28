package middlewares

import (
	"gateway/pkg/route"
	"gateway/pkg/util"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func AttachMetadata() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			tx, ok := c.Value("tx").(*pop.Connection)
			if !ok {
				return util.Error(c, "Auth", util.T(c, "auth.transaction.load.failed"), 0, nil)
			}

			// Allocate an empty Route
			r := &route.Route{}

			// To find the Route the parameter route_id is used.
			method := c.Request().Method
			host := c.Request().URL.Host
			port := c.Request().URL.Port()
			paths := strings.Split(c.Request().URL.Path, "")

			query := "type = '" + route.INCOMING + "' and status = 'incoming' and method = ? and host = ? and port = ?"

			service := "and service is null"
			if len(paths) > 0 {
				service = "and service = " + paths[0]
			}
			query = query + service

			resource := "and resource is null"
			if len(paths) > 1 {
				resource = "and resource = " + paths[1]
			}
			query = query + resource

			action := "and action is null"
			if len(paths) > 2 {
				action = "and action = " + paths[2]
			}
			query = query + action

			if err := tx.Where(
				query,
				method,
				host,
				port,
				service,
				resource,
				action,
			).First(r); err != nil {
				return c.Error(http.StatusNotFound, err)
			}

			c.Set("incomingRoute", r)

			return next(c)
		}
	}
}
