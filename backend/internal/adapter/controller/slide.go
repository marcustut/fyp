package controller

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

type slide struct {
	slideUsecase usecase.Slide
}

// Slide defines the interface of the slide controller.
type Slide interface {
	Get(ctx context.Context, id *model.ID) (*model.Slide, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.SlideWhereInput, orderBy *ent.SlideOrder) (*model.SlideConnection, error)
	Create(ctx context.Context, input model.CreateSlideInput) (*model.Slide, error)
	Update(ctx context.Context, input model.UpdateSlideInput) (*model.Slide, error)
}

// NewSlideController returns slide controller.
func NewSlideController(u usecase.Slide) Slide {
	return &slide{slideUsecase: u}
}

func (s *slide) Get(ctx context.Context, id *model.ID) (*model.Slide, error) {
	return s.slideUsecase.Get(ctx, id)
}

func (s *slide) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.SlideWhereInput, orderBy *ent.SlideOrder) (*model.SlideConnection, error) {
	return s.slideUsecase.List(ctx, after, first, before, last, where, orderBy)
}

func (s *slide) Create(ctx context.Context, input model.CreateSlideInput) (*model.Slide, error) {
	return s.slideUsecase.Create(ctx, input)
}

func (s *slide) Update(ctx context.Context, input model.UpdateSlideInput) (*model.Slide, error) {
	return s.slideUsecase.Update(ctx, input)
}
