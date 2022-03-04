package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"github.com/marcustut/fyp/backend/ent/mixin"
	"github.com/marcustut/fyp/backend/internal/const/regex"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// LinkMixin defines Fields
type LinkMixin struct {
	entMixin.Schema
}

// Fields of the Link.
func (LinkMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("link_id").
			Unique(),
		field.String("original_url").
			Match(regex.Regexes[regex.URL]),
		field.Int64("visited_count").
			Default(0),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("links").
			Unique().
			Required(),
	}
}

// Mixin of the Link.
func (Link) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid("LINK_"),
		LinkMixin{},
		mixin.NewTimestamp(),
	}
}
