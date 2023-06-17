// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/topic"
	"plms-backend/ent/topic_file_mats"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicFileMatsCreate is the builder for creating a Topic_File_Mats entity.
type TopicFileMatsCreate struct {
	config
	mutation *TopicFileMatsMutation
	hooks    []Hook
}

// SetTopicID sets the "topic_id" field.
func (tfmc *TopicFileMatsCreate) SetTopicID(i int) *TopicFileMatsCreate {
	tfmc.mutation.SetTopicID(i)
	return tfmc
}

// SetName sets the "name" field.
func (tfmc *TopicFileMatsCreate) SetName(s string) *TopicFileMatsCreate {
	tfmc.mutation.SetName(s)
	return tfmc
}

// SetFilePath sets the "file_path" field.
func (tfmc *TopicFileMatsCreate) SetFilePath(s string) *TopicFileMatsCreate {
	tfmc.mutation.SetFilePath(s)
	return tfmc
}

// SetCreatedAt sets the "created_at" field.
func (tfmc *TopicFileMatsCreate) SetCreatedAt(t time.Time) *TopicFileMatsCreate {
	tfmc.mutation.SetCreatedAt(t)
	return tfmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tfmc *TopicFileMatsCreate) SetNillableCreatedAt(t *time.Time) *TopicFileMatsCreate {
	if t != nil {
		tfmc.SetCreatedAt(*t)
	}
	return tfmc
}

// SetTopic sets the "topic" edge to the Topic entity.
func (tfmc *TopicFileMatsCreate) SetTopic(t *Topic) *TopicFileMatsCreate {
	return tfmc.SetTopicID(t.ID)
}

// Mutation returns the TopicFileMatsMutation object of the builder.
func (tfmc *TopicFileMatsCreate) Mutation() *TopicFileMatsMutation {
	return tfmc.mutation
}

// Save creates the Topic_File_Mats in the database.
func (tfmc *TopicFileMatsCreate) Save(ctx context.Context) (*Topic_File_Mats, error) {
	tfmc.defaults()
	return withHooks(ctx, tfmc.sqlSave, tfmc.mutation, tfmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tfmc *TopicFileMatsCreate) SaveX(ctx context.Context) *Topic_File_Mats {
	v, err := tfmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tfmc *TopicFileMatsCreate) Exec(ctx context.Context) error {
	_, err := tfmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfmc *TopicFileMatsCreate) ExecX(ctx context.Context) {
	if err := tfmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tfmc *TopicFileMatsCreate) defaults() {
	if _, ok := tfmc.mutation.CreatedAt(); !ok {
		v := topic_file_mats.DefaultCreatedAt
		tfmc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tfmc *TopicFileMatsCreate) check() error {
	if _, ok := tfmc.mutation.TopicID(); !ok {
		return &ValidationError{Name: "topic_id", err: errors.New(`ent: missing required field "Topic_File_Mats.topic_id"`)}
	}
	if _, ok := tfmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Topic_File_Mats.name"`)}
	}
	if v, ok := tfmc.mutation.Name(); ok {
		if err := topic_file_mats.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Topic_File_Mats.name": %w`, err)}
		}
	}
	if _, ok := tfmc.mutation.FilePath(); !ok {
		return &ValidationError{Name: "file_path", err: errors.New(`ent: missing required field "Topic_File_Mats.file_path"`)}
	}
	if _, ok := tfmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Topic_File_Mats.created_at"`)}
	}
	if _, ok := tfmc.mutation.TopicID(); !ok {
		return &ValidationError{Name: "topic", err: errors.New(`ent: missing required edge "Topic_File_Mats.topic"`)}
	}
	return nil
}

func (tfmc *TopicFileMatsCreate) sqlSave(ctx context.Context) (*Topic_File_Mats, error) {
	if err := tfmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tfmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tfmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tfmc.mutation.id = &_node.ID
	tfmc.mutation.done = true
	return _node, nil
}

func (tfmc *TopicFileMatsCreate) createSpec() (*Topic_File_Mats, *sqlgraph.CreateSpec) {
	var (
		_node = &Topic_File_Mats{config: tfmc.config}
		_spec = sqlgraph.NewCreateSpec(topic_file_mats.Table, sqlgraph.NewFieldSpec(topic_file_mats.FieldID, field.TypeInt))
	)
	if value, ok := tfmc.mutation.Name(); ok {
		_spec.SetField(topic_file_mats.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tfmc.mutation.FilePath(); ok {
		_spec.SetField(topic_file_mats.FieldFilePath, field.TypeString, value)
		_node.FilePath = value
	}
	if value, ok := tfmc.mutation.CreatedAt(); ok {
		_spec.SetField(topic_file_mats.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := tfmc.mutation.TopicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic_file_mats.TopicTable,
			Columns: []string{topic_file_mats.TopicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TopicID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TopicFileMatsCreateBulk is the builder for creating many Topic_File_Mats entities in bulk.
type TopicFileMatsCreateBulk struct {
	config
	builders []*TopicFileMatsCreate
}

// Save creates the Topic_File_Mats entities in the database.
func (tfmcb *TopicFileMatsCreateBulk) Save(ctx context.Context) ([]*Topic_File_Mats, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tfmcb.builders))
	nodes := make([]*Topic_File_Mats, len(tfmcb.builders))
	mutators := make([]Mutator, len(tfmcb.builders))
	for i := range tfmcb.builders {
		func(i int, root context.Context) {
			builder := tfmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopicFileMatsMutation)
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
					_, err = mutators[i+1].Mutate(root, tfmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tfmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tfmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tfmcb *TopicFileMatsCreateBulk) SaveX(ctx context.Context) []*Topic_File_Mats {
	v, err := tfmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tfmcb *TopicFileMatsCreateBulk) Exec(ctx context.Context) error {
	_, err := tfmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfmcb *TopicFileMatsCreateBulk) ExecX(ctx context.Context) {
	if err := tfmcb.Exec(ctx); err != nil {
		panic(err)
	}
}