package usecase

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type user struct {
	userRepository repository.User
}

// User ...
type User interface {
	repository.User
}

// NewUserUsecase creates a user usecase.
func NewUserUsecase(r repository.User) User {
	return &user{r}
}

func (s *user) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return s.userRepository.Get(ctx, id)
}

func (s *user) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.userRepository.GetByUsername(ctx, username)
}

func (s *user) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepository.GetByEmail(ctx, email)
}

func (s *user) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput, orderBy *ent.UserOrder) (*model.UserConnection, error) {
	return s.userRepository.List(ctx, after, first, before, last, where, orderBy)
}

func (s *user) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return s.userRepository.Create(ctx, input)
}

func (s *user) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return s.userRepository.Update(ctx, input)
}

func (s *user) Delete(ctx context.Context, id model.ID) (*model.User, error) {
	return s.userRepository.Delete(ctx, id)
}

func (s *user) DeleteByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.userRepository.DeleteByUsername(ctx, username)
}

func (s *user) DeleteByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepository.DeleteByEmail(ctx, email)
}
