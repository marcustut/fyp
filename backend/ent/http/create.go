// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/slide"
	"github.com/marcustut/fyp/backend/ent/user"
	"go.uber.org/zap"
)

// Create creates a new ent.Slide and stores it in the database.
func (h SlideHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d SlideCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Slide.Create()
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.CreatedAt != nil {
		b.SetCreatedAt(*d.CreatedAt)
	}
	if d.UpdatedAt != nil {
		b.SetUpdatedAt(*d.UpdatedAt)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create slide", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Slide.Query().Where(slide.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.String("id", string(id)))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.String("id", string(id)))
			BadRequest(w, msg)
		default:
			l.Error("could not read slide", zap.Error(err), zap.String("id", string(id)))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("slide rendered", zap.String("id", string(id)))
	easyjson.MarshalToHTTPResponseWriter(NewSlide3844259445View(ret), w)
}

// Create creates a new ent.User and stores it in the database.
func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d UserCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.Username == nil {
		errs["username"] = `missing required field: "username"`
	} else if err := user.UsernameValidator(*d.Username); err != nil {
		errs["username"] = strings.TrimPrefix(err.Error(), "user: ")
	}
	if d.Email == nil {
		errs["email"] = `missing required field: "email"`
	} else if err := user.EmailValidator(*d.Email); err != nil {
		errs["email"] = strings.TrimPrefix(err.Error(), "user: ")
	}
	if d.FullName != nil {
		if err := user.FullNameValidator(*d.FullName); err != nil {
			errs["full_name"] = strings.TrimPrefix(err.Error(), "user: ")
		}
	}
	if d.AvatarURL != nil {
		if err := user.AvatarURLValidator(*d.AvatarURL); err != nil {
			errs["avatar_url"] = strings.TrimPrefix(err.Error(), "user: ")
		}
	}
	if d.Bio != nil {
		if err := user.BioValidator(*d.Bio); err != nil {
			errs["bio"] = strings.TrimPrefix(err.Error(), "user: ")
		}
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.User.Create()
	if d.Username != nil {
		b.SetUsername(*d.Username)
	}
	if d.Email != nil {
		b.SetEmail(*d.Email)
	}
	if d.FullName != nil {
		b.SetFullName(*d.FullName)
	}
	if d.AvatarURL != nil {
		b.SetAvatarURL(*d.AvatarURL)
	}
	if d.Bio != nil {
		b.SetBio(*d.Bio)
	}
	if d.CreatedAt != nil {
		b.SetCreatedAt(*d.CreatedAt)
	}
	if d.UpdatedAt != nil {
		b.SetUpdatedAt(*d.UpdatedAt)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create user", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.User.Query().Where(user.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read user", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("user rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewUser843294600View(ret), w)
}
