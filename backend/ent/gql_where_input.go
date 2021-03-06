// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/marcustut/fyp/backend/ent/predicate"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/ent/slide"
	"github.com/marcustut/fyp/backend/ent/user"
)

// SlideWhereInput represents a where input for filtering Slide queries.
type SlideWhereInput struct {
	Not *SlideWhereInput   `json:"not,omitempty"`
	Or  []*SlideWhereInput `json:"or,omitempty"`
	And []*SlideWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *ulid.ID  `json:"id,omitempty"`
	IDNEQ   *ulid.ID  `json:"idNEQ,omitempty"`
	IDIn    []ulid.ID `json:"idIn,omitempty"`
	IDNotIn []ulid.ID `json:"idNotIn,omitempty"`
	IDGT    *ulid.ID  `json:"idGT,omitempty"`
	IDGTE   *ulid.ID  `json:"idGTE,omitempty"`
	IDLT    *ulid.ID  `json:"idLT,omitempty"`
	IDLTE   *ulid.ID  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`
}

// Filter applies the SlideWhereInput filter on the SlideQuery builder.
func (i *SlideWhereInput) Filter(q *SlideQuery) (*SlideQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering slides.
// An error is returned if the input is empty or invalid.
func (i *SlideWhereInput) P() (predicate.Slide, error) {
	var predicates []predicate.Slide
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, slide.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Slide, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, slide.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Slide, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, slide.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, slide.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, slide.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, slide.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, slide.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, slide.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, slide.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, slide.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, slide.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, slide.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, slide.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, slide.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, slide.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, slide.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, slide.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, slide.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, slide.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, slide.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, slide.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, slide.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, slide.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, slide.NameContainsFold(*i.NameContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, slide.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, slide.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, slide.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, slide.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, slide.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, slide.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, slide.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, slide.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, slide.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, slide.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, slide.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, slide.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, slide.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, slide.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, slide.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, slide.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/marcustut/fyp/backend/ent: empty predicate SlideWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return slide.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *ulid.ID  `json:"id,omitempty"`
	IDNEQ   *ulid.ID  `json:"idNEQ,omitempty"`
	IDIn    []ulid.ID `json:"idIn,omitempty"`
	IDNotIn []ulid.ID `json:"idNotIn,omitempty"`
	IDGT    *ulid.ID  `json:"idGT,omitempty"`
	IDGTE   *ulid.ID  `json:"idGTE,omitempty"`
	IDLT    *ulid.ID  `json:"idLT,omitempty"`
	IDLTE   *ulid.ID  `json:"idLTE,omitempty"`

	// "username" field predicates.
	Username             *string  `json:"username,omitempty"`
	UsernameNEQ          *string  `json:"usernameNEQ,omitempty"`
	UsernameIn           []string `json:"usernameIn,omitempty"`
	UsernameNotIn        []string `json:"usernameNotIn,omitempty"`
	UsernameGT           *string  `json:"usernameGT,omitempty"`
	UsernameGTE          *string  `json:"usernameGTE,omitempty"`
	UsernameLT           *string  `json:"usernameLT,omitempty"`
	UsernameLTE          *string  `json:"usernameLTE,omitempty"`
	UsernameContains     *string  `json:"usernameContains,omitempty"`
	UsernameHasPrefix    *string  `json:"usernameHasPrefix,omitempty"`
	UsernameHasSuffix    *string  `json:"usernameHasSuffix,omitempty"`
	UsernameEqualFold    *string  `json:"usernameEqualFold,omitempty"`
	UsernameContainsFold *string  `json:"usernameContainsFold,omitempty"`

	// "email" field predicates.
	Email             *string  `json:"email,omitempty"`
	EmailNEQ          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGT           *string  `json:"emailGT,omitempty"`
	EmailGTE          *string  `json:"emailGTE,omitempty"`
	EmailLT           *string  `json:"emailLT,omitempty"`
	EmailLTE          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`

	// "full_name" field predicates.
	FullName             *string  `json:"fullName,omitempty"`
	FullNameNEQ          *string  `json:"fullNameNEQ,omitempty"`
	FullNameIn           []string `json:"fullNameIn,omitempty"`
	FullNameNotIn        []string `json:"fullNameNotIn,omitempty"`
	FullNameGT           *string  `json:"fullNameGT,omitempty"`
	FullNameGTE          *string  `json:"fullNameGTE,omitempty"`
	FullNameLT           *string  `json:"fullNameLT,omitempty"`
	FullNameLTE          *string  `json:"fullNameLTE,omitempty"`
	FullNameContains     *string  `json:"fullNameContains,omitempty"`
	FullNameHasPrefix    *string  `json:"fullNameHasPrefix,omitempty"`
	FullNameHasSuffix    *string  `json:"fullNameHasSuffix,omitempty"`
	FullNameIsNil        bool     `json:"fullNameIsNil,omitempty"`
	FullNameNotNil       bool     `json:"fullNameNotNil,omitempty"`
	FullNameEqualFold    *string  `json:"fullNameEqualFold,omitempty"`
	FullNameContainsFold *string  `json:"fullNameContainsFold,omitempty"`

	// "avatar_url" field predicates.
	AvatarURL             *string  `json:"avatarURL,omitempty"`
	AvatarURLNEQ          *string  `json:"avatarURLNEQ,omitempty"`
	AvatarURLIn           []string `json:"avatarURLIn,omitempty"`
	AvatarURLNotIn        []string `json:"avatarURLNotIn,omitempty"`
	AvatarURLGT           *string  `json:"avatarURLGT,omitempty"`
	AvatarURLGTE          *string  `json:"avatarURLGTE,omitempty"`
	AvatarURLLT           *string  `json:"avatarURLLT,omitempty"`
	AvatarURLLTE          *string  `json:"avatarURLLTE,omitempty"`
	AvatarURLContains     *string  `json:"avatarURLContains,omitempty"`
	AvatarURLHasPrefix    *string  `json:"avatarURLHasPrefix,omitempty"`
	AvatarURLHasSuffix    *string  `json:"avatarURLHasSuffix,omitempty"`
	AvatarURLIsNil        bool     `json:"avatarURLIsNil,omitempty"`
	AvatarURLNotNil       bool     `json:"avatarURLNotNil,omitempty"`
	AvatarURLEqualFold    *string  `json:"avatarURLEqualFold,omitempty"`
	AvatarURLContainsFold *string  `json:"avatarURLContainsFold,omitempty"`

	// "bio" field predicates.
	Bio             *string  `json:"bio,omitempty"`
	BioNEQ          *string  `json:"bioNEQ,omitempty"`
	BioIn           []string `json:"bioIn,omitempty"`
	BioNotIn        []string `json:"bioNotIn,omitempty"`
	BioGT           *string  `json:"bioGT,omitempty"`
	BioGTE          *string  `json:"bioGTE,omitempty"`
	BioLT           *string  `json:"bioLT,omitempty"`
	BioLTE          *string  `json:"bioLTE,omitempty"`
	BioContains     *string  `json:"bioContains,omitempty"`
	BioHasPrefix    *string  `json:"bioHasPrefix,omitempty"`
	BioHasSuffix    *string  `json:"bioHasSuffix,omitempty"`
	BioIsNil        bool     `json:"bioIsNil,omitempty"`
	BioNotNil       bool     `json:"bioNotNil,omitempty"`
	BioEqualFold    *string  `json:"bioEqualFold,omitempty"`
	BioContainsFold *string  `json:"bioContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Username != nil {
		predicates = append(predicates, user.UsernameEQ(*i.Username))
	}
	if i.UsernameNEQ != nil {
		predicates = append(predicates, user.UsernameNEQ(*i.UsernameNEQ))
	}
	if len(i.UsernameIn) > 0 {
		predicates = append(predicates, user.UsernameIn(i.UsernameIn...))
	}
	if len(i.UsernameNotIn) > 0 {
		predicates = append(predicates, user.UsernameNotIn(i.UsernameNotIn...))
	}
	if i.UsernameGT != nil {
		predicates = append(predicates, user.UsernameGT(*i.UsernameGT))
	}
	if i.UsernameGTE != nil {
		predicates = append(predicates, user.UsernameGTE(*i.UsernameGTE))
	}
	if i.UsernameLT != nil {
		predicates = append(predicates, user.UsernameLT(*i.UsernameLT))
	}
	if i.UsernameLTE != nil {
		predicates = append(predicates, user.UsernameLTE(*i.UsernameLTE))
	}
	if i.UsernameContains != nil {
		predicates = append(predicates, user.UsernameContains(*i.UsernameContains))
	}
	if i.UsernameHasPrefix != nil {
		predicates = append(predicates, user.UsernameHasPrefix(*i.UsernameHasPrefix))
	}
	if i.UsernameHasSuffix != nil {
		predicates = append(predicates, user.UsernameHasSuffix(*i.UsernameHasSuffix))
	}
	if i.UsernameEqualFold != nil {
		predicates = append(predicates, user.UsernameEqualFold(*i.UsernameEqualFold))
	}
	if i.UsernameContainsFold != nil {
		predicates = append(predicates, user.UsernameContainsFold(*i.UsernameContainsFold))
	}
	if i.Email != nil {
		predicates = append(predicates, user.EmailEQ(*i.Email))
	}
	if i.EmailNEQ != nil {
		predicates = append(predicates, user.EmailNEQ(*i.EmailNEQ))
	}
	if len(i.EmailIn) > 0 {
		predicates = append(predicates, user.EmailIn(i.EmailIn...))
	}
	if len(i.EmailNotIn) > 0 {
		predicates = append(predicates, user.EmailNotIn(i.EmailNotIn...))
	}
	if i.EmailGT != nil {
		predicates = append(predicates, user.EmailGT(*i.EmailGT))
	}
	if i.EmailGTE != nil {
		predicates = append(predicates, user.EmailGTE(*i.EmailGTE))
	}
	if i.EmailLT != nil {
		predicates = append(predicates, user.EmailLT(*i.EmailLT))
	}
	if i.EmailLTE != nil {
		predicates = append(predicates, user.EmailLTE(*i.EmailLTE))
	}
	if i.EmailContains != nil {
		predicates = append(predicates, user.EmailContains(*i.EmailContains))
	}
	if i.EmailHasPrefix != nil {
		predicates = append(predicates, user.EmailHasPrefix(*i.EmailHasPrefix))
	}
	if i.EmailHasSuffix != nil {
		predicates = append(predicates, user.EmailHasSuffix(*i.EmailHasSuffix))
	}
	if i.EmailEqualFold != nil {
		predicates = append(predicates, user.EmailEqualFold(*i.EmailEqualFold))
	}
	if i.EmailContainsFold != nil {
		predicates = append(predicates, user.EmailContainsFold(*i.EmailContainsFold))
	}
	if i.FullName != nil {
		predicates = append(predicates, user.FullNameEQ(*i.FullName))
	}
	if i.FullNameNEQ != nil {
		predicates = append(predicates, user.FullNameNEQ(*i.FullNameNEQ))
	}
	if len(i.FullNameIn) > 0 {
		predicates = append(predicates, user.FullNameIn(i.FullNameIn...))
	}
	if len(i.FullNameNotIn) > 0 {
		predicates = append(predicates, user.FullNameNotIn(i.FullNameNotIn...))
	}
	if i.FullNameGT != nil {
		predicates = append(predicates, user.FullNameGT(*i.FullNameGT))
	}
	if i.FullNameGTE != nil {
		predicates = append(predicates, user.FullNameGTE(*i.FullNameGTE))
	}
	if i.FullNameLT != nil {
		predicates = append(predicates, user.FullNameLT(*i.FullNameLT))
	}
	if i.FullNameLTE != nil {
		predicates = append(predicates, user.FullNameLTE(*i.FullNameLTE))
	}
	if i.FullNameContains != nil {
		predicates = append(predicates, user.FullNameContains(*i.FullNameContains))
	}
	if i.FullNameHasPrefix != nil {
		predicates = append(predicates, user.FullNameHasPrefix(*i.FullNameHasPrefix))
	}
	if i.FullNameHasSuffix != nil {
		predicates = append(predicates, user.FullNameHasSuffix(*i.FullNameHasSuffix))
	}
	if i.FullNameIsNil {
		predicates = append(predicates, user.FullNameIsNil())
	}
	if i.FullNameNotNil {
		predicates = append(predicates, user.FullNameNotNil())
	}
	if i.FullNameEqualFold != nil {
		predicates = append(predicates, user.FullNameEqualFold(*i.FullNameEqualFold))
	}
	if i.FullNameContainsFold != nil {
		predicates = append(predicates, user.FullNameContainsFold(*i.FullNameContainsFold))
	}
	if i.AvatarURL != nil {
		predicates = append(predicates, user.AvatarURLEQ(*i.AvatarURL))
	}
	if i.AvatarURLNEQ != nil {
		predicates = append(predicates, user.AvatarURLNEQ(*i.AvatarURLNEQ))
	}
	if len(i.AvatarURLIn) > 0 {
		predicates = append(predicates, user.AvatarURLIn(i.AvatarURLIn...))
	}
	if len(i.AvatarURLNotIn) > 0 {
		predicates = append(predicates, user.AvatarURLNotIn(i.AvatarURLNotIn...))
	}
	if i.AvatarURLGT != nil {
		predicates = append(predicates, user.AvatarURLGT(*i.AvatarURLGT))
	}
	if i.AvatarURLGTE != nil {
		predicates = append(predicates, user.AvatarURLGTE(*i.AvatarURLGTE))
	}
	if i.AvatarURLLT != nil {
		predicates = append(predicates, user.AvatarURLLT(*i.AvatarURLLT))
	}
	if i.AvatarURLLTE != nil {
		predicates = append(predicates, user.AvatarURLLTE(*i.AvatarURLLTE))
	}
	if i.AvatarURLContains != nil {
		predicates = append(predicates, user.AvatarURLContains(*i.AvatarURLContains))
	}
	if i.AvatarURLHasPrefix != nil {
		predicates = append(predicates, user.AvatarURLHasPrefix(*i.AvatarURLHasPrefix))
	}
	if i.AvatarURLHasSuffix != nil {
		predicates = append(predicates, user.AvatarURLHasSuffix(*i.AvatarURLHasSuffix))
	}
	if i.AvatarURLIsNil {
		predicates = append(predicates, user.AvatarURLIsNil())
	}
	if i.AvatarURLNotNil {
		predicates = append(predicates, user.AvatarURLNotNil())
	}
	if i.AvatarURLEqualFold != nil {
		predicates = append(predicates, user.AvatarURLEqualFold(*i.AvatarURLEqualFold))
	}
	if i.AvatarURLContainsFold != nil {
		predicates = append(predicates, user.AvatarURLContainsFold(*i.AvatarURLContainsFold))
	}
	if i.Bio != nil {
		predicates = append(predicates, user.BioEQ(*i.Bio))
	}
	if i.BioNEQ != nil {
		predicates = append(predicates, user.BioNEQ(*i.BioNEQ))
	}
	if len(i.BioIn) > 0 {
		predicates = append(predicates, user.BioIn(i.BioIn...))
	}
	if len(i.BioNotIn) > 0 {
		predicates = append(predicates, user.BioNotIn(i.BioNotIn...))
	}
	if i.BioGT != nil {
		predicates = append(predicates, user.BioGT(*i.BioGT))
	}
	if i.BioGTE != nil {
		predicates = append(predicates, user.BioGTE(*i.BioGTE))
	}
	if i.BioLT != nil {
		predicates = append(predicates, user.BioLT(*i.BioLT))
	}
	if i.BioLTE != nil {
		predicates = append(predicates, user.BioLTE(*i.BioLTE))
	}
	if i.BioContains != nil {
		predicates = append(predicates, user.BioContains(*i.BioContains))
	}
	if i.BioHasPrefix != nil {
		predicates = append(predicates, user.BioHasPrefix(*i.BioHasPrefix))
	}
	if i.BioHasSuffix != nil {
		predicates = append(predicates, user.BioHasSuffix(*i.BioHasSuffix))
	}
	if i.BioIsNil {
		predicates = append(predicates, user.BioIsNil())
	}
	if i.BioNotNil {
		predicates = append(predicates, user.BioNotNil())
	}
	if i.BioEqualFold != nil {
		predicates = append(predicates, user.BioEqualFold(*i.BioEqualFold))
	}
	if i.BioContainsFold != nil {
		predicates = append(predicates, user.BioContainsFold(*i.BioContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, user.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, user.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, user.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, user.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, user.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, user.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, user.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, user.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, user.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, user.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, user.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, user.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, user.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, user.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, user.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, user.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/marcustut/fyp/backend/ent: empty predicate UserWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
