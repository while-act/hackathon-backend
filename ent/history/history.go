// Code generated by ent, DO NOT EDIT.

package history

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the history type in the database.
	Label = "history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOrganizationalLegal holds the string denoting the organizational_legal field in the database.
	FieldOrganizationalLegal = "organizational_legal"
	// FieldIndustryBranch holds the string denoting the industry_branch field in the database.
	FieldIndustryBranch = "industry_branch"
	// FieldFullTimeEmployees holds the string denoting the full_time_employees field in the database.
	FieldFullTimeEmployees = "full_time_employees"
	// FieldAvgSalary holds the string denoting the avg_salary field in the database.
	FieldAvgSalary = "avg_salary"
	// FieldDistrictTitle holds the string denoting the district_title field in the database.
	FieldDistrictTitle = "district_title"
	// FieldLandArea holds the string denoting the land_area field in the database.
	FieldLandArea = "land_area"
	// FieldIsBuy holds the string denoting the is_buy field in the database.
	FieldIsBuy = "is_buy"
	// FieldConstructionFacilitiesArea holds the string denoting the construction_facilities_area field in the database.
	FieldConstructionFacilitiesArea = "construction_facilities_area"
	// FieldBuildingType holds the string denoting the building_type field in the database.
	FieldBuildingType = "building_type"
	// FieldEquipment holds the string denoting the equipment field in the database.
	FieldEquipment = "equipment"
	// FieldAccountingSupport holds the string denoting the accounting_support field in the database.
	FieldAccountingSupport = "accounting_support"
	// FieldTaxationSystemOperations holds the string denoting the taxation_system_operations field in the database.
	FieldTaxationSystemOperations = "taxation_system_operations"
	// FieldOperationsNum holds the string denoting the operations_num field in the database.
	FieldOperationsNum = "operations_num"
	// FieldPatentCalc holds the string denoting the patent_calc field in the database.
	FieldPatentCalc = "patent_calc"
	// FieldBusinessActivityID holds the string denoting the business_activity_id field in the database.
	FieldBusinessActivityID = "business_activity_id"
	// FieldOther holds the string denoting the other field in the database.
	FieldOther = "other"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeIndustry holds the string denoting the industry edge name in mutations.
	EdgeIndustry = "industry"
	// EdgeTaxationSystems holds the string denoting the taxation_systems edge name in mutations.
	EdgeTaxationSystems = "taxation_systems"
	// EdgeBusinessActivity holds the string denoting the business_activity edge name in mutations.
	EdgeBusinessActivity = "business_activity"
	// EdgeDistrict holds the string denoting the district edge name in mutations.
	EdgeDistrict = "district"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// IndustryFieldID holds the string denoting the ID field of the Industry.
	IndustryFieldID = "branch"
	// TaxationSystemFieldID holds the string denoting the ID field of the TaxationSystem.
	TaxationSystemFieldID = "operations"
	// DistrictFieldID holds the string denoting the ID field of the District.
	DistrictFieldID = "title"
	// Table holds the table name of the history in the database.
	Table = "histories"
	// IndustryTable is the table that holds the industry relation/edge.
	IndustryTable = "histories"
	// IndustryInverseTable is the table name for the Industry entity.
	// It exists in this package in order to avoid circular dependency with the "industry" package.
	IndustryInverseTable = "industries"
	// IndustryColumn is the table column denoting the industry relation/edge.
	IndustryColumn = "industry_branch"
	// TaxationSystemsTable is the table that holds the taxation_systems relation/edge.
	TaxationSystemsTable = "histories"
	// TaxationSystemsInverseTable is the table name for the TaxationSystem entity.
	// It exists in this package in order to avoid circular dependency with the "taxationsystem" package.
	TaxationSystemsInverseTable = "taxation_systems"
	// TaxationSystemsColumn is the table column denoting the taxation_systems relation/edge.
	TaxationSystemsColumn = "taxation_system_operations"
	// BusinessActivityTable is the table that holds the business_activity relation/edge.
	BusinessActivityTable = "histories"
	// BusinessActivityInverseTable is the table name for the BusinessActivity entity.
	// It exists in this package in order to avoid circular dependency with the "businessactivity" package.
	BusinessActivityInverseTable = "business_activities"
	// BusinessActivityColumn is the table column denoting the business_activity relation/edge.
	BusinessActivityColumn = "business_activity_id"
	// DistrictTable is the table that holds the district relation/edge.
	DistrictTable = "histories"
	// DistrictInverseTable is the table name for the District entity.
	// It exists in this package in order to avoid circular dependency with the "district" package.
	DistrictInverseTable = "districts"
	// DistrictColumn is the table column denoting the district relation/edge.
	DistrictColumn = "district_title"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "histories"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
)

// Columns holds all SQL columns for history fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldOrganizationalLegal,
	FieldIndustryBranch,
	FieldFullTimeEmployees,
	FieldAvgSalary,
	FieldDistrictTitle,
	FieldLandArea,
	FieldIsBuy,
	FieldConstructionFacilitiesArea,
	FieldBuildingType,
	FieldEquipment,
	FieldAccountingSupport,
	FieldTaxationSystemOperations,
	FieldOperationsNum,
	FieldPatentCalc,
	FieldBusinessActivityID,
	FieldOther,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// FullTimeEmployeesValidator is a validator for the "full_time_employees" field. It is called by the builders before save.
	FullTimeEmployeesValidator func(int) error
	// LandAreaValidator is a validator for the "land_area" field. It is called by the builders before save.
	LandAreaValidator func(float64) error
	// ConstructionFacilitiesAreaValidator is a validator for the "construction_facilities_area" field. It is called by the builders before save.
	ConstructionFacilitiesAreaValidator func(float64) error
)

// OrderOption defines the ordering options for the History queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOrganizationalLegal orders the results by the organizational_legal field.
func ByOrganizationalLegal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationalLegal, opts...).ToFunc()
}

// ByIndustryBranch orders the results by the industry_branch field.
func ByIndustryBranch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndustryBranch, opts...).ToFunc()
}

// ByFullTimeEmployees orders the results by the full_time_employees field.
func ByFullTimeEmployees(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullTimeEmployees, opts...).ToFunc()
}

// ByAvgSalary orders the results by the avg_salary field.
func ByAvgSalary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvgSalary, opts...).ToFunc()
}

// ByDistrictTitle orders the results by the district_title field.
func ByDistrictTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDistrictTitle, opts...).ToFunc()
}

// ByLandArea orders the results by the land_area field.
func ByLandArea(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLandArea, opts...).ToFunc()
}

// ByIsBuy orders the results by the is_buy field.
func ByIsBuy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsBuy, opts...).ToFunc()
}

// ByConstructionFacilitiesArea orders the results by the construction_facilities_area field.
func ByConstructionFacilitiesArea(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConstructionFacilitiesArea, opts...).ToFunc()
}

// ByBuildingType orders the results by the building_type field.
func ByBuildingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuildingType, opts...).ToFunc()
}

// ByAccountingSupport orders the results by the accounting_support field.
func ByAccountingSupport(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountingSupport, opts...).ToFunc()
}

// ByTaxationSystemOperations orders the results by the taxation_system_operations field.
func ByTaxationSystemOperations(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxationSystemOperations, opts...).ToFunc()
}

// ByOperationsNum orders the results by the operations_num field.
func ByOperationsNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperationsNum, opts...).ToFunc()
}

// ByPatentCalc orders the results by the patent_calc field.
func ByPatentCalc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPatentCalc, opts...).ToFunc()
}

// ByBusinessActivityID orders the results by the business_activity_id field.
func ByBusinessActivityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessActivityID, opts...).ToFunc()
}

// ByOther orders the results by the other field.
func ByOther(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOther, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByIndustryField orders the results by industry field.
func ByIndustryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIndustryStep(), sql.OrderByField(field, opts...))
	}
}

// ByTaxationSystemsField orders the results by taxation_systems field.
func ByTaxationSystemsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaxationSystemsStep(), sql.OrderByField(field, opts...))
	}
}

// ByBusinessActivityField orders the results by business_activity field.
func ByBusinessActivityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessActivityStep(), sql.OrderByField(field, opts...))
	}
}

// ByDistrictField orders the results by district field.
func ByDistrictField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDistrictStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}
func newIndustryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IndustryInverseTable, IndustryFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, IndustryTable, IndustryColumn),
	)
}
func newTaxationSystemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaxationSystemsInverseTable, TaxationSystemFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TaxationSystemsTable, TaxationSystemsColumn),
	)
}
func newBusinessActivityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessActivityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BusinessActivityTable, BusinessActivityColumn),
	)
}
func newDistrictStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DistrictInverseTable, DistrictFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DistrictTable, DistrictColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
