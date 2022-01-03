package controller

import (
	"context"
	"github.com/marcustut/fyp/backend/internal/usecase"
)

type slide struct {
	slideUsecase usecase.Slide
}

// Slide defines the interface of the slide controller.
type Slide interface {
	Get(ctx context.Context)
	List(ctx context.Context)
	Create(ctx context.Context)
	Update(ctx context.Context)
}

func NewSlideController(u usecase.Slide) Slide {
	return &slide{slideUsecase: u}
}

func (s *slide) Get(ctx context.Context)    {}
func (s *slide) List(ctx context.Context)   {}
func (s *slide) Create(ctx context.Context) {}
func (s *slide) Update(ctx context.Context) {}
