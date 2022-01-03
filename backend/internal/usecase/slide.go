package usecase

import (
	"context"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
)

type slideUsecase struct {
	slideRepository controller.Slide
}

type Slide interface {
	controller.Slide
}

func NewSlideUsecase(r controller.Slide) controller.Slide {
	return &slideUsecase{r}
}

func (u *slideUsecase) Get(ctx context.Context)    {}
func (u *slideUsecase) List(ctx context.Context)   {}
func (u *slideUsecase) Create(ctx context.Context) {}
func (u *slideUsecase) Update(ctx context.Context) {}
