// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/tonymj76/savannah/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tonymj76/savannah/ent/commit"
	"github.com/tonymj76/savannah/ent/repository"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Commit is the client for interacting with the Commit builders.
	Commit *CommitClient
	// Repository is the client for interacting with the Repository builders.
	Repository *RepositoryClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Commit = NewCommitClient(c.config)
	c.Repository = NewRepositoryClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Commit:     NewCommitClient(cfg),
		Repository: NewRepositoryClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:        ctx,
		config:     cfg,
		Commit:     NewCommitClient(cfg),
		Repository: NewRepositoryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Commit.
//		Query().
//		Count(ctx)
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
	c.Commit.Use(hooks...)
	c.Repository.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Commit.Intercept(interceptors...)
	c.Repository.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CommitMutation:
		return c.Commit.mutate(ctx, m)
	case *RepositoryMutation:
		return c.Repository.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CommitClient is a client for the Commit schema.
type CommitClient struct {
	config
}

// NewCommitClient returns a client for the Commit from the given config.
func NewCommitClient(c config) *CommitClient {
	return &CommitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `commit.Hooks(f(g(h())))`.
func (c *CommitClient) Use(hooks ...Hook) {
	c.hooks.Commit = append(c.hooks.Commit, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `commit.Intercept(f(g(h())))`.
func (c *CommitClient) Intercept(interceptors ...Interceptor) {
	c.inters.Commit = append(c.inters.Commit, interceptors...)
}

// Create returns a builder for creating a Commit entity.
func (c *CommitClient) Create() *CommitCreate {
	mutation := newCommitMutation(c.config, OpCreate)
	return &CommitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Commit entities.
func (c *CommitClient) CreateBulk(builders ...*CommitCreate) *CommitCreateBulk {
	return &CommitCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CommitClient) MapCreateBulk(slice any, setFunc func(*CommitCreate, int)) *CommitCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CommitCreateBulk{err: fmt.Errorf("calling to CommitClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CommitCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CommitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Commit.
func (c *CommitClient) Update() *CommitUpdate {
	mutation := newCommitMutation(c.config, OpUpdate)
	return &CommitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommitClient) UpdateOne(co *Commit) *CommitUpdateOne {
	mutation := newCommitMutation(c.config, OpUpdateOne, withCommit(co))
	return &CommitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommitClient) UpdateOneID(id int) *CommitUpdateOne {
	mutation := newCommitMutation(c.config, OpUpdateOne, withCommitID(id))
	return &CommitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Commit.
func (c *CommitClient) Delete() *CommitDelete {
	mutation := newCommitMutation(c.config, OpDelete)
	return &CommitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommitClient) DeleteOne(co *Commit) *CommitDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CommitClient) DeleteOneID(id int) *CommitDeleteOne {
	builder := c.Delete().Where(commit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommitDeleteOne{builder}
}

// Query returns a query builder for Commit.
func (c *CommitClient) Query() *CommitQuery {
	return &CommitQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCommit},
		inters: c.Interceptors(),
	}
}

// Get returns a Commit entity by its id.
func (c *CommitClient) Get(ctx context.Context, id int) (*Commit, error) {
	return c.Query().Where(commit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommitClient) GetX(ctx context.Context, id int) *Commit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CommitClient) Hooks() []Hook {
	return c.hooks.Commit
}

// Interceptors returns the client interceptors.
func (c *CommitClient) Interceptors() []Interceptor {
	return c.inters.Commit
}

func (c *CommitClient) mutate(ctx context.Context, m *CommitMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CommitCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CommitUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CommitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CommitDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Commit mutation op: %q", m.Op())
	}
}

// RepositoryClient is a client for the Repository schema.
type RepositoryClient struct {
	config
}

// NewRepositoryClient returns a client for the Repository from the given config.
func NewRepositoryClient(c config) *RepositoryClient {
	return &RepositoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `repository.Hooks(f(g(h())))`.
func (c *RepositoryClient) Use(hooks ...Hook) {
	c.hooks.Repository = append(c.hooks.Repository, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `repository.Intercept(f(g(h())))`.
func (c *RepositoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Repository = append(c.inters.Repository, interceptors...)
}

// Create returns a builder for creating a Repository entity.
func (c *RepositoryClient) Create() *RepositoryCreate {
	mutation := newRepositoryMutation(c.config, OpCreate)
	return &RepositoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Repository entities.
func (c *RepositoryClient) CreateBulk(builders ...*RepositoryCreate) *RepositoryCreateBulk {
	return &RepositoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RepositoryClient) MapCreateBulk(slice any, setFunc func(*RepositoryCreate, int)) *RepositoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RepositoryCreateBulk{err: fmt.Errorf("calling to RepositoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RepositoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RepositoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Repository.
func (c *RepositoryClient) Update() *RepositoryUpdate {
	mutation := newRepositoryMutation(c.config, OpUpdate)
	return &RepositoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RepositoryClient) UpdateOne(r *Repository) *RepositoryUpdateOne {
	mutation := newRepositoryMutation(c.config, OpUpdateOne, withRepository(r))
	return &RepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RepositoryClient) UpdateOneID(id int) *RepositoryUpdateOne {
	mutation := newRepositoryMutation(c.config, OpUpdateOne, withRepositoryID(id))
	return &RepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Repository.
func (c *RepositoryClient) Delete() *RepositoryDelete {
	mutation := newRepositoryMutation(c.config, OpDelete)
	return &RepositoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RepositoryClient) DeleteOne(r *Repository) *RepositoryDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RepositoryClient) DeleteOneID(id int) *RepositoryDeleteOne {
	builder := c.Delete().Where(repository.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RepositoryDeleteOne{builder}
}

// Query returns a query builder for Repository.
func (c *RepositoryClient) Query() *RepositoryQuery {
	return &RepositoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRepository},
		inters: c.Interceptors(),
	}
}

// Get returns a Repository entity by its id.
func (c *RepositoryClient) Get(ctx context.Context, id int) (*Repository, error) {
	return c.Query().Where(repository.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RepositoryClient) GetX(ctx context.Context, id int) *Repository {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCommits queries the commits edge of a Repository.
func (c *RepositoryClient) QueryCommits(r *Repository) *CommitQuery {
	query := (&CommitClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(repository.Table, repository.FieldID, id),
			sqlgraph.To(commit.Table, commit.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, repository.CommitsTable, repository.CommitsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RepositoryClient) Hooks() []Hook {
	return c.hooks.Repository
}

// Interceptors returns the client interceptors.
func (c *RepositoryClient) Interceptors() []Interceptor {
	return c.inters.Repository
}

func (c *RepositoryClient) mutate(ctx context.Context, m *RepositoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RepositoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RepositoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RepositoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Repository mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Commit, Repository []ent.Hook
	}
	inters struct {
		Commit, Repository []ent.Interceptor
	}
)
