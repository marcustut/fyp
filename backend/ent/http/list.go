// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

// Read fetches the ent.Slide identified by a given url-parameter from the
// database and returns it to the client.
func (h *SlideHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Slide.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching slides from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("slides rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewSlide3844259445Views(es), w)
}

// Read fetches the ent.User identified by a given url-parameter from the
// database and returns it to the client.
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.User.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching users from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("users rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewUser843294600Views(es), w)
}
