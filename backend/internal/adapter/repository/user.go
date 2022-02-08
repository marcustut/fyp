package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/user"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	usecaseRepository "github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type userRepository struct {
	client *ent.Client
}

// NewUserRepository is specific implementation of the interface
func NewUserRepository(client *ent.Client) usecaseRepository.User {
	return &userRepository{client}
}

func (ur *userRepository) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	u, err := ur.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	u, err := ur.client.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := ur.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (ur *userRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput, orderBy *ent.UserOrder) (*model.UserConnection, error) {
	uc, err := ur.client.
		User.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter), ent.WithUserOrder(orderBy))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return uc, nil
}

func (ur *userRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	u, err := ur.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (ur *userRepository) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	u, err := ur.client.User.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (ur *userRepository) Delete(ctx context.Context, id model.ID) (*model.User, error) {
	u, err := ur.Get(ctx, &id)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	err = ur.client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (ur *userRepository) DeleteByUsername(ctx context.Context, username string) (*model.User, error) {
	u, err := ur.GetByUsername(ctx, username)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	err = ur.client.User.DeleteOne(u).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (ur *userRepository) DeleteByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := ur.GetByEmail(ctx, email)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	err = ur.client.User.DeleteOne(u).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}
