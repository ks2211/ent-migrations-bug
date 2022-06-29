package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Groups holds the schema definition for the Groups entity.
type Groups struct {
	ent.Schema
}

// Fields of the Groups.
func (Groups) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("description"),
		field.Int("tenant_id"),
	}
}

// Edges of the Groups.
func (Groups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("groups").
			Field("tenant_id").
			Required().
			Unique(),
		// edge.To("group_users", GroupUser.Type).
		// 	StorageKey(edge.Column("user_id")).
		// 	Annotations(
		// 		entsql.Annotation{
		// 			OnDelete: entsql.Cascade,
		// 		},
		// 	),
	}
}
