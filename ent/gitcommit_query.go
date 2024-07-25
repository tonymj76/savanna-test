// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tonymj76/savannah/ent/gitcommit"
	"github.com/tonymj76/savannah/ent/predicate"
)

// GitCommitQuery is the builder for querying GitCommit entities.
type GitCommitQuery struct {
	config
	ctx        *QueryContext
	order      []gitcommit.OrderOption
	inters     []Interceptor
	predicates []predicate.GitCommit
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GitCommitQuery builder.
func (gcq *GitCommitQuery) Where(ps ...predicate.GitCommit) *GitCommitQuery {
	gcq.predicates = append(gcq.predicates, ps...)
	return gcq
}

// Limit the number of records to be returned by this query.
func (gcq *GitCommitQuery) Limit(limit int) *GitCommitQuery {
	gcq.ctx.Limit = &limit
	return gcq
}

// Offset to start from.
func (gcq *GitCommitQuery) Offset(offset int) *GitCommitQuery {
	gcq.ctx.Offset = &offset
	return gcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gcq *GitCommitQuery) Unique(unique bool) *GitCommitQuery {
	gcq.ctx.Unique = &unique
	return gcq
}

// Order specifies how the records should be ordered.
func (gcq *GitCommitQuery) Order(o ...gitcommit.OrderOption) *GitCommitQuery {
	gcq.order = append(gcq.order, o...)
	return gcq
}

// First returns the first GitCommit entity from the query.
// Returns a *NotFoundError when no GitCommit was found.
func (gcq *GitCommitQuery) First(ctx context.Context) (*GitCommit, error) {
	nodes, err := gcq.Limit(1).All(setContextOp(ctx, gcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gitcommit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gcq *GitCommitQuery) FirstX(ctx context.Context) *GitCommit {
	node, err := gcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GitCommit ID from the query.
// Returns a *NotFoundError when no GitCommit ID was found.
func (gcq *GitCommitQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gcq.Limit(1).IDs(setContextOp(ctx, gcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gitcommit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gcq *GitCommitQuery) FirstIDX(ctx context.Context) int {
	id, err := gcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GitCommit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GitCommit entity is found.
// Returns a *NotFoundError when no GitCommit entities are found.
func (gcq *GitCommitQuery) Only(ctx context.Context) (*GitCommit, error) {
	nodes, err := gcq.Limit(2).All(setContextOp(ctx, gcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gitcommit.Label}
	default:
		return nil, &NotSingularError{gitcommit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gcq *GitCommitQuery) OnlyX(ctx context.Context) *GitCommit {
	node, err := gcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GitCommit ID in the query.
// Returns a *NotSingularError when more than one GitCommit ID is found.
// Returns a *NotFoundError when no entities are found.
func (gcq *GitCommitQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gcq.Limit(2).IDs(setContextOp(ctx, gcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gitcommit.Label}
	default:
		err = &NotSingularError{gitcommit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gcq *GitCommitQuery) OnlyIDX(ctx context.Context) int {
	id, err := gcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GitCommits.
func (gcq *GitCommitQuery) All(ctx context.Context) ([]*GitCommit, error) {
	ctx = setContextOp(ctx, gcq.ctx, "All")
	if err := gcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GitCommit, *GitCommitQuery]()
	return withInterceptors[[]*GitCommit](ctx, gcq, qr, gcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gcq *GitCommitQuery) AllX(ctx context.Context) []*GitCommit {
	nodes, err := gcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GitCommit IDs.
func (gcq *GitCommitQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gcq.ctx.Unique == nil && gcq.path != nil {
		gcq.Unique(true)
	}
	ctx = setContextOp(ctx, gcq.ctx, "IDs")
	if err = gcq.Select(gitcommit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gcq *GitCommitQuery) IDsX(ctx context.Context) []int {
	ids, err := gcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gcq *GitCommitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gcq.ctx, "Count")
	if err := gcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gcq, querierCount[*GitCommitQuery](), gcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gcq *GitCommitQuery) CountX(ctx context.Context) int {
	count, err := gcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gcq *GitCommitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gcq.ctx, "Exist")
	switch _, err := gcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gcq *GitCommitQuery) ExistX(ctx context.Context) bool {
	exist, err := gcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GitCommitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gcq *GitCommitQuery) Clone() *GitCommitQuery {
	if gcq == nil {
		return nil
	}
	return &GitCommitQuery{
		config:     gcq.config,
		ctx:        gcq.ctx.Clone(),
		order:      append([]gitcommit.OrderOption{}, gcq.order...),
		inters:     append([]Interceptor{}, gcq.inters...),
		predicates: append([]predicate.GitCommit{}, gcq.predicates...),
		// clone intermediate query.
		sql:  gcq.sql.Clone(),
		path: gcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Gitcommit any `json:"gitcommit,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GitCommit.Query().
//		GroupBy(gitcommit.FieldGitcommit).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gcq *GitCommitQuery) GroupBy(field string, fields ...string) *GitCommitGroupBy {
	gcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GitCommitGroupBy{build: gcq}
	grbuild.flds = &gcq.ctx.Fields
	grbuild.label = gitcommit.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Gitcommit any `json:"gitcommit,omitempty"`
//	}
//
//	client.GitCommit.Query().
//		Select(gitcommit.FieldGitcommit).
//		Scan(ctx, &v)
func (gcq *GitCommitQuery) Select(fields ...string) *GitCommitSelect {
	gcq.ctx.Fields = append(gcq.ctx.Fields, fields...)
	sbuild := &GitCommitSelect{GitCommitQuery: gcq}
	sbuild.label = gitcommit.Label
	sbuild.flds, sbuild.scan = &gcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GitCommitSelect configured with the given aggregations.
func (gcq *GitCommitQuery) Aggregate(fns ...AggregateFunc) *GitCommitSelect {
	return gcq.Select().Aggregate(fns...)
}

func (gcq *GitCommitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gcq); err != nil {
				return err
			}
		}
	}
	for _, f := range gcq.ctx.Fields {
		if !gitcommit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gcq.path != nil {
		prev, err := gcq.path(ctx)
		if err != nil {
			return err
		}
		gcq.sql = prev
	}
	return nil
}

func (gcq *GitCommitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GitCommit, error) {
	var (
		nodes   = []*GitCommit{}
		withFKs = gcq.withFKs
		_spec   = gcq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, gitcommit.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GitCommit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GitCommit{config: gcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gcq *GitCommitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gcq.querySpec()
	_spec.Node.Columns = gcq.ctx.Fields
	if len(gcq.ctx.Fields) > 0 {
		_spec.Unique = gcq.ctx.Unique != nil && *gcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gcq.driver, _spec)
}

func (gcq *GitCommitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(gitcommit.Table, gitcommit.Columns, sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt))
	_spec.From = gcq.sql
	if unique := gcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gcq.path != nil {
		_spec.Unique = true
	}
	if fields := gcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gitcommit.FieldID)
		for i := range fields {
			if fields[i] != gitcommit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gcq *GitCommitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gcq.driver.Dialect())
	t1 := builder.Table(gitcommit.Table)
	columns := gcq.ctx.Fields
	if len(columns) == 0 {
		columns = gitcommit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gcq.sql != nil {
		selector = gcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gcq.ctx.Unique != nil && *gcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gcq.predicates {
		p(selector)
	}
	for _, p := range gcq.order {
		p(selector)
	}
	if offset := gcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GitCommitGroupBy is the group-by builder for GitCommit entities.
type GitCommitGroupBy struct {
	selector
	build *GitCommitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gcgb *GitCommitGroupBy) Aggregate(fns ...AggregateFunc) *GitCommitGroupBy {
	gcgb.fns = append(gcgb.fns, fns...)
	return gcgb
}

// Scan applies the selector query and scans the result into the given value.
func (gcgb *GitCommitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gcgb.build.ctx, "GroupBy")
	if err := gcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GitCommitQuery, *GitCommitGroupBy](ctx, gcgb.build, gcgb, gcgb.build.inters, v)
}

func (gcgb *GitCommitGroupBy) sqlScan(ctx context.Context, root *GitCommitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gcgb.fns))
	for _, fn := range gcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gcgb.flds)+len(gcgb.fns))
		for _, f := range *gcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GitCommitSelect is the builder for selecting fields of GitCommit entities.
type GitCommitSelect struct {
	*GitCommitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gcs *GitCommitSelect) Aggregate(fns ...AggregateFunc) *GitCommitSelect {
	gcs.fns = append(gcs.fns, fns...)
	return gcs
}

// Scan applies the selector query and scans the result into the given value.
func (gcs *GitCommitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gcs.ctx, "Select")
	if err := gcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GitCommitQuery, *GitCommitSelect](ctx, gcs.GitCommitQuery, gcs, gcs.inters, v)
}

func (gcs *GitCommitSelect) sqlScan(ctx context.Context, root *GitCommitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gcs.fns))
	for _, fn := range gcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
