package gymPackage

import (
	"database/sql"
)

func ConfigPackageDi(dbConnection *sql.DB) PackageControllers {
	packageRepository := NewPackageRepository(dbConnection)
	packageServices := NewPackageServices(packageRepository)
	packageControllers := GetPackageControllers(packageServices)

	return packageControllers
}