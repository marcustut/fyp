package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/graph"
	"github.com/marcustut/fyp/backend/internal/adapter/handler"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/util/crypto/shortid"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input ent.CreateLinkInput) (*ent.Link, error) {
	l, err := r.controller.Link.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return l, nil
}

func (r *mutationResolver) CreateLinkOptionalLinkID(ctx context.Context, input graph.CreateLinkOptionalLinkIDInput) (*ent.Link, error) {
	// generate a link_id if not specified
	if input.LinkID == nil {
		sid, err := shortid.Generate()
		if err != nil {
			return nil, model.NewInternalServerError(err)
		}
		input.LinkID = &sid
	}

	// create the link
	l, err := r.controller.Link.Create(ctx, ent.CreateLinkInput{
		LinkID:      *input.LinkID,
		OriginalURL: input.OriginalURL,
		OwnerID:     input.OwnerID,
	})
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return l, nil
}

func (r *mutationResolver) UpdateLink(ctx context.Context, input ent.UpdateLinkInput) (*ent.Link, error) {
	l, err := r.controller.Link.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return l, nil
}

func (r *mutationResolver) DeleteLink(ctx context.Context, id ulid.ID) (*ent.Link, error) {
	l, err := r.controller.Link.Delete(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return l, nil
}

func (r *queryResolver) Link(ctx context.Context, id ulid.ID) (*ent.Link, error) {
	l, err := r.controller.Link.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return l, nil
}

func (r *queryResolver) Links(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.LinkWhereInput, orderBy *ent.LinkOrder) (*ent.LinkConnection, error) {
	lc, err := r.controller.Link.List(ctx, after, first, before, last, where, orderBy)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return lc, nil
}
