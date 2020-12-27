package schema

import (
	"time"

	"github.com/facebook/ent/schema/edge"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "usersx"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("billing_id").NotEmpty().Unique().Optional(),
		field.String("provider").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty().Sensitive().MinLen(8),
		field.String("api_key").NotEmpty().Sensitive().Optional(),
		field.Bool("confirmed").Default(false).Optional(),
		field.Time("confirmation_sent_at").Nillable().Optional(),
		field.String("confirmation_token").NotEmpty().Optional().Nillable().Unique(),

		field.Time("recovery_sent_at").Nillable().Optional(),
		field.String("recovery_token").NotEmpty().Optional().Nillable().Unique(),
		field.Time("otp_sent_at").Nillable().Optional(),
		field.String("otp").NotEmpty().Optional().Nillable().Unique(),

		field.String("email_change").NotEmpty().Optional(),
		field.Time("email_change_sent_at").Nillable().Optional(),
		field.String("email_change_token").NotEmpty().Optional().Nillable().Unique(),

		field.JSON("metadata", map[string]interface{}{}),
		field.Strings("roles").Optional(),
		field.JSON("teams", map[string]string{}).Optional(),

		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("last_signin_at").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("workspace", Workspace.Type).Unique(),
		edge.To("workspace_roles", WorkspaceRole.Type),
		edge.To("group_roles", GroupRole.Type),
		edge.To("user_roles", UserRole.Type),
	}
}
