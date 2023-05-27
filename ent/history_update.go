// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/while-act/hackathon-backend/ent/history"
	"github.com/while-act/hackathon-backend/ent/predicate"
)

// HistoryUpdate is the builder for updating History entities.
type HistoryUpdate struct {
	config
	hooks    []Hook
	mutation *HistoryMutation
}

// Where appends a list predicates to the HistoryUpdate builder.
func (hu *HistoryUpdate) Where(ps ...predicate.History) *HistoryUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetName sets the "name" field.
func (hu *HistoryUpdate) SetName(s string) *HistoryUpdate {
	hu.mutation.SetName(s)
	return hu
}

// Mutation returns the HistoryMutation object of the builder.
func (hu *HistoryUpdate) Mutation() *HistoryMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HistoryUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HistoryUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HistoryUpdate) check() error {
	if v, ok := hu.mutation.Name(); ok {
		if err := history.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "History.name": %w`, err)}
		}
	}
	if _, ok := hu.mutation.IndustryID(); hu.mutation.IndustryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "History.industry"`)
	}
	if _, ok := hu.mutation.DistrictID(); hu.mutation.DistrictCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "History.district"`)
	}
	if _, ok := hu.mutation.UsersID(); hu.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "History.users"`)
	}
	return nil
}

func (hu *HistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Name(); ok {
		_spec.SetField(history.FieldName, field.TypeString, value)
	}
	if hu.mutation.OperationsNumCleared() {
		_spec.ClearField(history.FieldOperationsNum, field.TypeInt)
	}
	if hu.mutation.OtherCleared() {
		_spec.ClearField(history.FieldOther, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HistoryUpdateOne is the builder for updating a single History entity.
type HistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HistoryMutation
}

// SetName sets the "name" field.
func (huo *HistoryUpdateOne) SetName(s string) *HistoryUpdateOne {
	huo.mutation.SetName(s)
	return huo
}

// Mutation returns the HistoryMutation object of the builder.
func (huo *HistoryUpdateOne) Mutation() *HistoryMutation {
	return huo.mutation
}

// Where appends a list predicates to the HistoryUpdate builder.
func (huo *HistoryUpdateOne) Where(ps ...predicate.History) *HistoryUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HistoryUpdateOne) Select(field string, fields ...string) *HistoryUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated History entity.
func (huo *HistoryUpdateOne) Save(ctx context.Context) (*History, error) {
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HistoryUpdateOne) SaveX(ctx context.Context) *History {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HistoryUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HistoryUpdateOne) check() error {
	if v, ok := huo.mutation.Name(); ok {
		if err := history.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "History.name": %w`, err)}
		}
	}
	if _, ok := huo.mutation.IndustryID(); huo.mutation.IndustryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "History.industry"`)
	}
	if _, ok := huo.mutation.DistrictID(); huo.mutation.DistrictCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "History.district"`)
	}
	if _, ok := huo.mutation.UsersID(); huo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "History.users"`)
	}
	return nil
}

func (huo *HistoryUpdateOne) sqlSave(ctx context.Context) (_node *History, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "History.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, history.FieldID)
		for _, f := range fields {
			if !history.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Name(); ok {
		_spec.SetField(history.FieldName, field.TypeString, value)
	}
	if huo.mutation.OperationsNumCleared() {
		_spec.ClearField(history.FieldOperationsNum, field.TypeInt)
	}
	if huo.mutation.OtherCleared() {
		_spec.ClearField(history.FieldOther, field.TypeString)
	}
	_node = &History{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
