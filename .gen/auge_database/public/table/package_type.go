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

var PackageType = newPackageTypeTable("public", "package_type", "")

type packageTypeTable struct {
	postgres.Table

	// Columns
	ID                     postgres.ColumnFloat
	PackageTypeDescription postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PackageTypeTable struct {
	packageTypeTable

	EXCLUDED packageTypeTable
}

// AS creates new PackageTypeTable with assigned alias
func (a PackageTypeTable) AS(alias string) *PackageTypeTable {
	return newPackageTypeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PackageTypeTable with assigned schema name
func (a PackageTypeTable) FromSchema(schemaName string) *PackageTypeTable {
	return newPackageTypeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PackageTypeTable with assigned table prefix
func (a PackageTypeTable) WithPrefix(prefix string) *PackageTypeTable {
	return newPackageTypeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PackageTypeTable with assigned table suffix
func (a PackageTypeTable) WithSuffix(suffix string) *PackageTypeTable {
	return newPackageTypeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPackageTypeTable(schemaName, tableName, alias string) *PackageTypeTable {
	return &PackageTypeTable{
		packageTypeTable: newPackageTypeTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newPackageTypeTableImpl("", "excluded", ""),
	}
}

func newPackageTypeTableImpl(schemaName, tableName, alias string) packageTypeTable {
	var (
		IDColumn                     = postgres.FloatColumn("id")
		PackageTypeDescriptionColumn = postgres.StringColumn("package_type_description")
		allColumns                   = postgres.ColumnList{IDColumn, PackageTypeDescriptionColumn}
		mutableColumns               = postgres.ColumnList{PackageTypeDescriptionColumn}
	)

	return packageTypeTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                     IDColumn,
		PackageTypeDescription: PackageTypeDescriptionColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}