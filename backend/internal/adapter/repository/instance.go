package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/instance"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	usecaseRepository "github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type instanceRepository struct {
	client *ent.Client
}

// NewInstanceRepository is specific implementation of the interface
func NewInstanceRepository(client *ent.Client) usecaseRepository.Instance {
	return &instanceRepository{client}
}

func (ir *instanceRepository) Get(ctx context.Context, id model.ID) (*model.Instance, error) {
	i, err := ir.client.Instance.Query().Where(instance.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return i, nil
}

func (ir *instanceRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.InstanceWhereInput, orderBy *ent.InstanceOrder) (*model.InstanceConnection, error) {
	ic, err := ir.client.
		Instance.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithInstanceFilter(where.Filter), ent.WithInstanceOrder(orderBy))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return ic, nil
}

func (ir *instanceRepository) Create(ctx context.Context, input model.CreateInstanceInput) (*model.Instance, error) {
	i, err := ir.client.Instance.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return i, nil
}

func (ir *instanceRepository) Update(ctx context.Context, input model.UpdateInstanceInput) (*model.Instance, error) {
	i, err := ir.client.Instance.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return i, nil
}

func (ir *instanceRepository) Delete(ctx context.Context, id model.ID) (*model.Instance, error) {
	i, err := ir.Get(ctx, id)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	err = ir.client.Instance.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return i, nil
}
