package middlewares

import (
	"fmt"
	"gateway/pkg/cache"
	"gateway/pkg/route"
	"gateway/pkg/util"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func GetIncomingRoute() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			tx, ok := c.Value("tx").(*pop.Connection)
			if !ok {
				return util.Error(c, "gateway.route.show.error.title", "gateway.route.show.error.message", 0, nil)
			}

			req := c.Request()

			key := fmt.Sprintf("gateway:route:%s:%s-%s-%s-%s", route.OUTGOING, req.Method, req.URL.Host, req.URL.Port(), strings.ReplaceAll(req.URL.Path, "/", "-"))
			incomingRoute := cache.Get(key).(*route.Route)
			if incomingRoute == nil {
				incomingRoute = route.GetIncomingRoute(req, tx)
			}

			c.Set("incomingRoute", incomingRoute)

			return next(c)
		}
	}
}
