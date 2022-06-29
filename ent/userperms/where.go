// Code generated by ent, DO NOT EDIT.

package userperms

import (
	"entgo.io/bug/ent/internal"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Admin applies equality check predicate on the "admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// License1 applies equality check predicate on the "license_1" field. It's identical to License1EQ.
func License1(v bool) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLicense1), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// AdminEQ applies the EQ predicate on the "admin" field.
func AdminEQ(v bool) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// AdminNEQ applies the NEQ predicate on the "admin" field.
func AdminNEQ(v bool) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdmin), v))
	})
}

// License1EQ applies the EQ predicate on the "license_1" field.
func License1EQ(v bool) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLicense1), v))
	})
}

// License1NEQ applies the NEQ predicate on the "license_1" field.
func License1NEQ(v bool) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLicense1), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserPerms {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserPerms {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserPerms
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserPerms
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPerms) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPerms) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPerms) predicate.UserPerms {
	return predicate.UserPerms(func(s *sql.Selector) {
		p(s.Not())
	})
}