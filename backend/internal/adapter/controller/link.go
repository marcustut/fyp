package controller

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

type link struct {
	linkUsecase usecase.Link
}

// Link defines the interface of the link controller.
type Link interface {
	Get(ctx context.Context, id model.ID) (*model.Link, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput, orderBy *ent.LinkOrder) (*model.LinkConnection, error)
	Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error)
	Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error)
	Delete(ctx context.Context, id model.ID) (*model.Link, error)
}

// NewLinkController returns link controller.
func NewLinkController(u usecase.Link) Link {
	return &link{linkUsecase: u}
}

func (l *link) Get(ctx context.Context, id model.ID) (*model.Link, error) {
	return l.linkUsecase.Get(ctx, id)
}

func (l *link) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput, orderBy *ent.LinkOrder) (*model.LinkConnection, error) {
	return l.linkUsecase.List(ctx, after, first, before, last, where, orderBy)
}

func (l *link) Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	return l.linkUsecase.Create(ctx, input)
}

func (l *link) Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	return l.linkUsecase.Update(ctx, input)
}

func (l *link) Delete(ctx context.Context, id model.ID) (*model.Link, error) {
	return l.linkUsecase.Delete(ctx, id)
}
