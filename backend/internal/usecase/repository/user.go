package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
)

// User is interface of repository
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput, orderBy *ent.UserOrder) (*model.UserConnection, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error)
	Delete(ctx context.Context, id model.ID) (*model.User, error)
	DeleteByUsername(ctx context.Context, username string) (*model.User, error)
	DeleteByEmail(ctx context.Context, email string) (*model.User, error)
}
