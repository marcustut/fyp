package controller

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

type instance struct {
	instanceUsecase usecase.Instance
}

// Instance defines the interface of the instance controller.
type Instance interface {
	Get(ctx context.Context, id model.ID) (*model.Instance, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.InstanceWhereInput, orderBy *ent.InstanceOrder) (*model.InstanceConnection, error)
	Create(ctx context.Context, input model.CreateInstanceInput) (*model.Instance, error)
	Update(ctx context.Context, input model.UpdateInstanceInput) (*model.Instance, error)
	Delete(ctx context.Context, id model.ID) (*model.Instance, error)
}

// NewInstanceController returns instance controller.
func NewInstanceController(u usecase.Instance) Instance {
	return &instance{instanceUsecase: u}
}

func (i *instance) Get(ctx context.Context, id model.ID) (*model.Instance, error) {
	return i.instanceUsecase.Get(ctx, id)
}

func (i *instance) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.InstanceWhereInput, orderBy *ent.InstanceOrder) (*model.InstanceConnection, error) {
	return i.instanceUsecase.List(ctx, after, first, before, last, where, orderBy)
}

func (i *instance) Create(ctx context.Context, input model.CreateInstanceInput) (*model.Instance, error) {
	return i.instanceUsecase.Create(ctx, input)
}

func (i *instance) Update(ctx context.Context, input model.UpdateInstanceInput) (*model.Instance, error) {
	return i.instanceUsecase.Update(ctx, input)
}

func (i *instance) Delete(ctx context.Context, id model.ID) (*model.Instance, error) {
	return i.instanceUsecase.Delete(ctx, id)
}
