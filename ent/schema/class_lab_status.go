package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class_Lab_Status holds the schema definition for the Class_Lab_Status entity.
type Class_Lab_Status struct {
	ent.Schema
}

// Fields of the Class_Lab_Status.
func (Class_Lab_Status) Fields() []ent.Field {
	return []ent.Field{
		field.Int("course_id"),
		field.Int("class_id"),
		field.Int("lab_id"),
		field.Bool("is_open").Default(false),
	}
}

// Edges of the Class_Lab_Status.
func (Class_Lab_Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("course_lab_statuses").Field("course_id").Unique().Required(),
		edge.From("class", Class.Type).Ref("class_lab_statuses").Field("class_id").Unique().Required(),
		edge.From("lab", Lab.Type).Ref("lab_statuses").Field("lab_id").Unique().Required(),
	}
}
