package controllers

import (
	"gateway/pkg/util"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func Fallback(c buffalo.Context) error {
	return util.Error(c, "gateway.local.fallback.error.title", "gateway.local.fallback.error.message", http.StatusNotFound, nil)
}
