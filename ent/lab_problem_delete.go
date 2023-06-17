// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/lab_problem"
	"plms-backend/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LabProblemDelete is the builder for deleting a Lab_Problem entity.
type LabProblemDelete struct {
	config
	hooks    []Hook
	mutation *LabProblemMutation
}

// Where appends a list predicates to the LabProblemDelete builder.
func (lpd *LabProblemDelete) Where(ps ...predicate.Lab_Problem) *LabProblemDelete {
	lpd.mutation.Where(ps...)
	return lpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lpd *LabProblemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lpd.sqlExec, lpd.mutation, lpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lpd *LabProblemDelete) ExecX(ctx context.Context) int {
	n, err := lpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lpd *LabProblemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(lab_problem.Table, sqlgraph.NewFieldSpec(lab_problem.FieldID, field.TypeInt))
	if ps := lpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lpd.mutation.done = true
	return affected, err
}

// LabProblemDeleteOne is the builder for deleting a single Lab_Problem entity.
type LabProblemDeleteOne struct {
	lpd *LabProblemDelete
}

// Where appends a list predicates to the LabProblemDelete builder.
func (lpdo *LabProblemDeleteOne) Where(ps ...predicate.Lab_Problem) *LabProblemDeleteOne {
	lpdo.lpd.mutation.Where(ps...)
	return lpdo
}

// Exec executes the deletion query.
func (lpdo *LabProblemDeleteOne) Exec(ctx context.Context) error {
	n, err := lpdo.lpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lab_problem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lpdo *LabProblemDeleteOne) ExecX(ctx context.Context) {
	if err := lpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
