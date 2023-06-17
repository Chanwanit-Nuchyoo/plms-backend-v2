// Code generated by ent, DO NOT EDIT.

package class

import (
	"plms-backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldID, id))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldOwnerID, v))
}

// CourseID applies equality check predicate on the "course_id" field. It's identical to CourseIDEQ.
func CourseID(v int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldCourseID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldDescription, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldEndTime, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldUpdatedAt, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v int) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...int) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...int) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldOwnerID, vs...))
}

// CourseIDEQ applies the EQ predicate on the "course_id" field.
func CourseIDEQ(v int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldCourseID, v))
}

// CourseIDNEQ applies the NEQ predicate on the "course_id" field.
func CourseIDNEQ(v int) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldCourseID, v))
}

// CourseIDIn applies the In predicate on the "course_id" field.
func CourseIDIn(vs ...int) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldCourseID, vs...))
}

// CourseIDNotIn applies the NotIn predicate on the "course_id" field.
func CourseIDNotIn(vs ...int) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldCourseID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldDescription, v))
}

// DefaultLangEQ applies the EQ predicate on the "default_lang" field.
func DefaultLangEQ(v DefaultLang) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldDefaultLang, v))
}

// DefaultLangNEQ applies the NEQ predicate on the "default_lang" field.
func DefaultLangNEQ(v DefaultLang) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldDefaultLang, v))
}

// DefaultLangIn applies the In predicate on the "default_lang" field.
func DefaultLangIn(vs ...DefaultLang) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldDefaultLang, vs...))
}

// DefaultLangNotIn applies the NotIn predicate on the "default_lang" field.
func DefaultLangNotIn(vs ...DefaultLang) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldDefaultLang, vs...))
}

// DotwEQ applies the EQ predicate on the "dotw" field.
func DotwEQ(v Dotw) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldDotw, v))
}

// DotwNEQ applies the NEQ predicate on the "dotw" field.
func DotwNEQ(v Dotw) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldDotw, v))
}

// DotwIn applies the In predicate on the "dotw" field.
func DotwIn(vs ...Dotw) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldDotw, vs...))
}

// DotwNotIn applies the NotIn predicate on the "dotw" field.
func DotwNotIn(vs ...Dotw) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldDotw, vs...))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldEndTime, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCourse applies the HasEdge predicate on the "course" edge.
func HasCourse() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseWith applies the HasEdge predicate on the "course" edge with a given conditions (other predicates).
func HasCourseWith(preds ...predicate.Course) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := newCourseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MembersTable, MembersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.Class_User_Role) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := newMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassLabStatuses applies the HasEdge predicate on the "class_lab_statuses" edge.
func HasClassLabStatuses() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassLabStatusesTable, ClassLabStatusesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassLabStatusesWith applies the HasEdge predicate on the "class_lab_statuses" edge with a given conditions (other predicates).
func HasClassLabStatusesWith(preds ...predicate.Class_Lab_Status) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := newClassLabStatusesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
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
func Not(p predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		p(s.Not())
	})
}