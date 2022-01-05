package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/marcustut/fyp/backend/internal/const/regex"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			MinLen(3).
			MaxLen(50).
			Unique(),
		field.String("email").
			Match(regex.Regexes[regex.Email]).
			MaxLen(320).
			Unique().
			Immutable(),
		field.String("full_name").
			MinLen(3).
			MaxLen(60).
			Optional(),
		field.String("avatar_url").
			Match(regex.Regexes[regex.URL]).
			MaxLen(2083).
			Optional(),
		field.String("bio").
			NotEmpty().
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP",
			}).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}).
			Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
