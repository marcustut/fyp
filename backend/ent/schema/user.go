package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"github.com/marcustut/fyp/backend/ent/mixin"
	"github.com/marcustut/fyp/backend/internal/const/regex"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// UserMixin defines Fields
type UserMixin struct {
	entMixin.Schema
}

// Fields of the User.
func (UserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			MinLen(3).
			MaxLen(50).
			Unique(),
		field.String("email").
			Match(regex.Regexes[regex.Email]).
			Unique().
			Immutable(),
		field.String("full_name").
			MinLen(3).
			MaxLen(60).
			Optional(),
		field.String("password_hash").
			Immutable(),
		field.String("avatar_url").
			Match(regex.Regexes[regex.URL]).
			MaxLen(2083).
			Optional(),
		field.String("bio").
			NotEmpty().
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("instances", Instance.Type),
		edge.To("slides", Slide.Type),
		edge.To("links", Link.Type),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid("USER_"),
		UserMixin{},
		mixin.NewTimestamp(),
	}
}
