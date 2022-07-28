package middlewares

import (
	limiter "github.com/alcalbg/buffalo-rate-limiter-mw"
	"github.com/gobuffalo/buffalo"
)

func Throttler() buffalo.MiddlewareFunc {
	return limiter.Limiter(5, []string{"CF-Connecting-I", "RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
}
