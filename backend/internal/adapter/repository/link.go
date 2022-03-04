package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/link"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	usecaseRepository "github.com/marcustut/fyp/backend/internal/usecase/repository"
)

type linkRepository struct {
	client *ent.Client
}

// NewLinkRepository is specific implementation of the interface
func NewLinkRepository(client *ent.Client) usecaseRepository.Link {
	return &linkRepository{client}
}

func (lr *linkRepository) Get(ctx context.Context, id model.ID) (*model.Link, error) {
	l, err := lr.client.Link.Query().Where(link.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return l, nil
}

func (lr *linkRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput, orderBy *ent.LinkOrder) (*model.LinkConnection, error) {
	lc, err := lr.client.
		Link.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithLinkFilter(where.Filter), ent.WithLinkOrder(orderBy))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return lc, nil
}

func (lr *linkRepository) Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	l, err := lr.client.Link.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return l, nil
}

func (lr *linkRepository) Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	l, err := lr.client.Link.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return l, nil
}

func (lr *linkRepository) Delete(ctx context.Context, id model.ID) (*model.Link, error) {
	l, err := lr.Get(ctx, id)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	err = lr.client.Link.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return l, nil
}
