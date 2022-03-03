package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
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
		field.Strings("path_token").
			Optional(),
		field.Int64("size").
			Optional().
			Nillable(),
		field.Enum("access_level").
			NamedValues(
				"Private", "PRIVATE",
				"Public", "PUBLIC",
				"View", "VIEW",
			).
			Default("PRIVATE").
			Annotations(
				entgql.Type("AccessLevel"),
			),
		field.Strings("shared_with").
			SchemaType(map[string]string{
				dialect.MySQL: "json DEFAULT (JSON_ARRAY())",
			}),
		field.Bool("deleted").
			Default(false),
	}
}

// Edges of the Slide.
func (Slide) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("instance", Instance.Type).
			Unique(),
		edge.From("user", User.Type).
			Ref("slides").
			Unique().
			Required(),
	}
}

// Mixin of the Slide.
func (Slide) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid("SLID_"),
		SlideMixin{},
		mixin.NewTimestamp(),
	}
}
