package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("course_id"),
		field.String("name").NotEmpty(),
		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("topics").Field("course_id").Unique().Required(),
		edge.To("files", Topic_File_Mats.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
