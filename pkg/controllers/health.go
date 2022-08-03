package controllers

import (
	"fmt"
	"gateway/pkg/route"
	"gateway/pkg/util"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Health(c buffalo.Context) error {
	var data map[string]string

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	routes := &route.Routes{}
	if err := tx.All(routes); err != nil {
		return err
	}

	for _, r := range *routes {
		_, err := util.Request(http.MethodGet, r.Upstream, nil)

		status := "Available"
		if err != nil {
			status = err.Error()
		}
		data[r.Upstream] = status
	}

	return util.Success(c, "gateway.local.health.title", "gateway.local.health.message", data, nil)
}
