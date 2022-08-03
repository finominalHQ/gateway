package route

import (
	"encoding/json"
	"gateway/models"
	"gateway/pkg/util"
	"net/http"
	"strings"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Route is used by pop to map your routes database table to your go code.
type Route struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	Name      nulls.String `json:"name" db:"name"`
	Desc      nulls.String `json:"desc" db:"desc"`
	Pattern   string       `json:"pattern" db:"pattern"`
	Upstream  string       `json:"upstream" db:"upstream"`
	Query     util.Json    `json:"query" db:"query"`
	Body      util.Json    `json:"body" db:"body"`
	Header    util.Json    `json:"header" db:"header"`
	Config    util.Json    `json:"config" db:"config"`
	Auth      AuthType     `json:"auth" db:"auth"`
	Type      TypeType     `json:"type" db:"type"`
	Status    StatusType   `json:"status" db:"status"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (r Route) String() string {
	ju, _ := json.Marshal(r)
	return string(ju)
}

// Routes is not required by pop and may be deleted
type Routes []Route

// String is not required by pop and may be deleted
func (r Routes) String() string {
	ju, _ := json.Marshal(r)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Route) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Route) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Route) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func GetRoutesByService(service string, routes *Routes) error {
	if err := models.DB.Where(
		"name = ?",
		service,
	).All(routes); err != nil {
		return err
	}

	return nil
}

func GetRequestRoute(req *http.Request) *Route {
	method := req.Method
	host := req.URL.Host
	port := req.URL.Port()
	paths := strings.Split(req.URL.Path, "")

	query := "status = '" + ACTIVE + "' and method = ? and host = ? and port = ?"

	service := "and service is null"
	if len(paths) > 0 {
		service = "and (service = " + paths[0] + " or service is null)"
	}
	query = query + service

	resource := "and resource is null"
	if len(paths) > 1 {
		resource = "and resource = " + paths[1]
	}
	query = query + resource

	action := "and action is null"
	if len(paths) > 2 {
		action = "and action = " + paths[2]
	}
	query = query + action

	route := &Route{}
	if err := models.DB.Where(
		query,
		method,
		host,
		port,
		service,
		resource,
		action,
	).First(route); err != nil {
		return nil
	}

	return route
}
