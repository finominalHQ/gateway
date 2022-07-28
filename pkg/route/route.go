package route

import (
	"encoding/json"
	"gateway/pkg/util"
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
	Method    string       `json:"method" db:"method"`
	Host      string       `json:"host" db:"host"`
	Port      nulls.String `json:"port" db:"port"`
	Service   nulls.String `json:"service" db:"service"`
	Resource  nulls.String `json:"resource" db:"resource"`
	Action    nulls.String `json:"action" db:"action"`
	Query     util.Json    `json:"query" db:"query"`
	Body      util.Json    `json:"body" db:"body"`
	Header    util.Json    `json:"header" db:"header"`
	Config    util.Json    `json:"config" db:"config"`
	Auth      AuthType     `json:"auth" db:"auth"`
	Type      TypeType     `json:"type" db:"type"`
	Status    StatusType   `json:"status" db:"status"`
	Ref       string       `json:"backend" db:"backend"`
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
