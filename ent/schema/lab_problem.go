package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lab_Problem holds the schema definition for the Lab_Problem entity.
type Lab_Problem struct {
	ent.Schema
}

// Fields of the Lab_Problem.
func (Lab_Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("lab_id"),
		field.String("name").NotEmpty(),
		field.String("prompt").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}).NotEmpty(),
		field.Float("full_score").Positive().Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Lab_Problem.
func (Lab_Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lab", Lab.Type).Ref("lab_problems").Field("lab_id").Unique().Required(),
		edge.To("testcases", Testcase.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
