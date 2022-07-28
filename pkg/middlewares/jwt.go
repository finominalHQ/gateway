package middlewares

import (
	"github.com/gobuffalo/buffalo"
	tokenauth "github.com/gobuffalo/mw-tokenauth"
	"github.com/golang-jwt/jwt"
)

func VerifyJWT() buffalo.MiddlewareFunc {
	return tokenauth.New(tokenauth.Options{
		SignMethod: jwt.SigningMethodHS512,
	})
}
