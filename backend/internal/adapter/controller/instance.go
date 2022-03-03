package controller

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

type instance struct {
	slideUsecase usecase.Instance
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
	return &instance{slideUsecase: u}
}

func (s *instance) Get(ctx context.Context, id model.ID) (*model.Instance, error) {
	return s.slideUsecase.Get(ctx, id)
}

func (s *instance) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.InstanceWhereInput, orderBy *ent.InstanceOrder) (*model.InstanceConnection, error) {
	return s.slideUsecase.List(ctx, after, first, before, last, where, orderBy)
}

func (s *instance) Create(ctx context.Context, input model.CreateInstanceInput) (*model.Instance, error) {
	return s.slideUsecase.Create(ctx, input)
}

func (s *instance) Update(ctx context.Context, input model.UpdateInstanceInput) (*model.Instance, error) {
	return s.slideUsecase.Update(ctx, input)
}

func (s *instance) Delete(ctx context.Context, id model.ID) (*model.Instance, error) {
	return s.slideUsecase.Delete(ctx, id)
}
