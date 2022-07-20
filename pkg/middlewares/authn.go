package middlewares

import (
	"fmt"
	"gateway/pkg/util"

	"github.com/gobuffalo/buffalo"
)

func AuthN() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			payload := map[string]string{
				"name":  "Toby",
				"email": "Toby@example.com",
			}

			res, err := util.Post("", payload)
			if err != nil {
				return fmt.Errorf("unable to authN")
			}

			if res["data"] != true {
				return fmt.Errorf("unauthorised")
			}

			return next(c)
		}
	}
}
