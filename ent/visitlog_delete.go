// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"long2ice/longurl/ent/predicate"
	"long2ice/longurl/ent/visitlog"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VisitLogDelete is the builder for deleting a VisitLog entity.
type VisitLogDelete struct {
	config
	hooks    []Hook
	mutation *VisitLogMutation
}

// Where appends a list predicates to the VisitLogDelete builder.
func (vld *VisitLogDelete) Where(ps ...predicate.VisitLog) *VisitLogDelete {
	vld.mutation.Where(ps...)
	return vld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vld *VisitLogDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vld.hooks) == 0 {
		affected, err = vld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VisitLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vld.mutation = mutation
			affected, err = vld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vld.hooks) - 1; i >= 0; i-- {
			if vld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vld *VisitLogDelete) ExecX(ctx context.Context) int {
	n, err := vld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vld *VisitLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: visitlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: visitlog.FieldID,
			},
		},
	}
	if ps := vld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vld.driver, _spec)
}

// VisitLogDeleteOne is the builder for deleting a single VisitLog entity.
type VisitLogDeleteOne struct {
	vld *VisitLogDelete
}

// Exec executes the deletion query.
func (vldo *VisitLogDeleteOne) Exec(ctx context.Context) error {
	n, err := vldo.vld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{visitlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vldo *VisitLogDeleteOne) ExecX(ctx context.Context) {
	vldo.vld.ExecX(ctx)
}
