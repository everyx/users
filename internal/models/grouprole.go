// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/users/internal/models/group"
	"github.com/adnaan/users/internal/models/grouprole"
	"github.com/adnaan/users/internal/models/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// GroupRole is the model entity for the GroupRole schema.
type GroupRole struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID uuid.UUID `json:"group_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupRoleQuery when eager-loading is set.
	Edges             GroupRoleEdges `json:"edges"`
	group_group_roles *uuid.UUID
	user_group_roles  *uuid.UUID
}

// GroupRoleEdges holds the relations/edges for other nodes in the graph.
type GroupRoleEdges struct {
	// Groups holds the value of the groups edge.
	Groups *Group
	// Users holds the value of the users edge.
	Users *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupRoleEdges) GroupsOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Groups == nil {
			// The edge groups was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupRoleEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[1] {
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
func (*GroupRole) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // name
		&uuid.UUID{},      // group_id
		&uuid.UUID{},      // user_id
		&[]byte{},         // metadata
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*GroupRole) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // group_group_roles
		&uuid.UUID{}, // user_group_roles
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupRole fields.
func (gr *GroupRole) assignValues(values ...interface{}) error {
	if m, n := len(values), len(grouprole.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		gr.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		gr.Name = value.String
	}
	if value, ok := values[1].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field group_id", values[1])
	} else if value != nil {
		gr.GroupID = *value
	}
	if value, ok := values[2].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field user_id", values[2])
	} else if value != nil {
		gr.UserID = *value
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &gr.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[4])
	} else if value.Valid {
		gr.CreatedAt = value.Time
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[5])
	} else if value.Valid {
		gr.UpdatedAt = value.Time
	}
	values = values[6:]
	if len(values) == len(grouprole.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field group_group_roles", values[0])
		} else if value != nil {
			gr.group_group_roles = value
		}
		if value, ok := values[1].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field user_group_roles", values[1])
		} else if value != nil {
			gr.user_group_roles = value
		}
	}
	return nil
}

// QueryGroups queries the groups edge of the GroupRole.
func (gr *GroupRole) QueryGroups() *GroupQuery {
	return (&GroupRoleClient{config: gr.config}).QueryGroups(gr)
}

// QueryUsers queries the users edge of the GroupRole.
func (gr *GroupRole) QueryUsers() *UserQuery {
	return (&GroupRoleClient{config: gr.config}).QueryUsers(gr)
}

// Update returns a builder for updating this GroupRole.
// Note that, you need to call GroupRole.Unwrap() before calling this method, if this GroupRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *GroupRole) Update() *GroupRoleUpdateOne {
	return (&GroupRoleClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *GroupRole) Unwrap() *GroupRole {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("models: GroupRole is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *GroupRole) String() string {
	var builder strings.Builder
	builder.WriteString("GroupRole(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", group_id=")
	builder.WriteString(fmt.Sprintf("%v", gr.GroupID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", gr.UserID))
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", gr.Metadata))
	builder.WriteString(", created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GroupRoles is a parsable slice of GroupRole.
type GroupRoles []*GroupRole

func (gr GroupRoles) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
