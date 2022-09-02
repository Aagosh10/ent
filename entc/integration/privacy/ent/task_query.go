// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/privacy/ent/predicate"
	"entgo.io/ent/entc/integration/privacy/ent/task"
	"entgo.io/ent/entc/integration/privacy/ent/team"
	"entgo.io/ent/entc/integration/privacy/ent/user"
	"entgo.io/ent/schema/field"
)

// TaskQuery is the builder for querying Task entities.
type TaskQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Task
	withTeams  *TeamQuery
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskQuery builder.
func (tq *TaskQuery) Where(ps ...predicate.Task) *TaskQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TaskQuery) Limit(limit int) *TaskQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TaskQuery) Offset(offset int) *TaskQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TaskQuery) Unique(unique bool) *TaskQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TaskQuery) Order(o ...OrderFunc) *TaskQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryTeams chains the current query on the "teams" edge.
func (tq *TaskQuery) QueryTeams() *TeamQuery {
	query := &TeamQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, task.TeamsTable, task.TeamsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (tq *TaskQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, task.OwnerTable, task.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Task entity from the query.
// Returns a *NotFoundError when no Task was found.
func (tq *TaskQuery) First(ctx context.Context) (*Task, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{task.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TaskQuery) FirstX(ctx context.Context) *Task {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Task ID from the query.
// Returns a *NotFoundError when no Task ID was found.
func (tq *TaskQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{task.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TaskQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Task entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Task entity is found.
// Returns a *NotFoundError when no Task entities are found.
func (tq *TaskQuery) Only(ctx context.Context) (*Task, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{task.Label}
	default:
		return nil, &NotSingularError{task.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TaskQuery) OnlyX(ctx context.Context) *Task {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Task ID in the query.
// Returns a *NotSingularError when more than one Task ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TaskQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{task.Label}
	default:
		err = &NotSingularError{task.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TaskQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tasks.
func (tq *TaskQuery) All(ctx context.Context) ([]*Task, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TaskQuery) AllX(ctx context.Context) []*Task {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Task IDs.
func (tq *TaskQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(task.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TaskQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TaskQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TaskQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TaskQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TaskQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TaskQuery) Clone() *TaskQuery {
	if tq == nil {
		return nil
	}
	return &TaskQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]OrderFunc{}, tq.order...),
		predicates: append([]predicate.Task{}, tq.predicates...),
		withTeams:  tq.withTeams.Clone(),
		withOwner:  tq.withOwner.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithTeams tells the query-builder to eager-load the nodes that are connected to
// the "teams" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TaskQuery) WithTeams(opts ...func(*TeamQuery)) *TaskQuery {
	query := &TeamQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withTeams = query
	return tq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TaskQuery) WithOwner(opts ...func(*UserQuery)) *TaskQuery {
	query := &UserQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withOwner = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Task.Query().
//		GroupBy(task.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TaskQuery) GroupBy(field string, fields ...string) *TaskGroupBy {
	grbuild := &TaskGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = task.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Task.Query().
//		Select(task.FieldTitle).
//		Scan(ctx, &v)
//
func (tq *TaskQuery) Select(fields ...string) *TaskSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TaskSelect{TaskQuery: tq}
	selbuild.label = task.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

func (tq *TaskQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !task.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	if task.Policy == nil {
		return errors.New("ent: uninitialized task.Policy (forgotten import ent/runtime?)")
	}
	if err := task.Policy.EvalQuery(ctx, tq); err != nil {
		return err
	}
	return nil
}

func (tq *TaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Task, error) {
	var (
		nodes       = []*Task{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withTeams != nil,
			tq.withOwner != nil,
		}
	)
	if tq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, task.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Task).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Task{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withTeams; query != nil {
		if err := tq.loadTeams(ctx, query, nodes,
			func(n *Task) { n.Edges.Teams = []*Team{} },
			func(n *Task, e *Team) { n.Edges.Teams = append(n.Edges.Teams, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withOwner; query != nil {
		if err := tq.loadOwner(ctx, query, nodes, nil,
			func(n *Task, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TaskQuery) loadTeams(ctx context.Context, query *TeamQuery, nodes []*Task, init func(*Task), assign func(*Task, *Team)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Task)
	nids := make(map[int]map[*Task]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(task.TeamsTable)
		s.Join(joinT).On(s.C(team.FieldID), joinT.C(task.TeamsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(task.TeamsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(task.TeamsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Task]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "teams" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (tq *TaskQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Task, init func(*Task), assign func(*Task, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Task)
	for i := range nodes {
		if nodes[i].user_tasks == nil {
			continue
		}
		fk := *nodes[i].user_tasks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TaskQuery) sqlExist(ctx context.Context) (bool, error) {
	_, err := tq.FirstID(ctx)
	switch {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (tq *TaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for i := range fields {
			if fields[i] != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(task.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = task.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskGroupBy is the group-by builder for Task entities.
type TaskGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TaskGroupBy) Aggregate(fns ...AggregateFunc) *TaskGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TaskGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *TaskGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tgb.fields {
		if !task.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TaskGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TaskSelect is the builder for selecting fields of Task entities.
type TaskSelect struct {
	*TaskQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TaskSelect) Scan(ctx context.Context, v any) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TaskQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *TaskSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
