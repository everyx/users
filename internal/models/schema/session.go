package schema

import (
	"time"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sessionsx"},
	}
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Text("data"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("expires_at").Nillable().Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
