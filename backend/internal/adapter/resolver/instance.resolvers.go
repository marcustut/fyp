package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/internal/adapter/handler"
)

func (r *mutationResolver) CreateInstance(ctx context.Context, input ent.CreateInstanceInput) (*ent.Instance, error) {
	i, err := r.controller.Instance.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return i, nil
}

func (r *mutationResolver) UpdateInstance(ctx context.Context, input ent.UpdateInstanceInput) (*ent.Instance, error) {
	i, err := r.controller.Instance.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return i, nil
}

func (r *mutationResolver) DeleteInstance(ctx context.Context, id ulid.ID) (*ent.Instance, error) {
	i, err := r.controller.Instance.Delete(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return i, nil
}

func (r *queryResolver) Instance(ctx context.Context, id ulid.ID) (*ent.Instance, error) {
	i, err := r.controller.Instance.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return i, nil
}

func (r *queryResolver) Instances(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.InstanceWhereInput, orderBy *ent.InstanceOrder) (*ent.InstanceConnection, error) {
	ic, err := r.controller.Instance.List(ctx, after, first, before, last, where, orderBy)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return ic, nil
}
