// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/users/internal/models/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
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
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LastSigninAt holds the value of the "last_signin_at" field.
	LastSigninAt *time.Time `json:"last_signin_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // provider
		&sql.NullString{}, // email
		&sql.NullString{}, // password
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
		return fmt.Errorf("unexpected type %T for field provider", values[0])
	} else if value.Valid {
		u.Provider = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field confirmed", values[3])
	} else if value.Valid {
		u.Confirmed = value.Bool
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_sent_at", values[4])
	} else if value.Valid {
		u.ConfirmationSentAt = new(time.Time)
		*u.ConfirmationSentAt = value.Time
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_token", values[5])
	} else if value.Valid {
		u.ConfirmationToken = new(string)
		*u.ConfirmationToken = value.String
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_sent_at", values[6])
	} else if value.Valid {
		u.RecoverySentAt = new(time.Time)
		*u.RecoverySentAt = value.Time
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_token", values[7])
	} else if value.Valid {
		u.RecoveryToken = new(string)
		*u.RecoveryToken = value.String
	}
	if value, ok := values[8].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field otp_sent_at", values[8])
	} else if value.Valid {
		u.OtpSentAt = new(time.Time)
		*u.OtpSentAt = value.Time
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field otp", values[9])
	} else if value.Valid {
		u.Otp = new(string)
		*u.Otp = value.String
	}
	if value, ok := values[10].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change", values[10])
	} else if value.Valid {
		u.EmailChange = value.String
	}
	if value, ok := values[11].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_sent_at", values[11])
	} else if value.Valid {
		u.EmailChangeSentAt = new(time.Time)
		*u.EmailChangeSentAt = value.Time
	}
	if value, ok := values[12].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_token", values[12])
	} else if value.Valid {
		u.EmailChangeToken = new(string)
		*u.EmailChangeToken = value.String
	}

	if value, ok := values[13].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[13])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	if value, ok := values[14].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[14])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[15].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[15])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	if value, ok := values[16].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field last_signin_at", values[16])
	} else if value.Valid {
		u.LastSigninAt = new(time.Time)
		*u.LastSigninAt = value.Time
	}
	return nil
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
	builder.WriteString(", provider=")
	builder.WriteString(u.Provider)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=<sensitive>")
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
