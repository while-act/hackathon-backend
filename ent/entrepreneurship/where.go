// Code generated by ent, DO NOT EDIT.

package entrepreneurship

import (
	"entgo.io/ent/dialect/sql"
	"github.com/while-act/hackathon-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldEQ(FieldType, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(sql.FieldContainsFold(FieldType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Entrepreneurship) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Entrepreneurship) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(func(s *sql.Selector) {
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
func Not(p predicate.Entrepreneurship) predicate.Entrepreneurship {
	return predicate.Entrepreneurship(func(s *sql.Selector) {
		p(s.Not())
	})
}
