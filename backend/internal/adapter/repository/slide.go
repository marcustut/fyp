package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/slide"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	usecaseRepository "github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type slideRepository struct {
	client *ent.Client
}

// NewSlideRepository is specific implementation of the interface
func NewSlideRepository(client *ent.Client) usecaseRepository.Slide {
	return &slideRepository{client}
}

func (sr *slideRepository) Get(ctx context.Context, id model.ID) (*model.Slide, error) {
	s, err := sr.client.Slide.Query().Where(slide.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return s, nil
}

func (sr *slideRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.SlideWhereInput, orderBy *ent.SlideOrder) (*model.SlideConnection, error) {
	sc, err := sr.client.
		Slide.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithSlideFilter(where.Filter), ent.WithSlideOrder(orderBy))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return sc, nil
}

func (sr *slideRepository) Create(ctx context.Context, input model.CreateSlideInput) (*model.Slide, error) {
	s, err := sr.client.Slide.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return s, nil
}

func (sr *slideRepository) Update(ctx context.Context, input model.UpdateSlideInput) (*model.Slide, error) {
	s, err := sr.client.Slide.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return s, nil
}

func (sr *slideRepository) Delete(ctx context.Context, id model.ID) (*model.Slide, error) {
	s, err := sr.Get(ctx, id)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	err = sr.client.Slide.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return s, nil
}
