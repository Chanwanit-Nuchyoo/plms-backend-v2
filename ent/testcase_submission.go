// Code generated by ent, DO NOT EDIT.

package ent

import (
	"plms-backend/ent/lab_problem_submission"
	"plms-backend/ent/testcase"
	"plms-backend/ent/testcase_submission"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Testcase_Submission is the model entity for the Testcase_Submission schema.
type Testcase_Submission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SubmissionID holds the value of the "submission_id" field.
	SubmissionID int `json:"submission_id,omitempty"`
	// TestcaseID holds the value of the "testcase_id" field.
	TestcaseID int `json:"testcase_id,omitempty"`
	// Input holds the value of the "input" field.
	Input string `json:"input,omitempty"`
	// ExpectedOutput holds the value of the "expected_output" field.
	ExpectedOutput string `json:"expected_output,omitempty"`
	// ActualOutput holds the value of the "actual_output" field.
	ActualOutput string `json:"actual_output,omitempty"`
	// IsPassed holds the value of the "is_passed" field.
	IsPassed bool `json:"is_passed,omitempty"`
	// Score holds the value of the "score" field.
	Score float64 `json:"score,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Testcase_SubmissionQuery when eager-loading is set.
	Edges        Testcase_SubmissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Testcase_SubmissionEdges holds the relations/edges for other nodes in the graph.
type Testcase_SubmissionEdges struct {
	// Testcase holds the value of the testcase edge.
	Testcase *Testcase `json:"testcase,omitempty"`
	// Submission holds the value of the submission edge.
	Submission *Lab_Problem_Submission `json:"submission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TestcaseOrErr returns the Testcase value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Testcase_SubmissionEdges) TestcaseOrErr() (*Testcase, error) {
	if e.loadedTypes[0] {
		if e.Testcase == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: testcase.Label}
		}
		return e.Testcase, nil
	}
	return nil, &NotLoadedError{edge: "testcase"}
}

// SubmissionOrErr returns the Submission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Testcase_SubmissionEdges) SubmissionOrErr() (*Lab_Problem_Submission, error) {
	if e.loadedTypes[1] {
		if e.Submission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: lab_problem_submission.Label}
		}
		return e.Submission, nil
	}
	return nil, &NotLoadedError{edge: "submission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Testcase_Submission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testcase_submission.FieldIsPassed:
			values[i] = new(sql.NullBool)
		case testcase_submission.FieldScore:
			values[i] = new(sql.NullFloat64)
		case testcase_submission.FieldID, testcase_submission.FieldSubmissionID, testcase_submission.FieldTestcaseID:
			values[i] = new(sql.NullInt64)
		case testcase_submission.FieldInput, testcase_submission.FieldExpectedOutput, testcase_submission.FieldActualOutput:
			values[i] = new(sql.NullString)
		case testcase_submission.FieldCreatedAt, testcase_submission.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Testcase_Submission fields.
func (ts *Testcase_Submission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testcase_submission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int(value.Int64)
		case testcase_submission.FieldSubmissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field submission_id", values[i])
			} else if value.Valid {
				ts.SubmissionID = int(value.Int64)
			}
		case testcase_submission.FieldTestcaseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field testcase_id", values[i])
			} else if value.Valid {
				ts.TestcaseID = int(value.Int64)
			}
		case testcase_submission.FieldInput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input", values[i])
			} else if value.Valid {
				ts.Input = value.String
			}
		case testcase_submission.FieldExpectedOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field expected_output", values[i])
			} else if value.Valid {
				ts.ExpectedOutput = value.String
			}
		case testcase_submission.FieldActualOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actual_output", values[i])
			} else if value.Valid {
				ts.ActualOutput = value.String
			}
		case testcase_submission.FieldIsPassed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_passed", values[i])
			} else if value.Valid {
				ts.IsPassed = value.Bool
			}
		case testcase_submission.FieldScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				ts.Score = value.Float64
			}
		case testcase_submission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ts.CreatedAt = value.Time
			}
		case testcase_submission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ts.UpdatedAt = value.Time
			}
		default:
			ts.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Testcase_Submission.
// This includes values selected through modifiers, order, etc.
func (ts *Testcase_Submission) Value(name string) (ent.Value, error) {
	return ts.selectValues.Get(name)
}

// QueryTestcase queries the "testcase" edge of the Testcase_Submission entity.
func (ts *Testcase_Submission) QueryTestcase() *TestcaseQuery {
	return NewTestcaseSubmissionClient(ts.config).QueryTestcase(ts)
}

// QuerySubmission queries the "submission" edge of the Testcase_Submission entity.
func (ts *Testcase_Submission) QuerySubmission() *LabProblemSubmissionQuery {
	return NewTestcaseSubmissionClient(ts.config).QuerySubmission(ts)
}

// Update returns a builder for updating this Testcase_Submission.
// Note that you need to call Testcase_Submission.Unwrap() before calling this method if this Testcase_Submission
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *Testcase_Submission) Update() *TestcaseSubmissionUpdateOne {
	return NewTestcaseSubmissionClient(ts.config).UpdateOne(ts)
}

// Unwrap unwraps the Testcase_Submission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *Testcase_Submission) Unwrap() *Testcase_Submission {
	_tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("ent: Testcase_Submission is not a transactional entity")
	}
	ts.config.driver = _tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *Testcase_Submission) String() string {
	var builder strings.Builder
	builder.WriteString("Testcase_Submission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ts.ID))
	builder.WriteString("submission_id=")
	builder.WriteString(fmt.Sprintf("%v", ts.SubmissionID))
	builder.WriteString(", ")
	builder.WriteString("testcase_id=")
	builder.WriteString(fmt.Sprintf("%v", ts.TestcaseID))
	builder.WriteString(", ")
	builder.WriteString("input=")
	builder.WriteString(ts.Input)
	builder.WriteString(", ")
	builder.WriteString("expected_output=")
	builder.WriteString(ts.ExpectedOutput)
	builder.WriteString(", ")
	builder.WriteString("actual_output=")
	builder.WriteString(ts.ActualOutput)
	builder.WriteString(", ")
	builder.WriteString("is_passed=")
	builder.WriteString(fmt.Sprintf("%v", ts.IsPassed))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", ts.Score))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ts.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ts.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Testcase_Submissions is a parsable slice of Testcase_Submission.
type Testcase_Submissions []*Testcase_Submission