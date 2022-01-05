// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marcustut/fyp/backend/ent/predicate"
	"github.com/marcustut/fyp/backend/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetFullName sets the "full_name" field.
func (uu *UserUpdate) SetFullName(s string) *UserUpdate {
	uu.mutation.SetFullName(s)
	return uu
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFullName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFullName(*s)
	}
	return uu
}

// ClearFullName clears the value of the "full_name" field.
func (uu *UserUpdate) ClearFullName() *UserUpdate {
	uu.mutation.ClearFullName()
	return uu
}

// SetAvatarURL sets the "avatar_url" field.
func (uu *UserUpdate) SetAvatarURL(s string) *UserUpdate {
	uu.mutation.SetAvatarURL(s)
	return uu
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatarURL(*s)
	}
	return uu
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uu *UserUpdate) ClearAvatarURL() *UserUpdate {
	uu.mutation.ClearAvatarURL()
	return uu
}

// SetBio sets the "bio" field.
func (uu *UserUpdate) SetBio(s string) *UserUpdate {
	uu.mutation.SetBio(s)
	return uu
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBio(s *string) *UserUpdate {
	if s != nil {
		uu.SetBio(*s)
	}
	return uu
}

// ClearBio clears the value of the "bio" field.
func (uu *UserUpdate) ClearBio() *UserUpdate {
	uu.mutation.ClearBio()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	if v, ok := uu.mutation.FullName(); ok {
		if err := user.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf("ent: validator failed for field \"full_name\": %w", err)}
		}
	}
	if v, ok := uu.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf("ent: validator failed for field \"avatar_url\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Bio(); ok {
		if err := user.BioValidator(v); err != nil {
			return &ValidationError{Name: "bio", err: fmt.Errorf("ent: validator failed for field \"bio\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uu.mutation.FullName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFullName,
		})
	}
	if uu.mutation.FullNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFullName,
		})
	}
	if value, ok := uu.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if uu.mutation.AvatarURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uu.mutation.Bio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBio,
		})
	}
	if uu.mutation.BioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldBio,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetFullName sets the "full_name" field.
func (uuo *UserUpdateOne) SetFullName(s string) *UserUpdateOne {
	uuo.mutation.SetFullName(s)
	return uuo
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFullName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFullName(*s)
	}
	return uuo
}

// ClearFullName clears the value of the "full_name" field.
func (uuo *UserUpdateOne) ClearFullName() *UserUpdateOne {
	uuo.mutation.ClearFullName()
	return uuo
}

// SetAvatarURL sets the "avatar_url" field.
func (uuo *UserUpdateOne) SetAvatarURL(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarURL(s)
	return uuo
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatarURL(*s)
	}
	return uuo
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uuo *UserUpdateOne) ClearAvatarURL() *UserUpdateOne {
	uuo.mutation.ClearAvatarURL()
	return uuo
}

// SetBio sets the "bio" field.
func (uuo *UserUpdateOne) SetBio(s string) *UserUpdateOne {
	uuo.mutation.SetBio(s)
	return uuo
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBio(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBio(*s)
	}
	return uuo
}

// ClearBio clears the value of the "bio" field.
func (uuo *UserUpdateOne) ClearBio() *UserUpdateOne {
	uuo.mutation.ClearBio()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.FullName(); ok {
		if err := user.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf("ent: validator failed for field \"full_name\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf("ent: validator failed for field \"avatar_url\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Bio(); ok {
		if err := user.BioValidator(v); err != nil {
			return &ValidationError{Name: "bio", err: fmt.Errorf("ent: validator failed for field \"bio\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uuo.mutation.FullName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFullName,
		})
	}
	if uuo.mutation.FullNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFullName,
		})
	}
	if value, ok := uuo.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if uuo.mutation.AvatarURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uuo.mutation.Bio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBio,
		})
	}
	if uuo.mutation.BioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldBio,
		})
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
