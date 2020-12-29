// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"log"

	"github.com/adnaan/users/internal/models/migrate"
	"github.com/google/uuid"

	"github.com/adnaan/users/internal/models/group"
	"github.com/adnaan/users/internal/models/grouprole"
	"github.com/adnaan/users/internal/models/permission"
	"github.com/adnaan/users/internal/models/session"
	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/adnaan/users/internal/models/workspace"
	"github.com/adnaan/users/internal/models/workspaceinvitation"
	"github.com/adnaan/users/internal/models/workspacerole"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// GroupRole is the client for interacting with the GroupRole builders.
	GroupRole *GroupRoleClient
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserRole is the client for interacting with the UserRole builders.
	UserRole *UserRoleClient
	// Workspace is the client for interacting with the Workspace builders.
	Workspace *WorkspaceClient
	// WorkspaceInvitation is the client for interacting with the WorkspaceInvitation builders.
	WorkspaceInvitation *WorkspaceInvitationClient
	// WorkspaceRole is the client for interacting with the WorkspaceRole builders.
	WorkspaceRole *WorkspaceRoleClient
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
	c.Group = NewGroupClient(c.config)
	c.GroupRole = NewGroupRoleClient(c.config)
	c.Permission = NewPermissionClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserRole = NewUserRoleClient(c.config)
	c.Workspace = NewWorkspaceClient(c.config)
	c.WorkspaceInvitation = NewWorkspaceInvitationClient(c.config)
	c.WorkspaceRole = NewWorkspaceRoleClient(c.config)
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
		return nil, fmt.Errorf("models: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("models: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Group:               NewGroupClient(cfg),
		GroupRole:           NewGroupRoleClient(cfg),
		Permission:          NewPermissionClient(cfg),
		Session:             NewSessionClient(cfg),
		User:                NewUserClient(cfg),
		UserRole:            NewUserRoleClient(cfg),
		Workspace:           NewWorkspaceClient(cfg),
		WorkspaceInvitation: NewWorkspaceInvitationClient(cfg),
		WorkspaceRole:       NewWorkspaceRoleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:              cfg,
		Group:               NewGroupClient(cfg),
		GroupRole:           NewGroupRoleClient(cfg),
		Permission:          NewPermissionClient(cfg),
		Session:             NewSessionClient(cfg),
		User:                NewUserClient(cfg),
		UserRole:            NewUserRoleClient(cfg),
		Workspace:           NewWorkspaceClient(cfg),
		WorkspaceInvitation: NewWorkspaceInvitationClient(cfg),
		WorkspaceRole:       NewWorkspaceRoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Group.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Group.Use(hooks...)
	c.GroupRole.Use(hooks...)
	c.Permission.Use(hooks...)
	c.Session.Use(hooks...)
	c.User.Use(hooks...)
	c.UserRole.Use(hooks...)
	c.Workspace.Use(hooks...)
	c.WorkspaceInvitation.Use(hooks...)
	c.WorkspaceRole.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id uuid.UUID) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id uuid.UUID) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{config: c.config}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id uuid.UUID) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id uuid.UUID) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroupRoles queries the group_roles edge of a Group.
func (c *GroupClient) QueryGroupRoles(gr *Group) *GroupRoleQuery {
	query := &GroupRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(grouprole.Table, grouprole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.GroupRolesTable, group.GroupRolesColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkspaces queries the workspaces edge of a Group.
func (c *GroupClient) QueryWorkspaces(gr *Group) *WorkspaceQuery {
	query := &WorkspaceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, group.WorkspacesTable, group.WorkspacesColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// GroupRoleClient is a client for the GroupRole schema.
type GroupRoleClient struct {
	config
}

// NewGroupRoleClient returns a client for the GroupRole from the given config.
func NewGroupRoleClient(c config) *GroupRoleClient {
	return &GroupRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `grouprole.Hooks(f(g(h())))`.
func (c *GroupRoleClient) Use(hooks ...Hook) {
	c.hooks.GroupRole = append(c.hooks.GroupRole, hooks...)
}

// Create returns a create builder for GroupRole.
func (c *GroupRoleClient) Create() *GroupRoleCreate {
	mutation := newGroupRoleMutation(c.config, OpCreate)
	return &GroupRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GroupRole entities.
func (c *GroupRoleClient) CreateBulk(builders ...*GroupRoleCreate) *GroupRoleCreateBulk {
	return &GroupRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GroupRole.
func (c *GroupRoleClient) Update() *GroupRoleUpdate {
	mutation := newGroupRoleMutation(c.config, OpUpdate)
	return &GroupRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupRoleClient) UpdateOne(gr *GroupRole) *GroupRoleUpdateOne {
	mutation := newGroupRoleMutation(c.config, OpUpdateOne, withGroupRole(gr))
	return &GroupRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupRoleClient) UpdateOneID(id uuid.UUID) *GroupRoleUpdateOne {
	mutation := newGroupRoleMutation(c.config, OpUpdateOne, withGroupRoleID(id))
	return &GroupRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GroupRole.
func (c *GroupRoleClient) Delete() *GroupRoleDelete {
	mutation := newGroupRoleMutation(c.config, OpDelete)
	return &GroupRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupRoleClient) DeleteOne(gr *GroupRole) *GroupRoleDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupRoleClient) DeleteOneID(id uuid.UUID) *GroupRoleDeleteOne {
	builder := c.Delete().Where(grouprole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupRoleDeleteOne{builder}
}

// Query returns a query builder for GroupRole.
func (c *GroupRoleClient) Query() *GroupRoleQuery {
	return &GroupRoleQuery{config: c.config}
}

// Get returns a GroupRole entity by its id.
func (c *GroupRoleClient) Get(ctx context.Context, id uuid.UUID) (*GroupRole, error) {
	return c.Query().Where(grouprole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupRoleClient) GetX(ctx context.Context, id uuid.UUID) *GroupRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroups queries the groups edge of a GroupRole.
func (c *GroupRoleClient) QueryGroups(gr *GroupRole) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprole.Table, grouprole.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grouprole.GroupsTable, grouprole.GroupsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a GroupRole.
func (c *GroupRoleClient) QueryUsers(gr *GroupRole) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprole.Table, grouprole.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grouprole.UsersTable, grouprole.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupRoleClient) Hooks() []Hook {
	return c.hooks.GroupRole
}

// PermissionClient is a client for the Permission schema.
type PermissionClient struct {
	config
}

// NewPermissionClient returns a client for the Permission from the given config.
func NewPermissionClient(c config) *PermissionClient {
	return &PermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `permission.Hooks(f(g(h())))`.
func (c *PermissionClient) Use(hooks ...Hook) {
	c.hooks.Permission = append(c.hooks.Permission, hooks...)
}

// Create returns a create builder for Permission.
func (c *PermissionClient) Create() *PermissionCreate {
	mutation := newPermissionMutation(c.config, OpCreate)
	return &PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Permission entities.
func (c *PermissionClient) CreateBulk(builders ...*PermissionCreate) *PermissionCreateBulk {
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Permission.
func (c *PermissionClient) Update() *PermissionUpdate {
	mutation := newPermissionMutation(c.config, OpUpdate)
	return &PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermissionClient) UpdateOne(pe *Permission) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermission(pe))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermissionClient) UpdateOneID(id uuid.UUID) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermissionID(id))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Permission.
func (c *PermissionClient) Delete() *PermissionDelete {
	mutation := newPermissionMutation(c.config, OpDelete)
	return &PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PermissionClient) DeleteOne(pe *Permission) *PermissionDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PermissionClient) DeleteOneID(id uuid.UUID) *PermissionDeleteOne {
	builder := c.Delete().Where(permission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermissionDeleteOne{builder}
}

// Query returns a query builder for Permission.
func (c *PermissionClient) Query() *PermissionQuery {
	return &PermissionQuery{config: c.config}
}

// Get returns a Permission entity by its id.
func (c *PermissionClient) Get(ctx context.Context, id uuid.UUID) (*Permission, error) {
	return c.Query().Where(permission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermissionClient) GetX(ctx context.Context, id uuid.UUID) *Permission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PermissionClient) Hooks() []Hook {
	return c.hooks.Permission
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a create builder for Session.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id string) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SessionClient) DeleteOneID(id string) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{config: c.config}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id string) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id string) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
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
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorkspace queries the workspace edge of a User.
func (c *UserClient) QueryWorkspace(u *User) *WorkspaceQuery {
	query := &WorkspaceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.WorkspaceTable, user.WorkspaceColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkspaceRoles queries the workspace_roles edge of a User.
func (c *UserClient) QueryWorkspaceRoles(u *User) *WorkspaceRoleQuery {
	query := &WorkspaceRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(workspacerole.Table, workspacerole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.WorkspaceRolesTable, user.WorkspaceRolesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroupRoles queries the group_roles edge of a User.
func (c *UserClient) QueryGroupRoles(u *User) *GroupRoleQuery {
	query := &GroupRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(grouprole.Table, grouprole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.GroupRolesTable, user.GroupRolesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserRoles queries the user_roles edge of a User.
func (c *UserClient) QueryUserRoles(u *User) *UserRoleQuery {
	query := &UserRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(userrole.Table, userrole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.UserRolesTable, user.UserRolesColumn),
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

// UserRoleClient is a client for the UserRole schema.
type UserRoleClient struct {
	config
}

// NewUserRoleClient returns a client for the UserRole from the given config.
func NewUserRoleClient(c config) *UserRoleClient {
	return &UserRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userrole.Hooks(f(g(h())))`.
func (c *UserRoleClient) Use(hooks ...Hook) {
	c.hooks.UserRole = append(c.hooks.UserRole, hooks...)
}

// Create returns a create builder for UserRole.
func (c *UserRoleClient) Create() *UserRoleCreate {
	mutation := newUserRoleMutation(c.config, OpCreate)
	return &UserRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserRole entities.
func (c *UserRoleClient) CreateBulk(builders ...*UserRoleCreate) *UserRoleCreateBulk {
	return &UserRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserRole.
func (c *UserRoleClient) Update() *UserRoleUpdate {
	mutation := newUserRoleMutation(c.config, OpUpdate)
	return &UserRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserRoleClient) UpdateOne(ur *UserRole) *UserRoleUpdateOne {
	mutation := newUserRoleMutation(c.config, OpUpdateOne, withUserRole(ur))
	return &UserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserRoleClient) UpdateOneID(id uuid.UUID) *UserRoleUpdateOne {
	mutation := newUserRoleMutation(c.config, OpUpdateOne, withUserRoleID(id))
	return &UserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserRole.
func (c *UserRoleClient) Delete() *UserRoleDelete {
	mutation := newUserRoleMutation(c.config, OpDelete)
	return &UserRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserRoleClient) DeleteOne(ur *UserRole) *UserRoleDeleteOne {
	return c.DeleteOneID(ur.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserRoleClient) DeleteOneID(id uuid.UUID) *UserRoleDeleteOne {
	builder := c.Delete().Where(userrole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserRoleDeleteOne{builder}
}

// Query returns a query builder for UserRole.
func (c *UserRoleClient) Query() *UserRoleQuery {
	return &UserRoleQuery{config: c.config}
}

// Get returns a UserRole entity by its id.
func (c *UserRoleClient) Get(ctx context.Context, id uuid.UUID) (*UserRole, error) {
	return c.Query().Where(userrole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserRoleClient) GetX(ctx context.Context, id uuid.UUID) *UserRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a UserRole.
func (c *UserRoleClient) QueryUsers(ur *UserRole) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ur.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(userrole.Table, userrole.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userrole.UsersTable, userrole.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(ur.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserRoleClient) Hooks() []Hook {
	return c.hooks.UserRole
}

// WorkspaceClient is a client for the Workspace schema.
type WorkspaceClient struct {
	config
}

// NewWorkspaceClient returns a client for the Workspace from the given config.
func NewWorkspaceClient(c config) *WorkspaceClient {
	return &WorkspaceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workspace.Hooks(f(g(h())))`.
func (c *WorkspaceClient) Use(hooks ...Hook) {
	c.hooks.Workspace = append(c.hooks.Workspace, hooks...)
}

// Create returns a create builder for Workspace.
func (c *WorkspaceClient) Create() *WorkspaceCreate {
	mutation := newWorkspaceMutation(c.config, OpCreate)
	return &WorkspaceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Workspace entities.
func (c *WorkspaceClient) CreateBulk(builders ...*WorkspaceCreate) *WorkspaceCreateBulk {
	return &WorkspaceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Workspace.
func (c *WorkspaceClient) Update() *WorkspaceUpdate {
	mutation := newWorkspaceMutation(c.config, OpUpdate)
	return &WorkspaceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkspaceClient) UpdateOne(w *Workspace) *WorkspaceUpdateOne {
	mutation := newWorkspaceMutation(c.config, OpUpdateOne, withWorkspace(w))
	return &WorkspaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkspaceClient) UpdateOneID(id uuid.UUID) *WorkspaceUpdateOne {
	mutation := newWorkspaceMutation(c.config, OpUpdateOne, withWorkspaceID(id))
	return &WorkspaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Workspace.
func (c *WorkspaceClient) Delete() *WorkspaceDelete {
	mutation := newWorkspaceMutation(c.config, OpDelete)
	return &WorkspaceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkspaceClient) DeleteOne(w *Workspace) *WorkspaceDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkspaceClient) DeleteOneID(id uuid.UUID) *WorkspaceDeleteOne {
	builder := c.Delete().Where(workspace.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkspaceDeleteOne{builder}
}

// Query returns a query builder for Workspace.
func (c *WorkspaceClient) Query() *WorkspaceQuery {
	return &WorkspaceQuery{config: c.config}
}

// Get returns a Workspace entity by its id.
func (c *WorkspaceClient) Get(ctx context.Context, id uuid.UUID) (*Workspace, error) {
	return c.Query().Where(workspace.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkspaceClient) GetX(ctx context.Context, id uuid.UUID) *Workspace {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Workspace.
func (c *WorkspaceClient) QueryOwner(w *Workspace) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, workspace.OwnerTable, workspace.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkspaceRoles queries the workspace_roles edge of a Workspace.
func (c *WorkspaceClient) QueryWorkspaceRoles(w *Workspace) *WorkspaceRoleQuery {
	query := &WorkspaceRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, id),
			sqlgraph.To(workspacerole.Table, workspacerole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workspace.WorkspaceRolesTable, workspace.WorkspaceRolesColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a Workspace.
func (c *WorkspaceClient) QueryGroups(w *Workspace) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workspace.GroupsTable, workspace.GroupsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkspaceClient) Hooks() []Hook {
	return c.hooks.Workspace
}

// WorkspaceInvitationClient is a client for the WorkspaceInvitation schema.
type WorkspaceInvitationClient struct {
	config
}

// NewWorkspaceInvitationClient returns a client for the WorkspaceInvitation from the given config.
func NewWorkspaceInvitationClient(c config) *WorkspaceInvitationClient {
	return &WorkspaceInvitationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workspaceinvitation.Hooks(f(g(h())))`.
func (c *WorkspaceInvitationClient) Use(hooks ...Hook) {
	c.hooks.WorkspaceInvitation = append(c.hooks.WorkspaceInvitation, hooks...)
}

// Create returns a create builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Create() *WorkspaceInvitationCreate {
	mutation := newWorkspaceInvitationMutation(c.config, OpCreate)
	return &WorkspaceInvitationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WorkspaceInvitation entities.
func (c *WorkspaceInvitationClient) CreateBulk(builders ...*WorkspaceInvitationCreate) *WorkspaceInvitationCreateBulk {
	return &WorkspaceInvitationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Update() *WorkspaceInvitationUpdate {
	mutation := newWorkspaceInvitationMutation(c.config, OpUpdate)
	return &WorkspaceInvitationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkspaceInvitationClient) UpdateOne(wi *WorkspaceInvitation) *WorkspaceInvitationUpdateOne {
	mutation := newWorkspaceInvitationMutation(c.config, OpUpdateOne, withWorkspaceInvitation(wi))
	return &WorkspaceInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkspaceInvitationClient) UpdateOneID(id uuid.UUID) *WorkspaceInvitationUpdateOne {
	mutation := newWorkspaceInvitationMutation(c.config, OpUpdateOne, withWorkspaceInvitationID(id))
	return &WorkspaceInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Delete() *WorkspaceInvitationDelete {
	mutation := newWorkspaceInvitationMutation(c.config, OpDelete)
	return &WorkspaceInvitationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkspaceInvitationClient) DeleteOne(wi *WorkspaceInvitation) *WorkspaceInvitationDeleteOne {
	return c.DeleteOneID(wi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkspaceInvitationClient) DeleteOneID(id uuid.UUID) *WorkspaceInvitationDeleteOne {
	builder := c.Delete().Where(workspaceinvitation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkspaceInvitationDeleteOne{builder}
}

// Query returns a query builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Query() *WorkspaceInvitationQuery {
	return &WorkspaceInvitationQuery{config: c.config}
}

// Get returns a WorkspaceInvitation entity by its id.
func (c *WorkspaceInvitationClient) Get(ctx context.Context, id uuid.UUID) (*WorkspaceInvitation, error) {
	return c.Query().Where(workspaceinvitation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkspaceInvitationClient) GetX(ctx context.Context, id uuid.UUID) *WorkspaceInvitation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WorkspaceInvitationClient) Hooks() []Hook {
	return c.hooks.WorkspaceInvitation
}

// WorkspaceRoleClient is a client for the WorkspaceRole schema.
type WorkspaceRoleClient struct {
	config
}

// NewWorkspaceRoleClient returns a client for the WorkspaceRole from the given config.
func NewWorkspaceRoleClient(c config) *WorkspaceRoleClient {
	return &WorkspaceRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workspacerole.Hooks(f(g(h())))`.
func (c *WorkspaceRoleClient) Use(hooks ...Hook) {
	c.hooks.WorkspaceRole = append(c.hooks.WorkspaceRole, hooks...)
}

// Create returns a create builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Create() *WorkspaceRoleCreate {
	mutation := newWorkspaceRoleMutation(c.config, OpCreate)
	return &WorkspaceRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WorkspaceRole entities.
func (c *WorkspaceRoleClient) CreateBulk(builders ...*WorkspaceRoleCreate) *WorkspaceRoleCreateBulk {
	return &WorkspaceRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Update() *WorkspaceRoleUpdate {
	mutation := newWorkspaceRoleMutation(c.config, OpUpdate)
	return &WorkspaceRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkspaceRoleClient) UpdateOne(wr *WorkspaceRole) *WorkspaceRoleUpdateOne {
	mutation := newWorkspaceRoleMutation(c.config, OpUpdateOne, withWorkspaceRole(wr))
	return &WorkspaceRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkspaceRoleClient) UpdateOneID(id uuid.UUID) *WorkspaceRoleUpdateOne {
	mutation := newWorkspaceRoleMutation(c.config, OpUpdateOne, withWorkspaceRoleID(id))
	return &WorkspaceRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Delete() *WorkspaceRoleDelete {
	mutation := newWorkspaceRoleMutation(c.config, OpDelete)
	return &WorkspaceRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkspaceRoleClient) DeleteOne(wr *WorkspaceRole) *WorkspaceRoleDeleteOne {
	return c.DeleteOneID(wr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkspaceRoleClient) DeleteOneID(id uuid.UUID) *WorkspaceRoleDeleteOne {
	builder := c.Delete().Where(workspacerole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkspaceRoleDeleteOne{builder}
}

// Query returns a query builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Query() *WorkspaceRoleQuery {
	return &WorkspaceRoleQuery{config: c.config}
}

// Get returns a WorkspaceRole entity by its id.
func (c *WorkspaceRoleClient) Get(ctx context.Context, id uuid.UUID) (*WorkspaceRole, error) {
	return c.Query().Where(workspacerole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkspaceRoleClient) GetX(ctx context.Context, id uuid.UUID) *WorkspaceRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorkspaces queries the workspaces edge of a WorkspaceRole.
func (c *WorkspaceRoleClient) QueryWorkspaces(wr *WorkspaceRole) *WorkspaceQuery {
	query := &WorkspaceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspacerole.Table, workspacerole.FieldID, id),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspacerole.WorkspacesTable, workspacerole.WorkspacesColumn),
		)
		fromV = sqlgraph.Neighbors(wr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a WorkspaceRole.
func (c *WorkspaceRoleClient) QueryUsers(wr *WorkspaceRole) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspacerole.Table, workspacerole.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspacerole.UsersTable, workspacerole.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(wr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkspaceRoleClient) Hooks() []Hook {
	return c.hooks.WorkspaceRole
}
