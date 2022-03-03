package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"github.com/marcustut/fyp/backend/ent/mixin"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// InstanceMixin defines Fields
type InstanceMixin struct {
	entMixin.Schema
}

// Fields of the Instance.
func (InstanceMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("instance_id").
			Immutable().
			Unique(),
		field.String("instance_type"),
		field.String("private_dns_name").
			Unique(),
		field.String("private_ip_address").
			Unique(),
		field.String("public_dns_name").
			Unique(),
		field.String("public_ip_address").
			Unique(),
		field.String("image_id"),
		field.String("architecture").
			Match(regexp.MustCompile("^(i386|x86_64|arm64|x86_64_mac)$")),
		field.String("availability_zone").
			Match(regexp.MustCompile("^(ap-southeast-1a|ap-southeast-1b|ap-southeast-1c|ap-southeast-1d)$")),
	}
}

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("instances").
			Unique().
			Required(),
		edge.From("slide", Slide.Type).
			Ref("instance").
			Unique().
			Required(),
	}
}

// Mixin of the Instance.
func (Instance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid("INST_"),
		InstanceMixin{},
		mixin.NewTimestamp(),
	}
}
