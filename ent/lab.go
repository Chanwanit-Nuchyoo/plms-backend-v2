// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/course"
	"plms-backend/ent/lab"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Lab is the model entity for the Lab schema.
type Lab struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CourseID holds the value of the "course_id" field.
	CourseID int `json:"course_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LabQuery when eager-loading is set.
	Edges        LabEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LabEdges holds the relations/edges for other nodes in the graph.
type LabEdges struct {
	// Course holds the value of the course edge.
	Course *Course `json:"course,omitempty"`
	// LabProblems holds the value of the lab_problems edge.
	LabProblems []*Lab_Problem `json:"lab_problems,omitempty"`
	// LabStatuses holds the value of the lab_statuses edge.
	LabStatuses []*Class_Lab_Status `json:"lab_statuses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LabEdges) CourseOrErr() (*Course, error) {
	if e.loadedTypes[0] {
		if e.Course == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: course.Label}
		}
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// LabProblemsOrErr returns the LabProblems value or an error if the edge
// was not loaded in eager-loading.
func (e LabEdges) LabProblemsOrErr() ([]*Lab_Problem, error) {
	if e.loadedTypes[1] {
		return e.LabProblems, nil
	}
	return nil, &NotLoadedError{edge: "lab_problems"}
}

// LabStatusesOrErr returns the LabStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e LabEdges) LabStatusesOrErr() ([]*Class_Lab_Status, error) {
	if e.loadedTypes[2] {
		return e.LabStatuses, nil
	}
	return nil, &NotLoadedError{edge: "lab_statuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Lab) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lab.FieldID, lab.FieldCourseID:
			values[i] = new(sql.NullInt64)
		case lab.FieldName:
			values[i] = new(sql.NullString)
		case lab.FieldCreatedAt, lab.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Lab fields.
func (l *Lab) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lab.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case lab.FieldCourseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_id", values[i])
			} else if value.Valid {
				l.CourseID = int(value.Int64)
			}
		case lab.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case lab.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case lab.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Lab.
// This includes values selected through modifiers, order, etc.
func (l *Lab) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryCourse queries the "course" edge of the Lab entity.
func (l *Lab) QueryCourse() *CourseQuery {
	return NewLabClient(l.config).QueryCourse(l)
}

// QueryLabProblems queries the "lab_problems" edge of the Lab entity.
func (l *Lab) QueryLabProblems() *LabProblemQuery {
	return NewLabClient(l.config).QueryLabProblems(l)
}

// QueryLabStatuses queries the "lab_statuses" edge of the Lab entity.
func (l *Lab) QueryLabStatuses() *ClassLabStatusQuery {
	return NewLabClient(l.config).QueryLabStatuses(l)
}

// Update returns a builder for updating this Lab.
// Note that you need to call Lab.Unwrap() before calling this method if this Lab
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Lab) Update() *LabUpdateOne {
	return NewLabClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Lab entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Lab) Unwrap() *Lab {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Lab is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Lab) String() string {
	var builder strings.Builder
	builder.WriteString("Lab(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("course_id=")
	builder.WriteString(fmt.Sprintf("%v", l.CourseID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Labs is a parsable slice of Lab.
type Labs []*Lab
