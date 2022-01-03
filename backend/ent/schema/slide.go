package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"time"
)

// Slide holds the schema definition for the Slide entity.
type Slide struct {
	ent.Schema
}

// Fields of the Slide.
func (Slide) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID { return ulid.MustNew("SLID_") }),
		field.String("name"),
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

// Edges of the Slide.
func (Slide) Edges() []ent.Edge {
	return nil
}
