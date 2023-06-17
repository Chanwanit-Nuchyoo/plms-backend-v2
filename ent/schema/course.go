package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
	owner *User
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id"),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now()).Optional(),
		field.Time("updated_at").Default(time.Now()).Optional(),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("courses").Unique().Field("owner_id").Required(),
		edge.To("classes", Class.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("labs", Lab.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("topics",Topic.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("course_lab_statuses", Class_Lab_Status.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
