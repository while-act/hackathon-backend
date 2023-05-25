// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/while-act/hackathon-backend/ent/district"
	"github.com/while-act/hackathon-backend/ent/equipment"
	"github.com/while-act/hackathon-backend/ent/history"
	"github.com/while-act/hackathon-backend/ent/industry"
	"github.com/while-act/hackathon-backend/ent/user"
)

// HistoryCreate is the builder for creating a History entity.
type HistoryCreate struct {
	config
	mutation *HistoryMutation
	hooks    []Hook
}

// SetIndustryBranch sets the "industry_branch" field.
func (hc *HistoryCreate) SetIndustryBranch(s string) *HistoryCreate {
	hc.mutation.SetIndustryBranch(s)
	return hc
}

// SetFullTimeEmployees sets the "full_time_employees" field.
func (hc *HistoryCreate) SetFullTimeEmployees(i int) *HistoryCreate {
	hc.mutation.SetFullTimeEmployees(i)
	return hc
}

// SetDistrictTitle sets the "district_title" field.
func (hc *HistoryCreate) SetDistrictTitle(s string) *HistoryCreate {
	hc.mutation.SetDistrictTitle(s)
	return hc
}

// SetLandArea sets the "land_area" field.
func (hc *HistoryCreate) SetLandArea(i int) *HistoryCreate {
	hc.mutation.SetLandArea(i)
	return hc
}

// SetConstructionFacilitiesArea sets the "construction_facilities_area" field.
func (hc *HistoryCreate) SetConstructionFacilitiesArea(i int) *HistoryCreate {
	hc.mutation.SetConstructionFacilitiesArea(i)
	return hc
}

// SetEquipmentType sets the "equipment_type" field.
func (hc *HistoryCreate) SetEquipmentType(s string) *HistoryCreate {
	hc.mutation.SetEquipmentType(s)
	return hc
}

// SetFacilityType sets the "facility_type" field.
func (hc *HistoryCreate) SetFacilityType(s string) *HistoryCreate {
	hc.mutation.SetFacilityType(s)
	return hc
}

// SetAccountingServices sets the "accounting_services" field.
func (hc *HistoryCreate) SetAccountingServices(b bool) *HistoryCreate {
	hc.mutation.SetAccountingServices(b)
	return hc
}

// SetPatent sets the "patent" field.
func (hc *HistoryCreate) SetPatent(b bool) *HistoryCreate {
	hc.mutation.SetPatent(b)
	return hc
}

// SetOther sets the "other" field.
func (hc *HistoryCreate) SetOther(s string) *HistoryCreate {
	hc.mutation.SetOther(s)
	return hc
}

// SetUserID sets the "user_id" field.
func (hc *HistoryCreate) SetUserID(i int) *HistoryCreate {
	hc.mutation.SetUserID(i)
	return hc
}

// SetID sets the "id" field.
func (hc *HistoryCreate) SetID(s string) *HistoryCreate {
	hc.mutation.SetID(s)
	return hc
}

// SetIndustryID sets the "industry" edge to the Industry entity by ID.
func (hc *HistoryCreate) SetIndustryID(id string) *HistoryCreate {
	hc.mutation.SetIndustryID(id)
	return hc
}

// SetIndustry sets the "industry" edge to the Industry entity.
func (hc *HistoryCreate) SetIndustry(i *Industry) *HistoryCreate {
	return hc.SetIndustryID(i.ID)
}

// SetDistrictID sets the "district" edge to the District entity by ID.
func (hc *HistoryCreate) SetDistrictID(id string) *HistoryCreate {
	hc.mutation.SetDistrictID(id)
	return hc
}

// SetDistrict sets the "district" edge to the District entity.
func (hc *HistoryCreate) SetDistrict(d *District) *HistoryCreate {
	return hc.SetDistrictID(d.ID)
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (hc *HistoryCreate) SetEquipmentID(id string) *HistoryCreate {
	hc.mutation.SetEquipmentID(id)
	return hc
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (hc *HistoryCreate) SetEquipment(e *Equipment) *HistoryCreate {
	return hc.SetEquipmentID(e.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (hc *HistoryCreate) SetUsersID(id int) *HistoryCreate {
	hc.mutation.SetUsersID(id)
	return hc
}

// SetUsers sets the "users" edge to the User entity.
func (hc *HistoryCreate) SetUsers(u *User) *HistoryCreate {
	return hc.SetUsersID(u.ID)
}

// Mutation returns the HistoryMutation object of the builder.
func (hc *HistoryCreate) Mutation() *HistoryMutation {
	return hc.mutation
}

// Save creates the History in the database.
func (hc *HistoryCreate) Save(ctx context.Context) (*History, error) {
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HistoryCreate) SaveX(ctx context.Context) *History {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HistoryCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HistoryCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HistoryCreate) check() error {
	if _, ok := hc.mutation.IndustryBranch(); !ok {
		return &ValidationError{Name: "industry_branch", err: errors.New(`ent: missing required field "History.industry_branch"`)}
	}
	if _, ok := hc.mutation.FullTimeEmployees(); !ok {
		return &ValidationError{Name: "full_time_employees", err: errors.New(`ent: missing required field "History.full_time_employees"`)}
	}
	if v, ok := hc.mutation.FullTimeEmployees(); ok {
		if err := history.FullTimeEmployeesValidator(v); err != nil {
			return &ValidationError{Name: "full_time_employees", err: fmt.Errorf(`ent: validator failed for field "History.full_time_employees": %w`, err)}
		}
	}
	if _, ok := hc.mutation.DistrictTitle(); !ok {
		return &ValidationError{Name: "district_title", err: errors.New(`ent: missing required field "History.district_title"`)}
	}
	if _, ok := hc.mutation.LandArea(); !ok {
		return &ValidationError{Name: "land_area", err: errors.New(`ent: missing required field "History.land_area"`)}
	}
	if v, ok := hc.mutation.LandArea(); ok {
		if err := history.LandAreaValidator(v); err != nil {
			return &ValidationError{Name: "land_area", err: fmt.Errorf(`ent: validator failed for field "History.land_area": %w`, err)}
		}
	}
	if _, ok := hc.mutation.ConstructionFacilitiesArea(); !ok {
		return &ValidationError{Name: "construction_facilities_area", err: errors.New(`ent: missing required field "History.construction_facilities_area"`)}
	}
	if v, ok := hc.mutation.ConstructionFacilitiesArea(); ok {
		if err := history.ConstructionFacilitiesAreaValidator(v); err != nil {
			return &ValidationError{Name: "construction_facilities_area", err: fmt.Errorf(`ent: validator failed for field "History.construction_facilities_area": %w`, err)}
		}
	}
	if _, ok := hc.mutation.EquipmentType(); !ok {
		return &ValidationError{Name: "equipment_type", err: errors.New(`ent: missing required field "History.equipment_type"`)}
	}
	if _, ok := hc.mutation.FacilityType(); !ok {
		return &ValidationError{Name: "facility_type", err: errors.New(`ent: missing required field "History.facility_type"`)}
	}
	if _, ok := hc.mutation.AccountingServices(); !ok {
		return &ValidationError{Name: "accounting_services", err: errors.New(`ent: missing required field "History.accounting_services"`)}
	}
	if _, ok := hc.mutation.Patent(); !ok {
		return &ValidationError{Name: "patent", err: errors.New(`ent: missing required field "History.patent"`)}
	}
	if _, ok := hc.mutation.Other(); !ok {
		return &ValidationError{Name: "other", err: errors.New(`ent: missing required field "History.other"`)}
	}
	if _, ok := hc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "History.user_id"`)}
	}
	if v, ok := hc.mutation.ID(); ok {
		if err := history.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "History.id": %w`, err)}
		}
	}
	if _, ok := hc.mutation.IndustryID(); !ok {
		return &ValidationError{Name: "industry", err: errors.New(`ent: missing required edge "History.industry"`)}
	}
	if _, ok := hc.mutation.DistrictID(); !ok {
		return &ValidationError{Name: "district", err: errors.New(`ent: missing required edge "History.district"`)}
	}
	if _, ok := hc.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment", err: errors.New(`ent: missing required edge "History.equipment"`)}
	}
	if _, ok := hc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "History.users"`)}
	}
	return nil
}

func (hc *HistoryCreate) sqlSave(ctx context.Context) (*History, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected History.ID type: %T", _spec.ID.Value)
		}
	}
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HistoryCreate) createSpec() (*History, *sqlgraph.CreateSpec) {
	var (
		_node = &History{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(history.Table, sqlgraph.NewFieldSpec(history.FieldID, field.TypeString))
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.FullTimeEmployees(); ok {
		_spec.SetField(history.FieldFullTimeEmployees, field.TypeInt, value)
		_node.FullTimeEmployees = value
	}
	if value, ok := hc.mutation.LandArea(); ok {
		_spec.SetField(history.FieldLandArea, field.TypeInt, value)
		_node.LandArea = value
	}
	if value, ok := hc.mutation.ConstructionFacilitiesArea(); ok {
		_spec.SetField(history.FieldConstructionFacilitiesArea, field.TypeInt, value)
		_node.ConstructionFacilitiesArea = value
	}
	if value, ok := hc.mutation.FacilityType(); ok {
		_spec.SetField(history.FieldFacilityType, field.TypeString, value)
		_node.FacilityType = value
	}
	if value, ok := hc.mutation.AccountingServices(); ok {
		_spec.SetField(history.FieldAccountingServices, field.TypeBool, value)
		_node.AccountingServices = value
	}
	if value, ok := hc.mutation.Patent(); ok {
		_spec.SetField(history.FieldPatent, field.TypeBool, value)
		_node.Patent = value
	}
	if value, ok := hc.mutation.Other(); ok {
		_spec.SetField(history.FieldOther, field.TypeString, value)
		_node.Other = value
	}
	if nodes := hc.mutation.IndustryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.IndustryTable,
			Columns: []string{history.IndustryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(industry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.IndustryBranch = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.DistrictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.DistrictTable,
			Columns: []string{history.DistrictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(district.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DistrictTitle = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.EquipmentTable,
			Columns: []string{history.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EquipmentType = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.UsersTable,
			Columns: []string{history.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HistoryCreateBulk is the builder for creating many History entities in bulk.
type HistoryCreateBulk struct {
	config
	builders []*HistoryCreate
}

// Save creates the History entities in the database.
func (hcb *HistoryCreateBulk) Save(ctx context.Context) ([]*History, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*History, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HistoryCreateBulk) SaveX(ctx context.Context) []*History {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HistoryCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
