// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/marcustut/fyp/backend/ent/schema"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/ent/slide"
	"github.com/marcustut/fyp/backend/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	slideMixin := schema.Slide{}.Mixin()
	slideMixinFields0 := slideMixin[0].Fields()
	_ = slideMixinFields0
	slideMixinFields2 := slideMixin[2].Fields()
	_ = slideMixinFields2
	slideFields := schema.Slide{}.Fields()
	_ = slideFields
	// slideDescCreatedAt is the schema descriptor for created_at field.
	slideDescCreatedAt := slideMixinFields2[0].Descriptor()
	// slide.DefaultCreatedAt holds the default value on creation for the created_at field.
	slide.DefaultCreatedAt = slideDescCreatedAt.Default.(func() time.Time)
	// slideDescUpdatedAt is the schema descriptor for updated_at field.
	slideDescUpdatedAt := slideMixinFields2[1].Descriptor()
	// slide.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	slide.DefaultUpdatedAt = slideDescUpdatedAt.Default.(func() time.Time)
	// slideDescID is the schema descriptor for id field.
	slideDescID := slideMixinFields0[0].Descriptor()
	// slide.DefaultID holds the default value on creation for the id field.
	slide.DefaultID = slideDescID.Default.(func() ulid.ID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userMixinFields2 := userMixin[2].Fields()
	_ = userMixinFields2
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userMixinFields1[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userMixinFields1[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescFullName is the schema descriptor for full_name field.
	userDescFullName := userMixinFields1[2].Descriptor()
	// user.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	user.FullNameValidator = func() func(string) error {
		validators := userDescFullName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(full_name string) error {
			for _, fn := range fns {
				if err := fn(full_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAvatarURL is the schema descriptor for avatar_url field.
	userDescAvatarURL := userMixinFields1[4].Descriptor()
	// user.AvatarURLValidator is a validator for the "avatar_url" field. It is called by the builders before save.
	user.AvatarURLValidator = func() func(string) error {
		validators := userDescAvatarURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(avatar_url string) error {
			for _, fn := range fns {
				if err := fn(avatar_url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescBio is the schema descriptor for bio field.
	userDescBio := userMixinFields1[5].Descriptor()
	// user.BioValidator is a validator for the "bio" field. It is called by the builders before save.
	user.BioValidator = userDescBio.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields2[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields2[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() ulid.ID)
}
