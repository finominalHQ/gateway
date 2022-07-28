package middlewares

import (
	"gateway/pkg/util"

	"github.com/gobuffalo/buffalo"
)

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func Translations() buffalo.MiddlewareFunc {
	return util.Translator.Middleware()
}
