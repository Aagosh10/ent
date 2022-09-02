// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweetlike"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// TweetLikeQuery is the builder for querying TweetLike entities.
type TweetLikeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TweetLike
	withTweet  *TweetQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TweetLikeQuery builder.
func (tlq *TweetLikeQuery) Where(ps ...predicate.TweetLike) *TweetLikeQuery {
	tlq.predicates = append(tlq.predicates, ps...)
	return tlq
}

// Limit adds a limit step to the query.
func (tlq *TweetLikeQuery) Limit(limit int) *TweetLikeQuery {
	tlq.limit = &limit
	return tlq
}

// Offset adds an offset step to the query.
func (tlq *TweetLikeQuery) Offset(offset int) *TweetLikeQuery {
	tlq.offset = &offset
	return tlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tlq *TweetLikeQuery) Unique(unique bool) *TweetLikeQuery {
	tlq.unique = &unique
	return tlq
}

// Order adds an order step to the query.
func (tlq *TweetLikeQuery) Order(o ...OrderFunc) *TweetLikeQuery {
	tlq.order = append(tlq.order, o...)
	return tlq
}

// QueryTweet chains the current query on the "tweet" edge.
func (tlq *TweetLikeQuery) QueryTweet() *TweetQuery {
	query := &TweetQuery{config: tlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweetlike.Table, tweetlike.TweetColumn, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweetlike.TweetTable, tweetlike.TweetColumn),
		)
		fromU = sqlgraph.SetNeighbors(tlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (tlq *TweetLikeQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: tlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweetlike.Table, tweetlike.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweetlike.UserTable, tweetlike.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TweetLike entity from the query.
// Returns a *NotFoundError when no TweetLike was found.
func (tlq *TweetLikeQuery) First(ctx context.Context) (*TweetLike, error) {
	nodes, err := tlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tweetlike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tlq *TweetLikeQuery) FirstX(ctx context.Context) *TweetLike {
	node, err := tlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single TweetLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TweetLike entity is found.
// Returns a *NotFoundError when no TweetLike entities are found.
func (tlq *TweetLikeQuery) Only(ctx context.Context) (*TweetLike, error) {
	nodes, err := tlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tweetlike.Label}
	default:
		return nil, &NotSingularError{tweetlike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tlq *TweetLikeQuery) OnlyX(ctx context.Context) *TweetLike {
	node, err := tlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of TweetLikes.
func (tlq *TweetLikeQuery) All(ctx context.Context) ([]*TweetLike, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tlq *TweetLikeQuery) AllX(ctx context.Context) []*TweetLike {
	nodes, err := tlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (tlq *TweetLikeQuery) Count(ctx context.Context) (int, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tlq *TweetLikeQuery) CountX(ctx context.Context) int {
	count, err := tlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tlq *TweetLikeQuery) Exist(ctx context.Context) (bool, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tlq *TweetLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := tlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TweetLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tlq *TweetLikeQuery) Clone() *TweetLikeQuery {
	if tlq == nil {
		return nil
	}
	return &TweetLikeQuery{
		config:     tlq.config,
		limit:      tlq.limit,
		offset:     tlq.offset,
		order:      append([]OrderFunc{}, tlq.order...),
		predicates: append([]predicate.TweetLike{}, tlq.predicates...),
		withTweet:  tlq.withTweet.Clone(),
		withUser:   tlq.withUser.Clone(),
		// clone intermediate query.
		sql:    tlq.sql.Clone(),
		path:   tlq.path,
		unique: tlq.unique,
	}
}

// WithTweet tells the query-builder to eager-load the nodes that are connected to
// the "tweet" edge. The optional arguments are used to configure the query builder of the edge.
func (tlq *TweetLikeQuery) WithTweet(opts ...func(*TweetQuery)) *TweetLikeQuery {
	query := &TweetQuery{config: tlq.config}
	for _, opt := range opts {
		opt(query)
	}
	tlq.withTweet = query
	return tlq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tlq *TweetLikeQuery) WithUser(opts ...func(*UserQuery)) *TweetLikeQuery {
	query := &UserQuery{config: tlq.config}
	for _, opt := range opts {
		opt(query)
	}
	tlq.withUser = query
	return tlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LikedAt time.Time `json:"liked_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TweetLike.Query().
//		GroupBy(tweetlike.FieldLikedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tlq *TweetLikeQuery) GroupBy(field string, fields ...string) *TweetLikeGroupBy {
	grbuild := &TweetLikeGroupBy{config: tlq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tlq.sqlQuery(ctx), nil
	}
	grbuild.label = tweetlike.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LikedAt time.Time `json:"liked_at,omitempty"`
//	}
//
//	client.TweetLike.Query().
//		Select(tweetlike.FieldLikedAt).
//		Scan(ctx, &v)
//
func (tlq *TweetLikeQuery) Select(fields ...string) *TweetLikeSelect {
	tlq.fields = append(tlq.fields, fields...)
	selbuild := &TweetLikeSelect{TweetLikeQuery: tlq}
	selbuild.label = tweetlike.Label
	selbuild.flds, selbuild.scan = &tlq.fields, selbuild.Scan
	return selbuild
}

func (tlq *TweetLikeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tlq.fields {
		if !tweetlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tlq.path != nil {
		prev, err := tlq.path(ctx)
		if err != nil {
			return err
		}
		tlq.sql = prev
	}
	if tweetlike.Policy == nil {
		return errors.New("ent: uninitialized tweetlike.Policy (forgotten import ent/runtime?)")
	}
	if err := tweetlike.Policy.EvalQuery(ctx, tlq); err != nil {
		return err
	}
	return nil
}

func (tlq *TweetLikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TweetLike, error) {
	var (
		nodes       = []*TweetLike{}
		_spec       = tlq.querySpec()
		loadedTypes = [2]bool{
			tlq.withTweet != nil,
			tlq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TweetLike).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TweetLike{config: tlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tlq.withTweet; query != nil {
		if err := tlq.loadTweet(ctx, query, nodes, nil,
			func(n *TweetLike, e *Tweet) { n.Edges.Tweet = e }); err != nil {
			return nil, err
		}
	}
	if query := tlq.withUser; query != nil {
		if err := tlq.loadUser(ctx, query, nodes, nil,
			func(n *TweetLike, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tlq *TweetLikeQuery) loadTweet(ctx context.Context, query *TweetQuery, nodes []*TweetLike, init func(*TweetLike), assign func(*TweetLike, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetLike)
	for i := range nodes {
		fk := nodes[i].TweetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tweet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tlq *TweetLikeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*TweetLike, init func(*TweetLike), assign func(*TweetLike, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetLike)
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

func (tlq *TweetLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tlq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, tlq.driver, _spec)
}

func (tlq *TweetLikeQuery) sqlExist(ctx context.Context) (bool, error) {
	_, err := tlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return !IsNotFound(err), nil
}

func (tlq *TweetLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweetlike.Table,
			Columns: tweetlike.Columns,
		},
		From:   tlq.sql,
		Unique: true,
	}
	if unique := tlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := tlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tlq *TweetLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tlq.driver.Dialect())
	t1 := builder.Table(tweetlike.Table)
	columns := tlq.fields
	if len(columns) == 0 {
		columns = tweetlike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tlq.sql != nil {
		selector = tlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tlq.unique != nil && *tlq.unique {
		selector.Distinct()
	}
	for _, p := range tlq.predicates {
		p(selector)
	}
	for _, p := range tlq.order {
		p(selector)
	}
	if offset := tlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TweetLikeGroupBy is the group-by builder for TweetLike entities.
type TweetLikeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlgb *TweetLikeGroupBy) Aggregate(fns ...AggregateFunc) *TweetLikeGroupBy {
	tlgb.fns = append(tlgb.fns, fns...)
	return tlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tlgb *TweetLikeGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tlgb.path(ctx)
	if err != nil {
		return err
	}
	tlgb.sql = query
	return tlgb.sqlScan(ctx, v)
}

func (tlgb *TweetLikeGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tlgb.fields {
		if !tweetlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tlgb *TweetLikeGroupBy) sqlQuery() *sql.Selector {
	selector := tlgb.sql.Select()
	aggregation := make([]string, 0, len(tlgb.fns))
	for _, fn := range tlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tlgb.fields)+len(tlgb.fns))
		for _, f := range tlgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tlgb.fields...)...)
}

// TweetLikeSelect is the builder for selecting fields of TweetLike entities.
type TweetLikeSelect struct {
	*TweetLikeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tls *TweetLikeSelect) Scan(ctx context.Context, v any) error {
	if err := tls.prepareQuery(ctx); err != nil {
		return err
	}
	tls.sql = tls.TweetLikeQuery.sqlQuery(ctx)
	return tls.sqlScan(ctx, v)
}

func (tls *TweetLikeSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := tls.sql.Query()
	if err := tls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
