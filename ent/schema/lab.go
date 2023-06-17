package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lab holds the schema definition for the Lab entity.
type Lab struct {
	ent.Schema
}

// Fields of the Lab.
func (Lab) Fields() []ent.Field {
	return []ent.Field{
		field.Int("course_id"),
		field.String("name").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Lab.
func (Lab) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("labs").Unique().Field("course_id").Required(),
		edge.To("lab_problems", Lab_Problem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("lab_statuses", Class_Lab_Status.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
