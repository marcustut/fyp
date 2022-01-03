package router

import (
	"github.com/gorilla/mux"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"net/http"
)

func NewSlideRouter(controller controller.Controller) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/slides/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.Slide.Get(r.Context())
	}).Methods("GET")

	return r
}
