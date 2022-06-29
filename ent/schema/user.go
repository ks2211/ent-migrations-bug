package schema

import (
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
		field.Int("id"),
		field.Int("age"),
		field.String("name"),
		field.Int("tenant_id"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("users").
			Field("tenant_id").
			Required().
			Unique(),
		edge.To("user_perms", UserPerms.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
		// edge.To("group_users", GroupUser.Type).
		// 	StorageKey(edge.Column("user_id")).
		// 	Annotations(
		// 		entsql.Annotation{
		// 			OnDelete: entsql.Cascade,
		// 		},
		// 	),
	}
}
