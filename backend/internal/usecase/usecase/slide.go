package usecase

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type slide struct {
	slideRepository repository.Slide
}

// Slide ...
type Slide interface {
	repository.Slide
}

// NewSlideUsecase creates a slide usecase.
func NewSlideUsecase(r repository.Slide) Slide {
	return &slide{r}
}

func (s *slide) Get(ctx context.Context, id *model.ID) (*model.Slide, error) {
	return s.slideRepository.Get(ctx, id)
}

func (s *slide) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.SlideWhereInput, orderBy *ent.SlideOrder) (*model.SlideConnection, error) {
	return s.slideRepository.List(ctx, after, first, before, last, where, orderBy)
}

func (s *slide) Create(ctx context.Context, input model.CreateSlideInput) (*model.Slide, error) {
	return s.slideRepository.Create(ctx, input)
}

func (s *slide) Update(ctx context.Context, input model.UpdateSlideInput) (*model.Slide, error) {
	return s.slideRepository.Update(ctx, input)
}
