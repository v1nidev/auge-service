// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EmergencyContactsColumns holds the columns for the "emergency_contacts" table.
	EmergencyContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "phone_number", Type: field.TypeString, Size: 20},
		{Name: "contact_name", Type: field.TypeString},
		{Name: "member_emergencycontact", Type: field.TypeInt, Unique: true},
	}
	// EmergencyContactsTable holds the schema information for the "emergency_contacts" table.
	EmergencyContactsTable = &schema.Table{
		Name:       "emergency_contacts",
		Columns:    EmergencyContactsColumns,
		PrimaryKey: []*schema.Column{EmergencyContactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "emergency_contacts_members_emergencycontact",
				Columns:    []*schema.Column{EmergencyContactsColumns[3]},
				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MembersColumns holds the columns for the "members" table.
	MembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "sex", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Size: 255},
	}
	// MembersTable holds the schema information for the "members" table.
	MembersTable = &schema.Table{
		Name:       "members",
		Columns:    MembersColumns,
		PrimaryKey: []*schema.Column{MembersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EmergencyContactsTable,
		MembersTable,
	}
)

func init() {
	EmergencyContactsTable.ForeignKeys[0].RefTable = MembersTable
	MembersTable.Annotation = &entsql.Annotation{
		Check: "sex in ('M', 'F')",
	}
}