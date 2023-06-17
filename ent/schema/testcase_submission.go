package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Testcase_Submission holds the schema definition for the Testcase_Submission entity.
type Testcase_Submission struct {
	ent.Schema
}

// Fields of the Testcase_Submission.
func (Testcase_Submission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("submission_id"),
		field.Int("testcase_id"),
		field.String("input").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.String("expected_output").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.String("actual_output").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.Bool("is_passed"),
		field.Float("score"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Testcase_Submission.
func (Testcase_Submission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("testcase", Testcase.Type).Ref("testcase_submissions").Field("testcase_id").Unique().Required(),
		edge.From("submission", Lab_Problem_Submission.Type).Ref("testcase_submissions").Field("submission_id").Unique().Required(),
	}
}
