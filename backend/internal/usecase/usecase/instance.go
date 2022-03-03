package usecase

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type instance struct {
	instanceRepository repository.Instance
}

// Instance ...
type Instance interface {
	repository.Instance
}

// NewInstanceUsecase creates a instance usecase.
func NewInstanceUsecase(r repository.Instance) Instance {
	return &instance{r}
}

func (s *instance) Get(ctx context.Context, id model.ID) (*model.Instance, error) {
	return s.instanceRepository.Get(ctx, id)
}

func (s *instance) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.InstanceWhereInput, orderBy *ent.InstanceOrder) (*model.InstanceConnection, error) {
	return s.instanceRepository.List(ctx, after, first, before, last, where, orderBy)
}

func (s *instance) Create(ctx context.Context, input model.CreateInstanceInput) (*model.Instance, error) {
	return s.instanceRepository.Create(ctx, input)
}

func (s *instance) Update(ctx context.Context, input model.UpdateInstanceInput) (*model.Instance, error) {
	return s.instanceRepository.Update(ctx, input)
}

func (s *instance) Delete(ctx context.Context, id model.ID) (*model.Instance, error) {
	return s.instanceRepository.Delete(ctx, id)
}
