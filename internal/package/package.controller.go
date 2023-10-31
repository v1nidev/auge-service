package gymPackage

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	// "github.com/v1nidev/auge-service/internal/package"
)

type PackagePayload struct {
	Id						string `json:"id"`
	Description 	string `json:"description"`
}

func Routes(packageControllers PackageControllers) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{packageId}", packageControllers.getPackages)
	return router
}

type PackageControllers struct {
	getPackages http.HandlerFunc
}
func GetPackageControllers(packageServices PackageServices) PackageControllers {
	getPackages := func (writer http.ResponseWriter, req *http.Request) {
		packageId := chi.URLParam(req, "packageId")

		packageServices.getAllPackages()

		pkg := PackagePayload{
			Id: packageId,
			Description: "FitFake",
		}
		render.JSON(writer, req, pkg)
	}

	return PackageControllers{
		getPackages,
	}
}
