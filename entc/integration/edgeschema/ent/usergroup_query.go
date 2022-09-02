// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usergroup"
	"entgo.io/ent/schema/field"
)

// UserGroupQuery is the builder for querying UserGroup entities.
type UserGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserGroup
	withUser   *UserQuery
	withGroup  *GroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserGroupQuery builder.
func (ugq *UserGroupQuery) Where(ps ...predicate.UserGroup) *UserGroupQuery {
	ugq.predicates = append(ugq.predicates, ps...)
	return ugq
}

// Limit adds a limit step to the query.
func (ugq *UserGroupQuery) Limit(limit int) *UserGroupQuery {
	ugq.limit = &limit
	return ugq
}

// Offset adds an offset step to the query.
func (ugq *UserGroupQuery) Offset(offset int) *UserGroupQuery {
	ugq.offset = &offset
	return ugq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ugq *UserGroupQuery) Unique(unique bool) *UserGroupQuery {
	ugq.unique = &unique
	return ugq
}

// Order adds an order step to the query.
func (ugq *UserGroupQuery) Order(o ...OrderFunc) *UserGroupQuery {
	ugq.order = append(ugq.order, o...)
	return ugq
}

// QueryUser chains the current query on the "user" edge.
func (ugq *UserGroupQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: ugq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ugq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ugq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usergroup.Table, usergroup.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usergroup.UserTable, usergroup.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ugq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (ugq *UserGroupQuery) QueryGroup() *GroupQuery {
	query := &GroupQuery{config: ugq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ugq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ugq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usergroup.Table, usergroup.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usergroup.GroupTable, usergroup.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(ugq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserGroup entity from the query.
// Returns a *NotFoundError when no UserGroup was found.
func (ugq *UserGroupQuery) First(ctx context.Context) (*UserGroup, error) {
	nodes, err := ugq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usergroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ugq *UserGroupQuery) FirstX(ctx context.Context) *UserGroup {
	node, err := ugq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserGroup ID from the query.
// Returns a *NotFoundError when no UserGroup ID was found.
func (ugq *UserGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ugq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usergroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ugq *UserGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := ugq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserGroup entity is found.
// Returns a *NotFoundError when no UserGroup entities are found.
func (ugq *UserGroupQuery) Only(ctx context.Context) (*UserGroup, error) {
	nodes, err := ugq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usergroup.Label}
	default:
		return nil, &NotSingularError{usergroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ugq *UserGroupQuery) OnlyX(ctx context.Context) *UserGroup {
	node, err := ugq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserGroup ID in the query.
// Returns a *NotSingularError when more than one UserGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (ugq *UserGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ugq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usergroup.Label}
	default:
		err = &NotSingularError{usergroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ugq *UserGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := ugq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserGroups.
func (ugq *UserGroupQuery) All(ctx context.Context) ([]*UserGroup, error) {
	if err := ugq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ugq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ugq *UserGroupQuery) AllX(ctx context.Context) []*UserGroup {
	nodes, err := ugq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserGroup IDs.
func (ugq *UserGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ugq.Select(usergroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ugq *UserGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := ugq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ugq *UserGroupQuery) Count(ctx context.Context) (int, error) {
	if err := ugq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ugq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ugq *UserGroupQuery) CountX(ctx context.Context) int {
	count, err := ugq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ugq *UserGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := ugq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ugq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ugq *UserGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := ugq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ugq *UserGroupQuery) Clone() *UserGroupQuery {
	if ugq == nil {
		return nil
	}
	return &UserGroupQuery{
		config:     ugq.config,
		limit:      ugq.limit,
		offset:     ugq.offset,
		order:      append([]OrderFunc{}, ugq.order...),
		predicates: append([]predicate.UserGroup{}, ugq.predicates...),
		withUser:   ugq.withUser.Clone(),
		withGroup:  ugq.withGroup.Clone(),
		// clone intermediate query.
		sql:    ugq.sql.Clone(),
		path:   ugq.path,
		unique: ugq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ugq *UserGroupQuery) WithUser(opts ...func(*UserQuery)) *UserGroupQuery {
	query := &UserQuery{config: ugq.config}
	for _, opt := range opts {
		opt(query)
	}
	ugq.withUser = query
	return ugq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (ugq *UserGroupQuery) WithGroup(opts ...func(*GroupQuery)) *UserGroupQuery {
	query := &GroupQuery{config: ugq.config}
	for _, opt := range opts {
		opt(query)
	}
	ugq.withGroup = query
	return ugq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JoinedAt time.Time `json:"joined_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserGroup.Query().
//		GroupBy(usergroup.FieldJoinedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ugq *UserGroupQuery) GroupBy(field string, fields ...string) *UserGroupGroupBy {
	grbuild := &UserGroupGroupBy{config: ugq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ugq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ugq.sqlQuery(ctx), nil
	}
	grbuild.label = usergroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JoinedAt time.Time `json:"joined_at,omitempty"`
//	}
//
//	client.UserGroup.Query().
//		Select(usergroup.FieldJoinedAt).
//		Scan(ctx, &v)
//
func (ugq *UserGroupQuery) Select(fields ...string) *UserGroupSelect {
	ugq.fields = append(ugq.fields, fields...)
	selbuild := &UserGroupSelect{UserGroupQuery: ugq}
	selbuild.label = usergroup.Label
	selbuild.flds, selbuild.scan = &ugq.fields, selbuild.Scan
	return selbuild
}

func (ugq *UserGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ugq.fields {
		if !usergroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ugq.path != nil {
		prev, err := ugq.path(ctx)
		if err != nil {
			return err
		}
		ugq.sql = prev
	}
	return nil
}

func (ugq *UserGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserGroup, error) {
	var (
		nodes       = []*UserGroup{}
		_spec       = ugq.querySpec()
		loadedTypes = [2]bool{
			ugq.withUser != nil,
			ugq.withGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserGroup{config: ugq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ugq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ugq.withUser; query != nil {
		if err := ugq.loadUser(ctx, query, nodes, nil,
			func(n *UserGroup, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ugq.withGroup; query != nil {
		if err := ugq.loadGroup(ctx, query, nodes, nil,
			func(n *UserGroup, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ugq *UserGroupQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserGroup, init func(*UserGroup), assign func(*UserGroup, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserGroup)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ugq *UserGroupQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*UserGroup, init func(*UserGroup), assign func(*UserGroup, *Group)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserGroup)
	for i := range nodes {
		fk := nodes[i].GroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ugq *UserGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ugq.querySpec()
	_spec.Node.Columns = ugq.fields
	if len(ugq.fields) > 0 {
		_spec.Unique = ugq.unique != nil && *ugq.unique
	}
	return sqlgraph.CountNodes(ctx, ugq.driver, _spec)
}

func (ugq *UserGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	_, err := ugq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return !IsNotFound(err), nil
}

func (ugq *UserGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usergroup.Table,
			Columns: usergroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usergroup.FieldID,
			},
		},
		From:   ugq.sql,
		Unique: true,
	}
	if unique := ugq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ugq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usergroup.FieldID)
		for i := range fields {
			if fields[i] != usergroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ugq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ugq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ugq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ugq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ugq *UserGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ugq.driver.Dialect())
	t1 := builder.Table(usergroup.Table)
	columns := ugq.fields
	if len(columns) == 0 {
		columns = usergroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ugq.sql != nil {
		selector = ugq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ugq.unique != nil && *ugq.unique {
		selector.Distinct()
	}
	for _, p := range ugq.predicates {
		p(selector)
	}
	for _, p := range ugq.order {
		p(selector)
	}
	if offset := ugq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ugq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupGroupBy is the group-by builder for UserGroup entities.
type UserGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uggb *UserGroupGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupGroupBy {
	uggb.fns = append(uggb.fns, fns...)
	return uggb
}

// Scan applies the group-by query and scans the result into the given value.
func (uggb *UserGroupGroupBy) Scan(ctx context.Context, v any) error {
	query, err := uggb.path(ctx)
	if err != nil {
		return err
	}
	uggb.sql = query
	return uggb.sqlScan(ctx, v)
}

func (uggb *UserGroupGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range uggb.fields {
		if !usergroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uggb *UserGroupGroupBy) sqlQuery() *sql.Selector {
	selector := uggb.sql.Select()
	aggregation := make([]string, 0, len(uggb.fns))
	for _, fn := range uggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uggb.fields)+len(uggb.fns))
		for _, f := range uggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uggb.fields...)...)
}

// UserGroupSelect is the builder for selecting fields of UserGroup entities.
type UserGroupSelect struct {
	*UserGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ugs *UserGroupSelect) Scan(ctx context.Context, v any) error {
	if err := ugs.prepareQuery(ctx); err != nil {
		return err
	}
	ugs.sql = ugs.UserGroupQuery.sqlQuery(ctx)
	return ugs.sqlScan(ctx, v)
}

func (ugs *UserGroupSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ugs.sql.Query()
	if err := ugs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
