// Code generated by ent, DO NOT EDIT.

package credential

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/c2micro/c2msrv/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldDeletedAt, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldPassword, v))
}

// Realm applies equality check predicate on the "realm" field. It's identical to RealmEQ.
func Realm(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldRealm, v))
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldHost, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldNote, v))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v uint32) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldColor, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldDeletedAt))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldUsername))
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldUsername))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldPassword, v))
}

// RealmEQ applies the EQ predicate on the "realm" field.
func RealmEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldRealm, v))
}

// RealmNEQ applies the NEQ predicate on the "realm" field.
func RealmNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldRealm, v))
}

// RealmIn applies the In predicate on the "realm" field.
func RealmIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldRealm, vs...))
}

// RealmNotIn applies the NotIn predicate on the "realm" field.
func RealmNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldRealm, vs...))
}

// RealmGT applies the GT predicate on the "realm" field.
func RealmGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldRealm, v))
}

// RealmGTE applies the GTE predicate on the "realm" field.
func RealmGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldRealm, v))
}

// RealmLT applies the LT predicate on the "realm" field.
func RealmLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldRealm, v))
}

// RealmLTE applies the LTE predicate on the "realm" field.
func RealmLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldRealm, v))
}

// RealmContains applies the Contains predicate on the "realm" field.
func RealmContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldRealm, v))
}

// RealmHasPrefix applies the HasPrefix predicate on the "realm" field.
func RealmHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldRealm, v))
}

// RealmHasSuffix applies the HasSuffix predicate on the "realm" field.
func RealmHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldRealm, v))
}

// RealmIsNil applies the IsNil predicate on the "realm" field.
func RealmIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldRealm))
}

// RealmNotNil applies the NotNil predicate on the "realm" field.
func RealmNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldRealm))
}

// RealmEqualFold applies the EqualFold predicate on the "realm" field.
func RealmEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldRealm, v))
}

// RealmContainsFold applies the ContainsFold predicate on the "realm" field.
func RealmContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldRealm, v))
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldHost, v))
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldHost, v))
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldHost, vs...))
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldHost, vs...))
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldHost, v))
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldHost, v))
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldHost, v))
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldHost, v))
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldHost, v))
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldHost, v))
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldHost, v))
}

// HostIsNil applies the IsNil predicate on the "host" field.
func HostIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldHost))
}

// HostNotNil applies the NotNil predicate on the "host" field.
func HostNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldHost))
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldHost, v))
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldHost, v))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldNote, v))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v uint32) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v uint32) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...uint32) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...uint32) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v uint32) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v uint32) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v uint32) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v uint32) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldColor, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Credential) predicate.Credential {
	return predicate.Credential(sql.NotPredicates(p))
}