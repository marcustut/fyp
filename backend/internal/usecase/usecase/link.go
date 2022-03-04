package usecase

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type link struct {
	instanceRepository repository.Link
}

// Link ...
type Link interface {
	repository.Link
}

// NewLinkUsecase creates a link usecase.
func NewLinkUsecase(r repository.Link) Link {
	return &link{r}
}

func (s *link) Get(ctx context.Context, id model.ID) (*model.Link, error) {
	return s.instanceRepository.Get(ctx, id)
}

func (s *link) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput, orderBy *ent.LinkOrder) (*model.LinkConnection, error) {
	return s.instanceRepository.List(ctx, after, first, before, last, where, orderBy)
}

func (s *link) Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	return s.instanceRepository.Create(ctx, input)
}

func (s *link) Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	return s.instanceRepository.Update(ctx, input)
}

func (s *link) Delete(ctx context.Context, id model.ID) (*model.Link, error) {
	return s.instanceRepository.Delete(ctx, id)
}
