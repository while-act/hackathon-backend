// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/while-act/hackathon-backend/ent/businessactivity"
	"github.com/while-act/hackathon-backend/ent/district"
	"github.com/while-act/hackathon-backend/ent/history"
	"github.com/while-act/hackathon-backend/ent/industry"
	"github.com/while-act/hackathon-backend/ent/taxationsystem"
	"github.com/while-act/hackathon-backend/ent/user"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// OrganizationalLegal holds the value of the "organizational_legal" field.
	OrganizationalLegal string `json:"organizational_legal,omitempty"`
	// IndustryBranch holds the value of the "industry_branch" field.
	IndustryBranch string `json:"industry_branch,omitempty"`
	// FullTimeEmployees holds the value of the "full_time_employees" field.
	FullTimeEmployees int `json:"full_time_employees,omitempty"`
	// AvgSalary holds the value of the "avg_salary" field.
	AvgSalary float64 `json:"avg_salary,omitempty"`
	// DistrictTitle holds the value of the "district_title" field.
	DistrictTitle string `json:"district_title,omitempty"`
	// LandArea holds the value of the "land_area" field.
	LandArea float64 `json:"land_area,omitempty"`
	// IsBuy holds the value of the "is_buy" field.
	IsBuy bool `json:"is_buy,omitempty"`
	// ConstructionFacilitiesArea holds the value of the "construction_facilities_area" field.
	ConstructionFacilitiesArea float64 `json:"construction_facilities_area,omitempty"`
	// BuildingType holds the value of the "building_type" field.
	BuildingType []string `json:"building_type,omitempty"`
	// Equipment holds the value of the "equipment" field.
	Equipment []dto.Equipment `json:"equipment,omitempty"`
	// AccountingSupport holds the value of the "accounting_support" field.
	AccountingSupport bool `json:"accounting_support,omitempty"`
	// TaxationSystemOperations holds the value of the "taxation_system_operations" field.
	TaxationSystemOperations int `json:"taxation_system_operations,omitempty"`
	// OperationType holds the value of the "operation_type" field.
	OperationType string `json:"operation_type,omitempty"`
	// PatentCalc holds the value of the "patent_calc" field.
	PatentCalc bool `json:"patent_calc,omitempty"`
	// BusinessActivityID holds the value of the "business_activity_id" field.
	BusinessActivityID int `json:"business_activity_id,omitempty"`
	// Other holds the value of the "other" field.
	Other string `json:"other,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistoryQuery when eager-loading is set.
	Edges        HistoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HistoryEdges holds the relations/edges for other nodes in the graph.
type HistoryEdges struct {
	// Industry holds the value of the industry edge.
	Industry *Industry `json:"industry,omitempty"`
	// TaxationSystems holds the value of the taxation_systems edge.
	TaxationSystems *TaxationSystem `json:"taxation_systems,omitempty"`
	// BusinessActivity holds the value of the business_activity edge.
	BusinessActivity *BusinessActivity `json:"business_activity,omitempty"`
	// District holds the value of the district edge.
	District *District `json:"district,omitempty"`
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// IndustryOrErr returns the Industry value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) IndustryOrErr() (*Industry, error) {
	if e.loadedTypes[0] {
		if e.Industry == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: industry.Label}
		}
		return e.Industry, nil
	}
	return nil, &NotLoadedError{edge: "industry"}
}

// TaxationSystemsOrErr returns the TaxationSystems value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) TaxationSystemsOrErr() (*TaxationSystem, error) {
	if e.loadedTypes[1] {
		if e.TaxationSystems == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: taxationsystem.Label}
		}
		return e.TaxationSystems, nil
	}
	return nil, &NotLoadedError{edge: "taxation_systems"}
}

// BusinessActivityOrErr returns the BusinessActivity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) BusinessActivityOrErr() (*BusinessActivity, error) {
	if e.loadedTypes[2] {
		if e.BusinessActivity == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: businessactivity.Label}
		}
		return e.BusinessActivity, nil
	}
	return nil, &NotLoadedError{edge: "business_activity"}
}

// DistrictOrErr returns the District value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) DistrictOrErr() (*District, error) {
	if e.loadedTypes[3] {
		if e.District == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: district.Label}
		}
		return e.District, nil
	}
	return nil, &NotLoadedError{edge: "district"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[4] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case history.FieldBuildingType, history.FieldEquipment:
			values[i] = new([]byte)
		case history.FieldIsBuy, history.FieldAccountingSupport, history.FieldPatentCalc:
			values[i] = new(sql.NullBool)
		case history.FieldAvgSalary, history.FieldLandArea, history.FieldConstructionFacilitiesArea:
			values[i] = new(sql.NullFloat64)
		case history.FieldID, history.FieldFullTimeEmployees, history.FieldTaxationSystemOperations, history.FieldBusinessActivityID, history.FieldUserID:
			values[i] = new(sql.NullInt64)
		case history.FieldName, history.FieldOrganizationalLegal, history.FieldIndustryBranch, history.FieldDistrictTitle, history.FieldOperationType, history.FieldOther:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case history.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = value.String
			}
		case history.FieldOrganizationalLegal:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organizational_legal", values[i])
			} else if value.Valid {
				h.OrganizationalLegal = value.String
			}
		case history.FieldIndustryBranch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field industry_branch", values[i])
			} else if value.Valid {
				h.IndustryBranch = value.String
			}
		case history.FieldFullTimeEmployees:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field full_time_employees", values[i])
			} else if value.Valid {
				h.FullTimeEmployees = int(value.Int64)
			}
		case history.FieldAvgSalary:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field avg_salary", values[i])
			} else if value.Valid {
				h.AvgSalary = value.Float64
			}
		case history.FieldDistrictTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field district_title", values[i])
			} else if value.Valid {
				h.DistrictTitle = value.String
			}
		case history.FieldLandArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field land_area", values[i])
			} else if value.Valid {
				h.LandArea = value.Float64
			}
		case history.FieldIsBuy:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_buy", values[i])
			} else if value.Valid {
				h.IsBuy = value.Bool
			}
		case history.FieldConstructionFacilitiesArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field construction_facilities_area", values[i])
			} else if value.Valid {
				h.ConstructionFacilitiesArea = value.Float64
			}
		case history.FieldBuildingType:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field building_type", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.BuildingType); err != nil {
					return fmt.Errorf("unmarshal field building_type: %w", err)
				}
			}
		case history.FieldEquipment:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field equipment", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.Equipment); err != nil {
					return fmt.Errorf("unmarshal field equipment: %w", err)
				}
			}
		case history.FieldAccountingSupport:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field accounting_support", values[i])
			} else if value.Valid {
				h.AccountingSupport = value.Bool
			}
		case history.FieldTaxationSystemOperations:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field taxation_system_operations", values[i])
			} else if value.Valid {
				h.TaxationSystemOperations = int(value.Int64)
			}
		case history.FieldOperationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation_type", values[i])
			} else if value.Valid {
				h.OperationType = value.String
			}
		case history.FieldPatentCalc:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field patent_calc", values[i])
			} else if value.Valid {
				h.PatentCalc = value.Bool
			}
		case history.FieldBusinessActivityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field business_activity_id", values[i])
			} else if value.Valid {
				h.BusinessActivityID = int(value.Int64)
			}
		case history.FieldOther:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other", values[i])
			} else if value.Valid {
				h.Other = value.String
			}
		case history.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				h.UserID = int(value.Int64)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the History.
// This includes values selected through modifiers, order, etc.
func (h *History) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryIndustry queries the "industry" edge of the History entity.
func (h *History) QueryIndustry() *IndustryQuery {
	return NewHistoryClient(h.config).QueryIndustry(h)
}

// QueryTaxationSystems queries the "taxation_systems" edge of the History entity.
func (h *History) QueryTaxationSystems() *TaxationSystemQuery {
	return NewHistoryClient(h.config).QueryTaxationSystems(h)
}

// QueryBusinessActivity queries the "business_activity" edge of the History entity.
func (h *History) QueryBusinessActivity() *BusinessActivityQuery {
	return NewHistoryClient(h.config).QueryBusinessActivity(h)
}

// QueryDistrict queries the "district" edge of the History entity.
func (h *History) QueryDistrict() *DistrictQuery {
	return NewHistoryClient(h.config).QueryDistrict(h)
}

// QueryUsers queries the "users" edge of the History entity.
func (h *History) QueryUsers() *UserQuery {
	return NewHistoryClient(h.config).QueryUsers(h)
}

// Update returns a builder for updating this History.
// Note that you need to call History.Unwrap() before calling this method if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return NewHistoryClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the History entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("name=")
	builder.WriteString(h.Name)
	builder.WriteString(", ")
	builder.WriteString("organizational_legal=")
	builder.WriteString(h.OrganizationalLegal)
	builder.WriteString(", ")
	builder.WriteString("industry_branch=")
	builder.WriteString(h.IndustryBranch)
	builder.WriteString(", ")
	builder.WriteString("full_time_employees=")
	builder.WriteString(fmt.Sprintf("%v", h.FullTimeEmployees))
	builder.WriteString(", ")
	builder.WriteString("avg_salary=")
	builder.WriteString(fmt.Sprintf("%v", h.AvgSalary))
	builder.WriteString(", ")
	builder.WriteString("district_title=")
	builder.WriteString(h.DistrictTitle)
	builder.WriteString(", ")
	builder.WriteString("land_area=")
	builder.WriteString(fmt.Sprintf("%v", h.LandArea))
	builder.WriteString(", ")
	builder.WriteString("is_buy=")
	builder.WriteString(fmt.Sprintf("%v", h.IsBuy))
	builder.WriteString(", ")
	builder.WriteString("construction_facilities_area=")
	builder.WriteString(fmt.Sprintf("%v", h.ConstructionFacilitiesArea))
	builder.WriteString(", ")
	builder.WriteString("building_type=")
	builder.WriteString(fmt.Sprintf("%v", h.BuildingType))
	builder.WriteString(", ")
	builder.WriteString("equipment=")
	builder.WriteString(fmt.Sprintf("%v", h.Equipment))
	builder.WriteString(", ")
	builder.WriteString("accounting_support=")
	builder.WriteString(fmt.Sprintf("%v", h.AccountingSupport))
	builder.WriteString(", ")
	builder.WriteString("taxation_system_operations=")
	builder.WriteString(fmt.Sprintf("%v", h.TaxationSystemOperations))
	builder.WriteString(", ")
	builder.WriteString("operation_type=")
	builder.WriteString(h.OperationType)
	builder.WriteString(", ")
	builder.WriteString("patent_calc=")
	builder.WriteString(fmt.Sprintf("%v", h.PatentCalc))
	builder.WriteString(", ")
	builder.WriteString("business_activity_id=")
	builder.WriteString(fmt.Sprintf("%v", h.BusinessActivityID))
	builder.WriteString(", ")
	builder.WriteString("other=")
	builder.WriteString(h.Other)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", h.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History
