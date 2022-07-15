package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("path").MaxLen(20).NotEmpty(),
		field.Int("current_times").Default(0),
		field.Int("max_times").Default(0),
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

// Indexes of the Url.
func (Url) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("path"),
		index.Fields("url").Annotations(entsql.Prefix(200)),
	}
}
