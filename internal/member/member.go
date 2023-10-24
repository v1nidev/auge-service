package member

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Member struct {
	Id		string `json:"id"`
	Name 	string `json:"name"`
	Sex 	string `json:"sex"`
	Email string `json:"email"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{memberId}", GetAMember)
	return router
}

func GetAMember(writer http.ResponseWriter, req *http.Request) {
	memberId := chi.URLParam(req, "memberId")
	member := Member{
		Id: memberId,
		Name: "Vinicius",
		Sex: "M",
		Email: "a@x.com",
	}
	render.JSON(writer, req, member)
}