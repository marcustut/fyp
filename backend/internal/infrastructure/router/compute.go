package router

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gorilla/mux"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
)

// NewCompute ...
func NewCompute(ctrl controller.Controller, cfg aws.Config) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	r.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	r.HandleFunc("/terminate", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")

	return r
}
