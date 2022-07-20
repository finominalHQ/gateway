package route

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"gateway/actions"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Route)
// DB Table: Plural (routes)
// Resource: Plural (Routes)
// Path: Plural (/routes)
// View Template Folder: Plural (/templates/routes/)

// RoutesResource is the resource for the Route model
type RoutesResource struct {
	buffalo.Resource
}

// List gets all Routes. This function is mapped to the path
// GET /routes
func (v RoutesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	routes := &Routes{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Routes from the DB
	if err := q.All(routes); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, actions.R.JSON(routes))
	}).Respond(c)
}

// Show gets the data for one Route. This function is mapped to
// the path GET /routes/{route_id}
func (v RoutesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Route
	route := &Route{}

	// To find the Route the parameter route_id is used.
	if err := tx.Find(route, c.Param("route_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, actions.R.JSON(route))
	}).Respond(c)
}

// Create adds a Route to the DB. This function is mapped to the
// path POST /routes
func (v RoutesResource) Create(c buffalo.Context) error {
	// Allocate an empty Route
	route := &Route{}

	// Bind route to the html form elements
	if err := c.Bind(route); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(route)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, actions.R.JSON(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, actions.R.JSON(route))
	}).Respond(c)
}

// Update changes a Route in the DB. This function is mapped to
// the path PUT /routes/{route_id}
func (v RoutesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Route
	route := &Route{}

	if err := tx.Find(route, c.Param("route_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Route to the html form elements
	if err := c.Bind(route); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(route)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, actions.R.JSON(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, actions.R.JSON(route))
	}).Respond(c)
}

// Destroy deletes a Route from the DB. This function is mapped
// to the path DELETE /routes/{route_id}
func (v RoutesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Route
	route := &Route{}

	// To find the Route the parameter route_id is used.
	if err := tx.Find(route, c.Param("route_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(route); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, actions.R.JSON(route))
	}).Respond(c)
}