package repository

import (
	"context"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/entity/model"
)

// Slide is interface of repository
type Slide interface {
	Get(ctx context.Context, id *model.ID) (*model.Slide, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.SlideWhereInput, orderBy *ent.SlideOrder) (*model.SlideConnection, error)
	Create(ctx context.Context, input model.CreateSlideInput) (*model.Slide, error)
	Update(ctx context.Context, input model.UpdateSlideInput) (*model.Slide, error)
}
