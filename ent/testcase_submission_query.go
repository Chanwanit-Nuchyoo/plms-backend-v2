// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/lab_problem_submission"
	"plms-backend/ent/predicate"
	"plms-backend/ent/testcase"
	"plms-backend/ent/testcase_submission"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestcaseSubmissionQuery is the builder for querying Testcase_Submission entities.
type TestcaseSubmissionQuery struct {
	config
	ctx            *QueryContext
	order          []testcase_submission.OrderOption
	inters         []Interceptor
	predicates     []predicate.Testcase_Submission
	withTestcase   *TestcaseQuery
	withSubmission *LabProblemSubmissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestcaseSubmissionQuery builder.
func (tsq *TestcaseSubmissionQuery) Where(ps ...predicate.Testcase_Submission) *TestcaseSubmissionQuery {
	tsq.predicates = append(tsq.predicates, ps...)
	return tsq
}

// Limit the number of records to be returned by this query.
func (tsq *TestcaseSubmissionQuery) Limit(limit int) *TestcaseSubmissionQuery {
	tsq.ctx.Limit = &limit
	return tsq
}

// Offset to start from.
func (tsq *TestcaseSubmissionQuery) Offset(offset int) *TestcaseSubmissionQuery {
	tsq.ctx.Offset = &offset
	return tsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsq *TestcaseSubmissionQuery) Unique(unique bool) *TestcaseSubmissionQuery {
	tsq.ctx.Unique = &unique
	return tsq
}

// Order specifies how the records should be ordered.
func (tsq *TestcaseSubmissionQuery) Order(o ...testcase_submission.OrderOption) *TestcaseSubmissionQuery {
	tsq.order = append(tsq.order, o...)
	return tsq
}

// QueryTestcase chains the current query on the "testcase" edge.
func (tsq *TestcaseSubmissionQuery) QueryTestcase() *TestcaseQuery {
	query := (&TestcaseClient{config: tsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testcase_submission.Table, testcase_submission.FieldID, selector),
			sqlgraph.To(testcase.Table, testcase.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, testcase_submission.TestcaseTable, testcase_submission.TestcaseColumn),
		)
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubmission chains the current query on the "submission" edge.
func (tsq *TestcaseSubmissionQuery) QuerySubmission() *LabProblemSubmissionQuery {
	query := (&LabProblemSubmissionClient{config: tsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testcase_submission.Table, testcase_submission.FieldID, selector),
			sqlgraph.To(lab_problem_submission.Table, lab_problem_submission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, testcase_submission.SubmissionTable, testcase_submission.SubmissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Testcase_Submission entity from the query.
// Returns a *NotFoundError when no Testcase_Submission was found.
func (tsq *TestcaseSubmissionQuery) First(ctx context.Context) (*Testcase_Submission, error) {
	nodes, err := tsq.Limit(1).All(setContextOp(ctx, tsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testcase_submission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) FirstX(ctx context.Context) *Testcase_Submission {
	node, err := tsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Testcase_Submission ID from the query.
// Returns a *NotFoundError when no Testcase_Submission ID was found.
func (tsq *TestcaseSubmissionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(1).IDs(setContextOp(ctx, tsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testcase_submission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) FirstIDX(ctx context.Context) int {
	id, err := tsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Testcase_Submission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Testcase_Submission entity is found.
// Returns a *NotFoundError when no Testcase_Submission entities are found.
func (tsq *TestcaseSubmissionQuery) Only(ctx context.Context) (*Testcase_Submission, error) {
	nodes, err := tsq.Limit(2).All(setContextOp(ctx, tsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testcase_submission.Label}
	default:
		return nil, &NotSingularError{testcase_submission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) OnlyX(ctx context.Context) *Testcase_Submission {
	node, err := tsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Testcase_Submission ID in the query.
// Returns a *NotSingularError when more than one Testcase_Submission ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsq *TestcaseSubmissionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(2).IDs(setContextOp(ctx, tsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testcase_submission.Label}
	default:
		err = &NotSingularError{testcase_submission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) OnlyIDX(ctx context.Context) int {
	id, err := tsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Testcase_Submissions.
func (tsq *TestcaseSubmissionQuery) All(ctx context.Context) ([]*Testcase_Submission, error) {
	ctx = setContextOp(ctx, tsq.ctx, "All")
	if err := tsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Testcase_Submission, *TestcaseSubmissionQuery]()
	return withInterceptors[[]*Testcase_Submission](ctx, tsq, qr, tsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) AllX(ctx context.Context) []*Testcase_Submission {
	nodes, err := tsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Testcase_Submission IDs.
func (tsq *TestcaseSubmissionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tsq.ctx.Unique == nil && tsq.path != nil {
		tsq.Unique(true)
	}
	ctx = setContextOp(ctx, tsq.ctx, "IDs")
	if err = tsq.Select(testcase_submission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) IDsX(ctx context.Context) []int {
	ids, err := tsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsq *TestcaseSubmissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Count")
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsq, querierCount[*TestcaseSubmissionQuery](), tsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) CountX(ctx context.Context) int {
	count, err := tsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsq *TestcaseSubmissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Exist")
	switch _, err := tsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsq *TestcaseSubmissionQuery) ExistX(ctx context.Context) bool {
	exist, err := tsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestcaseSubmissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsq *TestcaseSubmissionQuery) Clone() *TestcaseSubmissionQuery {
	if tsq == nil {
		return nil
	}
	return &TestcaseSubmissionQuery{
		config:         tsq.config,
		ctx:            tsq.ctx.Clone(),
		order:          append([]testcase_submission.OrderOption{}, tsq.order...),
		inters:         append([]Interceptor{}, tsq.inters...),
		predicates:     append([]predicate.Testcase_Submission{}, tsq.predicates...),
		withTestcase:   tsq.withTestcase.Clone(),
		withSubmission: tsq.withSubmission.Clone(),
		// clone intermediate query.
		sql:  tsq.sql.Clone(),
		path: tsq.path,
	}
}

// WithTestcase tells the query-builder to eager-load the nodes that are connected to
// the "testcase" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *TestcaseSubmissionQuery) WithTestcase(opts ...func(*TestcaseQuery)) *TestcaseSubmissionQuery {
	query := (&TestcaseClient{config: tsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tsq.withTestcase = query
	return tsq
}

// WithSubmission tells the query-builder to eager-load the nodes that are connected to
// the "submission" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *TestcaseSubmissionQuery) WithSubmission(opts ...func(*LabProblemSubmissionQuery)) *TestcaseSubmissionQuery {
	query := (&LabProblemSubmissionClient{config: tsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tsq.withSubmission = query
	return tsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SubmissionID int `json:"submission_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestcaseSubmission.Query().
//		GroupBy(testcase_submission.FieldSubmissionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tsq *TestcaseSubmissionQuery) GroupBy(field string, fields ...string) *TestcaseSubmissionGroupBy {
	tsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestcaseSubmissionGroupBy{build: tsq}
	grbuild.flds = &tsq.ctx.Fields
	grbuild.label = testcase_submission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SubmissionID int `json:"submission_id,omitempty"`
//	}
//
//	client.TestcaseSubmission.Query().
//		Select(testcase_submission.FieldSubmissionID).
//		Scan(ctx, &v)
func (tsq *TestcaseSubmissionQuery) Select(fields ...string) *TestcaseSubmissionSelect {
	tsq.ctx.Fields = append(tsq.ctx.Fields, fields...)
	sbuild := &TestcaseSubmissionSelect{TestcaseSubmissionQuery: tsq}
	sbuild.label = testcase_submission.Label
	sbuild.flds, sbuild.scan = &tsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestcaseSubmissionSelect configured with the given aggregations.
func (tsq *TestcaseSubmissionQuery) Aggregate(fns ...AggregateFunc) *TestcaseSubmissionSelect {
	return tsq.Select().Aggregate(fns...)
}

func (tsq *TestcaseSubmissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsq.ctx.Fields {
		if !testcase_submission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tsq.path != nil {
		prev, err := tsq.path(ctx)
		if err != nil {
			return err
		}
		tsq.sql = prev
	}
	return nil
}

func (tsq *TestcaseSubmissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Testcase_Submission, error) {
	var (
		nodes       = []*Testcase_Submission{}
		_spec       = tsq.querySpec()
		loadedTypes = [2]bool{
			tsq.withTestcase != nil,
			tsq.withSubmission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Testcase_Submission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Testcase_Submission{config: tsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tsq.withTestcase; query != nil {
		if err := tsq.loadTestcase(ctx, query, nodes, nil,
			func(n *Testcase_Submission, e *Testcase) { n.Edges.Testcase = e }); err != nil {
			return nil, err
		}
	}
	if query := tsq.withSubmission; query != nil {
		if err := tsq.loadSubmission(ctx, query, nodes, nil,
			func(n *Testcase_Submission, e *Lab_Problem_Submission) { n.Edges.Submission = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tsq *TestcaseSubmissionQuery) loadTestcase(ctx context.Context, query *TestcaseQuery, nodes []*Testcase_Submission, init func(*Testcase_Submission), assign func(*Testcase_Submission, *Testcase)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Testcase_Submission)
	for i := range nodes {
		fk := nodes[i].TestcaseID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(testcase.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "testcase_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tsq *TestcaseSubmissionQuery) loadSubmission(ctx context.Context, query *LabProblemSubmissionQuery, nodes []*Testcase_Submission, init func(*Testcase_Submission), assign func(*Testcase_Submission, *Lab_Problem_Submission)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Testcase_Submission)
	for i := range nodes {
		fk := nodes[i].SubmissionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(lab_problem_submission.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "submission_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tsq *TestcaseSubmissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsq.querySpec()
	_spec.Node.Columns = tsq.ctx.Fields
	if len(tsq.ctx.Fields) > 0 {
		_spec.Unique = tsq.ctx.Unique != nil && *tsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsq.driver, _spec)
}

func (tsq *TestcaseSubmissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testcase_submission.Table, testcase_submission.Columns, sqlgraph.NewFieldSpec(testcase_submission.FieldID, field.TypeInt))
	_spec.From = tsq.sql
	if unique := tsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsq.path != nil {
		_spec.Unique = true
	}
	if fields := tsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testcase_submission.FieldID)
		for i := range fields {
			if fields[i] != testcase_submission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tsq.withTestcase != nil {
			_spec.Node.AddColumnOnce(testcase_submission.FieldTestcaseID)
		}
		if tsq.withSubmission != nil {
			_spec.Node.AddColumnOnce(testcase_submission.FieldSubmissionID)
		}
	}
	if ps := tsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsq *TestcaseSubmissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsq.driver.Dialect())
	t1 := builder.Table(testcase_submission.Table)
	columns := tsq.ctx.Fields
	if len(columns) == 0 {
		columns = testcase_submission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsq.sql != nil {
		selector = tsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsq.ctx.Unique != nil && *tsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tsq.predicates {
		p(selector)
	}
	for _, p := range tsq.order {
		p(selector)
	}
	if offset := tsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestcaseSubmissionGroupBy is the group-by builder for Testcase_Submission entities.
type TestcaseSubmissionGroupBy struct {
	selector
	build *TestcaseSubmissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsgb *TestcaseSubmissionGroupBy) Aggregate(fns ...AggregateFunc) *TestcaseSubmissionGroupBy {
	tsgb.fns = append(tsgb.fns, fns...)
	return tsgb
}

// Scan applies the selector query and scans the result into the given value.
func (tsgb *TestcaseSubmissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsgb.build.ctx, "GroupBy")
	if err := tsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestcaseSubmissionQuery, *TestcaseSubmissionGroupBy](ctx, tsgb.build, tsgb, tsgb.build.inters, v)
}

func (tsgb *TestcaseSubmissionGroupBy) sqlScan(ctx context.Context, root *TestcaseSubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsgb.fns))
	for _, fn := range tsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsgb.flds)+len(tsgb.fns))
		for _, f := range *tsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TestcaseSubmissionSelect is the builder for selecting fields of TestcaseSubmission entities.
type TestcaseSubmissionSelect struct {
	*TestcaseSubmissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tss *TestcaseSubmissionSelect) Aggregate(fns ...AggregateFunc) *TestcaseSubmissionSelect {
	tss.fns = append(tss.fns, fns...)
	return tss
}

// Scan applies the selector query and scans the result into the given value.
func (tss *TestcaseSubmissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tss.ctx, "Select")
	if err := tss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestcaseSubmissionQuery, *TestcaseSubmissionSelect](ctx, tss.TestcaseSubmissionQuery, tss, tss.inters, v)
}

func (tss *TestcaseSubmissionSelect) sqlScan(ctx context.Context, root *TestcaseSubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tss.fns))
	for _, fn := range tss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
