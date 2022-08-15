package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Version holds the schema definition for the Version entity.
type Version struct {
	ent.Schema
}

// Fields of the Version.
func (Version) Fields() []ent.Field {
	return []ent.Field{
		field.String("version"),
		field.Int("build"),
		field.Float("apk_size").Default(0).Optional(),
		field.Float("ipa_size").Default(0).Optional(),
		field.String("ipa_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}).Default("").Optional(),
		field.String("apk_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}).Default("").Optional(),
		field.String("plist_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}).Default("").Optional(),
		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)",
		}).Default("").Optional(),
		field.String("access").Default("public").Optional(),
		field.String("access_code").Default("").Optional(),
		field.Int("status").Default(3).Optional().Comment("0 全部禁止访问； 1 允许访问iOS下载包； 2 允许访问安卓下载包；3 允许访问苹果、安卓下载包"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Version.
func (Version) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("versions"),
		edge.From("app", App.Type).Ref("versions"),
		edge.To("histories", History.Type),
	}
}
