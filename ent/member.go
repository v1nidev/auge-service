// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/v1nidev/auge-service/ent/emergencycontact"
	"github.com/v1nidev/auge-service/ent/member"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Sex holds the value of the "sex" field.
	Sex string `json:"sex,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberQuery when eager-loading is set.
	Edges        MemberEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MemberEdges holds the relations/edges for other nodes in the graph.
type MemberEdges struct {
	// Emergencycontact holds the value of the emergencycontact edge.
	Emergencycontact *EmergencyContact `json:"emergencycontact,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmergencycontactOrErr returns the Emergencycontact value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberEdges) EmergencycontactOrErr() (*EmergencyContact, error) {
	if e.loadedTypes[0] {
		if e.Emergencycontact == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: emergencycontact.Label}
		}
		return e.Emergencycontact, nil
	}
	return nil, &NotLoadedError{edge: "emergencycontact"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			values[i] = new(sql.NullInt64)
		case member.FieldName, member.FieldSex, member.FieldEmail:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case member.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case member.FieldSex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sex", values[i])
			} else if value.Valid {
				m.Sex = value.String
			}
		case member.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				m.Email = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Member.
// This includes values selected through modifiers, order, etc.
func (m *Member) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryEmergencycontact queries the "emergencycontact" edge of the Member entity.
func (m *Member) QueryEmergencycontact() *EmergencyContactQuery {
	return NewMemberClient(m.config).QueryEmergencycontact(m)
}

// Update returns a builder for updating this Member.
// Note that you need to call Member.Unwrap() before calling this method if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return NewMemberClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Member entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("sex=")
	builder.WriteString(m.Sex)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(m.Email)
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member
