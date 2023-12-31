// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"atlas/ent/migrate"

	"atlas/ent/paymentattempt"
	"atlas/ent/paymentintent"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// PaymentAttempt is the client for interacting with the PaymentAttempt builders.
	PaymentAttempt *PaymentAttemptClient
	// PaymentIntent is the client for interacting with the PaymentIntent builders.
	PaymentIntent *PaymentIntentClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.PaymentAttempt = NewPaymentAttemptClient(c.config)
	c.PaymentIntent = NewPaymentIntentClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		PaymentAttempt: NewPaymentAttemptClient(cfg),
		PaymentIntent:  NewPaymentIntentClient(cfg),
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
		ctx:            ctx,
		config:         cfg,
		PaymentAttempt: NewPaymentAttemptClient(cfg),
		PaymentIntent:  NewPaymentIntentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		PaymentAttempt.
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
	c.PaymentAttempt.Use(hooks...)
	c.PaymentIntent.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.PaymentAttempt.Intercept(interceptors...)
	c.PaymentIntent.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PaymentAttemptMutation:
		return c.PaymentAttempt.mutate(ctx, m)
	case *PaymentIntentMutation:
		return c.PaymentIntent.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PaymentAttemptClient is a client for the PaymentAttempt schema.
type PaymentAttemptClient struct {
	config
}

// NewPaymentAttemptClient returns a client for the PaymentAttempt from the given config.
func NewPaymentAttemptClient(c config) *PaymentAttemptClient {
	return &PaymentAttemptClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `paymentattempt.Hooks(f(g(h())))`.
func (c *PaymentAttemptClient) Use(hooks ...Hook) {
	c.hooks.PaymentAttempt = append(c.hooks.PaymentAttempt, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `paymentattempt.Intercept(f(g(h())))`.
func (c *PaymentAttemptClient) Intercept(interceptors ...Interceptor) {
	c.inters.PaymentAttempt = append(c.inters.PaymentAttempt, interceptors...)
}

// Create returns a builder for creating a PaymentAttempt entity.
func (c *PaymentAttemptClient) Create() *PaymentAttemptCreate {
	mutation := newPaymentAttemptMutation(c.config, OpCreate)
	return &PaymentAttemptCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PaymentAttempt entities.
func (c *PaymentAttemptClient) CreateBulk(builders ...*PaymentAttemptCreate) *PaymentAttemptCreateBulk {
	return &PaymentAttemptCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PaymentAttempt.
func (c *PaymentAttemptClient) Update() *PaymentAttemptUpdate {
	mutation := newPaymentAttemptMutation(c.config, OpUpdate)
	return &PaymentAttemptUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaymentAttemptClient) UpdateOne(pa *PaymentAttempt) *PaymentAttemptUpdateOne {
	mutation := newPaymentAttemptMutation(c.config, OpUpdateOne, withPaymentAttempt(pa))
	return &PaymentAttemptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaymentAttemptClient) UpdateOneID(id int) *PaymentAttemptUpdateOne {
	mutation := newPaymentAttemptMutation(c.config, OpUpdateOne, withPaymentAttemptID(id))
	return &PaymentAttemptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PaymentAttempt.
func (c *PaymentAttemptClient) Delete() *PaymentAttemptDelete {
	mutation := newPaymentAttemptMutation(c.config, OpDelete)
	return &PaymentAttemptDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PaymentAttemptClient) DeleteOne(pa *PaymentAttempt) *PaymentAttemptDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PaymentAttemptClient) DeleteOneID(id int) *PaymentAttemptDeleteOne {
	builder := c.Delete().Where(paymentattempt.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaymentAttemptDeleteOne{builder}
}

// Query returns a query builder for PaymentAttempt.
func (c *PaymentAttemptClient) Query() *PaymentAttemptQuery {
	return &PaymentAttemptQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePaymentAttempt},
		inters: c.Interceptors(),
	}
}

// Get returns a PaymentAttempt entity by its id.
func (c *PaymentAttemptClient) Get(ctx context.Context, id int) (*PaymentAttempt, error) {
	return c.Query().Where(paymentattempt.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaymentAttemptClient) GetX(ctx context.Context, id int) *PaymentAttempt {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPaymentIntent queries the payment_intent edge of a PaymentAttempt.
func (c *PaymentAttemptClient) QueryPaymentIntent(pa *PaymentAttempt) *PaymentIntentQuery {
	query := (&PaymentIntentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(paymentattempt.Table, paymentattempt.FieldID, id),
			sqlgraph.To(paymentintent.Table, paymentintent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, paymentattempt.PaymentIntentTable, paymentattempt.PaymentIntentColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PaymentAttemptClient) Hooks() []Hook {
	return c.hooks.PaymentAttempt
}

// Interceptors returns the client interceptors.
func (c *PaymentAttemptClient) Interceptors() []Interceptor {
	return c.inters.PaymentAttempt
}

func (c *PaymentAttemptClient) mutate(ctx context.Context, m *PaymentAttemptMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PaymentAttemptCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PaymentAttemptUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PaymentAttemptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PaymentAttemptDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PaymentAttempt mutation op: %q", m.Op())
	}
}

// PaymentIntentClient is a client for the PaymentIntent schema.
type PaymentIntentClient struct {
	config
}

// NewPaymentIntentClient returns a client for the PaymentIntent from the given config.
func NewPaymentIntentClient(c config) *PaymentIntentClient {
	return &PaymentIntentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `paymentintent.Hooks(f(g(h())))`.
func (c *PaymentIntentClient) Use(hooks ...Hook) {
	c.hooks.PaymentIntent = append(c.hooks.PaymentIntent, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `paymentintent.Intercept(f(g(h())))`.
func (c *PaymentIntentClient) Intercept(interceptors ...Interceptor) {
	c.inters.PaymentIntent = append(c.inters.PaymentIntent, interceptors...)
}

// Create returns a builder for creating a PaymentIntent entity.
func (c *PaymentIntentClient) Create() *PaymentIntentCreate {
	mutation := newPaymentIntentMutation(c.config, OpCreate)
	return &PaymentIntentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PaymentIntent entities.
func (c *PaymentIntentClient) CreateBulk(builders ...*PaymentIntentCreate) *PaymentIntentCreateBulk {
	return &PaymentIntentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PaymentIntent.
func (c *PaymentIntentClient) Update() *PaymentIntentUpdate {
	mutation := newPaymentIntentMutation(c.config, OpUpdate)
	return &PaymentIntentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaymentIntentClient) UpdateOne(pi *PaymentIntent) *PaymentIntentUpdateOne {
	mutation := newPaymentIntentMutation(c.config, OpUpdateOne, withPaymentIntent(pi))
	return &PaymentIntentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaymentIntentClient) UpdateOneID(id int) *PaymentIntentUpdateOne {
	mutation := newPaymentIntentMutation(c.config, OpUpdateOne, withPaymentIntentID(id))
	return &PaymentIntentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PaymentIntent.
func (c *PaymentIntentClient) Delete() *PaymentIntentDelete {
	mutation := newPaymentIntentMutation(c.config, OpDelete)
	return &PaymentIntentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PaymentIntentClient) DeleteOne(pi *PaymentIntent) *PaymentIntentDeleteOne {
	return c.DeleteOneID(pi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PaymentIntentClient) DeleteOneID(id int) *PaymentIntentDeleteOne {
	builder := c.Delete().Where(paymentintent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaymentIntentDeleteOne{builder}
}

// Query returns a query builder for PaymentIntent.
func (c *PaymentIntentClient) Query() *PaymentIntentQuery {
	return &PaymentIntentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePaymentIntent},
		inters: c.Interceptors(),
	}
}

// Get returns a PaymentIntent entity by its id.
func (c *PaymentIntentClient) Get(ctx context.Context, id int) (*PaymentIntent, error) {
	return c.Query().Where(paymentintent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaymentIntentClient) GetX(ctx context.Context, id int) *PaymentIntent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPaymentAttempts queries the payment_attempts edge of a PaymentIntent.
func (c *PaymentIntentClient) QueryPaymentAttempts(pi *PaymentIntent) *PaymentAttemptQuery {
	query := (&PaymentAttemptClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(paymentintent.Table, paymentintent.FieldID, id),
			sqlgraph.To(paymentattempt.Table, paymentattempt.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, paymentintent.PaymentAttemptsTable, paymentintent.PaymentAttemptsColumn),
		)
		fromV = sqlgraph.Neighbors(pi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PaymentIntentClient) Hooks() []Hook {
	return c.hooks.PaymentIntent
}

// Interceptors returns the client interceptors.
func (c *PaymentIntentClient) Interceptors() []Interceptor {
	return c.inters.PaymentIntent
}

func (c *PaymentIntentClient) mutate(ctx context.Context, m *PaymentIntentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PaymentIntentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PaymentIntentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PaymentIntentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PaymentIntentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PaymentIntent mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		PaymentAttempt, PaymentIntent []ent.Hook
	}
	inters struct {
		PaymentAttempt, PaymentIntent []ent.Interceptor
	}
)
