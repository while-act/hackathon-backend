// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wtkeqrf0/while.act/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// PasswordHash applies equality check predicate on the "password_hash" field. It's identical to PasswordHashEQ.
func PasswordHash(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordHash, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// CompanyInn applies equality check predicate on the "company_inn" field. It's identical to CompanyInnEQ.
func CompanyInn(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCompanyInn, v))
}

// FatherName applies equality check predicate on the "father_name" field. It's identical to FatherNameEQ.
func FatherName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFatherName, v))
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPosition, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountry, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCity, v))
}

// Biography applies equality check predicate on the "biography" field. It's identical to BiographyEQ.
func Biography(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBiography, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdateTime, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRole, v))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRole, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// PasswordHashEQ applies the EQ predicate on the "password_hash" field.
func PasswordHashEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordHash, v))
}

// PasswordHashNEQ applies the NEQ predicate on the "password_hash" field.
func PasswordHashNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswordHash, v))
}

// PasswordHashIn applies the In predicate on the "password_hash" field.
func PasswordHashIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswordHash, vs...))
}

// PasswordHashNotIn applies the NotIn predicate on the "password_hash" field.
func PasswordHashNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswordHash, vs...))
}

// PasswordHashGT applies the GT predicate on the "password_hash" field.
func PasswordHashGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswordHash, v))
}

// PasswordHashGTE applies the GTE predicate on the "password_hash" field.
func PasswordHashGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswordHash, v))
}

// PasswordHashLT applies the LT predicate on the "password_hash" field.
func PasswordHashLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswordHash, v))
}

// PasswordHashLTE applies the LTE predicate on the "password_hash" field.
func PasswordHashLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswordHash, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldLastName, v))
}

// CompanyInnEQ applies the EQ predicate on the "company_inn" field.
func CompanyInnEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCompanyInn, v))
}

// CompanyInnNEQ applies the NEQ predicate on the "company_inn" field.
func CompanyInnNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCompanyInn, v))
}

// CompanyInnIn applies the In predicate on the "company_inn" field.
func CompanyInnIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCompanyInn, vs...))
}

// CompanyInnNotIn applies the NotIn predicate on the "company_inn" field.
func CompanyInnNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCompanyInn, vs...))
}

// CompanyInnGT applies the GT predicate on the "company_inn" field.
func CompanyInnGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCompanyInn, v))
}

// CompanyInnGTE applies the GTE predicate on the "company_inn" field.
func CompanyInnGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCompanyInn, v))
}

// CompanyInnLT applies the LT predicate on the "company_inn" field.
func CompanyInnLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCompanyInn, v))
}

// CompanyInnLTE applies the LTE predicate on the "company_inn" field.
func CompanyInnLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCompanyInn, v))
}

// CompanyInnContains applies the Contains predicate on the "company_inn" field.
func CompanyInnContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCompanyInn, v))
}

// CompanyInnHasPrefix applies the HasPrefix predicate on the "company_inn" field.
func CompanyInnHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCompanyInn, v))
}

// CompanyInnHasSuffix applies the HasSuffix predicate on the "company_inn" field.
func CompanyInnHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCompanyInn, v))
}

// CompanyInnEqualFold applies the EqualFold predicate on the "company_inn" field.
func CompanyInnEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCompanyInn, v))
}

// CompanyInnContainsFold applies the ContainsFold predicate on the "company_inn" field.
func CompanyInnContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCompanyInn, v))
}

// FatherNameEQ applies the EQ predicate on the "father_name" field.
func FatherNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFatherName, v))
}

// FatherNameNEQ applies the NEQ predicate on the "father_name" field.
func FatherNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFatherName, v))
}

// FatherNameIn applies the In predicate on the "father_name" field.
func FatherNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFatherName, vs...))
}

// FatherNameNotIn applies the NotIn predicate on the "father_name" field.
func FatherNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFatherName, vs...))
}

// FatherNameGT applies the GT predicate on the "father_name" field.
func FatherNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFatherName, v))
}

// FatherNameGTE applies the GTE predicate on the "father_name" field.
func FatherNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFatherName, v))
}

// FatherNameLT applies the LT predicate on the "father_name" field.
func FatherNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFatherName, v))
}

// FatherNameLTE applies the LTE predicate on the "father_name" field.
func FatherNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFatherName, v))
}

// FatherNameContains applies the Contains predicate on the "father_name" field.
func FatherNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFatherName, v))
}

// FatherNameHasPrefix applies the HasPrefix predicate on the "father_name" field.
func FatherNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFatherName, v))
}

// FatherNameHasSuffix applies the HasSuffix predicate on the "father_name" field.
func FatherNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFatherName, v))
}

// FatherNameIsNil applies the IsNil predicate on the "father_name" field.
func FatherNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldFatherName))
}

// FatherNameNotNil applies the NotNil predicate on the "father_name" field.
func FatherNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldFatherName))
}

// FatherNameEqualFold applies the EqualFold predicate on the "father_name" field.
func FatherNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFatherName, v))
}

// FatherNameContainsFold applies the ContainsFold predicate on the "father_name" field.
func FatherNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFatherName, v))
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPosition, v))
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPosition, v))
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPosition, vs...))
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPosition, vs...))
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPosition, v))
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPosition, v))
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPosition, v))
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPosition, v))
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPosition, v))
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPosition, v))
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPosition, v))
}

// PositionIsNil applies the IsNil predicate on the "position" field.
func PositionIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPosition))
}

// PositionNotNil applies the NotNil predicate on the "position" field.
func PositionNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPosition))
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPosition, v))
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPosition, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCountry))
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCountry))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCountry, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCity, v))
}

// CityIsNil applies the IsNil predicate on the "city" field.
func CityIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCity))
}

// CityNotNil applies the NotNil predicate on the "city" field.
func CityNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCity))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCity, v))
}

// BiographyEQ applies the EQ predicate on the "biography" field.
func BiographyEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBiography, v))
}

// BiographyNEQ applies the NEQ predicate on the "biography" field.
func BiographyNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBiography, v))
}

// BiographyIn applies the In predicate on the "biography" field.
func BiographyIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldBiography, vs...))
}

// BiographyNotIn applies the NotIn predicate on the "biography" field.
func BiographyNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBiography, vs...))
}

// BiographyGT applies the GT predicate on the "biography" field.
func BiographyGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldBiography, v))
}

// BiographyGTE applies the GTE predicate on the "biography" field.
func BiographyGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBiography, v))
}

// BiographyLT applies the LT predicate on the "biography" field.
func BiographyLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldBiography, v))
}

// BiographyLTE applies the LTE predicate on the "biography" field.
func BiographyLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBiography, v))
}

// BiographyContains applies the Contains predicate on the "biography" field.
func BiographyContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldBiography, v))
}

// BiographyHasPrefix applies the HasPrefix predicate on the "biography" field.
func BiographyHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldBiography, v))
}

// BiographyHasSuffix applies the HasSuffix predicate on the "biography" field.
func BiographyHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldBiography, v))
}

// BiographyIsNil applies the IsNil predicate on the "biography" field.
func BiographyIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBiography))
}

// BiographyNotNil applies the NotNil predicate on the "biography" field.
func BiographyNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBiography))
}

// BiographyEqualFold applies the EqualFold predicate on the "biography" field.
func BiographyEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldBiography, v))
}

// BiographyContainsFold applies the ContainsFold predicate on the "biography" field.
func BiographyContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldBiography, v))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
