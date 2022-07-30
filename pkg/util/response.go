package util

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

var R *render.Engine

func init() {
	R = render.New(render.Options{
		DefaultContentType: "application/json",
	})
}

func Response(c buffalo.Context, status int, title, message string, data, meta any) error {
	return c.Render(status, R.JSON(map[string]any{
		"title":   T(c, title),
		"message": T(c, message),
		"data":    data,
		"meta":    meta,
	}))
}

func Success(c buffalo.Context, title, message string, data, meta any) error {
	return Response(c, http.StatusOK, title, message, data, meta)
}

func Error(c buffalo.Context, title, message string, status int, data any) error {
	if status == 0 {
		status = http.StatusBadRequest
	}
	return Response(c, status, title, message, data, nil)
}
