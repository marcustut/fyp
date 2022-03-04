// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marcustut/fyp/backend/ent/link"
	"github.com/marcustut/fyp/backend/ent/predicate"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/ent/user"
)

// LinkUpdate is the builder for updating Link entities.
type LinkUpdate struct {
	config
	hooks    []Hook
	mutation *LinkMutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (lu *LinkUpdate) Where(ps ...predicate.Link) *LinkUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetLinkID sets the "link_id" field.
func (lu *LinkUpdate) SetLinkID(s string) *LinkUpdate {
	lu.mutation.SetLinkID(s)
	return lu
}

// SetOriginalURL sets the "original_url" field.
func (lu *LinkUpdate) SetOriginalURL(s string) *LinkUpdate {
	lu.mutation.SetOriginalURL(s)
	return lu
}

// SetVisitedCount sets the "visited_count" field.
func (lu *LinkUpdate) SetVisitedCount(i int64) *LinkUpdate {
	lu.mutation.ResetVisitedCount()
	lu.mutation.SetVisitedCount(i)
	return lu
}

// SetNillableVisitedCount sets the "visited_count" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableVisitedCount(i *int64) *LinkUpdate {
	if i != nil {
		lu.SetVisitedCount(*i)
	}
	return lu
}

// AddVisitedCount adds i to the "visited_count" field.
func (lu *LinkUpdate) AddVisitedCount(i int64) *LinkUpdate {
	lu.mutation.AddVisitedCount(i)
	return lu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (lu *LinkUpdate) SetOwnerID(id ulid.ID) *LinkUpdate {
	lu.mutation.SetOwnerID(id)
	return lu
}

// SetOwner sets the "owner" edge to the User entity.
func (lu *LinkUpdate) SetOwner(u *User) *LinkUpdate {
	return lu.SetOwnerID(u.ID)
}

// Mutation returns the LinkMutation object of the builder.
func (lu *LinkUpdate) Mutation() *LinkMutation {
	return lu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (lu *LinkUpdate) ClearOwner() *LinkUpdate {
	lu.mutation.ClearOwner()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LinkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LinkUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LinkUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LinkUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LinkUpdate) check() error {
	if v, ok := lu.mutation.OriginalURL(); ok {
		if err := link.OriginalURLValidator(v); err != nil {
			return &ValidationError{Name: "original_url", err: fmt.Errorf("ent: validator failed for field \"original_url\": %w", err)}
		}
	}
	if _, ok := lu.mutation.OwnerID(); lu.mutation.OwnerCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (lu *LinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: link.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.LinkID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldLinkID,
		})
	}
	if value, ok := lu.mutation.OriginalURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldOriginalURL,
		})
	}
	if value, ok := lu.mutation.VisitedCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: link.FieldVisitedCount,
		})
	}
	if value, ok := lu.mutation.AddedVisitedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: link.FieldVisitedCount,
		})
	}
	if lu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
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
	if nodes := lu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LinkUpdateOne is the builder for updating a single Link entity.
type LinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LinkMutation
}

// SetLinkID sets the "link_id" field.
func (luo *LinkUpdateOne) SetLinkID(s string) *LinkUpdateOne {
	luo.mutation.SetLinkID(s)
	return luo
}

// SetOriginalURL sets the "original_url" field.
func (luo *LinkUpdateOne) SetOriginalURL(s string) *LinkUpdateOne {
	luo.mutation.SetOriginalURL(s)
	return luo
}

// SetVisitedCount sets the "visited_count" field.
func (luo *LinkUpdateOne) SetVisitedCount(i int64) *LinkUpdateOne {
	luo.mutation.ResetVisitedCount()
	luo.mutation.SetVisitedCount(i)
	return luo
}

// SetNillableVisitedCount sets the "visited_count" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableVisitedCount(i *int64) *LinkUpdateOne {
	if i != nil {
		luo.SetVisitedCount(*i)
	}
	return luo
}

// AddVisitedCount adds i to the "visited_count" field.
func (luo *LinkUpdateOne) AddVisitedCount(i int64) *LinkUpdateOne {
	luo.mutation.AddVisitedCount(i)
	return luo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (luo *LinkUpdateOne) SetOwnerID(id ulid.ID) *LinkUpdateOne {
	luo.mutation.SetOwnerID(id)
	return luo
}

// SetOwner sets the "owner" edge to the User entity.
func (luo *LinkUpdateOne) SetOwner(u *User) *LinkUpdateOne {
	return luo.SetOwnerID(u.ID)
}

// Mutation returns the LinkMutation object of the builder.
func (luo *LinkUpdateOne) Mutation() *LinkMutation {
	return luo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (luo *LinkUpdateOne) ClearOwner() *LinkUpdateOne {
	luo.mutation.ClearOwner()
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LinkUpdateOne) Select(field string, fields ...string) *LinkUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Link entity.
func (luo *LinkUpdateOne) Save(ctx context.Context) (*Link, error) {
	var (
		err  error
		node *Link
	)
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LinkUpdateOne) SaveX(ctx context.Context) *Link {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LinkUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LinkUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LinkUpdateOne) check() error {
	if v, ok := luo.mutation.OriginalURL(); ok {
		if err := link.OriginalURLValidator(v); err != nil {
			return &ValidationError{Name: "original_url", err: fmt.Errorf("ent: validator failed for field \"original_url\": %w", err)}
		}
	}
	if _, ok := luo.mutation.OwnerID(); luo.mutation.OwnerCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (luo *LinkUpdateOne) sqlSave(ctx context.Context) (_node *Link, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: link.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Link.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for _, f := range fields {
			if !link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.LinkID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldLinkID,
		})
	}
	if value, ok := luo.mutation.OriginalURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldOriginalURL,
		})
	}
	if value, ok := luo.mutation.VisitedCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: link.FieldVisitedCount,
		})
	}
	if value, ok := luo.mutation.AddedVisitedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: link.FieldVisitedCount,
		})
	}
	if luo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
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
	if nodes := luo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
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
	_node = &Link{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
