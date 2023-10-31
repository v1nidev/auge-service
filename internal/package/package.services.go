package gymPackage

type PackageServices struct {
	getAllPackages func()
	getAPackage    func(id int)
}

func NewPackageServices(packageRepository PackageRepository) PackageServices {
	getAllPackages := func() {
		// business logic
		packageRepository.getAllPackages()
	}

	getAPackage := func(id int) {
	}

	return PackageServices{
		getAllPackages,
		getAPackage,
	}
}