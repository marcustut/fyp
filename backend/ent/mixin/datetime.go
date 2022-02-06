package mixin

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// DatetimeMixin defines an ent Mixin
type DatetimeMixin struct {
	mixin.Schema
}

// NewDatetime creates a Mixin that includes created_at and updated_at
func NewDatetime() *DatetimeMixin {
	return &DatetimeMixin{}
}

// Fields provides the created_at and updated_at field.
func (m DatetimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP",
			}).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			).
			Immutable(),
	}
}
