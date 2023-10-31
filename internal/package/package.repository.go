package gymPackage

import (
	"fmt"
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"

	// dot import so go code would resemble as much as native SQL
	// dot import is not mandatory
	. "github.com/v1nidev/auge-service/.gen/auge_database/public/table"
	"github.com/v1nidev/auge-service/.gen/auge_database/public/model"
)

type PackageRepository struct {
	getAllPackages func()
}
func NewPackageRepository(db *sql.DB) PackageRepository {
	getAllPackages := func()  {
		statement := SELECT(
			PackageType.ID,
			PackageType.PackageTypeDescription,
		).FROM(PackageType)

		// query, _ := statement.Sql()
		// debugSql := statement.DebugSql()

		var dest []model.PackageType

		// fmt.Printf("query:\n")
		// fmt.Printf("%+v\n", query)

		err := statement.Query(db, &dest)
		fmt.Printf("Err:\n")
		fmt.Printf("%+v\n", err)
		fmt.Printf("Package Types:\n")
		fmt.Printf("%+v\n", dest)

		// return dest
	}

	return PackageRepository{
		getAllPackages,
	}
}
