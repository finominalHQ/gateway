package util

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/golang-jwt/jwt"
)

func JWTNew(id string, expiry int) string {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expiry)).Unix()
	// add more claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(envy.Get("JWT_SECRET", "")))

	return tokenString
}

func JWTId(c buffalo.Context) string {
	claims := c.Value("claims").(jwt.MapClaims)
	return claims["id"].(string)
}
