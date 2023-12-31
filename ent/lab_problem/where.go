// Code generated by ent, DO NOT EDIT.

package lab_problem

import (
	"plms-backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLTE(FieldID, id))
}

// LabID applies equality check predicate on the "lab_id" field. It's identical to LabIDEQ.
func LabID(v int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldLabID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldName, v))
}

// Prompt applies equality check predicate on the "prompt" field. It's identical to PromptEQ.
func Prompt(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldPrompt, v))
}

// FullScore applies equality check predicate on the "full_score" field. It's identical to FullScoreEQ.
func FullScore(v float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldFullScore, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldUpdatedAt, v))
}

// LabIDEQ applies the EQ predicate on the "lab_id" field.
func LabIDEQ(v int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldLabID, v))
}

// LabIDNEQ applies the NEQ predicate on the "lab_id" field.
func LabIDNEQ(v int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNEQ(FieldLabID, v))
}

// LabIDIn applies the In predicate on the "lab_id" field.
func LabIDIn(vs ...int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIn(FieldLabID, vs...))
}

// LabIDNotIn applies the NotIn predicate on the "lab_id" field.
func LabIDNotIn(vs ...int) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotIn(FieldLabID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldContainsFold(FieldName, v))
}

// PromptEQ applies the EQ predicate on the "prompt" field.
func PromptEQ(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldPrompt, v))
}

// PromptNEQ applies the NEQ predicate on the "prompt" field.
func PromptNEQ(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNEQ(FieldPrompt, v))
}

// PromptIn applies the In predicate on the "prompt" field.
func PromptIn(vs ...string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIn(FieldPrompt, vs...))
}

// PromptNotIn applies the NotIn predicate on the "prompt" field.
func PromptNotIn(vs ...string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotIn(FieldPrompt, vs...))
}

// PromptGT applies the GT predicate on the "prompt" field.
func PromptGT(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGT(FieldPrompt, v))
}

// PromptGTE applies the GTE predicate on the "prompt" field.
func PromptGTE(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGTE(FieldPrompt, v))
}

// PromptLT applies the LT predicate on the "prompt" field.
func PromptLT(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLT(FieldPrompt, v))
}

// PromptLTE applies the LTE predicate on the "prompt" field.
func PromptLTE(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLTE(FieldPrompt, v))
}

// PromptContains applies the Contains predicate on the "prompt" field.
func PromptContains(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldContains(FieldPrompt, v))
}

// PromptHasPrefix applies the HasPrefix predicate on the "prompt" field.
func PromptHasPrefix(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldHasPrefix(FieldPrompt, v))
}

// PromptHasSuffix applies the HasSuffix predicate on the "prompt" field.
func PromptHasSuffix(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldHasSuffix(FieldPrompt, v))
}

// PromptEqualFold applies the EqualFold predicate on the "prompt" field.
func PromptEqualFold(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEqualFold(FieldPrompt, v))
}

// PromptContainsFold applies the ContainsFold predicate on the "prompt" field.
func PromptContainsFold(v string) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldContainsFold(FieldPrompt, v))
}

// FullScoreEQ applies the EQ predicate on the "full_score" field.
func FullScoreEQ(v float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldFullScore, v))
}

// FullScoreNEQ applies the NEQ predicate on the "full_score" field.
func FullScoreNEQ(v float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNEQ(FieldFullScore, v))
}

// FullScoreIn applies the In predicate on the "full_score" field.
func FullScoreIn(vs ...float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIn(FieldFullScore, vs...))
}

// FullScoreNotIn applies the NotIn predicate on the "full_score" field.
func FullScoreNotIn(vs ...float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotIn(FieldFullScore, vs...))
}

// FullScoreGT applies the GT predicate on the "full_score" field.
func FullScoreGT(v float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGT(FieldFullScore, v))
}

// FullScoreGTE applies the GTE predicate on the "full_score" field.
func FullScoreGTE(v float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGTE(FieldFullScore, v))
}

// FullScoreLT applies the LT predicate on the "full_score" field.
func FullScoreLT(v float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLT(FieldFullScore, v))
}

// FullScoreLTE applies the LTE predicate on the "full_score" field.
func FullScoreLTE(v float64) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLTE(FieldFullScore, v))
}

// FullScoreIsNil applies the IsNil predicate on the "full_score" field.
func FullScoreIsNil() predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIsNull(FieldFullScore))
}

// FullScoreNotNil applies the NotNil predicate on the "full_score" field.
func FullScoreNotNil() predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotNull(FieldFullScore))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Lab_Problem {
	return predicate.Lab_Problem(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasLab applies the HasEdge predicate on the "lab" edge.
func HasLab() predicate.Lab_Problem {
	return predicate.Lab_Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LabTable, LabColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLabWith applies the HasEdge predicate on the "lab" edge with a given conditions (other predicates).
func HasLabWith(preds ...predicate.Lab) predicate.Lab_Problem {
	return predicate.Lab_Problem(func(s *sql.Selector) {
		step := newLabStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTestcases applies the HasEdge predicate on the "testcases" edge.
func HasTestcases() predicate.Lab_Problem {
	return predicate.Lab_Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcasesTable, TestcasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestcasesWith applies the HasEdge predicate on the "testcases" edge with a given conditions (other predicates).
func HasTestcasesWith(preds ...predicate.Testcase) predicate.Lab_Problem {
	return predicate.Lab_Problem(func(s *sql.Selector) {
		step := newTestcasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Lab_Problem) predicate.Lab_Problem {
	return predicate.Lab_Problem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Lab_Problem) predicate.Lab_Problem {
	return predicate.Lab_Problem(func(s *sql.Selector) {
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
func Not(p predicate.Lab_Problem) predicate.Lab_Problem {
	return predicate.Lab_Problem(func(s *sql.Selector) {
		p(s.Not())
	})
}
