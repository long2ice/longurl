package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Url holds the schema definition for the Url entity.
type Url struct {
	ent.Schema
}

// Fields of the Url.
func (Url) Fields() []ent.Field {
	return []ent.Field{
		field.Text("url").NotEmpty(),
		field.String("path").Unique().NotEmpty(),
		field.Time("expire_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Url.
func (Url) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("logs", VisitLog.Type),
	}
}

// Annotations of the Url.
func (Url) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "url"},
	}
}
