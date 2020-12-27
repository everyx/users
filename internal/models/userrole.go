// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// UserRole is the model entity for the UserRole schema.
type UserRole struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserRoleQuery when eager-loading is set.
	Edges           UserRoleEdges `json:"edges"`
	user_user_roles *uuid.UUID
}

// UserRoleEdges holds the relations/edges for other nodes in the graph.
type UserRoleEdges struct {
	// Users holds the value of the users edge.
	Users *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserRoleEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// The edge users was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserRole) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // name
		&uuid.UUID{},      // user_id
		&[]byte{},         // metadata
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*UserRole) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // user_user_roles
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserRole fields.
func (ur *UserRole) assignValues(values ...interface{}) error {
	if m, n := len(values), len(userrole.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		ur.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		ur.Name = value.String
	}
	if value, ok := values[1].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field user_id", values[1])
	} else if value != nil {
		ur.UserID = *value
	}

	if value, ok := values[2].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[2])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &ur.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		ur.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[4])
	} else if value.Valid {
		ur.UpdatedAt = value.Time
	}
	values = values[5:]
	if len(values) == len(userrole.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field user_user_roles", values[0])
		} else if value != nil {
			ur.user_user_roles = value
		}
	}
	return nil
}

// QueryUsers queries the users edge of the UserRole.
func (ur *UserRole) QueryUsers() *UserQuery {
	return (&UserRoleClient{config: ur.config}).QueryUsers(ur)
}

// Update returns a builder for updating this UserRole.
// Note that, you need to call UserRole.Unwrap() before calling this method, if this UserRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (ur *UserRole) Update() *UserRoleUpdateOne {
	return (&UserRoleClient{config: ur.config}).UpdateOne(ur)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ur *UserRole) Unwrap() *UserRole {
	tx, ok := ur.config.driver.(*txDriver)
	if !ok {
		panic("models: UserRole is not a transactional entity")
	}
	ur.config.driver = tx.drv
	return ur
}

// String implements the fmt.Stringer.
func (ur *UserRole) String() string {
	var builder strings.Builder
	builder.WriteString("UserRole(")
	builder.WriteString(fmt.Sprintf("id=%v", ur.ID))
	builder.WriteString(", name=")
	builder.WriteString(ur.Name)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", ur.UserID))
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", ur.Metadata))
	builder.WriteString(", created_at=")
	builder.WriteString(ur.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ur.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserRoles is a parsable slice of UserRole.
type UserRoles []*UserRole

func (ur UserRoles) config(cfg config) {
	for _i := range ur {
		ur[_i].config = cfg
	}
}
