package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
)

// Link is interface of repository
type Link interface {
	Get(ctx context.Context, id model.ID) (*model.Link, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput, orderBy *ent.LinkOrder) (*model.LinkConnection, error)
	Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error)
	Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error)
	Delete(ctx context.Context, id model.ID) (*model.Link, error)
}
