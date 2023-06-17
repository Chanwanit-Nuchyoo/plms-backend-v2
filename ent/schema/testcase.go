package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Testcase holds the schema definition for the Testcase entity.
type Testcase struct {
	ent.Schema
}

// Fields of the Testcase.
func (Testcase) Fields() []ent.Field {
	return []ent.Field{
		field.Int("lab_problem_id"),
		field.String("input").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.String("output").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Testcase.
func (Testcase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lab_problem", Lab_Problem.Type).Ref("testcases").Field("lab_problem_id").Unique().Required(),
		edge.To("testcase_submissions", Testcase_Submission.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
