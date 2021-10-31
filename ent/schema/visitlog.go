package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// VisitLog holds the schema definition for the VisitLog entity.
type VisitLog struct {
	ent.Schema
}

// Fields of the VisitLog.
func (VisitLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("platform"),
		field.String("os"),
		field.String("engine_name"),
		field.String("engine_version"),
		field.String("browser_name"),
		field.String("browser_version"),
		field.String("mozilla"),
		field.Bool("bot"),
		field.Bool("mobile"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the VisitLog.
func (VisitLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("url", Url.Type).
			Ref("logs").
			Unique(),
	}
}

// Annotations of the Url.
func (VisitLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "log"},
	}
}
