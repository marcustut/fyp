// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/marcustut/fyp/backend/ent/schema/ulid"
)

// CreateSlideInput represents a mutation input for creating slides.
type CreateSlideInput struct {
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// Mutate applies the CreateSlideInput on the SlideCreate builder.
func (i *CreateSlideInput) Mutate(m *SlideCreate) {
	m.SetName(i.Name)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateSlideInput on the create builder.
func (c *SlideCreate) SetInput(i CreateSlideInput) *SlideCreate {
	i.Mutate(c)
	return c
}

// UpdateSlideInput represents a mutation input for updating slides.
type UpdateSlideInput struct {
	ID   ulid.ID
	Name *string
}

// Mutate applies the UpdateSlideInput on the SlideMutation.
func (i *UpdateSlideInput) Mutate(m *SlideMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateSlideInput on the update builder.
func (u *SlideUpdate) SetInput(i UpdateSlideInput) *SlideUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateSlideInput on the update-one builder.
func (u *SlideUpdateOne) SetInput(i UpdateSlideInput) *SlideUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Username     string
	Email        string
	FullName     *string
	PasswordHash string
	AvatarURL    *string
	Bio          *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

// Mutate applies the CreateUserInput on the UserCreate builder.
func (i *CreateUserInput) Mutate(m *UserCreate) {
	m.SetUsername(i.Username)
	m.SetEmail(i.Email)
	if v := i.FullName; v != nil {
		m.SetFullName(*v)
	}
	m.SetPasswordHash(i.PasswordHash)
	if v := i.AvatarURL; v != nil {
		m.SetAvatarURL(*v)
	}
	if v := i.Bio; v != nil {
		m.SetBio(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateUserInput on the create builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c)
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	ID             ulid.ID
	Username       *string
	FullName       *string
	ClearFullName  bool
	AvatarURL      *string
	ClearAvatarURL bool
	Bio            *string
	ClearBio       bool
}

// Mutate applies the UpdateUserInput on the UserMutation.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if i.ClearFullName {
		m.ClearFullName()
	}
	if v := i.FullName; v != nil {
		m.SetFullName(*v)
	}
	if i.ClearAvatarURL {
		m.ClearAvatarURL()
	}
	if v := i.AvatarURL; v != nil {
		m.SetAvatarURL(*v)
	}
	if i.ClearBio {
		m.ClearBio()
	}
	if v := i.Bio; v != nil {
		m.SetBio(*v)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the update builder.
func (u *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateUserInput on the update-one builder.
func (u *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
