package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lab_Problem_Submission holds the schema definition for the Lab_Problem_Submission entity.
type Lab_Problem_Submission struct {
	ent.Schema
}

// Fields of the Lab_Problem_Submission.
func (Lab_Problem_Submission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id"),
		field.Int("lab_problem_id"),
		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.Float("score").Positive(),
		field.String("feedback").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}).Optional(),
		field.Float("instructor_score").Positive(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Lab_Problem_Submission.
func (Lab_Problem_Submission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("lab_problem_submissions").Field("owner_id").Unique().Required(),
		edge.To("testcase_submissions", Testcase_Submission.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
