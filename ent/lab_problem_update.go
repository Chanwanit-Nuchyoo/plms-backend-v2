// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/lab"
	"plms-backend/ent/lab_problem"
	"plms-backend/ent/predicate"
	"plms-backend/ent/testcase"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LabProblemUpdate is the builder for updating Lab_Problem entities.
type LabProblemUpdate struct {
	config
	hooks    []Hook
	mutation *LabProblemMutation
}

// Where appends a list predicates to the LabProblemUpdate builder.
func (lpu *LabProblemUpdate) Where(ps ...predicate.Lab_Problem) *LabProblemUpdate {
	lpu.mutation.Where(ps...)
	return lpu
}

// SetLabID sets the "lab_id" field.
func (lpu *LabProblemUpdate) SetLabID(i int) *LabProblemUpdate {
	lpu.mutation.SetLabID(i)
	return lpu
}

// SetName sets the "name" field.
func (lpu *LabProblemUpdate) SetName(s string) *LabProblemUpdate {
	lpu.mutation.SetName(s)
	return lpu
}

// SetPrompt sets the "prompt" field.
func (lpu *LabProblemUpdate) SetPrompt(s string) *LabProblemUpdate {
	lpu.mutation.SetPrompt(s)
	return lpu
}

// SetFullScore sets the "full_score" field.
func (lpu *LabProblemUpdate) SetFullScore(f float64) *LabProblemUpdate {
	lpu.mutation.ResetFullScore()
	lpu.mutation.SetFullScore(f)
	return lpu
}

// SetNillableFullScore sets the "full_score" field if the given value is not nil.
func (lpu *LabProblemUpdate) SetNillableFullScore(f *float64) *LabProblemUpdate {
	if f != nil {
		lpu.SetFullScore(*f)
	}
	return lpu
}

// AddFullScore adds f to the "full_score" field.
func (lpu *LabProblemUpdate) AddFullScore(f float64) *LabProblemUpdate {
	lpu.mutation.AddFullScore(f)
	return lpu
}

// ClearFullScore clears the value of the "full_score" field.
func (lpu *LabProblemUpdate) ClearFullScore() *LabProblemUpdate {
	lpu.mutation.ClearFullScore()
	return lpu
}

// SetCreatedAt sets the "created_at" field.
func (lpu *LabProblemUpdate) SetCreatedAt(t time.Time) *LabProblemUpdate {
	lpu.mutation.SetCreatedAt(t)
	return lpu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lpu *LabProblemUpdate) SetNillableCreatedAt(t *time.Time) *LabProblemUpdate {
	if t != nil {
		lpu.SetCreatedAt(*t)
	}
	return lpu
}

// SetUpdatedAt sets the "updated_at" field.
func (lpu *LabProblemUpdate) SetUpdatedAt(t time.Time) *LabProblemUpdate {
	lpu.mutation.SetUpdatedAt(t)
	return lpu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lpu *LabProblemUpdate) SetNillableUpdatedAt(t *time.Time) *LabProblemUpdate {
	if t != nil {
		lpu.SetUpdatedAt(*t)
	}
	return lpu
}

// SetLab sets the "lab" edge to the Lab entity.
func (lpu *LabProblemUpdate) SetLab(l *Lab) *LabProblemUpdate {
	return lpu.SetLabID(l.ID)
}

// AddTestcaseIDs adds the "testcases" edge to the Testcase entity by IDs.
func (lpu *LabProblemUpdate) AddTestcaseIDs(ids ...int) *LabProblemUpdate {
	lpu.mutation.AddTestcaseIDs(ids...)
	return lpu
}

// AddTestcases adds the "testcases" edges to the Testcase entity.
func (lpu *LabProblemUpdate) AddTestcases(t ...*Testcase) *LabProblemUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return lpu.AddTestcaseIDs(ids...)
}

// Mutation returns the LabProblemMutation object of the builder.
func (lpu *LabProblemUpdate) Mutation() *LabProblemMutation {
	return lpu.mutation
}

// ClearLab clears the "lab" edge to the Lab entity.
func (lpu *LabProblemUpdate) ClearLab() *LabProblemUpdate {
	lpu.mutation.ClearLab()
	return lpu
}

// ClearTestcases clears all "testcases" edges to the Testcase entity.
func (lpu *LabProblemUpdate) ClearTestcases() *LabProblemUpdate {
	lpu.mutation.ClearTestcases()
	return lpu
}

// RemoveTestcaseIDs removes the "testcases" edge to Testcase entities by IDs.
func (lpu *LabProblemUpdate) RemoveTestcaseIDs(ids ...int) *LabProblemUpdate {
	lpu.mutation.RemoveTestcaseIDs(ids...)
	return lpu
}

// RemoveTestcases removes "testcases" edges to Testcase entities.
func (lpu *LabProblemUpdate) RemoveTestcases(t ...*Testcase) *LabProblemUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return lpu.RemoveTestcaseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lpu *LabProblemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lpu.sqlSave, lpu.mutation, lpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lpu *LabProblemUpdate) SaveX(ctx context.Context) int {
	affected, err := lpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lpu *LabProblemUpdate) Exec(ctx context.Context) error {
	_, err := lpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpu *LabProblemUpdate) ExecX(ctx context.Context) {
	if err := lpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpu *LabProblemUpdate) check() error {
	if v, ok := lpu.mutation.Name(); ok {
		if err := lab_problem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Lab_Problem.name": %w`, err)}
		}
	}
	if v, ok := lpu.mutation.Prompt(); ok {
		if err := lab_problem.PromptValidator(v); err != nil {
			return &ValidationError{Name: "prompt", err: fmt.Errorf(`ent: validator failed for field "Lab_Problem.prompt": %w`, err)}
		}
	}
	if v, ok := lpu.mutation.FullScore(); ok {
		if err := lab_problem.FullScoreValidator(v); err != nil {
			return &ValidationError{Name: "full_score", err: fmt.Errorf(`ent: validator failed for field "Lab_Problem.full_score": %w`, err)}
		}
	}
	if _, ok := lpu.mutation.LabID(); lpu.mutation.LabCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Lab_Problem.lab"`)
	}
	return nil
}

func (lpu *LabProblemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lab_problem.Table, lab_problem.Columns, sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt))
	if ps := lpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lpu.mutation.Name(); ok {
		_spec.SetField(lab_problem.FieldName, field.TypeString, value)
	}
	if value, ok := lpu.mutation.Prompt(); ok {
		_spec.SetField(lab_problem.FieldPrompt, field.TypeString, value)
	}
	if value, ok := lpu.mutation.FullScore(); ok {
		_spec.SetField(lab_problem.FieldFullScore, field.TypeFloat64, value)
	}
	if value, ok := lpu.mutation.AddedFullScore(); ok {
		_spec.AddField(lab_problem.FieldFullScore, field.TypeFloat64, value)
	}
	if lpu.mutation.FullScoreCleared() {
		_spec.ClearField(lab_problem.FieldFullScore, field.TypeFloat64)
	}
	if value, ok := lpu.mutation.CreatedAt(); ok {
		_spec.SetField(lab_problem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lpu.mutation.UpdatedAt(); ok {
		_spec.SetField(lab_problem.FieldUpdatedAt, field.TypeTime, value)
	}
	if lpu.mutation.LabCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab_problem.LabTable,
			Columns: []string{lab_problem.LabColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpu.mutation.LabIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab_problem.LabTable,
			Columns: []string{lab_problem.LabColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpu.mutation.TestcasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab_problem.TestcasesTable,
			Columns: []string{lab_problem.TestcasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpu.mutation.RemovedTestcasesIDs(); len(nodes) > 0 && !lpu.mutation.TestcasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab_problem.TestcasesTable,
			Columns: []string{lab_problem.TestcasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpu.mutation.TestcasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab_problem.TestcasesTable,
			Columns: []string{lab_problem.TestcasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lab_problem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lpu.mutation.done = true
	return n, nil
}

// LabProblemUpdateOne is the builder for updating a single Lab_Problem entity.
type LabProblemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabProblemMutation
}

// SetLabID sets the "lab_id" field.
func (lpuo *LabProblemUpdateOne) SetLabID(i int) *LabProblemUpdateOne {
	lpuo.mutation.SetLabID(i)
	return lpuo
}

// SetName sets the "name" field.
func (lpuo *LabProblemUpdateOne) SetName(s string) *LabProblemUpdateOne {
	lpuo.mutation.SetName(s)
	return lpuo
}

// SetPrompt sets the "prompt" field.
func (lpuo *LabProblemUpdateOne) SetPrompt(s string) *LabProblemUpdateOne {
	lpuo.mutation.SetPrompt(s)
	return lpuo
}

// SetFullScore sets the "full_score" field.
func (lpuo *LabProblemUpdateOne) SetFullScore(f float64) *LabProblemUpdateOne {
	lpuo.mutation.ResetFullScore()
	lpuo.mutation.SetFullScore(f)
	return lpuo
}

// SetNillableFullScore sets the "full_score" field if the given value is not nil.
func (lpuo *LabProblemUpdateOne) SetNillableFullScore(f *float64) *LabProblemUpdateOne {
	if f != nil {
		lpuo.SetFullScore(*f)
	}
	return lpuo
}

// AddFullScore adds f to the "full_score" field.
func (lpuo *LabProblemUpdateOne) AddFullScore(f float64) *LabProblemUpdateOne {
	lpuo.mutation.AddFullScore(f)
	return lpuo
}

// ClearFullScore clears the value of the "full_score" field.
func (lpuo *LabProblemUpdateOne) ClearFullScore() *LabProblemUpdateOne {
	lpuo.mutation.ClearFullScore()
	return lpuo
}

// SetCreatedAt sets the "created_at" field.
func (lpuo *LabProblemUpdateOne) SetCreatedAt(t time.Time) *LabProblemUpdateOne {
	lpuo.mutation.SetCreatedAt(t)
	return lpuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lpuo *LabProblemUpdateOne) SetNillableCreatedAt(t *time.Time) *LabProblemUpdateOne {
	if t != nil {
		lpuo.SetCreatedAt(*t)
	}
	return lpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (lpuo *LabProblemUpdateOne) SetUpdatedAt(t time.Time) *LabProblemUpdateOne {
	lpuo.mutation.SetUpdatedAt(t)
	return lpuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lpuo *LabProblemUpdateOne) SetNillableUpdatedAt(t *time.Time) *LabProblemUpdateOne {
	if t != nil {
		lpuo.SetUpdatedAt(*t)
	}
	return lpuo
}

// SetLab sets the "lab" edge to the Lab entity.
func (lpuo *LabProblemUpdateOne) SetLab(l *Lab) *LabProblemUpdateOne {
	return lpuo.SetLabID(l.ID)
}

// AddTestcaseIDs adds the "testcases" edge to the Testcase entity by IDs.
func (lpuo *LabProblemUpdateOne) AddTestcaseIDs(ids ...int) *LabProblemUpdateOne {
	lpuo.mutation.AddTestcaseIDs(ids...)
	return lpuo
}

// AddTestcases adds the "testcases" edges to the Testcase entity.
func (lpuo *LabProblemUpdateOne) AddTestcases(t ...*Testcase) *LabProblemUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return lpuo.AddTestcaseIDs(ids...)
}

// Mutation returns the LabProblemMutation object of the builder.
func (lpuo *LabProblemUpdateOne) Mutation() *LabProblemMutation {
	return lpuo.mutation
}

// ClearLab clears the "lab" edge to the Lab entity.
func (lpuo *LabProblemUpdateOne) ClearLab() *LabProblemUpdateOne {
	lpuo.mutation.ClearLab()
	return lpuo
}

// ClearTestcases clears all "testcases" edges to the Testcase entity.
func (lpuo *LabProblemUpdateOne) ClearTestcases() *LabProblemUpdateOne {
	lpuo.mutation.ClearTestcases()
	return lpuo
}

// RemoveTestcaseIDs removes the "testcases" edge to Testcase entities by IDs.
func (lpuo *LabProblemUpdateOne) RemoveTestcaseIDs(ids ...int) *LabProblemUpdateOne {
	lpuo.mutation.RemoveTestcaseIDs(ids...)
	return lpuo
}

// RemoveTestcases removes "testcases" edges to Testcase entities.
func (lpuo *LabProblemUpdateOne) RemoveTestcases(t ...*Testcase) *LabProblemUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return lpuo.RemoveTestcaseIDs(ids...)
}

// Where appends a list predicates to the LabProblemUpdate builder.
func (lpuo *LabProblemUpdateOne) Where(ps ...predicate.Lab_Problem) *LabProblemUpdateOne {
	lpuo.mutation.Where(ps...)
	return lpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lpuo *LabProblemUpdateOne) Select(field string, fields ...string) *LabProblemUpdateOne {
	lpuo.fields = append([]string{field}, fields...)
	return lpuo
}

// Save executes the query and returns the updated Lab_Problem entity.
func (lpuo *LabProblemUpdateOne) Save(ctx context.Context) (*Lab_Problem, error) {
	return withHooks(ctx, lpuo.sqlSave, lpuo.mutation, lpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lpuo *LabProblemUpdateOne) SaveX(ctx context.Context) *Lab_Problem {
	node, err := lpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lpuo *LabProblemUpdateOne) Exec(ctx context.Context) error {
	_, err := lpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpuo *LabProblemUpdateOne) ExecX(ctx context.Context) {
	if err := lpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpuo *LabProblemUpdateOne) check() error {
	if v, ok := lpuo.mutation.Name(); ok {
		if err := lab_problem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Lab_Problem.name": %w`, err)}
		}
	}
	if v, ok := lpuo.mutation.Prompt(); ok {
		if err := lab_problem.PromptValidator(v); err != nil {
			return &ValidationError{Name: "prompt", err: fmt.Errorf(`ent: validator failed for field "Lab_Problem.prompt": %w`, err)}
		}
	}
	if v, ok := lpuo.mutation.FullScore(); ok {
		if err := lab_problem.FullScoreValidator(v); err != nil {
			return &ValidationError{Name: "full_score", err: fmt.Errorf(`ent: validator failed for field "Lab_Problem.full_score": %w`, err)}
		}
	}
	if _, ok := lpuo.mutation.LabID(); lpuo.mutation.LabCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Lab_Problem.lab"`)
	}
	return nil
}

func (lpuo *LabProblemUpdateOne) sqlSave(ctx context.Context) (_node *Lab_Problem, err error) {
	if err := lpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lab_problem.Table, lab_problem.Columns, sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt))
	id, ok := lpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Lab_Problem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lab_problem.FieldID)
		for _, f := range fields {
			if !lab_problem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lab_problem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lpuo.mutation.Name(); ok {
		_spec.SetField(lab_problem.FieldName, field.TypeString, value)
	}
	if value, ok := lpuo.mutation.Prompt(); ok {
		_spec.SetField(lab_problem.FieldPrompt, field.TypeString, value)
	}
	if value, ok := lpuo.mutation.FullScore(); ok {
		_spec.SetField(lab_problem.FieldFullScore, field.TypeFloat64, value)
	}
	if value, ok := lpuo.mutation.AddedFullScore(); ok {
		_spec.AddField(lab_problem.FieldFullScore, field.TypeFloat64, value)
	}
	if lpuo.mutation.FullScoreCleared() {
		_spec.ClearField(lab_problem.FieldFullScore, field.TypeFloat64)
	}
	if value, ok := lpuo.mutation.CreatedAt(); ok {
		_spec.SetField(lab_problem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(lab_problem.FieldUpdatedAt, field.TypeTime, value)
	}
	if lpuo.mutation.LabCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab_problem.LabTable,
			Columns: []string{lab_problem.LabColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpuo.mutation.LabIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab_problem.LabTable,
			Columns: []string{lab_problem.LabColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpuo.mutation.TestcasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab_problem.TestcasesTable,
			Columns: []string{lab_problem.TestcasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpuo.mutation.RemovedTestcasesIDs(); len(nodes) > 0 && !lpuo.mutation.TestcasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab_problem.TestcasesTable,
			Columns: []string{lab_problem.TestcasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpuo.mutation.TestcasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab_problem.TestcasesTable,
			Columns: []string{lab_problem.TestcasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Lab_Problem{config: lpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lab_problem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lpuo.mutation.done = true
	return _node, nil
}
