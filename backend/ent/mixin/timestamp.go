package mixin

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimestampMixin defines an ent Mixin
type TimestampMixin struct {
	mixin.Schema
}

// NewTimestamp creates a Mixin that includes created_at and updated_at
func NewTimestamp() *TimestampMixin {
	return &TimestampMixin{}
}

// Fields provides the created_at and updated_at field.
func (m TimestampMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "timestamp DEFAULT CURRENT_TIMESTAMP",
			}).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			).
			Immutable(),
	}
}
