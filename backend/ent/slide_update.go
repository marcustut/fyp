// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marcustut/fyp/backend/ent/instance"
	"github.com/marcustut/fyp/backend/ent/predicate"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/ent/slide"
	"github.com/marcustut/fyp/backend/ent/user"
)

// SlideUpdate is the builder for updating Slide entities.
type SlideUpdate struct {
	config
	hooks    []Hook
	mutation *SlideMutation
}

// Where appends a list predicates to the SlideUpdate builder.
func (su *SlideUpdate) Where(ps ...predicate.Slide) *SlideUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SlideUpdate) SetName(s string) *SlideUpdate {
	su.mutation.SetName(s)
	return su
}

// SetPathToken sets the "path_token" field.
func (su *SlideUpdate) SetPathToken(s []string) *SlideUpdate {
	su.mutation.SetPathToken(s)
	return su
}

// ClearPathToken clears the value of the "path_token" field.
func (su *SlideUpdate) ClearPathToken() *SlideUpdate {
	su.mutation.ClearPathToken()
	return su
}

// SetSize sets the "size" field.
func (su *SlideUpdate) SetSize(i int64) *SlideUpdate {
	su.mutation.ResetSize()
	su.mutation.SetSize(i)
	return su
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (su *SlideUpdate) SetNillableSize(i *int64) *SlideUpdate {
	if i != nil {
		su.SetSize(*i)
	}
	return su
}

// AddSize adds i to the "size" field.
func (su *SlideUpdate) AddSize(i int64) *SlideUpdate {
	su.mutation.AddSize(i)
	return su
}

// ClearSize clears the value of the "size" field.
func (su *SlideUpdate) ClearSize() *SlideUpdate {
	su.mutation.ClearSize()
	return su
}

// SetAccessLevel sets the "access_level" field.
func (su *SlideUpdate) SetAccessLevel(sl slide.AccessLevel) *SlideUpdate {
	su.mutation.SetAccessLevel(sl)
	return su
}

// SetNillableAccessLevel sets the "access_level" field if the given value is not nil.
func (su *SlideUpdate) SetNillableAccessLevel(sl *slide.AccessLevel) *SlideUpdate {
	if sl != nil {
		su.SetAccessLevel(*sl)
	}
	return su
}

// SetSharedWith sets the "shared_with" field.
func (su *SlideUpdate) SetSharedWith(s []string) *SlideUpdate {
	su.mutation.SetSharedWith(s)
	return su
}

// SetDeleted sets the "deleted" field.
func (su *SlideUpdate) SetDeleted(b bool) *SlideUpdate {
	su.mutation.SetDeleted(b)
	return su
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (su *SlideUpdate) SetNillableDeleted(b *bool) *SlideUpdate {
	if b != nil {
		su.SetDeleted(*b)
	}
	return su
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (su *SlideUpdate) SetInstanceID(id ulid.ID) *SlideUpdate {
	su.mutation.SetInstanceID(id)
	return su
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (su *SlideUpdate) SetNillableInstanceID(id *ulid.ID) *SlideUpdate {
	if id != nil {
		su = su.SetInstanceID(*id)
	}
	return su
}

// SetInstance sets the "instance" edge to the Instance entity.
func (su *SlideUpdate) SetInstance(i *Instance) *SlideUpdate {
	return su.SetInstanceID(i.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SlideUpdate) SetUserID(id ulid.ID) *SlideUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SlideUpdate) SetUser(u *User) *SlideUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the SlideMutation object of the builder.
func (su *SlideUpdate) Mutation() *SlideMutation {
	return su.mutation
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (su *SlideUpdate) ClearInstance() *SlideUpdate {
	su.mutation.ClearInstance()
	return su
}

// ClearUser clears the "user" edge to the User entity.
func (su *SlideUpdate) ClearUser() *SlideUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SlideUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SlideMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SlideUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SlideUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SlideUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SlideUpdate) check() error {
	if v, ok := su.mutation.AccessLevel(); ok {
		if err := slide.AccessLevelValidator(v); err != nil {
			return &ValidationError{Name: "access_level", err: fmt.Errorf("ent: validator failed for field \"access_level\": %w", err)}
		}
	}
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (su *SlideUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   slide.Table,
			Columns: slide.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: slide.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: slide.FieldName,
		})
	}
	if value, ok := su.mutation.PathToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: slide.FieldPathToken,
		})
	}
	if su.mutation.PathTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: slide.FieldPathToken,
		})
	}
	if value, ok := su.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: slide.FieldSize,
		})
	}
	if value, ok := su.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: slide.FieldSize,
		})
	}
	if su.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: slide.FieldSize,
		})
	}
	if value, ok := su.mutation.AccessLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: slide.FieldAccessLevel,
		})
	}
	if value, ok := su.mutation.SharedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: slide.FieldSharedWith,
		})
	}
	if value, ok := su.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: slide.FieldDeleted,
		})
	}
	if su.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   slide.InstanceTable,
			Columns: []string{slide.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   slide.InstanceTable,
			Columns: []string{slide.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   slide.UserTable,
			Columns: []string{slide.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   slide.UserTable,
			Columns: []string{slide.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{slide.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SlideUpdateOne is the builder for updating a single Slide entity.
type SlideUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SlideMutation
}

// SetName sets the "name" field.
func (suo *SlideUpdateOne) SetName(s string) *SlideUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetPathToken sets the "path_token" field.
func (suo *SlideUpdateOne) SetPathToken(s []string) *SlideUpdateOne {
	suo.mutation.SetPathToken(s)
	return suo
}

// ClearPathToken clears the value of the "path_token" field.
func (suo *SlideUpdateOne) ClearPathToken() *SlideUpdateOne {
	suo.mutation.ClearPathToken()
	return suo
}

// SetSize sets the "size" field.
func (suo *SlideUpdateOne) SetSize(i int64) *SlideUpdateOne {
	suo.mutation.ResetSize()
	suo.mutation.SetSize(i)
	return suo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (suo *SlideUpdateOne) SetNillableSize(i *int64) *SlideUpdateOne {
	if i != nil {
		suo.SetSize(*i)
	}
	return suo
}

// AddSize adds i to the "size" field.
func (suo *SlideUpdateOne) AddSize(i int64) *SlideUpdateOne {
	suo.mutation.AddSize(i)
	return suo
}

// ClearSize clears the value of the "size" field.
func (suo *SlideUpdateOne) ClearSize() *SlideUpdateOne {
	suo.mutation.ClearSize()
	return suo
}

// SetAccessLevel sets the "access_level" field.
func (suo *SlideUpdateOne) SetAccessLevel(sl slide.AccessLevel) *SlideUpdateOne {
	suo.mutation.SetAccessLevel(sl)
	return suo
}

// SetNillableAccessLevel sets the "access_level" field if the given value is not nil.
func (suo *SlideUpdateOne) SetNillableAccessLevel(sl *slide.AccessLevel) *SlideUpdateOne {
	if sl != nil {
		suo.SetAccessLevel(*sl)
	}
	return suo
}

// SetSharedWith sets the "shared_with" field.
func (suo *SlideUpdateOne) SetSharedWith(s []string) *SlideUpdateOne {
	suo.mutation.SetSharedWith(s)
	return suo
}

// SetDeleted sets the "deleted" field.
func (suo *SlideUpdateOne) SetDeleted(b bool) *SlideUpdateOne {
	suo.mutation.SetDeleted(b)
	return suo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (suo *SlideUpdateOne) SetNillableDeleted(b *bool) *SlideUpdateOne {
	if b != nil {
		suo.SetDeleted(*b)
	}
	return suo
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (suo *SlideUpdateOne) SetInstanceID(id ulid.ID) *SlideUpdateOne {
	suo.mutation.SetInstanceID(id)
	return suo
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (suo *SlideUpdateOne) SetNillableInstanceID(id *ulid.ID) *SlideUpdateOne {
	if id != nil {
		suo = suo.SetInstanceID(*id)
	}
	return suo
}

// SetInstance sets the "instance" edge to the Instance entity.
func (suo *SlideUpdateOne) SetInstance(i *Instance) *SlideUpdateOne {
	return suo.SetInstanceID(i.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SlideUpdateOne) SetUserID(id ulid.ID) *SlideUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SlideUpdateOne) SetUser(u *User) *SlideUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the SlideMutation object of the builder.
func (suo *SlideUpdateOne) Mutation() *SlideMutation {
	return suo.mutation
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (suo *SlideUpdateOne) ClearInstance() *SlideUpdateOne {
	suo.mutation.ClearInstance()
	return suo
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SlideUpdateOne) ClearUser() *SlideUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SlideUpdateOne) Select(field string, fields ...string) *SlideUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Slide entity.
func (suo *SlideUpdateOne) Save(ctx context.Context) (*Slide, error) {
	var (
		err  error
		node *Slide
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SlideMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SlideUpdateOne) SaveX(ctx context.Context) *Slide {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SlideUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SlideUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SlideUpdateOne) check() error {
	if v, ok := suo.mutation.AccessLevel(); ok {
		if err := slide.AccessLevelValidator(v); err != nil {
			return &ValidationError{Name: "access_level", err: fmt.Errorf("ent: validator failed for field \"access_level\": %w", err)}
		}
	}
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (suo *SlideUpdateOne) sqlSave(ctx context.Context) (_node *Slide, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   slide.Table,
			Columns: slide.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: slide.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Slide.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, slide.FieldID)
		for _, f := range fields {
			if !slide.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != slide.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: slide.FieldName,
		})
	}
	if value, ok := suo.mutation.PathToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: slide.FieldPathToken,
		})
	}
	if suo.mutation.PathTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: slide.FieldPathToken,
		})
	}
	if value, ok := suo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: slide.FieldSize,
		})
	}
	if value, ok := suo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: slide.FieldSize,
		})
	}
	if suo.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: slide.FieldSize,
		})
	}
	if value, ok := suo.mutation.AccessLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: slide.FieldAccessLevel,
		})
	}
	if value, ok := suo.mutation.SharedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: slide.FieldSharedWith,
		})
	}
	if value, ok := suo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: slide.FieldDeleted,
		})
	}
	if suo.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   slide.InstanceTable,
			Columns: []string{slide.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   slide.InstanceTable,
			Columns: []string{slide.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   slide.UserTable,
			Columns: []string{slide.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   slide.UserTable,
			Columns: []string{slide.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Slide{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{slide.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
