// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/workspace"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BillingID holds the value of the "billing_id" field.
	BillingID string `json:"billing_id,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// APIKey holds the value of the "api_key" field.
	APIKey string `json:"-"`
	// Confirmed holds the value of the "confirmed" field.
	Confirmed bool `json:"confirmed,omitempty"`
	// ConfirmationSentAt holds the value of the "confirmation_sent_at" field.
	ConfirmationSentAt *time.Time `json:"confirmation_sent_at,omitempty"`
	// ConfirmationToken holds the value of the "confirmation_token" field.
	ConfirmationToken *string `json:"confirmation_token,omitempty"`
	// RecoverySentAt holds the value of the "recovery_sent_at" field.
	RecoverySentAt *time.Time `json:"recovery_sent_at,omitempty"`
	// RecoveryToken holds the value of the "recovery_token" field.
	RecoveryToken *string `json:"recovery_token,omitempty"`
	// OtpSentAt holds the value of the "otp_sent_at" field.
	OtpSentAt *time.Time `json:"otp_sent_at,omitempty"`
	// Otp holds the value of the "otp" field.
	Otp *string `json:"otp,omitempty"`
	// EmailChange holds the value of the "email_change" field.
	EmailChange string `json:"email_change,omitempty"`
	// EmailChangeSentAt holds the value of the "email_change_sent_at" field.
	EmailChangeSentAt *time.Time `json:"email_change_sent_at,omitempty"`
	// EmailChangeToken holds the value of the "email_change_token" field.
	EmailChangeToken *string `json:"email_change_token,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Roles holds the value of the "roles" field.
	Roles []string `json:"roles,omitempty"`
	// Teams holds the value of the "teams" field.
	Teams map[string]string `json:"teams,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LastSigninAt holds the value of the "last_signin_at" field.
	LastSigninAt *time.Time `json:"last_signin_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Workspace holds the value of the workspace edge.
	Workspace *Workspace
	// WorkspaceRoles holds the value of the workspace_roles edge.
	WorkspaceRoles []*WorkspaceRole
	// GroupRoles holds the value of the group_roles edge.
	GroupRoles []*GroupRole
	// UserRoles holds the value of the user_roles edge.
	UserRoles []*UserRole
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WorkspaceOrErr returns the Workspace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) WorkspaceOrErr() (*Workspace, error) {
	if e.loadedTypes[0] {
		if e.Workspace == nil {
			// The edge workspace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workspace.Label}
		}
		return e.Workspace, nil
	}
	return nil, &NotLoadedError{edge: "workspace"}
}

// WorkspaceRolesOrErr returns the WorkspaceRoles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WorkspaceRolesOrErr() ([]*WorkspaceRole, error) {
	if e.loadedTypes[1] {
		return e.WorkspaceRoles, nil
	}
	return nil, &NotLoadedError{edge: "workspace_roles"}
}

// GroupRolesOrErr returns the GroupRoles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupRolesOrErr() ([]*GroupRole, error) {
	if e.loadedTypes[2] {
		return e.GroupRoles, nil
	}
	return nil, &NotLoadedError{edge: "group_roles"}
}

// UserRolesOrErr returns the UserRoles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[3] {
		return e.UserRoles, nil
	}
	return nil, &NotLoadedError{edge: "user_roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // billing_id
		&sql.NullString{}, // provider
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&sql.NullString{}, // api_key
		&sql.NullBool{},   // confirmed
		&sql.NullTime{},   // confirmation_sent_at
		&sql.NullString{}, // confirmation_token
		&sql.NullTime{},   // recovery_sent_at
		&sql.NullString{}, // recovery_token
		&sql.NullTime{},   // otp_sent_at
		&sql.NullString{}, // otp
		&sql.NullString{}, // email_change
		&sql.NullTime{},   // email_change_sent_at
		&sql.NullString{}, // email_change_token
		&[]byte{},         // metadata
		&[]byte{},         // roles
		&[]byte{},         // teams
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // last_signin_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		u.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field billing_id", values[0])
	} else if value.Valid {
		u.BillingID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field provider", values[1])
	} else if value.Valid {
		u.Provider = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[2])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[3])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field api_key", values[4])
	} else if value.Valid {
		u.APIKey = value.String
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field confirmed", values[5])
	} else if value.Valid {
		u.Confirmed = value.Bool
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_sent_at", values[6])
	} else if value.Valid {
		u.ConfirmationSentAt = new(time.Time)
		*u.ConfirmationSentAt = value.Time
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_token", values[7])
	} else if value.Valid {
		u.ConfirmationToken = new(string)
		*u.ConfirmationToken = value.String
	}
	if value, ok := values[8].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_sent_at", values[8])
	} else if value.Valid {
		u.RecoverySentAt = new(time.Time)
		*u.RecoverySentAt = value.Time
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_token", values[9])
	} else if value.Valid {
		u.RecoveryToken = new(string)
		*u.RecoveryToken = value.String
	}
	if value, ok := values[10].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field otp_sent_at", values[10])
	} else if value.Valid {
		u.OtpSentAt = new(time.Time)
		*u.OtpSentAt = value.Time
	}
	if value, ok := values[11].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field otp", values[11])
	} else if value.Valid {
		u.Otp = new(string)
		*u.Otp = value.String
	}
	if value, ok := values[12].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change", values[12])
	} else if value.Valid {
		u.EmailChange = value.String
	}
	if value, ok := values[13].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_sent_at", values[13])
	} else if value.Valid {
		u.EmailChangeSentAt = new(time.Time)
		*u.EmailChangeSentAt = value.Time
	}
	if value, ok := values[14].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_token", values[14])
	} else if value.Valid {
		u.EmailChangeToken = new(string)
		*u.EmailChangeToken = value.String
	}

	if value, ok := values[15].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[15])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}

	if value, ok := values[16].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field roles", values[16])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Roles); err != nil {
			return fmt.Errorf("unmarshal field roles: %v", err)
		}
	}

	if value, ok := values[17].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field teams", values[17])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Teams); err != nil {
			return fmt.Errorf("unmarshal field teams: %v", err)
		}
	}
	if value, ok := values[18].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[18])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[19].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[19])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	if value, ok := values[20].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field last_signin_at", values[20])
	} else if value.Valid {
		u.LastSigninAt = new(time.Time)
		*u.LastSigninAt = value.Time
	}
	return nil
}

// QueryWorkspace queries the workspace edge of the User.
func (u *User) QueryWorkspace() *WorkspaceQuery {
	return (&UserClient{config: u.config}).QueryWorkspace(u)
}

// QueryWorkspaceRoles queries the workspace_roles edge of the User.
func (u *User) QueryWorkspaceRoles() *WorkspaceRoleQuery {
	return (&UserClient{config: u.config}).QueryWorkspaceRoles(u)
}

// QueryGroupRoles queries the group_roles edge of the User.
func (u *User) QueryGroupRoles() *GroupRoleQuery {
	return (&UserClient{config: u.config}).QueryGroupRoles(u)
}

// QueryUserRoles queries the user_roles edge of the User.
func (u *User) QueryUserRoles() *UserRoleQuery {
	return (&UserClient{config: u.config}).QueryUserRoles(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("models: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", billing_id=")
	builder.WriteString(u.BillingID)
	builder.WriteString(", provider=")
	builder.WriteString(u.Provider)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", api_key=<sensitive>")
	builder.WriteString(", confirmed=")
	builder.WriteString(fmt.Sprintf("%v", u.Confirmed))
	if v := u.ConfirmationSentAt; v != nil {
		builder.WriteString(", confirmation_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := u.ConfirmationToken; v != nil {
		builder.WriteString(", confirmation_token=")
		builder.WriteString(*v)
	}
	if v := u.RecoverySentAt; v != nil {
		builder.WriteString(", recovery_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := u.RecoveryToken; v != nil {
		builder.WriteString(", recovery_token=")
		builder.WriteString(*v)
	}
	if v := u.OtpSentAt; v != nil {
		builder.WriteString(", otp_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := u.Otp; v != nil {
		builder.WriteString(", otp=")
		builder.WriteString(*v)
	}
	builder.WriteString(", email_change=")
	builder.WriteString(u.EmailChange)
	if v := u.EmailChangeSentAt; v != nil {
		builder.WriteString(", email_change_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := u.EmailChangeToken; v != nil {
		builder.WriteString(", email_change_token=")
		builder.WriteString(*v)
	}
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", u.Metadata))
	builder.WriteString(", roles=")
	builder.WriteString(fmt.Sprintf("%v", u.Roles))
	builder.WriteString(", teams=")
	builder.WriteString(fmt.Sprintf("%v", u.Teams))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	if v := u.LastSigninAt; v != nil {
		builder.WriteString(", last_signin_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
