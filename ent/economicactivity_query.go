// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/while.act/ent/economicactivity"
	"github.com/wtkeqrf0/while.act/ent/predicate"
)

// EconomicActivityQuery is the builder for querying EconomicActivity entities.
type EconomicActivityQuery struct {
	config
	ctx        *QueryContext
	order      []economicactivity.OrderOption
	inters     []Interceptor
	predicates []predicate.EconomicActivity
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EconomicActivityQuery builder.
func (eaq *EconomicActivityQuery) Where(ps ...predicate.EconomicActivity) *EconomicActivityQuery {
	eaq.predicates = append(eaq.predicates, ps...)
	return eaq
}

// Limit the number of records to be returned by this query.
func (eaq *EconomicActivityQuery) Limit(limit int) *EconomicActivityQuery {
	eaq.ctx.Limit = &limit
	return eaq
}

// Offset to start from.
func (eaq *EconomicActivityQuery) Offset(offset int) *EconomicActivityQuery {
	eaq.ctx.Offset = &offset
	return eaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eaq *EconomicActivityQuery) Unique(unique bool) *EconomicActivityQuery {
	eaq.ctx.Unique = &unique
	return eaq
}

// Order specifies how the records should be ordered.
func (eaq *EconomicActivityQuery) Order(o ...economicactivity.OrderOption) *EconomicActivityQuery {
	eaq.order = append(eaq.order, o...)
	return eaq
}

// First returns the first EconomicActivity entity from the query.
// Returns a *NotFoundError when no EconomicActivity was found.
func (eaq *EconomicActivityQuery) First(ctx context.Context) (*EconomicActivity, error) {
	nodes, err := eaq.Limit(1).All(setContextOp(ctx, eaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{economicactivity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eaq *EconomicActivityQuery) FirstX(ctx context.Context) *EconomicActivity {
	node, err := eaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EconomicActivity ID from the query.
// Returns a *NotFoundError when no EconomicActivity ID was found.
func (eaq *EconomicActivityQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eaq.Limit(1).IDs(setContextOp(ctx, eaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{economicactivity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eaq *EconomicActivityQuery) FirstIDX(ctx context.Context) string {
	id, err := eaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EconomicActivity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EconomicActivity entity is found.
// Returns a *NotFoundError when no EconomicActivity entities are found.
func (eaq *EconomicActivityQuery) Only(ctx context.Context) (*EconomicActivity, error) {
	nodes, err := eaq.Limit(2).All(setContextOp(ctx, eaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{economicactivity.Label}
	default:
		return nil, &NotSingularError{economicactivity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eaq *EconomicActivityQuery) OnlyX(ctx context.Context) *EconomicActivity {
	node, err := eaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EconomicActivity ID in the query.
// Returns a *NotSingularError when more than one EconomicActivity ID is found.
// Returns a *NotFoundError when no entities are found.
func (eaq *EconomicActivityQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eaq.Limit(2).IDs(setContextOp(ctx, eaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{economicactivity.Label}
	default:
		err = &NotSingularError{economicactivity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eaq *EconomicActivityQuery) OnlyIDX(ctx context.Context) string {
	id, err := eaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EconomicActivities.
func (eaq *EconomicActivityQuery) All(ctx context.Context) ([]*EconomicActivity, error) {
	ctx = setContextOp(ctx, eaq.ctx, "All")
	if err := eaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EconomicActivity, *EconomicActivityQuery]()
	return withInterceptors[[]*EconomicActivity](ctx, eaq, qr, eaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eaq *EconomicActivityQuery) AllX(ctx context.Context) []*EconomicActivity {
	nodes, err := eaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EconomicActivity IDs.
func (eaq *EconomicActivityQuery) IDs(ctx context.Context) (ids []string, err error) {
	if eaq.ctx.Unique == nil && eaq.path != nil {
		eaq.Unique(true)
	}
	ctx = setContextOp(ctx, eaq.ctx, "IDs")
	if err = eaq.Select(economicactivity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eaq *EconomicActivityQuery) IDsX(ctx context.Context) []string {
	ids, err := eaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eaq *EconomicActivityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eaq.ctx, "Count")
	if err := eaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eaq, querierCount[*EconomicActivityQuery](), eaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eaq *EconomicActivityQuery) CountX(ctx context.Context) int {
	count, err := eaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eaq *EconomicActivityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eaq.ctx, "Exist")
	switch _, err := eaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eaq *EconomicActivityQuery) ExistX(ctx context.Context) bool {
	exist, err := eaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EconomicActivityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eaq *EconomicActivityQuery) Clone() *EconomicActivityQuery {
	if eaq == nil {
		return nil
	}
	return &EconomicActivityQuery{
		config:     eaq.config,
		ctx:        eaq.ctx.Clone(),
		order:      append([]economicactivity.OrderOption{}, eaq.order...),
		inters:     append([]Interceptor{}, eaq.inters...),
		predicates: append([]predicate.EconomicActivity{}, eaq.predicates...),
		// clone intermediate query.
		sql:  eaq.sql.Clone(),
		path: eaq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Subs string `json:"subs,omitempty" example:"Автомобильная промышленность"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EconomicActivity.Query().
//		GroupBy(economicactivity.FieldSubs).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eaq *EconomicActivityQuery) GroupBy(field string, fields ...string) *EconomicActivityGroupBy {
	eaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EconomicActivityGroupBy{build: eaq}
	grbuild.flds = &eaq.ctx.Fields
	grbuild.label = economicactivity.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Subs string `json:"subs,omitempty" example:"Автомобильная промышленность"`
//	}
//
//	client.EconomicActivity.Query().
//		Select(economicactivity.FieldSubs).
//		Scan(ctx, &v)
func (eaq *EconomicActivityQuery) Select(fields ...string) *EconomicActivitySelect {
	eaq.ctx.Fields = append(eaq.ctx.Fields, fields...)
	sbuild := &EconomicActivitySelect{EconomicActivityQuery: eaq}
	sbuild.label = economicactivity.Label
	sbuild.flds, sbuild.scan = &eaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EconomicActivitySelect configured with the given aggregations.
func (eaq *EconomicActivityQuery) Aggregate(fns ...AggregateFunc) *EconomicActivitySelect {
	return eaq.Select().Aggregate(fns...)
}

func (eaq *EconomicActivityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eaq); err != nil {
				return err
			}
		}
	}
	for _, f := range eaq.ctx.Fields {
		if !economicactivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eaq.path != nil {
		prev, err := eaq.path(ctx)
		if err != nil {
			return err
		}
		eaq.sql = prev
	}
	return nil
}

func (eaq *EconomicActivityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EconomicActivity, error) {
	var (
		nodes = []*EconomicActivity{}
		_spec = eaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EconomicActivity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EconomicActivity{config: eaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (eaq *EconomicActivityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eaq.querySpec()
	_spec.Node.Columns = eaq.ctx.Fields
	if len(eaq.ctx.Fields) > 0 {
		_spec.Unique = eaq.ctx.Unique != nil && *eaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eaq.driver, _spec)
}

func (eaq *EconomicActivityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(economicactivity.Table, economicactivity.Columns, sqlgraph.NewFieldSpec(economicactivity.FieldID, field.TypeString))
	_spec.From = eaq.sql
	if unique := eaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eaq.path != nil {
		_spec.Unique = true
	}
	if fields := eaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, economicactivity.FieldID)
		for i := range fields {
			if fields[i] != economicactivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eaq *EconomicActivityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eaq.driver.Dialect())
	t1 := builder.Table(economicactivity.Table)
	columns := eaq.ctx.Fields
	if len(columns) == 0 {
		columns = economicactivity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eaq.sql != nil {
		selector = eaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eaq.ctx.Unique != nil && *eaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eaq.predicates {
		p(selector)
	}
	for _, p := range eaq.order {
		p(selector)
	}
	if offset := eaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EconomicActivityGroupBy is the group-by builder for EconomicActivity entities.
type EconomicActivityGroupBy struct {
	selector
	build *EconomicActivityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eagb *EconomicActivityGroupBy) Aggregate(fns ...AggregateFunc) *EconomicActivityGroupBy {
	eagb.fns = append(eagb.fns, fns...)
	return eagb
}

// Scan applies the selector query and scans the result into the given value.
func (eagb *EconomicActivityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eagb.build.ctx, "GroupBy")
	if err := eagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EconomicActivityQuery, *EconomicActivityGroupBy](ctx, eagb.build, eagb, eagb.build.inters, v)
}

func (eagb *EconomicActivityGroupBy) sqlScan(ctx context.Context, root *EconomicActivityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eagb.fns))
	for _, fn := range eagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eagb.flds)+len(eagb.fns))
		for _, f := range *eagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EconomicActivitySelect is the builder for selecting fields of EconomicActivity entities.
type EconomicActivitySelect struct {
	*EconomicActivityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eas *EconomicActivitySelect) Aggregate(fns ...AggregateFunc) *EconomicActivitySelect {
	eas.fns = append(eas.fns, fns...)
	return eas
}

// Scan applies the selector query and scans the result into the given value.
func (eas *EconomicActivitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eas.ctx, "Select")
	if err := eas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EconomicActivityQuery, *EconomicActivitySelect](ctx, eas.EconomicActivityQuery, eas, eas.inters, v)
}

func (eas *EconomicActivitySelect) sqlScan(ctx context.Context, root *EconomicActivityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eas.fns))
	for _, fn := range eas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
