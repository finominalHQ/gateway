package middlewares

import (
	"fmt"
	"gateway/pkg/cache"
	"gateway/pkg/route"
	"gateway/pkg/util"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func GetOutgoingRoute() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			tx, ok := c.Value("tx").(*pop.Connection)
			if !ok {
				return util.Error(c, "gateway.route.show.error.title", "gateway.route.show.error.message", 0, nil)
			}

			incoming, ok := c.Value("incomingRoute").(*route.Route)
			if !ok {
				return util.Error(c, "Auth", util.T(c, " "), 0, nil)
			}

			req := c.Request()

			key := fmt.Sprintf("gateway:route:%s:%s-%s-%s-%s", route.OUTGOING, req.Method, req.URL.Host, req.URL.Port(), req.URL.Path)
			outgoingRoute := cache.Get(key).(*route.Route)
			if outgoingRoute == nil {
				outgoingRoute = route.GetOutgoingRoute(incoming, tx)
			}

			c.Set("outgoingRoute", outgoingRoute)

			return next(c)
		}
	}
}
