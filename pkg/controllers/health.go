package controllers

import (
	"fmt"
	"gateway/pkg/route"
	"gateway/pkg/util"

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
		url := fmt.Sprintf("%s:%s/%s/%s", r.Host, r.Port.String, r.Resource.String, r.Action.String)
		_, err := util.Request(r.Method, url, nil)

		status := "Available"
		if err != nil {
			status = err.Error()
		}
		data[url] = status
	}

	return util.Success(c, "gateway.local.health.title", "gateway.local.health.message", data, nil)
}
