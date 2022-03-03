package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
)

// Instance is interface of repository
type Instance interface {
	Get(ctx context.Context, id model.ID) (*model.Instance, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.InstanceWhereInput, orderBy *ent.InstanceOrder) (*model.InstanceConnection, error)
	Create(ctx context.Context, input model.CreateInstanceInput) (*model.Instance, error)
	Update(ctx context.Context, input model.UpdateInstanceInput) (*model.Instance, error)
	Delete(ctx context.Context, id model.ID) (*model.Instance, error)
}
