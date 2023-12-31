// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/class_lab_status"
	"plms-backend/ent/course"
	"plms-backend/ent/lab"
	"plms-backend/ent/lab_problem"
	"plms-backend/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LabUpdate is the builder for updating Lab entities.
type LabUpdate struct {
	config
	hooks    []Hook
	mutation *LabMutation
}

// Where appends a list predicates to the LabUpdate builder.
func (lu *LabUpdate) Where(ps ...predicate.Lab) *LabUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetCourseID sets the "course_id" field.
func (lu *LabUpdate) SetCourseID(i int) *LabUpdate {
	lu.mutation.SetCourseID(i)
	return lu
}

// SetName sets the "name" field.
func (lu *LabUpdate) SetName(s string) *LabUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LabUpdate) SetCreatedAt(t time.Time) *LabUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LabUpdate) SetNillableCreatedAt(t *time.Time) *LabUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LabUpdate) SetUpdatedAt(t time.Time) *LabUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lu *LabUpdate) SetNillableUpdatedAt(t *time.Time) *LabUpdate {
	if t != nil {
		lu.SetUpdatedAt(*t)
	}
	return lu
}

// SetCourse sets the "course" edge to the Course entity.
func (lu *LabUpdate) SetCourse(c *Course) *LabUpdate {
	return lu.SetCourseID(c.ID)
}

// AddLabProblemIDs adds the "lab_problems" edge to the Lab_Problem entity by IDs.
func (lu *LabUpdate) AddLabProblemIDs(ids ...int) *LabUpdate {
	lu.mutation.AddLabProblemIDs(ids...)
	return lu
}

// AddLabProblems adds the "lab_problems" edges to the Lab_Problem entity.
func (lu *LabUpdate) AddLabProblems(l ...*Lab_Problem) *LabUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.AddLabProblemIDs(ids...)
}

// AddLabStatusIDs adds the "lab_statuses" edge to the Class_Lab_Status entity by IDs.
func (lu *LabUpdate) AddLabStatusIDs(ids ...int) *LabUpdate {
	lu.mutation.AddLabStatusIDs(ids...)
	return lu
}

// AddLabStatuses adds the "lab_statuses" edges to the Class_Lab_Status entity.
func (lu *LabUpdate) AddLabStatuses(c ...*Class_Lab_Status) *LabUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.AddLabStatusIDs(ids...)
}

// Mutation returns the LabMutation object of the builder.
func (lu *LabUpdate) Mutation() *LabMutation {
	return lu.mutation
}

// ClearCourse clears the "course" edge to the Course entity.
func (lu *LabUpdate) ClearCourse() *LabUpdate {
	lu.mutation.ClearCourse()
	return lu
}

// ClearLabProblems clears all "lab_problems" edges to the Lab_Problem entity.
func (lu *LabUpdate) ClearLabProblems() *LabUpdate {
	lu.mutation.ClearLabProblems()
	return lu
}

// RemoveLabProblemIDs removes the "lab_problems" edge to Lab_Problem entities by IDs.
func (lu *LabUpdate) RemoveLabProblemIDs(ids ...int) *LabUpdate {
	lu.mutation.RemoveLabProblemIDs(ids...)
	return lu
}

// RemoveLabProblems removes "lab_problems" edges to Lab_Problem entities.
func (lu *LabUpdate) RemoveLabProblems(l ...*Lab_Problem) *LabUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.RemoveLabProblemIDs(ids...)
}

// ClearLabStatuses clears all "lab_statuses" edges to the Class_Lab_Status entity.
func (lu *LabUpdate) ClearLabStatuses() *LabUpdate {
	lu.mutation.ClearLabStatuses()
	return lu
}

// RemoveLabStatusIDs removes the "lab_statuses" edge to Class_Lab_Status entities by IDs.
func (lu *LabUpdate) RemoveLabStatusIDs(ids ...int) *LabUpdate {
	lu.mutation.RemoveLabStatusIDs(ids...)
	return lu
}

// RemoveLabStatuses removes "lab_statuses" edges to Class_Lab_Status entities.
func (lu *LabUpdate) RemoveLabStatuses(c ...*Class_Lab_Status) *LabUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.RemoveLabStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LabUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LabUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LabUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LabUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LabUpdate) check() error {
	if v, ok := lu.mutation.Name(); ok {
		if err := lab.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Lab.name": %w`, err)}
		}
	}
	if _, ok := lu.mutation.CourseID(); lu.mutation.CourseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Lab.course"`)
	}
	return nil
}

func (lu *LabUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lab.Table, lab.Columns, sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(lab.FieldName, field.TypeString, value)
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(lab.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(lab.FieldUpdatedAt, field.TypeTime, value)
	}
	if lu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab.CourseTable,
			Columns: []string{lab.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab.CourseTable,
			Columns: []string{lab.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.LabProblemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabProblemsTable,
			Columns: []string{lab.LabProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedLabProblemsIDs(); len(nodes) > 0 && !lu.mutation.LabProblemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabProblemsTable,
			Columns: []string{lab.LabProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.LabProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabProblemsTable,
			Columns: []string{lab.LabProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.LabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabStatusesTable,
			Columns: []string{lab.LabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class_lab_status.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedLabStatusesIDs(); len(nodes) > 0 && !lu.mutation.LabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabStatusesTable,
			Columns: []string{lab.LabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class_lab_status.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.LabStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabStatusesTable,
			Columns: []string{lab.LabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class_lab_status.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lab.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LabUpdateOne is the builder for updating a single Lab entity.
type LabUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabMutation
}

// SetCourseID sets the "course_id" field.
func (luo *LabUpdateOne) SetCourseID(i int) *LabUpdateOne {
	luo.mutation.SetCourseID(i)
	return luo
}

// SetName sets the "name" field.
func (luo *LabUpdateOne) SetName(s string) *LabUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LabUpdateOne) SetCreatedAt(t time.Time) *LabUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LabUpdateOne) SetNillableCreatedAt(t *time.Time) *LabUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LabUpdateOne) SetUpdatedAt(t time.Time) *LabUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (luo *LabUpdateOne) SetNillableUpdatedAt(t *time.Time) *LabUpdateOne {
	if t != nil {
		luo.SetUpdatedAt(*t)
	}
	return luo
}

// SetCourse sets the "course" edge to the Course entity.
func (luo *LabUpdateOne) SetCourse(c *Course) *LabUpdateOne {
	return luo.SetCourseID(c.ID)
}

// AddLabProblemIDs adds the "lab_problems" edge to the Lab_Problem entity by IDs.
func (luo *LabUpdateOne) AddLabProblemIDs(ids ...int) *LabUpdateOne {
	luo.mutation.AddLabProblemIDs(ids...)
	return luo
}

// AddLabProblems adds the "lab_problems" edges to the Lab_Problem entity.
func (luo *LabUpdateOne) AddLabProblems(l ...*Lab_Problem) *LabUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.AddLabProblemIDs(ids...)
}

// AddLabStatusIDs adds the "lab_statuses" edge to the Class_Lab_Status entity by IDs.
func (luo *LabUpdateOne) AddLabStatusIDs(ids ...int) *LabUpdateOne {
	luo.mutation.AddLabStatusIDs(ids...)
	return luo
}

// AddLabStatuses adds the "lab_statuses" edges to the Class_Lab_Status entity.
func (luo *LabUpdateOne) AddLabStatuses(c ...*Class_Lab_Status) *LabUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.AddLabStatusIDs(ids...)
}

// Mutation returns the LabMutation object of the builder.
func (luo *LabUpdateOne) Mutation() *LabMutation {
	return luo.mutation
}

// ClearCourse clears the "course" edge to the Course entity.
func (luo *LabUpdateOne) ClearCourse() *LabUpdateOne {
	luo.mutation.ClearCourse()
	return luo
}

// ClearLabProblems clears all "lab_problems" edges to the Lab_Problem entity.
func (luo *LabUpdateOne) ClearLabProblems() *LabUpdateOne {
	luo.mutation.ClearLabProblems()
	return luo
}

// RemoveLabProblemIDs removes the "lab_problems" edge to Lab_Problem entities by IDs.
func (luo *LabUpdateOne) RemoveLabProblemIDs(ids ...int) *LabUpdateOne {
	luo.mutation.RemoveLabProblemIDs(ids...)
	return luo
}

// RemoveLabProblems removes "lab_problems" edges to Lab_Problem entities.
func (luo *LabUpdateOne) RemoveLabProblems(l ...*Lab_Problem) *LabUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.RemoveLabProblemIDs(ids...)
}

// ClearLabStatuses clears all "lab_statuses" edges to the Class_Lab_Status entity.
func (luo *LabUpdateOne) ClearLabStatuses() *LabUpdateOne {
	luo.mutation.ClearLabStatuses()
	return luo
}

// RemoveLabStatusIDs removes the "lab_statuses" edge to Class_Lab_Status entities by IDs.
func (luo *LabUpdateOne) RemoveLabStatusIDs(ids ...int) *LabUpdateOne {
	luo.mutation.RemoveLabStatusIDs(ids...)
	return luo
}

// RemoveLabStatuses removes "lab_statuses" edges to Class_Lab_Status entities.
func (luo *LabUpdateOne) RemoveLabStatuses(c ...*Class_Lab_Status) *LabUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.RemoveLabStatusIDs(ids...)
}

// Where appends a list predicates to the LabUpdate builder.
func (luo *LabUpdateOne) Where(ps ...predicate.Lab) *LabUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LabUpdateOne) Select(field string, fields ...string) *LabUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Lab entity.
func (luo *LabUpdateOne) Save(ctx context.Context) (*Lab, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LabUpdateOne) SaveX(ctx context.Context) *Lab {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LabUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LabUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LabUpdateOne) check() error {
	if v, ok := luo.mutation.Name(); ok {
		if err := lab.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Lab.name": %w`, err)}
		}
	}
	if _, ok := luo.mutation.CourseID(); luo.mutation.CourseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Lab.course"`)
	}
	return nil
}

func (luo *LabUpdateOne) sqlSave(ctx context.Context) (_node *Lab, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lab.Table, lab.Columns, sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Lab.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lab.FieldID)
		for _, f := range fields {
			if !lab.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lab.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(lab.FieldName, field.TypeString, value)
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(lab.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(lab.FieldUpdatedAt, field.TypeTime, value)
	}
	if luo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab.CourseTable,
			Columns: []string{lab.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lab.CourseTable,
			Columns: []string{lab.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.LabProblemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabProblemsTable,
			Columns: []string{lab.LabProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedLabProblemsIDs(); len(nodes) > 0 && !luo.mutation.LabProblemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabProblemsTable,
			Columns: []string{lab.LabProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.LabProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabProblemsTable,
			Columns: []string{lab.LabProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.LabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabStatusesTable,
			Columns: []string{lab.LabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class_lab_status.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedLabStatusesIDs(); len(nodes) > 0 && !luo.mutation.LabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabStatusesTable,
			Columns: []string{lab.LabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class_lab_status.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.LabStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lab.LabStatusesTable,
			Columns: []string{lab.LabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class_lab_status.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Lab{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lab.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
