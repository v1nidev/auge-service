//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var EmergencyPhone = newEmergencyPhoneTable("public", "emergency_phone", "")

type emergencyPhoneTable struct {
	postgres.Table

	// Columns
	PhoneNumber postgres.ColumnString
	Name        postgres.ColumnString
	IDMember    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EmergencyPhoneTable struct {
	emergencyPhoneTable

	EXCLUDED emergencyPhoneTable
}

// AS creates new EmergencyPhoneTable with assigned alias
func (a EmergencyPhoneTable) AS(alias string) *EmergencyPhoneTable {
	return newEmergencyPhoneTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EmergencyPhoneTable with assigned schema name
func (a EmergencyPhoneTable) FromSchema(schemaName string) *EmergencyPhoneTable {
	return newEmergencyPhoneTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EmergencyPhoneTable with assigned table prefix
func (a EmergencyPhoneTable) WithPrefix(prefix string) *EmergencyPhoneTable {
	return newEmergencyPhoneTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EmergencyPhoneTable with assigned table suffix
func (a EmergencyPhoneTable) WithSuffix(suffix string) *EmergencyPhoneTable {
	return newEmergencyPhoneTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEmergencyPhoneTable(schemaName, tableName, alias string) *EmergencyPhoneTable {
	return &EmergencyPhoneTable{
		emergencyPhoneTable: newEmergencyPhoneTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newEmergencyPhoneTableImpl("", "excluded", ""),
	}
}

func newEmergencyPhoneTableImpl(schemaName, tableName, alias string) emergencyPhoneTable {
	var (
		PhoneNumberColumn = postgres.StringColumn("phone_number")
		NameColumn        = postgres.StringColumn("name")
		IDMemberColumn    = postgres.StringColumn("id_member")
		allColumns        = postgres.ColumnList{PhoneNumberColumn, NameColumn, IDMemberColumn}
		mutableColumns    = postgres.ColumnList{PhoneNumberColumn, NameColumn}
	)

	return emergencyPhoneTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		PhoneNumber: PhoneNumberColumn,
		Name:        NameColumn,
		IDMember:    IDMemberColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}