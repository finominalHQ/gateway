package middlewares

import (
	"gateway/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
)

// Wraps each request in a transaction.
//   c.Value("tx").(*pop.Connection)
// Remove to disable this.
func WrapWithTransaction() buffalo.MiddlewareFunc {
	return popmw.Transaction(models.DB)
}
