// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/lab_problem_submission"
	"plms-backend/ent/testcase"
	"plms-backend/ent/testcase_submission"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestcaseSubmissionCreate is the builder for creating a Testcase_Submission entity.
type TestcaseSubmissionCreate struct {
	config
	mutation *TestcaseSubmissionMutation
	hooks    []Hook
}

// SetSubmissionID sets the "submission_id" field.
func (tsc *TestcaseSubmissionCreate) SetSubmissionID(i int) *TestcaseSubmissionCreate {
	tsc.mutation.SetSubmissionID(i)
	return tsc
}

// SetTestcaseID sets the "testcase_id" field.
func (tsc *TestcaseSubmissionCreate) SetTestcaseID(i int) *TestcaseSubmissionCreate {
	tsc.mutation.SetTestcaseID(i)
	return tsc
}

// SetInput sets the "input" field.
func (tsc *TestcaseSubmissionCreate) SetInput(s string) *TestcaseSubmissionCreate {
	tsc.mutation.SetInput(s)
	return tsc
}

// SetExpectedOutput sets the "expected_output" field.
func (tsc *TestcaseSubmissionCreate) SetExpectedOutput(s string) *TestcaseSubmissionCreate {
	tsc.mutation.SetExpectedOutput(s)
	return tsc
}

// SetActualOutput sets the "actual_output" field.
func (tsc *TestcaseSubmissionCreate) SetActualOutput(s string) *TestcaseSubmissionCreate {
	tsc.mutation.SetActualOutput(s)
	return tsc
}

// SetIsPassed sets the "is_passed" field.
func (tsc *TestcaseSubmissionCreate) SetIsPassed(b bool) *TestcaseSubmissionCreate {
	tsc.mutation.SetIsPassed(b)
	return tsc
}

// SetScore sets the "score" field.
func (tsc *TestcaseSubmissionCreate) SetScore(f float64) *TestcaseSubmissionCreate {
	tsc.mutation.SetScore(f)
	return tsc
}

// SetCreatedAt sets the "created_at" field.
func (tsc *TestcaseSubmissionCreate) SetCreatedAt(t time.Time) *TestcaseSubmissionCreate {
	tsc.mutation.SetCreatedAt(t)
	return tsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tsc *TestcaseSubmissionCreate) SetNillableCreatedAt(t *time.Time) *TestcaseSubmissionCreate {
	if t != nil {
		tsc.SetCreatedAt(*t)
	}
	return tsc
}

// SetUpdatedAt sets the "updated_at" field.
func (tsc *TestcaseSubmissionCreate) SetUpdatedAt(t time.Time) *TestcaseSubmissionCreate {
	tsc.mutation.SetUpdatedAt(t)
	return tsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tsc *TestcaseSubmissionCreate) SetNillableUpdatedAt(t *time.Time) *TestcaseSubmissionCreate {
	if t != nil {
		tsc.SetUpdatedAt(*t)
	}
	return tsc
}

// SetTestcase sets the "testcase" edge to the Testcase entity.
func (tsc *TestcaseSubmissionCreate) SetTestcase(t *Testcase) *TestcaseSubmissionCreate {
	return tsc.SetTestcaseID(t.ID)
}

// SetSubmission sets the "submission" edge to the Lab_Problem_Submission entity.
func (tsc *TestcaseSubmissionCreate) SetSubmission(l *Lab_Problem_Submission) *TestcaseSubmissionCreate {
	return tsc.SetSubmissionID(l.ID)
}

// Mutation returns the TestcaseSubmissionMutation object of the builder.
func (tsc *TestcaseSubmissionCreate) Mutation() *TestcaseSubmissionMutation {
	return tsc.mutation
}

// Save creates the Testcase_Submission in the database.
func (tsc *TestcaseSubmissionCreate) Save(ctx context.Context) (*Testcase_Submission, error) {
	tsc.defaults()
	return withHooks(ctx, tsc.sqlSave, tsc.mutation, tsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tsc *TestcaseSubmissionCreate) SaveX(ctx context.Context) *Testcase_Submission {
	v, err := tsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tsc *TestcaseSubmissionCreate) Exec(ctx context.Context) error {
	_, err := tsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsc *TestcaseSubmissionCreate) ExecX(ctx context.Context) {
	if err := tsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tsc *TestcaseSubmissionCreate) defaults() {
	if _, ok := tsc.mutation.CreatedAt(); !ok {
		v := testcase_submission.DefaultCreatedAt
		tsc.mutation.SetCreatedAt(v)
	}
	if _, ok := tsc.mutation.UpdatedAt(); !ok {
		v := testcase_submission.DefaultUpdatedAt
		tsc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsc *TestcaseSubmissionCreate) check() error {
	if _, ok := tsc.mutation.SubmissionID(); !ok {
		return &ValidationError{Name: "submission_id", err: errors.New(`ent: missing required field "Testcase_Submission.submission_id"`)}
	}
	if _, ok := tsc.mutation.TestcaseID(); !ok {
		return &ValidationError{Name: "testcase_id", err: errors.New(`ent: missing required field "Testcase_Submission.testcase_id"`)}
	}
	if _, ok := tsc.mutation.Input(); !ok {
		return &ValidationError{Name: "input", err: errors.New(`ent: missing required field "Testcase_Submission.input"`)}
	}
	if _, ok := tsc.mutation.ExpectedOutput(); !ok {
		return &ValidationError{Name: "expected_output", err: errors.New(`ent: missing required field "Testcase_Submission.expected_output"`)}
	}
	if _, ok := tsc.mutation.ActualOutput(); !ok {
		return &ValidationError{Name: "actual_output", err: errors.New(`ent: missing required field "Testcase_Submission.actual_output"`)}
	}
	if _, ok := tsc.mutation.IsPassed(); !ok {
		return &ValidationError{Name: "is_passed", err: errors.New(`ent: missing required field "Testcase_Submission.is_passed"`)}
	}
	if _, ok := tsc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "Testcase_Submission.score"`)}
	}
	if _, ok := tsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Testcase_Submission.created_at"`)}
	}
	if _, ok := tsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Testcase_Submission.updated_at"`)}
	}
	if _, ok := tsc.mutation.TestcaseID(); !ok {
		return &ValidationError{Name: "testcase", err: errors.New(`ent: missing required edge "Testcase_Submission.testcase"`)}
	}
	if _, ok := tsc.mutation.SubmissionID(); !ok {
		return &ValidationError{Name: "submission", err: errors.New(`ent: missing required edge "Testcase_Submission.submission"`)}
	}
	return nil
}

func (tsc *TestcaseSubmissionCreate) sqlSave(ctx context.Context) (*Testcase_Submission, error) {
	if err := tsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tsc.mutation.id = &_node.ID
	tsc.mutation.done = true
	return _node, nil
}

func (tsc *TestcaseSubmissionCreate) createSpec() (*Testcase_Submission, *sqlgraph.CreateSpec) {
	var (
		_node = &Testcase_Submission{config: tsc.config}
		_spec = sqlgraph.NewCreateSpec(testcase_submission.Table, sqlgraph.NewFieldSpec(testcase_submission.FieldID, field.TypeInt))
	)
	if value, ok := tsc.mutation.Input(); ok {
		_spec.SetField(testcase_submission.FieldInput, field.TypeString, value)
		_node.Input = value
	}
	if value, ok := tsc.mutation.ExpectedOutput(); ok {
		_spec.SetField(testcase_submission.FieldExpectedOutput, field.TypeString, value)
		_node.ExpectedOutput = value
	}
	if value, ok := tsc.mutation.ActualOutput(); ok {
		_spec.SetField(testcase_submission.FieldActualOutput, field.TypeString, value)
		_node.ActualOutput = value
	}
	if value, ok := tsc.mutation.IsPassed(); ok {
		_spec.SetField(testcase_submission.FieldIsPassed, field.TypeBool, value)
		_node.IsPassed = value
	}
	if value, ok := tsc.mutation.Score(); ok {
		_spec.SetField(testcase_submission.FieldScore, field.TypeFloat64, value)
		_node.Score = value
	}
	if value, ok := tsc.mutation.CreatedAt(); ok {
		_spec.SetField(testcase_submission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tsc.mutation.UpdatedAt(); ok {
		_spec.SetField(testcase_submission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tsc.mutation.TestcaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testcase_submission.TestcaseTable,
			Columns: []string{testcase_submission.TestcaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TestcaseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tsc.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testcase_submission.SubmissionTable,
			Columns: []string{testcase_submission.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lab_problem_submission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SubmissionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TestcaseSubmissionCreateBulk is the builder for creating many Testcase_Submission entities in bulk.
type TestcaseSubmissionCreateBulk struct {
	config
	builders []*TestcaseSubmissionCreate
}

// Save creates the Testcase_Submission entities in the database.
func (tscb *TestcaseSubmissionCreateBulk) Save(ctx context.Context) ([]*Testcase_Submission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tscb.builders))
	nodes := make([]*Testcase_Submission, len(tscb.builders))
	mutators := make([]Mutator, len(tscb.builders))
	for i := range tscb.builders {
		func(i int, root context.Context) {
			builder := tscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestcaseSubmissionMutation)
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
					_, err = mutators[i+1].Mutate(root, tscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tscb *TestcaseSubmissionCreateBulk) SaveX(ctx context.Context) []*Testcase_Submission {
	v, err := tscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tscb *TestcaseSubmissionCreateBulk) Exec(ctx context.Context) error {
	_, err := tscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tscb *TestcaseSubmissionCreateBulk) ExecX(ctx context.Context) {
	if err := tscb.Exec(ctx); err != nil {
		panic(err)
	}
}