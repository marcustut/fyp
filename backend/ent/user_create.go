// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marcustut/fyp/backend/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetFullName sets the "full_name" field.
func (uc *UserCreate) SetFullName(s string) *UserCreate {
	uc.mutation.SetFullName(s)
	return uc
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableFullName(s *string) *UserCreate {
	if s != nil {
		uc.SetFullName(*s)
	}
	return uc
}

// SetAvatarURL sets the "avatar_url" field.
func (uc *UserCreate) SetAvatarURL(s string) *UserCreate {
	uc.mutation.SetAvatarURL(s)
	return uc
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatarURL(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatarURL(*s)
	}
	return uc
}

// SetBio sets the "bio" field.
func (uc *UserCreate) SetBio(s string) *UserCreate {
	uc.mutation.SetBio(s)
	return uc
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uc *UserCreate) SetNillableBio(s *string) *UserCreate {
	if s != nil {
		uc.SetBio(*s)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "username"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "email": %w`, err)}
		}
	}
	if v, ok := uc.mutation.FullName(); ok {
		if err := user.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "full_name": %w`, err)}
		}
	}
	if v, ok := uc.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "avatar_url": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Bio(); ok {
		if err := user.BioValidator(v); err != nil {
			return &ValidationError{Name: "bio", err: fmt.Errorf(`ent: validator failed for field "bio": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.FullName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFullName,
		})
		_node.FullName = value
	}
	if value, ok := uc.mutation.AvatarURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
		_node.AvatarURL = value
	}
	if value, ok := uc.mutation.Bio(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBio,
		})
		_node.Bio = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}