package actions

import (
	"gateway/pkg/middlewares"
	"gateway/pkg/util"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

var (
	app *buffalo.App
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          util.ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_gateway_session",
		})

		// Automatically redirect to SSL
		app.Use(middlewares.ForceSSL())

		// Automatically redirect to SSL
		app.Use(middlewares.Translations())

		// Log request parameters (filters apply).
		app.Use(middlewares.Logger())

		// Force content type to JSON
		app.Use(middlewares.ForceContentType())

		// Wraps each request in a transaction.
		app.Use(middlewares.WrapWithTransaction())

		// check for jwt token
		app.Use(middlewares.VerifyJWT())

		// global request throttle
		app.Use(middlewares.Throttler())

		// Attach incoming route to route
		app.Use(middlewares.AttachMetadata())

		// fall back response handler
		app.Use(middlewares.Fallback())
	}

	return app
}
