package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{ 
		field.Int("class_id"),
		field.Int("author_id"),
		field.String("content").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}).NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),                          
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("posts").Field("class_id").Unique().Required(),
		edge.From("author", User.Type).Ref("owned_posts").Field("author_id").Unique().Required(),
	}
}
