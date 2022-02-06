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

func (r *slideRepository) Get(ctx context.Context, id *model.ID) (*model.Slide, error) {
	s, err := r.client.Slide.Query().Where(slide.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *slideRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.SlideWhereInput, orderBy *ent.SlideOrder) (*model.SlideConnection, error) {
	sc, err := r.client.
		Slide.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithSlideFilter(where.Filter), ent.WithSlideOrder(orderBy))
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func (r *slideRepository) Create(ctx context.Context, input model.CreateSlideInput) (*model.Slide, error) {
	s, err := r.client.Slide.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return s, nil
}

func (r *slideRepository) Update(ctx context.Context, input model.UpdateSlideInput) (*model.Slide, error) {
	s, err := r.client.Slide.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return s, nil
}
