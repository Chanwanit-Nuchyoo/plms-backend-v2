package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id"),
		field.Int("course_id"),
		field.String("name").NotEmpty(),
		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.Enum("default_lang").Values("c_cpp", "python"),
		field.Enum("dotw").Values("sunday","monday","tuesday","wednesday", "thursday", "friday"),
		field.Time("start_time").SchemaType(map[string]string{
			dialect.MySQL: "time",
		}),
		field.Time("end_time").SchemaType(map[string]string{
			dialect.MySQL: "time",
		}),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("owned_classes").Field("owner_id").Unique().Required(),
		edge.From("course", Course.Type).Ref("classes").Field("course_id").Unique().Required(),
		edge.To("members", Class_User_Role.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("posts", Post.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("class_lab_statuses", Class_Lab_Status.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
