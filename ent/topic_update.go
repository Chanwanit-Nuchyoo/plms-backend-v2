// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/course"
	"plms-backend/ent/predicate"
	"plms-backend/ent/topic"
	"plms-backend/ent/topic_file_mats"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// Where appends a list predicates to the TopicUpdate builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetCourseID sets the "course_id" field.
func (tu *TopicUpdate) SetCourseID(i int) *TopicUpdate {
	tu.mutation.SetCourseID(i)
	return tu
}

// SetName sets the "name" field.
func (tu *TopicUpdate) SetName(s string) *TopicUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TopicUpdate) SetDescription(s string) *TopicUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TopicUpdate) SetCreatedAt(t time.Time) *TopicUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableCreatedAt(t *time.Time) *TopicUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TopicUpdate) SetUpdatedAt(t time.Time) *TopicUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableUpdatedAt(t *time.Time) *TopicUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// SetCourse sets the "course" edge to the Course entity.
func (tu *TopicUpdate) SetCourse(c *Course) *TopicUpdate {
	return tu.SetCourseID(c.ID)
}

// AddFileIDs adds the "files" edge to the Topic_File_Mats entity by IDs.
func (tu *TopicUpdate) AddFileIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddFileIDs(ids...)
	return tu
}

// AddFiles adds the "files" edges to the Topic_File_Mats entity.
func (tu *TopicUpdate) AddFiles(t ...*Topic_File_Mats) *TopicUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddFileIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// ClearCourse clears the "course" edge to the Course entity.
func (tu *TopicUpdate) ClearCourse() *TopicUpdate {
	tu.mutation.ClearCourse()
	return tu
}

// ClearFiles clears all "files" edges to the Topic_File_Mats entity.
func (tu *TopicUpdate) ClearFiles() *TopicUpdate {
	tu.mutation.ClearFiles()
	return tu
}

// RemoveFileIDs removes the "files" edge to Topic_File_Mats entities by IDs.
func (tu *TopicUpdate) RemoveFileIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveFileIDs(ids...)
	return tu
}

// RemoveFiles removes "files" edges to Topic_File_Mats entities.
func (tu *TopicUpdate) RemoveFiles(t ...*Topic_File_Mats) *TopicUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TopicUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := topic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Topic.name": %w`, err)}
		}
	}
	if _, ok := tu.mutation.CourseID(); tu.mutation.CourseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Topic.course"`)
	}
	return nil
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(topic.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(topic.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(topic.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.CourseTable,
			Columns: []string{topic.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.CourseTable,
			Columns: []string{topic.CourseColumn},
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
	if tu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.FilesTable,
			Columns: []string{topic.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic_file_mats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !tu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.FilesTable,
			Columns: []string{topic.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic_file_mats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.FilesTable,
			Columns: []string{topic.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic_file_mats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TopicMutation
}

// SetCourseID sets the "course_id" field.
func (tuo *TopicUpdateOne) SetCourseID(i int) *TopicUpdateOne {
	tuo.mutation.SetCourseID(i)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TopicUpdateOne) SetName(s string) *TopicUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TopicUpdateOne) SetDescription(s string) *TopicUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TopicUpdateOne) SetCreatedAt(t time.Time) *TopicUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableCreatedAt(t *time.Time) *TopicUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TopicUpdateOne) SetUpdatedAt(t time.Time) *TopicUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableUpdatedAt(t *time.Time) *TopicUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// SetCourse sets the "course" edge to the Course entity.
func (tuo *TopicUpdateOne) SetCourse(c *Course) *TopicUpdateOne {
	return tuo.SetCourseID(c.ID)
}

// AddFileIDs adds the "files" edge to the Topic_File_Mats entity by IDs.
func (tuo *TopicUpdateOne) AddFileIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddFileIDs(ids...)
	return tuo
}

// AddFiles adds the "files" edges to the Topic_File_Mats entity.
func (tuo *TopicUpdateOne) AddFiles(t ...*Topic_File_Mats) *TopicUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddFileIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// ClearCourse clears the "course" edge to the Course entity.
func (tuo *TopicUpdateOne) ClearCourse() *TopicUpdateOne {
	tuo.mutation.ClearCourse()
	return tuo
}

// ClearFiles clears all "files" edges to the Topic_File_Mats entity.
func (tuo *TopicUpdateOne) ClearFiles() *TopicUpdateOne {
	tuo.mutation.ClearFiles()
	return tuo
}

// RemoveFileIDs removes the "files" edge to Topic_File_Mats entities by IDs.
func (tuo *TopicUpdateOne) RemoveFileIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveFileIDs(ids...)
	return tuo
}

// RemoveFiles removes "files" edges to Topic_File_Mats entities.
func (tuo *TopicUpdateOne) RemoveFiles(t ...*Topic_File_Mats) *TopicUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveFileIDs(ids...)
}

// Where appends a list predicates to the TopicUpdate builder.
func (tuo *TopicUpdateOne) Where(ps ...predicate.Topic) *TopicUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TopicUpdateOne) Select(field string, fields ...string) *TopicUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Topic entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TopicUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := topic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Topic.name": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.CourseID(); tuo.mutation.CourseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Topic.course"`)
	}
	return nil
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (_node *Topic, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Topic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topic.FieldID)
		for _, f := range fields {
			if !topic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(topic.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(topic.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(topic.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.CourseTable,
			Columns: []string{topic.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.CourseTable,
			Columns: []string{topic.CourseColumn},
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
	if tuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.FilesTable,
			Columns: []string{topic.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic_file_mats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !tuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.FilesTable,
			Columns: []string{topic.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic_file_mats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.FilesTable,
			Columns: []string{topic.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic_file_mats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Topic{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
