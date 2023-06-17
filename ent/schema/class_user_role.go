package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class_User_Role holds the schema definition for the Class_User_Role entity.
type Class_User_Role struct {
	ent.Schema
}

// Fields of the Class_User_Role.
func (Class_User_Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("class_id"),
		field.Enum("role").Values("student","ta","instructor"),
	}
}

// Edges of the Class_User_Role.
func (Class_User_Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("joined_classes").Field("user_id").Unique().Required(),
		edge.From("class", Class.Type).Ref("members").Field("class_id").Unique().Required(),
	}
}
