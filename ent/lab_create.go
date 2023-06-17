// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/class_lab_status"
	"plms-backend/ent/course"
	"plms-backend/ent/lab"
	"plms-backend/ent/lab_problem"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LabCreate is the builder for creating a Lab entity.
type LabCreate struct {
	config
	mutation *LabMutation
	hooks    []Hook
}

// SetCourseID sets the "course_id" field.
func (lc *LabCreate) SetCourseID(i int) *LabCreate {
	lc.mutation.SetCourseID(i)
	return lc
}

// SetName sets the "name" field.
func (lc *LabCreate) SetName(s string) *LabCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LabCreate) SetCreatedAt(t time.Time) *LabCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LabCreate) SetNillableCreatedAt(t *time.Time) *LabCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LabCreate) SetUpdatedAt(t time.Time) *LabCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LabCreate) SetNillableUpdatedAt(t *time.Time) *LabCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetCourse sets the "course" edge to the Course entity.
func (lc *LabCreate) SetCourse(c *Course) *LabCreate {
	return lc.SetCourseID(c.ID)
}

// AddLabProblemIDs adds the "lab_problems" edge to the Lab_Problem entity by IDs.
func (lc *LabCreate) AddLabProblemIDs(ids ...int) *LabCreate {
	lc.mutation.AddLabProblemIDs(ids...)
	return lc
}

// AddLabProblems adds the "lab_problems" edges to the Lab_Problem entity.
func (lc *LabCreate) AddLabProblems(l ...*Lab_Problem) *LabCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lc.AddLabProblemIDs(ids...)
}

// AddLabStatusIDs adds the "lab_statuses" edge to the Class_Lab_Status entity by IDs.
func (lc *LabCreate) AddLabStatusIDs(ids ...int) *LabCreate {
	lc.mutation.AddLabStatusIDs(ids...)
	return lc
}

// AddLabStatuses adds the "lab_statuses" edges to the Class_Lab_Status entity.
func (lc *LabCreate) AddLabStatuses(c ...*Class_Lab_Status) *LabCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lc.AddLabStatusIDs(ids...)
}

// Mutation returns the LabMutation object of the builder.
func (lc *LabCreate) Mutation() *LabMutation {
	return lc.mutation
}

// Save creates the Lab in the database.
func (lc *LabCreate) Save(ctx context.Context) (*Lab, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LabCreate) SaveX(ctx context.Context) *Lab {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LabCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LabCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LabCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := lab.DefaultCreatedAt
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := lab.DefaultUpdatedAt
		lc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LabCreate) check() error {
	if _, ok := lc.mutation.CourseID(); !ok {
		return &ValidationError{Name: "course_id", err: errors.New(`ent: missing required field "Lab.course_id"`)}
	}
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Lab.name"`)}
	}
	if v, ok := lc.mutation.Name(); ok {
		if err := lab.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Lab.name": %w`, err)}
		}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Lab.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Lab.updated_at"`)}
	}
	if _, ok := lc.mutation.CourseID(); !ok {
		return &ValidationError{Name: "course", err: errors.New(`ent: missing required edge "Lab.course"`)}
	}
	return nil
}

func (lc *LabCreate) sqlSave(ctx context.Context) (*Lab, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LabCreate) createSpec() (*Lab, *sqlgraph.CreateSpec) {
	var (
		_node = &Lab{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(lab.Table, sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.Name(); ok {
		_spec.SetField(lab.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(lab.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(lab.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := lc.mutation.CourseIDs(); len(nodes) > 0 {
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
		_node.CourseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.LabProblemsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.LabStatusesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LabCreateBulk is the builder for creating many Lab entities in bulk.
type LabCreateBulk struct {
	config
	builders []*LabCreate
}

// Save creates the Lab entities in the database.
func (lcb *LabCreateBulk) Save(ctx context.Context) ([]*Lab, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lab, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LabMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LabCreateBulk) SaveX(ctx context.Context) []*Lab {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LabCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LabCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
