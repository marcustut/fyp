// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/marcustut/fyp/backend/ent/migrate"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"

	"github.com/marcustut/fyp/backend/ent/instance"
	"github.com/marcustut/fyp/backend/ent/link"
	"github.com/marcustut/fyp/backend/ent/slide"
	"github.com/marcustut/fyp/backend/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Instance is the client for interacting with the Instance builders.
	Instance *InstanceClient
	// Link is the client for interacting with the Link builders.
	Link *LinkClient
	// Slide is the client for interacting with the Slide builders.
	Slide *SlideClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Instance = NewInstanceClient(c.config)
	c.Link = NewLinkClient(c.config)
	c.Slide = NewSlideClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Instance: NewInstanceClient(cfg),
		Link:     NewLinkClient(cfg),
		Slide:    NewSlideClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Instance: NewInstanceClient(cfg),
		Link:     NewLinkClient(cfg),
		Slide:    NewSlideClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Instance.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Instance.Use(hooks...)
	c.Link.Use(hooks...)
	c.Slide.Use(hooks...)
	c.User.Use(hooks...)
}

// InstanceClient is a client for the Instance schema.
type InstanceClient struct {
	config
}

// NewInstanceClient returns a client for the Instance from the given config.
func NewInstanceClient(c config) *InstanceClient {
	return &InstanceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instance.Hooks(f(g(h())))`.
func (c *InstanceClient) Use(hooks ...Hook) {
	c.hooks.Instance = append(c.hooks.Instance, hooks...)
}

// Create returns a create builder for Instance.
func (c *InstanceClient) Create() *InstanceCreate {
	mutation := newInstanceMutation(c.config, OpCreate)
	return &InstanceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Instance entities.
func (c *InstanceClient) CreateBulk(builders ...*InstanceCreate) *InstanceCreateBulk {
	return &InstanceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Instance.
func (c *InstanceClient) Update() *InstanceUpdate {
	mutation := newInstanceMutation(c.config, OpUpdate)
	return &InstanceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstanceClient) UpdateOne(i *Instance) *InstanceUpdateOne {
	mutation := newInstanceMutation(c.config, OpUpdateOne, withInstance(i))
	return &InstanceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstanceClient) UpdateOneID(id ulid.ID) *InstanceUpdateOne {
	mutation := newInstanceMutation(c.config, OpUpdateOne, withInstanceID(id))
	return &InstanceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Instance.
func (c *InstanceClient) Delete() *InstanceDelete {
	mutation := newInstanceMutation(c.config, OpDelete)
	return &InstanceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *InstanceClient) DeleteOne(i *Instance) *InstanceDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *InstanceClient) DeleteOneID(id ulid.ID) *InstanceDeleteOne {
	builder := c.Delete().Where(instance.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstanceDeleteOne{builder}
}

// Query returns a query builder for Instance.
func (c *InstanceClient) Query() *InstanceQuery {
	return &InstanceQuery{
		config: c.config,
	}
}

// Get returns a Instance entity by its id.
func (c *InstanceClient) Get(ctx context.Context, id ulid.ID) (*Instance, error) {
	return c.Query().Where(instance.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstanceClient) GetX(ctx context.Context, id ulid.ID) *Instance {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Instance.
func (c *InstanceClient) QueryUser(i *Instance) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instance.UserTable, instance.UserColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySlide queries the slide edge of a Instance.
func (c *InstanceClient) QuerySlide(i *Instance) *SlideQuery {
	query := &SlideQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, id),
			sqlgraph.To(slide.Table, slide.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, instance.SlideTable, instance.SlideColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstanceClient) Hooks() []Hook {
	return c.hooks.Instance
}

// LinkClient is a client for the Link schema.
type LinkClient struct {
	config
}

// NewLinkClient returns a client for the Link from the given config.
func NewLinkClient(c config) *LinkClient {
	return &LinkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `link.Hooks(f(g(h())))`.
func (c *LinkClient) Use(hooks ...Hook) {
	c.hooks.Link = append(c.hooks.Link, hooks...)
}

// Create returns a create builder for Link.
func (c *LinkClient) Create() *LinkCreate {
	mutation := newLinkMutation(c.config, OpCreate)
	return &LinkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Link entities.
func (c *LinkClient) CreateBulk(builders ...*LinkCreate) *LinkCreateBulk {
	return &LinkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Link.
func (c *LinkClient) Update() *LinkUpdate {
	mutation := newLinkMutation(c.config, OpUpdate)
	return &LinkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LinkClient) UpdateOne(l *Link) *LinkUpdateOne {
	mutation := newLinkMutation(c.config, OpUpdateOne, withLink(l))
	return &LinkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LinkClient) UpdateOneID(id ulid.ID) *LinkUpdateOne {
	mutation := newLinkMutation(c.config, OpUpdateOne, withLinkID(id))
	return &LinkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Link.
func (c *LinkClient) Delete() *LinkDelete {
	mutation := newLinkMutation(c.config, OpDelete)
	return &LinkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LinkClient) DeleteOne(l *Link) *LinkDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LinkClient) DeleteOneID(id ulid.ID) *LinkDeleteOne {
	builder := c.Delete().Where(link.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LinkDeleteOne{builder}
}

// Query returns a query builder for Link.
func (c *LinkClient) Query() *LinkQuery {
	return &LinkQuery{
		config: c.config,
	}
}

// Get returns a Link entity by its id.
func (c *LinkClient) Get(ctx context.Context, id ulid.ID) (*Link, error) {
	return c.Query().Where(link.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LinkClient) GetX(ctx context.Context, id ulid.ID) *Link {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Link.
func (c *LinkClient) QueryOwner(l *Link) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(link.Table, link.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, link.OwnerTable, link.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LinkClient) Hooks() []Hook {
	return c.hooks.Link
}

// SlideClient is a client for the Slide schema.
type SlideClient struct {
	config
}

// NewSlideClient returns a client for the Slide from the given config.
func NewSlideClient(c config) *SlideClient {
	return &SlideClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `slide.Hooks(f(g(h())))`.
func (c *SlideClient) Use(hooks ...Hook) {
	c.hooks.Slide = append(c.hooks.Slide, hooks...)
}

// Create returns a create builder for Slide.
func (c *SlideClient) Create() *SlideCreate {
	mutation := newSlideMutation(c.config, OpCreate)
	return &SlideCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Slide entities.
func (c *SlideClient) CreateBulk(builders ...*SlideCreate) *SlideCreateBulk {
	return &SlideCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Slide.
func (c *SlideClient) Update() *SlideUpdate {
	mutation := newSlideMutation(c.config, OpUpdate)
	return &SlideUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SlideClient) UpdateOne(s *Slide) *SlideUpdateOne {
	mutation := newSlideMutation(c.config, OpUpdateOne, withSlide(s))
	return &SlideUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SlideClient) UpdateOneID(id ulid.ID) *SlideUpdateOne {
	mutation := newSlideMutation(c.config, OpUpdateOne, withSlideID(id))
	return &SlideUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Slide.
func (c *SlideClient) Delete() *SlideDelete {
	mutation := newSlideMutation(c.config, OpDelete)
	return &SlideDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SlideClient) DeleteOne(s *Slide) *SlideDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SlideClient) DeleteOneID(id ulid.ID) *SlideDeleteOne {
	builder := c.Delete().Where(slide.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SlideDeleteOne{builder}
}

// Query returns a query builder for Slide.
func (c *SlideClient) Query() *SlideQuery {
	return &SlideQuery{
		config: c.config,
	}
}

// Get returns a Slide entity by its id.
func (c *SlideClient) Get(ctx context.Context, id ulid.ID) (*Slide, error) {
	return c.Query().Where(slide.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SlideClient) GetX(ctx context.Context, id ulid.ID) *Slide {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstance queries the instance edge of a Slide.
func (c *SlideClient) QueryInstance(s *Slide) *InstanceQuery {
	query := &InstanceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(slide.Table, slide.FieldID, id),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, slide.InstanceTable, slide.InstanceColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Slide.
func (c *SlideClient) QueryUser(s *Slide) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(slide.Table, slide.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, slide.UserTable, slide.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SlideClient) Hooks() []Hook {
	return c.hooks.Slide
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id ulid.ID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id ulid.ID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id ulid.ID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id ulid.ID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstances queries the instances edge of a User.
func (c *UserClient) QueryInstances(u *User) *InstanceQuery {
	query := &InstanceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.InstancesTable, user.InstancesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySlides queries the slides edge of a User.
func (c *UserClient) QuerySlides(u *User) *SlideQuery {
	query := &SlideQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(slide.Table, slide.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SlidesTable, user.SlidesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLinks queries the links edge of a User.
func (c *UserClient) QueryLinks(u *User) *LinkQuery {
	query := &LinkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(link.Table, link.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.LinksTable, user.LinksColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
