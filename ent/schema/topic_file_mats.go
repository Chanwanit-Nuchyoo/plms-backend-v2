package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Topic_File_Mats holds the schema definition for the Topic_File_Mats entity.
type Topic_File_Mats struct {
	ent.Schema
}

// Fields of the Topic_File_Mats.
func (Topic_File_Mats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("topic_id"),
		field.String("name").NotEmpty(),
		field.String("file_path"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Topic_File_Mats.
func (Topic_File_Mats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("topic", Topic.Type).Ref("files").Field("topic_id").Unique().Required(),
	}
}
