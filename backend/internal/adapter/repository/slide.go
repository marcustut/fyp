package repository

import (
	"context"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
)

type slideRepository struct {
}

func NewSlideRepository() controller.Slide {
	return &slideRepository{}
}

func (r *slideRepository) Get(ctx context.Context)    {}
func (r *slideRepository) List(ctx context.Context)   {}
func (r *slideRepository) Create(ctx context.Context) {}
func (r *slideRepository) Update(ctx context.Context) {}
