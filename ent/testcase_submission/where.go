// Code generated by ent, DO NOT EDIT.

package testcase_submission

import (
	"plms-backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLTE(FieldID, id))
}

// SubmissionID applies equality check predicate on the "submission_id" field. It's identical to SubmissionIDEQ.
func SubmissionID(v int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldSubmissionID, v))
}

// TestcaseID applies equality check predicate on the "testcase_id" field. It's identical to TestcaseIDEQ.
func TestcaseID(v int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldTestcaseID, v))
}

// Input applies equality check predicate on the "input" field. It's identical to InputEQ.
func Input(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldInput, v))
}

// ExpectedOutput applies equality check predicate on the "expected_output" field. It's identical to ExpectedOutputEQ.
func ExpectedOutput(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldExpectedOutput, v))
}

// ActualOutput applies equality check predicate on the "actual_output" field. It's identical to ActualOutputEQ.
func ActualOutput(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldActualOutput, v))
}

// IsPassed applies equality check predicate on the "is_passed" field. It's identical to IsPassedEQ.
func IsPassed(v bool) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldIsPassed, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldScore, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldUpdatedAt, v))
}

// SubmissionIDEQ applies the EQ predicate on the "submission_id" field.
func SubmissionIDEQ(v int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldSubmissionID, v))
}

// SubmissionIDNEQ applies the NEQ predicate on the "submission_id" field.
func SubmissionIDNEQ(v int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldSubmissionID, v))
}

// SubmissionIDIn applies the In predicate on the "submission_id" field.
func SubmissionIDIn(vs ...int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldSubmissionID, vs...))
}

// SubmissionIDNotIn applies the NotIn predicate on the "submission_id" field.
func SubmissionIDNotIn(vs ...int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldSubmissionID, vs...))
}

// TestcaseIDEQ applies the EQ predicate on the "testcase_id" field.
func TestcaseIDEQ(v int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldTestcaseID, v))
}

// TestcaseIDNEQ applies the NEQ predicate on the "testcase_id" field.
func TestcaseIDNEQ(v int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldTestcaseID, v))
}

// TestcaseIDIn applies the In predicate on the "testcase_id" field.
func TestcaseIDIn(vs ...int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldTestcaseID, vs...))
}

// TestcaseIDNotIn applies the NotIn predicate on the "testcase_id" field.
func TestcaseIDNotIn(vs ...int) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldTestcaseID, vs...))
}

// InputEQ applies the EQ predicate on the "input" field.
func InputEQ(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldInput, v))
}

// InputNEQ applies the NEQ predicate on the "input" field.
func InputNEQ(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldInput, v))
}

// InputIn applies the In predicate on the "input" field.
func InputIn(vs ...string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldInput, vs...))
}

// InputNotIn applies the NotIn predicate on the "input" field.
func InputNotIn(vs ...string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldInput, vs...))
}

// InputGT applies the GT predicate on the "input" field.
func InputGT(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGT(FieldInput, v))
}

// InputGTE applies the GTE predicate on the "input" field.
func InputGTE(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGTE(FieldInput, v))
}

// InputLT applies the LT predicate on the "input" field.
func InputLT(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLT(FieldInput, v))
}

// InputLTE applies the LTE predicate on the "input" field.
func InputLTE(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLTE(FieldInput, v))
}

// InputContains applies the Contains predicate on the "input" field.
func InputContains(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldContains(FieldInput, v))
}

// InputHasPrefix applies the HasPrefix predicate on the "input" field.
func InputHasPrefix(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldHasPrefix(FieldInput, v))
}

// InputHasSuffix applies the HasSuffix predicate on the "input" field.
func InputHasSuffix(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldHasSuffix(FieldInput, v))
}

// InputEqualFold applies the EqualFold predicate on the "input" field.
func InputEqualFold(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEqualFold(FieldInput, v))
}

// InputContainsFold applies the ContainsFold predicate on the "input" field.
func InputContainsFold(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldContainsFold(FieldInput, v))
}

// ExpectedOutputEQ applies the EQ predicate on the "expected_output" field.
func ExpectedOutputEQ(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldExpectedOutput, v))
}

// ExpectedOutputNEQ applies the NEQ predicate on the "expected_output" field.
func ExpectedOutputNEQ(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldExpectedOutput, v))
}

// ExpectedOutputIn applies the In predicate on the "expected_output" field.
func ExpectedOutputIn(vs ...string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldExpectedOutput, vs...))
}

// ExpectedOutputNotIn applies the NotIn predicate on the "expected_output" field.
func ExpectedOutputNotIn(vs ...string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldExpectedOutput, vs...))
}

// ExpectedOutputGT applies the GT predicate on the "expected_output" field.
func ExpectedOutputGT(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGT(FieldExpectedOutput, v))
}

// ExpectedOutputGTE applies the GTE predicate on the "expected_output" field.
func ExpectedOutputGTE(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGTE(FieldExpectedOutput, v))
}

// ExpectedOutputLT applies the LT predicate on the "expected_output" field.
func ExpectedOutputLT(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLT(FieldExpectedOutput, v))
}

// ExpectedOutputLTE applies the LTE predicate on the "expected_output" field.
func ExpectedOutputLTE(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLTE(FieldExpectedOutput, v))
}

// ExpectedOutputContains applies the Contains predicate on the "expected_output" field.
func ExpectedOutputContains(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldContains(FieldExpectedOutput, v))
}

// ExpectedOutputHasPrefix applies the HasPrefix predicate on the "expected_output" field.
func ExpectedOutputHasPrefix(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldHasPrefix(FieldExpectedOutput, v))
}

// ExpectedOutputHasSuffix applies the HasSuffix predicate on the "expected_output" field.
func ExpectedOutputHasSuffix(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldHasSuffix(FieldExpectedOutput, v))
}

// ExpectedOutputEqualFold applies the EqualFold predicate on the "expected_output" field.
func ExpectedOutputEqualFold(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEqualFold(FieldExpectedOutput, v))
}

// ExpectedOutputContainsFold applies the ContainsFold predicate on the "expected_output" field.
func ExpectedOutputContainsFold(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldContainsFold(FieldExpectedOutput, v))
}

// ActualOutputEQ applies the EQ predicate on the "actual_output" field.
func ActualOutputEQ(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldActualOutput, v))
}

// ActualOutputNEQ applies the NEQ predicate on the "actual_output" field.
func ActualOutputNEQ(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldActualOutput, v))
}

// ActualOutputIn applies the In predicate on the "actual_output" field.
func ActualOutputIn(vs ...string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldActualOutput, vs...))
}

// ActualOutputNotIn applies the NotIn predicate on the "actual_output" field.
func ActualOutputNotIn(vs ...string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldActualOutput, vs...))
}

// ActualOutputGT applies the GT predicate on the "actual_output" field.
func ActualOutputGT(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGT(FieldActualOutput, v))
}

// ActualOutputGTE applies the GTE predicate on the "actual_output" field.
func ActualOutputGTE(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGTE(FieldActualOutput, v))
}

// ActualOutputLT applies the LT predicate on the "actual_output" field.
func ActualOutputLT(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLT(FieldActualOutput, v))
}

// ActualOutputLTE applies the LTE predicate on the "actual_output" field.
func ActualOutputLTE(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLTE(FieldActualOutput, v))
}

// ActualOutputContains applies the Contains predicate on the "actual_output" field.
func ActualOutputContains(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldContains(FieldActualOutput, v))
}

// ActualOutputHasPrefix applies the HasPrefix predicate on the "actual_output" field.
func ActualOutputHasPrefix(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldHasPrefix(FieldActualOutput, v))
}

// ActualOutputHasSuffix applies the HasSuffix predicate on the "actual_output" field.
func ActualOutputHasSuffix(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldHasSuffix(FieldActualOutput, v))
}

// ActualOutputEqualFold applies the EqualFold predicate on the "actual_output" field.
func ActualOutputEqualFold(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEqualFold(FieldActualOutput, v))
}

// ActualOutputContainsFold applies the ContainsFold predicate on the "actual_output" field.
func ActualOutputContainsFold(v string) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldContainsFold(FieldActualOutput, v))
}

// IsPassedEQ applies the EQ predicate on the "is_passed" field.
func IsPassedEQ(v bool) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldIsPassed, v))
}

// IsPassedNEQ applies the NEQ predicate on the "is_passed" field.
func IsPassedNEQ(v bool) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldIsPassed, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v float64) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLTE(FieldScore, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasTestcase applies the HasEdge predicate on the "testcase" edge.
func HasTestcase() predicate.Testcase_Submission {
	return predicate.Testcase_Submission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestcaseTable, TestcaseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestcaseWith applies the HasEdge predicate on the "testcase" edge with a given conditions (other predicates).
func HasTestcaseWith(preds ...predicate.Testcase) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(func(s *sql.Selector) {
		step := newTestcaseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubmission applies the HasEdge predicate on the "submission" edge.
func HasSubmission() predicate.Testcase_Submission {
	return predicate.Testcase_Submission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubmissionTable, SubmissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmissionWith applies the HasEdge predicate on the "submission" edge with a given conditions (other predicates).
func HasSubmissionWith(preds ...predicate.Lab_Problem_Submission) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(func(s *sql.Selector) {
		step := newSubmissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Testcase_Submission) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Testcase_Submission) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Testcase_Submission) predicate.Testcase_Submission {
	return predicate.Testcase_Submission(func(s *sql.Selector) {
		p(s.Not())
	})
}
