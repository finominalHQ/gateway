package middlewares

import (
	"gateway/pkg/util"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func Fallback() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			// println("pre-routex")
			// println(c.Value("routex").(string))
			// println("post-routex")
			return util.Error(c, "Fallback Response", "Unable to find a matching record", http.StatusNotFound, nil)
		}
	}
}
