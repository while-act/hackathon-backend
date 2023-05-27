// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/while-act/hackathon-backend/ent/predicate"
	"github.com/while-act/hackathon-backend/ent/taxationsystem"
)

// TaxationSystemDelete is the builder for deleting a TaxationSystem entity.
type TaxationSystemDelete struct {
	config
	hooks    []Hook
	mutation *TaxationSystemMutation
}

// Where appends a list predicates to the TaxationSystemDelete builder.
func (tsd *TaxationSystemDelete) Where(ps ...predicate.TaxationSystem) *TaxationSystemDelete {
	tsd.mutation.Where(ps...)
	return tsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tsd *TaxationSystemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tsd.sqlExec, tsd.mutation, tsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tsd *TaxationSystemDelete) ExecX(ctx context.Context) int {
	n, err := tsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tsd *TaxationSystemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(taxationsystem.Table, sqlgraph.NewFieldSpec(taxationsystem.FieldID, field.TypeInt))
	if ps := tsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tsd.mutation.done = true
	return affected, err
}

// TaxationSystemDeleteOne is the builder for deleting a single TaxationSystem entity.
type TaxationSystemDeleteOne struct {
	tsd *TaxationSystemDelete
}

// Where appends a list predicates to the TaxationSystemDelete builder.
func (tsdo *TaxationSystemDeleteOne) Where(ps ...predicate.TaxationSystem) *TaxationSystemDeleteOne {
	tsdo.tsd.mutation.Where(ps...)
	return tsdo
}

// Exec executes the deletion query.
func (tsdo *TaxationSystemDeleteOne) Exec(ctx context.Context) error {
	n, err := tsdo.tsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taxationsystem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tsdo *TaxationSystemDeleteOne) ExecX(ctx context.Context) {
	if err := tsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
