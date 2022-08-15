// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"yayar/internal/data/ent/migrate"

	"yayar/internal/data/ent/app"
	"yayar/internal/data/ent/history"
	"yayar/internal/data/ent/user"
	"yayar/internal/data/ent/version"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// App is the client for interacting with the App builders.
	App *AppClient
	// History is the client for interacting with the History builders.
	History *HistoryClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Version is the client for interacting with the Version builders.
	Version *VersionClient
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
	c.App = NewAppClient(c.config)
	c.History = NewHistoryClient(c.config)
	c.User = NewUserClient(c.config)
	c.Version = NewVersionClient(c.config)
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
		ctx:     ctx,
		config:  cfg,
		App:     NewAppClient(cfg),
		History: NewHistoryClient(cfg),
		User:    NewUserClient(cfg),
		Version: NewVersionClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		App:     NewAppClient(cfg),
		History: NewHistoryClient(cfg),
		User:    NewUserClient(cfg),
		Version: NewVersionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		App.
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
	c.App.Use(hooks...)
	c.History.Use(hooks...)
	c.User.Use(hooks...)
	c.Version.Use(hooks...)
}

// AppClient is a client for the App schema.
type AppClient struct {
	config
}

// NewAppClient returns a client for the App from the given config.
func NewAppClient(c config) *AppClient {
	return &AppClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `app.Hooks(f(g(h())))`.
func (c *AppClient) Use(hooks ...Hook) {
	c.hooks.App = append(c.hooks.App, hooks...)
}

// Create returns a create builder for App.
func (c *AppClient) Create() *AppCreate {
	mutation := newAppMutation(c.config, OpCreate)
	return &AppCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of App entities.
func (c *AppClient) CreateBulk(builders ...*AppCreate) *AppCreateBulk {
	return &AppCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for App.
func (c *AppClient) Update() *AppUpdate {
	mutation := newAppMutation(c.config, OpUpdate)
	return &AppUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppClient) UpdateOne(a *App) *AppUpdateOne {
	mutation := newAppMutation(c.config, OpUpdateOne, withApp(a))
	return &AppUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppClient) UpdateOneID(id int) *AppUpdateOne {
	mutation := newAppMutation(c.config, OpUpdateOne, withAppID(id))
	return &AppUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for App.
func (c *AppClient) Delete() *AppDelete {
	mutation := newAppMutation(c.config, OpDelete)
	return &AppDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AppClient) DeleteOne(a *App) *AppDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AppClient) DeleteOneID(id int) *AppDeleteOne {
	builder := c.Delete().Where(app.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppDeleteOne{builder}
}

// Query returns a query builder for App.
func (c *AppClient) Query() *AppQuery {
	return &AppQuery{
		config: c.config,
	}
}

// Get returns a App entity by its id.
func (c *AppClient) Get(ctx context.Context, id int) (*App, error) {
	return c.Query().Where(app.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppClient) GetX(ctx context.Context, id int) *App {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a App.
func (c *AppClient) QueryUser(a *App) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, app.UserTable, app.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVersions queries the versions edge of a App.
func (c *AppClient) QueryVersions(a *App) *VersionQuery {
	query := &VersionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, id),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, app.VersionsTable, app.VersionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AppClient) Hooks() []Hook {
	return c.hooks.App
}

// HistoryClient is a client for the History schema.
type HistoryClient struct {
	config
}

// NewHistoryClient returns a client for the History from the given config.
func NewHistoryClient(c config) *HistoryClient {
	return &HistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `history.Hooks(f(g(h())))`.
func (c *HistoryClient) Use(hooks ...Hook) {
	c.hooks.History = append(c.hooks.History, hooks...)
}

// Create returns a create builder for History.
func (c *HistoryClient) Create() *HistoryCreate {
	mutation := newHistoryMutation(c.config, OpCreate)
	return &HistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of History entities.
func (c *HistoryClient) CreateBulk(builders ...*HistoryCreate) *HistoryCreateBulk {
	return &HistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for History.
func (c *HistoryClient) Update() *HistoryUpdate {
	mutation := newHistoryMutation(c.config, OpUpdate)
	return &HistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HistoryClient) UpdateOne(h *History) *HistoryUpdateOne {
	mutation := newHistoryMutation(c.config, OpUpdateOne, withHistory(h))
	return &HistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HistoryClient) UpdateOneID(id int) *HistoryUpdateOne {
	mutation := newHistoryMutation(c.config, OpUpdateOne, withHistoryID(id))
	return &HistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for History.
func (c *HistoryClient) Delete() *HistoryDelete {
	mutation := newHistoryMutation(c.config, OpDelete)
	return &HistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *HistoryClient) DeleteOne(h *History) *HistoryDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *HistoryClient) DeleteOneID(id int) *HistoryDeleteOne {
	builder := c.Delete().Where(history.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HistoryDeleteOne{builder}
}

// Query returns a query builder for History.
func (c *HistoryClient) Query() *HistoryQuery {
	return &HistoryQuery{
		config: c.config,
	}
}

// Get returns a History entity by its id.
func (c *HistoryClient) Get(ctx context.Context, id int) (*History, error) {
	return c.Query().Where(history.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HistoryClient) GetX(ctx context.Context, id int) *History {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVersion queries the version edge of a History.
func (c *HistoryClient) QueryVersion(h *History) *VersionQuery {
	query := &VersionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(history.Table, history.FieldID, id),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, history.VersionTable, history.VersionPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HistoryClient) Hooks() []Hook {
	return c.hooks.History
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
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
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
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
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
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryApps queries the apps edge of a User.
func (c *UserClient) QueryApps(u *User) *AppQuery {
	query := &AppQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.AppsTable, user.AppsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVersions queries the versions edge of a User.
func (c *UserClient) QueryVersions(u *User) *VersionQuery {
	query := &VersionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.VersionsTable, user.VersionsPrimaryKey...),
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

// VersionClient is a client for the Version schema.
type VersionClient struct {
	config
}

// NewVersionClient returns a client for the Version from the given config.
func NewVersionClient(c config) *VersionClient {
	return &VersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `version.Hooks(f(g(h())))`.
func (c *VersionClient) Use(hooks ...Hook) {
	c.hooks.Version = append(c.hooks.Version, hooks...)
}

// Create returns a create builder for Version.
func (c *VersionClient) Create() *VersionCreate {
	mutation := newVersionMutation(c.config, OpCreate)
	return &VersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Version entities.
func (c *VersionClient) CreateBulk(builders ...*VersionCreate) *VersionCreateBulk {
	return &VersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Version.
func (c *VersionClient) Update() *VersionUpdate {
	mutation := newVersionMutation(c.config, OpUpdate)
	return &VersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VersionClient) UpdateOne(v *Version) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersion(v))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VersionClient) UpdateOneID(id int) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersionID(id))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Version.
func (c *VersionClient) Delete() *VersionDelete {
	mutation := newVersionMutation(c.config, OpDelete)
	return &VersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VersionClient) DeleteOne(v *Version) *VersionDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VersionClient) DeleteOneID(id int) *VersionDeleteOne {
	builder := c.Delete().Where(version.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VersionDeleteOne{builder}
}

// Query returns a query builder for Version.
func (c *VersionClient) Query() *VersionQuery {
	return &VersionQuery{
		config: c.config,
	}
}

// Get returns a Version entity by its id.
func (c *VersionClient) Get(ctx context.Context, id int) (*Version, error) {
	return c.Query().Where(version.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VersionClient) GetX(ctx context.Context, id int) *Version {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Version.
func (c *VersionClient) QueryUser(v *Version) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, version.UserTable, version.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryApp queries the app edge of a Version.
func (c *VersionClient) QueryApp(v *Version) *AppQuery {
	query := &AppQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, id),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, version.AppTable, version.AppPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHistories queries the histories edge of a Version.
func (c *VersionClient) QueryHistories(v *Version) *HistoryQuery {
	query := &HistoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, id),
			sqlgraph.To(history.Table, history.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, version.HistoriesTable, version.HistoriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VersionClient) Hooks() []Hook {
	return c.hooks.Version
}
