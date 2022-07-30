package middlewares

import (
	"context"
	"gateway/pkg/route"
	"gateway/pkg/util"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/gobuffalo/buffalo"
)

// Proxy request to the right backend
func Proxy(c context.Context) buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		incoming := c.Value("incomingRoute").(*route.Route)
		outgoing := c.Value("outgoingRoute").(*route.Route)

		// do not proxy request if no matching outgoing route is found
		if incoming == nil || outgoing == nil {
			return func(c buffalo.Context) error {
				return next(c)
			}
		}

		rp := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				// mutate req header
				for k, v := range incoming.Header {
					req.Header.Set(k, v.(string))
				}

				// mutate req query
				existingQueries := req.URL.Query()
				for k, v := range incoming.Query {
					existingQueries.Add(k, v.(string))
				}
				req.URL.RawQuery = existingQueries.Encode()

				// mutate req body
				b, _ := ioutil.ReadAll(req.Body)
				defer req.Body.Close()
				existingBody := make(util.Json)
				util.JsonParse(b, &existingBody)
				for k, v := range incoming.Body {
					existingBody[k] = v
				}

				// mutate req destination
				req.URL.Host = outgoing.Host
				// req.URL.Port = func() string { return r.Port.String }
				req.URL.Path = outgoing.Service.String + outgoing.Resource.String + outgoing.Action.String
			},
		}
		return buffalo.WrapHandler(rp)
	}
}
