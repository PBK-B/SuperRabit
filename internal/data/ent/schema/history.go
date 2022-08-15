package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.String("device").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}),
		field.String("ip"),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("version", Version.Type).Ref("histories"),
	}
}
