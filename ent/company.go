// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wtkeqrf0/while.act/ent/company"
	"github.com/wtkeqrf0/while.act/ent/user"
)

// Company is the model entity for the Company schema.
type Company struct {
	config `example:"-" json:"-"`
	// ID of the ent.
	ID string `json:"inn,omitempty" example:"7707083893"`
	// Name holds the value of the "name" field.
	Name *string `json:"company_name,omitempty" example:"ООО \"Парк\""`
	// Website holds the value of the "website" field.
	Website *string `json:"website,omitempty" example:"https://www.rusprofile.ru"`
	// EconomicActivityBranch holds the value of the "economic_activity_branch" field.
	EconomicActivityBranch string `json:"economicActivityBranch,omitempty" example:"Радиоэлектроника и приборостроение - Приборостроение"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyQuery when eager-loading is set.
	Edges        CompanyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CompanyEdges holds the relations/edges for other nodes in the graph.
type CompanyEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompanyEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldID, company.FieldName, company.FieldWebsite, company.FieldEconomicActivityBranch:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case company.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = new(string)
				*c.Name = value.String
			}
		case company.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				c.Website = new(string)
				*c.Website = value.String
			}
		case company.FieldEconomicActivityBranch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field economic_activity_branch", values[i])
			} else if value.Valid {
				c.EconomicActivityBranch = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Company.
// This includes values selected through modifiers, order, etc.
func (c *Company) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Company entity.
func (c *Company) QueryUsers() *UserQuery {
	return NewCompanyClient(c.config).QueryUsers(c)
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return NewCompanyClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Company is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	if v := c.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.Website; v != nil {
		builder.WriteString("website=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("economic_activity_branch=")
	builder.WriteString(c.EconomicActivityBranch)
	builder.WriteByte(')')
	return builder.String()
}

// Companies is a parsable slice of Company.
type Companies []*Company