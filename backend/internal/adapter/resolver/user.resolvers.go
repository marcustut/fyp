package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/internal/adapter/handler"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	u, err := r.controller.User.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*ent.User, error) {
	u, err := r.controller.User.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id ulid.ID) (*ent.User, error) {
	u, err := r.controller.User.Delete(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) DeleteUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	u, err := r.controller.User.DeleteByUsername(ctx, username)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) DeleteUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.controller.User.DeleteByEmail(ctx, email)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) User(ctx context.Context, id ulid.ID) (*ent.User, error) {
	u, err := r.controller.User.Get(ctx, &id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*ent.User, error) {
	u, err := r.controller.User.GetByUsername(ctx, username)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.controller.User.GetByEmail(ctx, email)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	us, err := r.controller.User.List(ctx, after, first, before, last, where, orderBy)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return us, nil
}
