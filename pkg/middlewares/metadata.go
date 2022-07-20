package middlewares

import (
	"github.com/gobuffalo/buffalo"
)

func AttachMetadata() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {

			println("hello")
			c.Set("routex", "hellp")

			return next(c)
			// tx, ok := c.Value("tx").(*pop.Connection)
			// if !ok {
			// 	return fmt.Errorf("no transaction found")
			// }

			// // Allocate an empty Route
			// route := &struct{}{}

			// // To find the Route the parameter route_id is used.
			// method := c.Request().Method
			// host := c.Request().URL.Host
			// port := c.Request().URL.Port()
			// paths := strings.Split(c.Request().URL.Path, "")

			// query := "method = ? and host = ? and port = ?"

			// service := "and service is null"
			// if len(paths) > 1 {
			// 	service = "and service = " + paths[0]
			// }
			// query = query + service

			// resource := "and resource is null"
			// if len(paths) > 2 {
			// 	resource = "and resource = " + paths[1]
			// }
			// query = query + resource

			// action := "and action is null"
			// if len(paths) > 3 {
			// 	action = "and action = " + paths[2]
			// }
			// query = query + action

			// if err := tx.Where(
			// 	query,
			// 	method,
			// 	host,
			// 	port,
			// 	service,
			// 	resource,
			// 	action,
			// ).First(route); err != nil {
			// 	return c.Error(http.StatusNotFound, err)
			// }

			// println("hello")
			// c = context.WithValue(c, "routex", "hello-val").(buffalo.Context)
			// c = context.WithValue(c, "route", route).(buffalo.Context)

			// return next(c)
		}
	}
}
