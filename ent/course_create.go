// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/class"
	"plms-backend/ent/class_lab_status"
	"plms-backend/ent/course"
	"plms-backend/ent/lab"
	"plms-backend/ent/topic"
	"plms-backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseCreate is the builder for creating a Course entity.
type CourseCreate struct {
	config
	mutation *CourseMutation
	hooks    []Hook
}

// SetOwnerID sets the "owner_id" field.
func (cc *CourseCreate) SetOwnerID(i int) *CourseCreate {
	cc.mutation.SetOwnerID(i)
	return cc
}

// SetName sets the "name" field.
func (cc *CourseCreate) SetName(s string) *CourseCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CourseCreate) SetDescription(s string) *CourseCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CourseCreate) SetNillableDescription(s *string) *CourseCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CourseCreate) SetCreatedAt(t time.Time) *CourseCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CourseCreate) SetNillableCreatedAt(t *time.Time) *CourseCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CourseCreate) SetUpdatedAt(t time.Time) *CourseCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CourseCreate) SetNillableUpdatedAt(t *time.Time) *CourseCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CourseCreate) SetOwner(u *User) *CourseCreate {
	return cc.SetOwnerID(u.ID)
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (cc *CourseCreate) AddClassIDs(ids ...int) *CourseCreate {
	cc.mutation.AddClassIDs(ids...)
	return cc
}

// AddClasses adds the "classes" edges to the Class entity.
func (cc *CourseCreate) AddClasses(c ...*Class) *CourseCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddClassIDs(ids...)
}

// AddLabIDs adds the "labs" edge to the Lab entity by IDs.
func (cc *CourseCreate) AddLabIDs(ids ...int) *CourseCreate {
	cc.mutation.AddLabIDs(ids...)
	return cc
}

// AddLabs adds the "labs" edges to the Lab entity.
func (cc *CourseCreate) AddLabs(l ...*Lab) *CourseCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cc.AddLabIDs(ids...)
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (cc *CourseCreate) AddTopicIDs(ids ...int) *CourseCreate {
	cc.mutation.AddTopicIDs(ids...)
	return cc
}

// AddTopics adds the "topics" edges to the Topic entity.
func (cc *CourseCreate) AddTopics(t ...*Topic) *CourseCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTopicIDs(ids...)
}

// AddCourseLabStatusIDs adds the "course_lab_statuses" edge to the Class_Lab_Status entity by IDs.
func (cc *CourseCreate) AddCourseLabStatusIDs(ids ...int) *CourseCreate {
	cc.mutation.AddCourseLabStatusIDs(ids...)
	return cc
}

// AddCourseLabStatuses adds the "course_lab_statuses" edges to the Class_Lab_Status entity.
func (cc *CourseCreate) AddCourseLabStatuses(c ...*Class_Lab_Status) *CourseCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCourseLabStatusIDs(ids...)
}

// Mutation returns the CourseMutation object of the builder.
func (cc *CourseCreate) Mutation() *CourseMutation {
	return cc.mutation
}

// Save creates the Course in the database.
func (cc *CourseCreate) Save(ctx context.Context) (*Course, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CourseCreate) SaveX(ctx context.Context) *Course {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CourseCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CourseCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CourseCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := course.DefaultCreatedAt
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := course.DefaultUpdatedAt
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CourseCreate) check() error {
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Course.owner_id"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Course.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := course.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Course.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Course.owner"`)}
	}
	return nil
}

func (cc *CourseCreate) sqlSave(ctx context.Context) (*Course, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CourseCreate) createSpec() (*Course, *sqlgraph.CreateSpec) {
	var (
		_node = &Course{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(course.Table, sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(course.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(course.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(course.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(course.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.OwnerTable,
			Columns: []string{course.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ClassesTable,
			Columns: []string{course.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.LabsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.LabsTable,
			Columns: []string{course.LabsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.TopicsTable,
			Columns: []string{course.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CourseLabStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.CourseLabStatusesTable,
			Columns: []string{course.CourseLabStatusesColumn},
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

// CourseCreateBulk is the builder for creating many Course entities in bulk.
type CourseCreateBulk struct {
	config
	builders []*CourseCreate
}

// Save creates the Course entities in the database.
func (ccb *CourseCreateBulk) Save(ctx context.Context) ([]*Course, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Course, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CourseMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CourseCreateBulk) SaveX(ctx context.Context) []*Course {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CourseCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CourseCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
