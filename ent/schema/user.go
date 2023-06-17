package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty().MinLen(8).MaxLen(16),
		field.String("email").Unique(),
		field.String("password").NotEmpty().MinLen(8).MaxLen(16),
		field.String("name").NotEmpty(),
		field.String("image_path").Optional(),
		field.Time("last_online").Optional(),
		field.Time("created_at").Default(time.Now()).Optional(),
		field.Time("updated_at").Default(time.Now()).Optional(),
		field.Bool("is_admin").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("courses", Course.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("owned_classes", Class.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("joined_classes", Class_User_Role.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("owned_posts", Post.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("lab_problem_submissions",Lab_Problem_Submission.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
