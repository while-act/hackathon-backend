// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/while-act/hackathon-backend/ent/company"
	"github.com/while-act/hackathon-backend/ent/user"
)

// CompanyCreate is the builder for creating a Company entity.
type CompanyCreate struct {
	config
	mutation *CompanyMutation
	hooks    []Hook
}

// SetInn sets the "inn" field.
func (cc *CompanyCreate) SetInn(s string) *CompanyCreate {
	cc.mutation.SetInn(s)
	return cc
}

// SetName sets the "name" field.
func (cc *CompanyCreate) SetName(s string) *CompanyCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableName(s *string) *CompanyCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetWebsite sets the "website" field.
func (cc *CompanyCreate) SetWebsite(s string) *CompanyCreate {
	cc.mutation.SetWebsite(s)
	return cc
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableWebsite(s *string) *CompanyCreate {
	if s != nil {
		cc.SetWebsite(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CompanyCreate) SetID(i int) *CompanyCreate {
	cc.mutation.SetID(i)
	return cc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cc *CompanyCreate) SetUsersID(id int) *CompanyCreate {
	cc.mutation.SetUsersID(id)
	return cc
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (cc *CompanyCreate) SetNillableUsersID(id *int) *CompanyCreate {
	if id != nil {
		cc = cc.SetUsersID(*id)
	}
	return cc
}

// SetUsers sets the "users" edge to the User entity.
func (cc *CompanyCreate) SetUsers(u *User) *CompanyCreate {
	return cc.SetUsersID(u.ID)
}

// Mutation returns the CompanyMutation object of the builder.
func (cc *CompanyCreate) Mutation() *CompanyMutation {
	return cc.mutation
}

// Save creates the Company in the database.
func (cc *CompanyCreate) Save(ctx context.Context) (*Company, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompanyCreate) SaveX(ctx context.Context) *Company {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompanyCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompanyCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompanyCreate) check() error {
	if _, ok := cc.mutation.Inn(); !ok {
		return &ValidationError{Name: "inn", err: errors.New(`ent: missing required field "Company.inn"`)}
	}
	if v, ok := cc.mutation.Inn(); ok {
		if err := company.InnValidator(v); err != nil {
			return &ValidationError{Name: "inn", err: fmt.Errorf(`ent: validator failed for field "Company.inn": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := company.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Company.name": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Website(); ok {
		if err := company.WebsiteValidator(v); err != nil {
			return &ValidationError{Name: "website", err: fmt.Errorf(`ent: validator failed for field "Company.website": %w`, err)}
		}
	}
	return nil
}

func (cc *CompanyCreate) sqlSave(ctx context.Context) (*Company, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CompanyCreate) createSpec() (*Company, *sqlgraph.CreateSpec) {
	var (
		_node = &Company{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(company.Table, sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Inn(); ok {
		_spec.SetField(company.FieldInn, field.TypeString, value)
		_node.Inn = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(company.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := cc.mutation.Website(); ok {
		_spec.SetField(company.FieldWebsite, field.TypeString, value)
		_node.Website = &value
	}
	if nodes := cc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   company.UsersTable,
			Columns: []string{company.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompanyCreateBulk is the builder for creating many Company entities in bulk.
type CompanyCreateBulk struct {
	config
	builders []*CompanyCreate
}

// Save creates the Company entities in the database.
func (ccb *CompanyCreateBulk) Save(ctx context.Context) ([]*Company, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Company, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompanyCreateBulk) SaveX(ctx context.Context) []*Company {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CompanyCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CompanyCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
