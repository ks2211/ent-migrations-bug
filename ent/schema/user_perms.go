package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserPerms holds the schema definition for the UserPerms entity.
type UserPerms struct {
	ent.Schema
}

// Fields of the UserPerms.
func (UserPerms) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Bool("admin"),
		field.Bool("license_1"),
		field.Int("user_id"),
	}
}

// Edges of the UserPerms.
func (UserPerms) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_perms").
			Field("user_id").
			Required().
			Unique(),
	}
}
