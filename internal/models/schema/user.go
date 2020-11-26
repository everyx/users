package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty().Sensitive().MinLen(8),

		field.Bool("confirmed").Default(false).Optional(),
		field.Time("confirmation_sent_at").Nillable().Optional(),
		field.String("confirmation_token").NotEmpty().Optional().Nillable().Unique(),

		field.Time("recovery_sent_at").Nillable().Optional(),
		field.String("recovery_token").NotEmpty().Optional().Nillable().Unique(),

		field.String("email_change").NotEmpty().Optional(),
		field.Time("email_change_sent_at").Nillable().Optional(),
		field.String("email_change_token").NotEmpty().Optional().Nillable().Unique(),

		field.JSON("metadata", map[string]interface{}{}),

		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("last_signin_at").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
