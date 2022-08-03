package controllers

import (
	"gateway/pkg/util"

	"github.com/gobuffalo/buffalo"
)

func Error404() buffalo.ErrorHandler {
	return func(s int, e error, c buffalo.Context) error {
		return c.Render(s, util.R.JSON(map[string]any{
			"title":   "gateway.app.404.error.title",
			"message": "gateway.app.404.error.message",
			"data":    e.Error(),
			"meta":    nil,
		}))
	}
}
