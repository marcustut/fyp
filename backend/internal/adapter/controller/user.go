package controller

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

type user struct {
	UserUsecase usecase.User
}

// User defines the interface of the user controller.
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

// NewUserController returns user controller.
func NewUserController(u usecase.User) User {
	return &user{UserUsecase: u}
}

func (s *user) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return s.UserUsecase.Get(ctx, id)
}

func (s *user) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.UserUsecase.GetByUsername(ctx, username)
}

func (s *user) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.UserUsecase.GetByEmail(ctx, email)
}

func (s *user) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput, orderBy *ent.UserOrder) (*model.UserConnection, error) {
	return s.UserUsecase.List(ctx, after, first, before, last, where, orderBy)
}

func (s *user) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return s.UserUsecase.Create(ctx, input)
}

func (s *user) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return s.UserUsecase.Update(ctx, input)
}

func (s *user) Delete(ctx context.Context, id model.ID) (*model.User, error) {
	return s.UserUsecase.Delete(ctx, id)
}

func (s *user) DeleteByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.UserUsecase.DeleteByUsername(ctx, username)
}

func (s *user) DeleteByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.UserUsecase.DeleteByEmail(ctx, email)
}
