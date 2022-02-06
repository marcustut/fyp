package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/internal/adapter/handler"
)

func (r *mutationResolver) CreateSlide(ctx context.Context, input ent.CreateSlideInput) (*ent.Slide, error) {
	s, err := r.controller.Slide.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return s, nil
}

func (r *mutationResolver) UpdateSlide(ctx context.Context, input ent.UpdateSlideInput) (*ent.Slide, error) {
	s, err := r.controller.Slide.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return s, nil
}

func (r *queryResolver) Slide(ctx context.Context, id ulid.ID) (*ent.Slide, error) {
	s, err := r.controller.Slide.Get(ctx, &id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return s, nil
}
