package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"github.com/marcustut/fyp/backend/ent/mixin"
)

// Slide holds the schema definition for the Slide entity.
type Slide struct {
	ent.Schema
}

// SlideMixin defines Fields
type SlideMixin struct {
	entMixin.Schema
}

// Fields of the Slide.
func (SlideMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Slide.
func (Slide) Edges() []ent.Edge {
	return nil
}

// Mixin of the Slide.
func (Slide) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid("SLID_"),
		SlideMixin{},
		mixin.NewTimestamp(),
	}
}
