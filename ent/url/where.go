// Code generated by entc, DO NOT EDIT.

package url

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/long2ice/longurl/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// CurrentTimes applies equality check predicate on the "current_times" field. It's identical to CurrentTimesEQ.
func CurrentTimes(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrentTimes), v))
	})
}

// MaxTimes applies equality check predicate on the "max_times" field. It's identical to MaxTimesEQ.
func MaxTimes(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxTimes), v))
	})
}

// ExpireAt applies equality check predicate on the "expire_at" field. It's identical to ExpireAtEQ.
func ExpireAt(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpireAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// CurrentTimesEQ applies the EQ predicate on the "current_times" field.
func CurrentTimesEQ(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrentTimes), v))
	})
}

// CurrentTimesNEQ applies the NEQ predicate on the "current_times" field.
func CurrentTimesNEQ(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCurrentTimes), v))
	})
}

// CurrentTimesIn applies the In predicate on the "current_times" field.
func CurrentTimesIn(vs ...int) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCurrentTimes), v...))
	})
}

// CurrentTimesNotIn applies the NotIn predicate on the "current_times" field.
func CurrentTimesNotIn(vs ...int) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCurrentTimes), v...))
	})
}

// CurrentTimesGT applies the GT predicate on the "current_times" field.
func CurrentTimesGT(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCurrentTimes), v))
	})
}

// CurrentTimesGTE applies the GTE predicate on the "current_times" field.
func CurrentTimesGTE(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCurrentTimes), v))
	})
}

// CurrentTimesLT applies the LT predicate on the "current_times" field.
func CurrentTimesLT(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCurrentTimes), v))
	})
}

// CurrentTimesLTE applies the LTE predicate on the "current_times" field.
func CurrentTimesLTE(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCurrentTimes), v))
	})
}

// MaxTimesEQ applies the EQ predicate on the "max_times" field.
func MaxTimesEQ(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxTimes), v))
	})
}

// MaxTimesNEQ applies the NEQ predicate on the "max_times" field.
func MaxTimesNEQ(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxTimes), v))
	})
}

// MaxTimesIn applies the In predicate on the "max_times" field.
func MaxTimesIn(vs ...int) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMaxTimes), v...))
	})
}

// MaxTimesNotIn applies the NotIn predicate on the "max_times" field.
func MaxTimesNotIn(vs ...int) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMaxTimes), v...))
	})
}

// MaxTimesGT applies the GT predicate on the "max_times" field.
func MaxTimesGT(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxTimes), v))
	})
}

// MaxTimesGTE applies the GTE predicate on the "max_times" field.
func MaxTimesGTE(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxTimes), v))
	})
}

// MaxTimesLT applies the LT predicate on the "max_times" field.
func MaxTimesLT(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxTimes), v))
	})
}

// MaxTimesLTE applies the LTE predicate on the "max_times" field.
func MaxTimesLTE(v int) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxTimes), v))
	})
}

// ExpireAtEQ applies the EQ predicate on the "expire_at" field.
func ExpireAtEQ(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpireAt), v))
	})
}

// ExpireAtNEQ applies the NEQ predicate on the "expire_at" field.
func ExpireAtNEQ(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpireAt), v))
	})
}

// ExpireAtIn applies the In predicate on the "expire_at" field.
func ExpireAtIn(vs ...time.Time) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpireAt), v...))
	})
}

// ExpireAtNotIn applies the NotIn predicate on the "expire_at" field.
func ExpireAtNotIn(vs ...time.Time) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpireAt), v...))
	})
}

// ExpireAtGT applies the GT predicate on the "expire_at" field.
func ExpireAtGT(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpireAt), v))
	})
}

// ExpireAtGTE applies the GTE predicate on the "expire_at" field.
func ExpireAtGTE(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpireAt), v))
	})
}

// ExpireAtLT applies the LT predicate on the "expire_at" field.
func ExpireAtLT(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpireAt), v))
	})
}

// ExpireAtLTE applies the LTE predicate on the "expire_at" field.
func ExpireAtLTE(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpireAt), v))
	})
}

// ExpireAtIsNil applies the IsNil predicate on the "expire_at" field.
func ExpireAtIsNil() predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExpireAt)))
	})
}

// ExpireAtNotNil applies the NotNil predicate on the "expire_at" field.
func ExpireAtNotNil() predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExpireAt)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Url {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasLogs applies the HasEdge predicate on the "logs" edge.
func HasLogs() predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLogsWith applies the HasEdge predicate on the "logs" edge with a given conditions (other predicates).
func HasLogsWith(preds ...predicate.VisitLog) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
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
func Not(p predicate.Url) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		p(s.Not())
	})
}
