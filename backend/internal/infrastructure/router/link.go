package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
)

// NewLink ...
func NewLink(ctrl controller.Controller) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/{link_id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		linkID := params["link_id"]

		// query links db to find if such link_id exists
		lc, err := ctrl.Link.List(r.Context(), nil, nil, nil, nil, &ent.LinkWhereInput{
			LinkID: &linkID,
		}, nil)
		if err != nil || lc.TotalCount == 0 {
			w.Header().Set("Location", config.C.Services.Link.ErrorReturnURL)
			w.WriteHeader(http.StatusFound)
			return
		}

		// redirect to original_url
		w.Header().Set("Location", fmt.Sprint(lc.Edges[0].Node.OriginalURL))
		w.WriteHeader(http.StatusFound)
	}).Methods("GET")

	return r
}
