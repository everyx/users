// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/adnaan/users/internal/models/group"
	"github.com/adnaan/users/internal/models/grouprole"
	"github.com/adnaan/users/internal/models/permission"
	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/session"
	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/adnaan/users/internal/models/workspace"
	"github.com/adnaan/users/internal/models/workspacerole"
	"github.com/google/uuid"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGroup         = "Group"
	TypeGroupRole     = "GroupRole"
	TypePermission    = "Permission"
	TypeSession       = "Session"
	TypeUser          = "User"
	TypeUserRole      = "UserRole"
	TypeWorkspace     = "Workspace"
	TypeWorkspaceRole = "WorkspaceRole"
)

// GroupMutation represents an operation that mutate the Groups
// nodes in the graph.
type GroupMutation struct {
	config
	op                 Op
	typ                string
	id                 *uuid.UUID
	name               *string
	description        *string
	metadata           *map[string]interface{}
	created_at         *time.Time
	updated_at         *time.Time
	clearedFields      map[string]struct{}
	group_roles        map[uuid.UUID]struct{}
	removedgroup_roles map[uuid.UUID]struct{}
	clearedgroup_roles bool
	workspaces         *uuid.UUID
	clearedworkspaces  bool
	done               bool
	oldValue           func(context.Context) (*Group, error)
	predicates         []predicate.Group
}

var _ ent.Mutation = (*GroupMutation)(nil)

// groupOption allows to manage the mutation configuration using functional options.
type groupOption func(*GroupMutation)

// newGroupMutation creates new mutation for Group.
func newGroupMutation(c config, op Op, opts ...groupOption) *GroupMutation {
	m := &GroupMutation{
		config:        c,
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGroupID sets the id field of the mutation.
func withGroupID(id uuid.UUID) groupOption {
	return func(m *GroupMutation) {
		var (
			err   error
			once  sync.Once
			value *Group
		)
		m.oldValue = func(ctx context.Context) (*Group, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Group.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGroup sets the old Group of the mutation.
func withGroup(node *Group) groupOption {
	return func(m *GroupMutation) {
		m.oldValue = func(context.Context) (*Group, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Group creation.
func (m *GroupMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *GroupMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *GroupMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *GroupMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Group.
// If the Group object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *GroupMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the description field.
func (m *GroupMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the description value in the mutation.
func (m *GroupMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old description value of the Group.
// If the Group object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDescription is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of description.
func (m *GroupMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[group.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the field description was cleared in this mutation.
func (m *GroupMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[group.FieldDescription]
	return ok
}

// ResetDescription reset all changes of the "description" field.
func (m *GroupMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, group.FieldDescription)
}

// SetMetadata sets the metadata field.
func (m *GroupMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *GroupMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old metadata value of the Group.
// If the Group object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupMutation) OldMetadata(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMetadata is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ClearMetadata clears the value of metadata.
func (m *GroupMutation) ClearMetadata() {
	m.metadata = nil
	m.clearedFields[group.FieldMetadata] = struct{}{}
}

// MetadataCleared returns if the field metadata was cleared in this mutation.
func (m *GroupMutation) MetadataCleared() bool {
	_, ok := m.clearedFields[group.FieldMetadata]
	return ok
}

// ResetMetadata reset all changes of the "metadata" field.
func (m *GroupMutation) ResetMetadata() {
	m.metadata = nil
	delete(m.clearedFields, group.FieldMetadata)
}

// SetCreatedAt sets the created_at field.
func (m *GroupMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *GroupMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Group.
// If the Group object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *GroupMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *GroupMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *GroupMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Group.
// If the Group object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *GroupMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddGroupRoleIDs adds the group_roles edge to GroupRole by ids.
func (m *GroupMutation) AddGroupRoleIDs(ids ...uuid.UUID) {
	if m.group_roles == nil {
		m.group_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.group_roles[ids[i]] = struct{}{}
	}
}

// ClearGroupRoles clears the group_roles edge to GroupRole.
func (m *GroupMutation) ClearGroupRoles() {
	m.clearedgroup_roles = true
}

// GroupRolesCleared returns if the edge group_roles was cleared.
func (m *GroupMutation) GroupRolesCleared() bool {
	return m.clearedgroup_roles
}

// RemoveGroupRoleIDs removes the group_roles edge to GroupRole by ids.
func (m *GroupMutation) RemoveGroupRoleIDs(ids ...uuid.UUID) {
	if m.removedgroup_roles == nil {
		m.removedgroup_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedgroup_roles[ids[i]] = struct{}{}
	}
}

// RemovedGroupRoles returns the removed ids of group_roles.
func (m *GroupMutation) RemovedGroupRolesIDs() (ids []uuid.UUID) {
	for id := range m.removedgroup_roles {
		ids = append(ids, id)
	}
	return
}

// GroupRolesIDs returns the group_roles ids in the mutation.
func (m *GroupMutation) GroupRolesIDs() (ids []uuid.UUID) {
	for id := range m.group_roles {
		ids = append(ids, id)
	}
	return
}

// ResetGroupRoles reset all changes of the "group_roles" edge.
func (m *GroupMutation) ResetGroupRoles() {
	m.group_roles = nil
	m.clearedgroup_roles = false
	m.removedgroup_roles = nil
}

// SetWorkspacesID sets the workspaces edge to Workspace by id.
func (m *GroupMutation) SetWorkspacesID(id uuid.UUID) {
	m.workspaces = &id
}

// ClearWorkspaces clears the workspaces edge to Workspace.
func (m *GroupMutation) ClearWorkspaces() {
	m.clearedworkspaces = true
}

// WorkspacesCleared returns if the edge workspaces was cleared.
func (m *GroupMutation) WorkspacesCleared() bool {
	return m.clearedworkspaces
}

// WorkspacesID returns the workspaces id in the mutation.
func (m *GroupMutation) WorkspacesID() (id uuid.UUID, exists bool) {
	if m.workspaces != nil {
		return *m.workspaces, true
	}
	return
}

// WorkspacesIDs returns the workspaces ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// WorkspacesID instead. It exists only for internal usage by the builders.
func (m *GroupMutation) WorkspacesIDs() (ids []uuid.UUID) {
	if id := m.workspaces; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetWorkspaces reset all changes of the "workspaces" edge.
func (m *GroupMutation) ResetWorkspaces() {
	m.workspaces = nil
	m.clearedworkspaces = false
}

// Op returns the operation name.
func (m *GroupMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, group.FieldName)
	}
	if m.description != nil {
		fields = append(fields, group.FieldDescription)
	}
	if m.metadata != nil {
		fields = append(fields, group.FieldMetadata)
	}
	if m.created_at != nil {
		fields = append(fields, group.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, group.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case group.FieldName:
		return m.Name()
	case group.FieldDescription:
		return m.Description()
	case group.FieldMetadata:
		return m.Metadata()
	case group.FieldCreatedAt:
		return m.CreatedAt()
	case group.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *GroupMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case group.FieldName:
		return m.OldName(ctx)
	case group.FieldDescription:
		return m.OldDescription(ctx)
	case group.FieldMetadata:
		return m.OldMetadata(ctx)
	case group.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case group.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Group field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case group.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case group.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case group.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case group.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case group.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *GroupMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *GroupMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(group.FieldDescription) {
		fields = append(fields, group.FieldDescription)
	}
	if m.FieldCleared(group.FieldMetadata) {
		fields = append(fields, group.FieldMetadata)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	switch name {
	case group.FieldDescription:
		m.ClearDescription()
		return nil
	case group.FieldMetadata:
		m.ClearMetadata()
		return nil
	}
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	case group.FieldName:
		m.ResetName()
		return nil
	case group.FieldDescription:
		m.ResetDescription()
		return nil
	case group.FieldMetadata:
		m.ResetMetadata()
		return nil
	case group.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case group.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.group_roles != nil {
		edges = append(edges, group.EdgeGroupRoles)
	}
	if m.workspaces != nil {
		edges = append(edges, group.EdgeWorkspaces)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeGroupRoles:
		ids := make([]ent.Value, 0, len(m.group_roles))
		for id := range m.group_roles {
			ids = append(ids, id)
		}
		return ids
	case group.EdgeWorkspaces:
		if id := m.workspaces; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedgroup_roles != nil {
		edges = append(edges, group.EdgeGroupRoles)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeGroupRoles:
		ids := make([]ent.Value, 0, len(m.removedgroup_roles))
		for id := range m.removedgroup_roles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedgroup_roles {
		edges = append(edges, group.EdgeGroupRoles)
	}
	if m.clearedworkspaces {
		edges = append(edges, group.EdgeWorkspaces)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	switch name {
	case group.EdgeGroupRoles:
		return m.clearedgroup_roles
	case group.EdgeWorkspaces:
		return m.clearedworkspaces
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	switch name {
	case group.EdgeWorkspaces:
		m.ClearWorkspaces()
		return nil
	}
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	switch name {
	case group.EdgeGroupRoles:
		m.ResetGroupRoles()
		return nil
	case group.EdgeWorkspaces:
		m.ResetWorkspaces()
		return nil
	}
	return fmt.Errorf("unknown Group edge %s", name)
}

// GroupRoleMutation represents an operation that mutate the GroupRoles
// nodes in the graph.
type GroupRoleMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	group_id      *uuid.UUID
	user_id       *uuid.UUID
	metadata      *map[string]interface{}
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	groups        *uuid.UUID
	clearedgroups bool
	users         *uuid.UUID
	clearedusers  bool
	done          bool
	oldValue      func(context.Context) (*GroupRole, error)
	predicates    []predicate.GroupRole
}

var _ ent.Mutation = (*GroupRoleMutation)(nil)

// grouproleOption allows to manage the mutation configuration using functional options.
type grouproleOption func(*GroupRoleMutation)

// newGroupRoleMutation creates new mutation for GroupRole.
func newGroupRoleMutation(c config, op Op, opts ...grouproleOption) *GroupRoleMutation {
	m := &GroupRoleMutation{
		config:        c,
		op:            op,
		typ:           TypeGroupRole,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGroupRoleID sets the id field of the mutation.
func withGroupRoleID(id uuid.UUID) grouproleOption {
	return func(m *GroupRoleMutation) {
		var (
			err   error
			once  sync.Once
			value *GroupRole
		)
		m.oldValue = func(ctx context.Context) (*GroupRole, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().GroupRole.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGroupRole sets the old GroupRole of the mutation.
func withGroupRole(node *GroupRole) grouproleOption {
	return func(m *GroupRoleMutation) {
		m.oldValue = func(context.Context) (*GroupRole, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupRoleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupRoleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on GroupRole creation.
func (m *GroupRoleMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *GroupRoleMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *GroupRoleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *GroupRoleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the GroupRole.
// If the GroupRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupRoleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *GroupRoleMutation) ResetName() {
	m.name = nil
}

// SetGroupID sets the group_id field.
func (m *GroupRoleMutation) SetGroupID(u uuid.UUID) {
	m.group_id = &u
}

// GroupID returns the group_id value in the mutation.
func (m *GroupRoleMutation) GroupID() (r uuid.UUID, exists bool) {
	v := m.group_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old group_id value of the GroupRole.
// If the GroupRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupRoleMutation) OldGroupID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldGroupID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldGroupID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGroupID: %w", err)
	}
	return oldValue.GroupID, nil
}

// ResetGroupID reset all changes of the "group_id" field.
func (m *GroupRoleMutation) ResetGroupID() {
	m.group_id = nil
}

// SetUserID sets the user_id field.
func (m *GroupRoleMutation) SetUserID(u uuid.UUID) {
	m.user_id = &u
}

// UserID returns the user_id value in the mutation.
func (m *GroupRoleMutation) UserID() (r uuid.UUID, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old user_id value of the GroupRole.
// If the GroupRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupRoleMutation) OldUserID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID reset all changes of the "user_id" field.
func (m *GroupRoleMutation) ResetUserID() {
	m.user_id = nil
}

// SetMetadata sets the metadata field.
func (m *GroupRoleMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *GroupRoleMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old metadata value of the GroupRole.
// If the GroupRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupRoleMutation) OldMetadata(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMetadata is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ClearMetadata clears the value of metadata.
func (m *GroupRoleMutation) ClearMetadata() {
	m.metadata = nil
	m.clearedFields[grouprole.FieldMetadata] = struct{}{}
}

// MetadataCleared returns if the field metadata was cleared in this mutation.
func (m *GroupRoleMutation) MetadataCleared() bool {
	_, ok := m.clearedFields[grouprole.FieldMetadata]
	return ok
}

// ResetMetadata reset all changes of the "metadata" field.
func (m *GroupRoleMutation) ResetMetadata() {
	m.metadata = nil
	delete(m.clearedFields, grouprole.FieldMetadata)
}

// SetCreatedAt sets the created_at field.
func (m *GroupRoleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *GroupRoleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the GroupRole.
// If the GroupRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupRoleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *GroupRoleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *GroupRoleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *GroupRoleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the GroupRole.
// If the GroupRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupRoleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *GroupRoleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetGroupsID sets the groups edge to Group by id.
func (m *GroupRoleMutation) SetGroupsID(id uuid.UUID) {
	m.groups = &id
}

// ClearGroups clears the groups edge to Group.
func (m *GroupRoleMutation) ClearGroups() {
	m.clearedgroups = true
}

// GroupsCleared returns if the edge groups was cleared.
func (m *GroupRoleMutation) GroupsCleared() bool {
	return m.clearedgroups
}

// GroupsID returns the groups id in the mutation.
func (m *GroupRoleMutation) GroupsID() (id uuid.UUID, exists bool) {
	if m.groups != nil {
		return *m.groups, true
	}
	return
}

// GroupsIDs returns the groups ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// GroupsID instead. It exists only for internal usage by the builders.
func (m *GroupRoleMutation) GroupsIDs() (ids []uuid.UUID) {
	if id := m.groups; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetGroups reset all changes of the "groups" edge.
func (m *GroupRoleMutation) ResetGroups() {
	m.groups = nil
	m.clearedgroups = false
}

// SetUsersID sets the users edge to User by id.
func (m *GroupRoleMutation) SetUsersID(id uuid.UUID) {
	m.users = &id
}

// ClearUsers clears the users edge to User.
func (m *GroupRoleMutation) ClearUsers() {
	m.clearedusers = true
}

// UsersCleared returns if the edge users was cleared.
func (m *GroupRoleMutation) UsersCleared() bool {
	return m.clearedusers
}

// UsersID returns the users id in the mutation.
func (m *GroupRoleMutation) UsersID() (id uuid.UUID, exists bool) {
	if m.users != nil {
		return *m.users, true
	}
	return
}

// UsersIDs returns the users ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// UsersID instead. It exists only for internal usage by the builders.
func (m *GroupRoleMutation) UsersIDs() (ids []uuid.UUID) {
	if id := m.users; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUsers reset all changes of the "users" edge.
func (m *GroupRoleMutation) ResetUsers() {
	m.users = nil
	m.clearedusers = false
}

// Op returns the operation name.
func (m *GroupRoleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (GroupRole).
func (m *GroupRoleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *GroupRoleMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.name != nil {
		fields = append(fields, grouprole.FieldName)
	}
	if m.group_id != nil {
		fields = append(fields, grouprole.FieldGroupID)
	}
	if m.user_id != nil {
		fields = append(fields, grouprole.FieldUserID)
	}
	if m.metadata != nil {
		fields = append(fields, grouprole.FieldMetadata)
	}
	if m.created_at != nil {
		fields = append(fields, grouprole.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, grouprole.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *GroupRoleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case grouprole.FieldName:
		return m.Name()
	case grouprole.FieldGroupID:
		return m.GroupID()
	case grouprole.FieldUserID:
		return m.UserID()
	case grouprole.FieldMetadata:
		return m.Metadata()
	case grouprole.FieldCreatedAt:
		return m.CreatedAt()
	case grouprole.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *GroupRoleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case grouprole.FieldName:
		return m.OldName(ctx)
	case grouprole.FieldGroupID:
		return m.OldGroupID(ctx)
	case grouprole.FieldUserID:
		return m.OldUserID(ctx)
	case grouprole.FieldMetadata:
		return m.OldMetadata(ctx)
	case grouprole.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case grouprole.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown GroupRole field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupRoleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case grouprole.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case grouprole.FieldGroupID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case grouprole.FieldUserID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case grouprole.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case grouprole.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case grouprole.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown GroupRole field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *GroupRoleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *GroupRoleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupRoleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown GroupRole numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *GroupRoleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(grouprole.FieldMetadata) {
		fields = append(fields, grouprole.FieldMetadata)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *GroupRoleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupRoleMutation) ClearField(name string) error {
	switch name {
	case grouprole.FieldMetadata:
		m.ClearMetadata()
		return nil
	}
	return fmt.Errorf("unknown GroupRole nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *GroupRoleMutation) ResetField(name string) error {
	switch name {
	case grouprole.FieldName:
		m.ResetName()
		return nil
	case grouprole.FieldGroupID:
		m.ResetGroupID()
		return nil
	case grouprole.FieldUserID:
		m.ResetUserID()
		return nil
	case grouprole.FieldMetadata:
		m.ResetMetadata()
		return nil
	case grouprole.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case grouprole.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown GroupRole field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *GroupRoleMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.groups != nil {
		edges = append(edges, grouprole.EdgeGroups)
	}
	if m.users != nil {
		edges = append(edges, grouprole.EdgeUsers)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *GroupRoleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case grouprole.EdgeGroups:
		if id := m.groups; id != nil {
			return []ent.Value{*id}
		}
	case grouprole.EdgeUsers:
		if id := m.users; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *GroupRoleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *GroupRoleMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *GroupRoleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedgroups {
		edges = append(edges, grouprole.EdgeGroups)
	}
	if m.clearedusers {
		edges = append(edges, grouprole.EdgeUsers)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *GroupRoleMutation) EdgeCleared(name string) bool {
	switch name {
	case grouprole.EdgeGroups:
		return m.clearedgroups
	case grouprole.EdgeUsers:
		return m.clearedusers
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *GroupRoleMutation) ClearEdge(name string) error {
	switch name {
	case grouprole.EdgeGroups:
		m.ClearGroups()
		return nil
	case grouprole.EdgeUsers:
		m.ClearUsers()
		return nil
	}
	return fmt.Errorf("unknown GroupRole unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *GroupRoleMutation) ResetEdge(name string) error {
	switch name {
	case grouprole.EdgeGroups:
		m.ResetGroups()
		return nil
	case grouprole.EdgeUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown GroupRole edge %s", name)
}

// PermissionMutation represents an operation that mutate the Permissions
// nodes in the graph.
type PermissionMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	role_id       *uuid.UUID
	action        *string
	target        *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Permission, error)
	predicates    []predicate.Permission
}

var _ ent.Mutation = (*PermissionMutation)(nil)

// permissionOption allows to manage the mutation configuration using functional options.
type permissionOption func(*PermissionMutation)

// newPermissionMutation creates new mutation for Permission.
func newPermissionMutation(c config, op Op, opts ...permissionOption) *PermissionMutation {
	m := &PermissionMutation{
		config:        c,
		op:            op,
		typ:           TypePermission,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPermissionID sets the id field of the mutation.
func withPermissionID(id uuid.UUID) permissionOption {
	return func(m *PermissionMutation) {
		var (
			err   error
			once  sync.Once
			value *Permission
		)
		m.oldValue = func(ctx context.Context) (*Permission, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Permission.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPermission sets the old Permission of the mutation.
func withPermission(node *Permission) permissionOption {
	return func(m *PermissionMutation) {
		m.oldValue = func(context.Context) (*Permission, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PermissionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PermissionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Permission creation.
func (m *PermissionMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *PermissionMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetRoleID sets the role_id field.
func (m *PermissionMutation) SetRoleID(u uuid.UUID) {
	m.role_id = &u
}

// RoleID returns the role_id value in the mutation.
func (m *PermissionMutation) RoleID() (r uuid.UUID, exists bool) {
	v := m.role_id
	if v == nil {
		return
	}
	return *v, true
}

// OldRoleID returns the old role_id value of the Permission.
// If the Permission object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PermissionMutation) OldRoleID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRoleID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRoleID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRoleID: %w", err)
	}
	return oldValue.RoleID, nil
}

// ResetRoleID reset all changes of the "role_id" field.
func (m *PermissionMutation) ResetRoleID() {
	m.role_id = nil
}

// SetAction sets the action field.
func (m *PermissionMutation) SetAction(s string) {
	m.action = &s
}

// Action returns the action value in the mutation.
func (m *PermissionMutation) Action() (r string, exists bool) {
	v := m.action
	if v == nil {
		return
	}
	return *v, true
}

// OldAction returns the old action value of the Permission.
// If the Permission object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PermissionMutation) OldAction(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAction is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAction requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAction: %w", err)
	}
	return oldValue.Action, nil
}

// ResetAction reset all changes of the "action" field.
func (m *PermissionMutation) ResetAction() {
	m.action = nil
}

// SetTarget sets the target field.
func (m *PermissionMutation) SetTarget(s string) {
	m.target = &s
}

// Target returns the target value in the mutation.
func (m *PermissionMutation) Target() (r string, exists bool) {
	v := m.target
	if v == nil {
		return
	}
	return *v, true
}

// OldTarget returns the old target value of the Permission.
// If the Permission object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PermissionMutation) OldTarget(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTarget is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTarget requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTarget: %w", err)
	}
	return oldValue.Target, nil
}

// ResetTarget reset all changes of the "target" field.
func (m *PermissionMutation) ResetTarget() {
	m.target = nil
}

// SetCreatedAt sets the created_at field.
func (m *PermissionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *PermissionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Permission.
// If the Permission object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PermissionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *PermissionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *PermissionMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *PermissionMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Permission.
// If the Permission object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PermissionMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *PermissionMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Op returns the operation name.
func (m *PermissionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Permission).
func (m *PermissionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *PermissionMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.role_id != nil {
		fields = append(fields, permission.FieldRoleID)
	}
	if m.action != nil {
		fields = append(fields, permission.FieldAction)
	}
	if m.target != nil {
		fields = append(fields, permission.FieldTarget)
	}
	if m.created_at != nil {
		fields = append(fields, permission.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, permission.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *PermissionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case permission.FieldRoleID:
		return m.RoleID()
	case permission.FieldAction:
		return m.Action()
	case permission.FieldTarget:
		return m.Target()
	case permission.FieldCreatedAt:
		return m.CreatedAt()
	case permission.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *PermissionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case permission.FieldRoleID:
		return m.OldRoleID(ctx)
	case permission.FieldAction:
		return m.OldAction(ctx)
	case permission.FieldTarget:
		return m.OldTarget(ctx)
	case permission.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case permission.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Permission field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PermissionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case permission.FieldRoleID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRoleID(v)
		return nil
	case permission.FieldAction:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAction(v)
		return nil
	case permission.FieldTarget:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTarget(v)
		return nil
	case permission.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case permission.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Permission field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *PermissionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *PermissionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PermissionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Permission numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *PermissionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *PermissionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *PermissionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Permission nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *PermissionMutation) ResetField(name string) error {
	switch name {
	case permission.FieldRoleID:
		m.ResetRoleID()
		return nil
	case permission.FieldAction:
		m.ResetAction()
		return nil
	case permission.FieldTarget:
		m.ResetTarget()
		return nil
	case permission.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case permission.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Permission field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *PermissionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *PermissionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *PermissionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *PermissionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *PermissionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *PermissionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *PermissionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Permission unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *PermissionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Permission edge %s", name)
}

// SessionMutation represents an operation that mutate the Sessions
// nodes in the graph.
type SessionMutation struct {
	config
	op            Op
	typ           string
	id            *string
	data          *string
	created_at    *time.Time
	updated_at    *time.Time
	expires_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Session, error)
	predicates    []predicate.Session
}

var _ ent.Mutation = (*SessionMutation)(nil)

// sessionOption allows to manage the mutation configuration using functional options.
type sessionOption func(*SessionMutation)

// newSessionMutation creates new mutation for Session.
func newSessionMutation(c config, op Op, opts ...sessionOption) *SessionMutation {
	m := &SessionMutation{
		config:        c,
		op:            op,
		typ:           TypeSession,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSessionID sets the id field of the mutation.
func withSessionID(id string) sessionOption {
	return func(m *SessionMutation) {
		var (
			err   error
			once  sync.Once
			value *Session
		)
		m.oldValue = func(ctx context.Context) (*Session, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Session.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSession sets the old Session of the mutation.
func withSession(node *Session) sessionOption {
	return func(m *SessionMutation) {
		m.oldValue = func(context.Context) (*Session, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SessionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SessionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Session creation.
func (m *SessionMutation) SetID(id string) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *SessionMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetData sets the data field.
func (m *SessionMutation) SetData(s string) {
	m.data = &s
}

// Data returns the data value in the mutation.
func (m *SessionMutation) Data() (r string, exists bool) {
	v := m.data
	if v == nil {
		return
	}
	return *v, true
}

// OldData returns the old data value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldData(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldData is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldData requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldData: %w", err)
	}
	return oldValue.Data, nil
}

// ResetData reset all changes of the "data" field.
func (m *SessionMutation) ResetData() {
	m.data = nil
}

// SetCreatedAt sets the created_at field.
func (m *SessionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *SessionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *SessionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *SessionMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *SessionMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *SessionMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetExpiresAt sets the expires_at field.
func (m *SessionMutation) SetExpiresAt(t time.Time) {
	m.expires_at = &t
}

// ExpiresAt returns the expires_at value in the mutation.
func (m *SessionMutation) ExpiresAt() (r time.Time, exists bool) {
	v := m.expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldExpiresAt returns the old expires_at value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldExpiresAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldExpiresAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpiresAt: %w", err)
	}
	return oldValue.ExpiresAt, nil
}

// ClearExpiresAt clears the value of expires_at.
func (m *SessionMutation) ClearExpiresAt() {
	m.expires_at = nil
	m.clearedFields[session.FieldExpiresAt] = struct{}{}
}

// ExpiresAtCleared returns if the field expires_at was cleared in this mutation.
func (m *SessionMutation) ExpiresAtCleared() bool {
	_, ok := m.clearedFields[session.FieldExpiresAt]
	return ok
}

// ResetExpiresAt reset all changes of the "expires_at" field.
func (m *SessionMutation) ResetExpiresAt() {
	m.expires_at = nil
	delete(m.clearedFields, session.FieldExpiresAt)
}

// Op returns the operation name.
func (m *SessionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Session).
func (m *SessionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *SessionMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.data != nil {
		fields = append(fields, session.FieldData)
	}
	if m.created_at != nil {
		fields = append(fields, session.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, session.FieldUpdatedAt)
	}
	if m.expires_at != nil {
		fields = append(fields, session.FieldExpiresAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *SessionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case session.FieldData:
		return m.Data()
	case session.FieldCreatedAt:
		return m.CreatedAt()
	case session.FieldUpdatedAt:
		return m.UpdatedAt()
	case session.FieldExpiresAt:
		return m.ExpiresAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *SessionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case session.FieldData:
		return m.OldData(ctx)
	case session.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case session.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case session.FieldExpiresAt:
		return m.OldExpiresAt(ctx)
	}
	return nil, fmt.Errorf("unknown Session field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *SessionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case session.FieldData:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetData(v)
		return nil
	case session.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case session.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case session.FieldExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpiresAt(v)
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *SessionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *SessionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *SessionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Session numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *SessionMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(session.FieldExpiresAt) {
		fields = append(fields, session.FieldExpiresAt)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *SessionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *SessionMutation) ClearField(name string) error {
	switch name {
	case session.FieldExpiresAt:
		m.ClearExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Session nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *SessionMutation) ResetField(name string) error {
	switch name {
	case session.FieldData:
		m.ResetData()
		return nil
	case session.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case session.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case session.FieldExpiresAt:
		m.ResetExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *SessionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *SessionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *SessionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *SessionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *SessionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *SessionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *SessionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Session unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *SessionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Session edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op                     Op
	typ                    string
	id                     *uuid.UUID
	billing_id             *string
	provider               *string
	email                  *string
	password               *string
	api_key                *string
	confirmed              *bool
	confirmation_sent_at   *time.Time
	confirmation_token     *string
	recovery_sent_at       *time.Time
	recovery_token         *string
	otp_sent_at            *time.Time
	otp                    *string
	email_change           *string
	email_change_sent_at   *time.Time
	email_change_token     *string
	metadata               *map[string]interface{}
	roles                  *[]string
	teams                  *map[string]string
	created_at             *time.Time
	updated_at             *time.Time
	last_signin_at         *time.Time
	clearedFields          map[string]struct{}
	workspace              *uuid.UUID
	clearedworkspace       bool
	workspace_roles        map[uuid.UUID]struct{}
	removedworkspace_roles map[uuid.UUID]struct{}
	clearedworkspace_roles bool
	group_roles            map[uuid.UUID]struct{}
	removedgroup_roles     map[uuid.UUID]struct{}
	clearedgroup_roles     bool
	user_roles             map[uuid.UUID]struct{}
	removeduser_roles      map[uuid.UUID]struct{}
	cleareduser_roles      bool
	done                   bool
	oldValue               func(context.Context) (*User, error)
	predicates             []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows to manage the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for User.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the id field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on User creation.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetBillingID sets the billing_id field.
func (m *UserMutation) SetBillingID(s string) {
	m.billing_id = &s
}

// BillingID returns the billing_id value in the mutation.
func (m *UserMutation) BillingID() (r string, exists bool) {
	v := m.billing_id
	if v == nil {
		return
	}
	return *v, true
}

// OldBillingID returns the old billing_id value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldBillingID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBillingID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBillingID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBillingID: %w", err)
	}
	return oldValue.BillingID, nil
}

// ClearBillingID clears the value of billing_id.
func (m *UserMutation) ClearBillingID() {
	m.billing_id = nil
	m.clearedFields[user.FieldBillingID] = struct{}{}
}

// BillingIDCleared returns if the field billing_id was cleared in this mutation.
func (m *UserMutation) BillingIDCleared() bool {
	_, ok := m.clearedFields[user.FieldBillingID]
	return ok
}

// ResetBillingID reset all changes of the "billing_id" field.
func (m *UserMutation) ResetBillingID() {
	m.billing_id = nil
	delete(m.clearedFields, user.FieldBillingID)
}

// SetProvider sets the provider field.
func (m *UserMutation) SetProvider(s string) {
	m.provider = &s
}

// Provider returns the provider value in the mutation.
func (m *UserMutation) Provider() (r string, exists bool) {
	v := m.provider
	if v == nil {
		return
	}
	return *v, true
}

// OldProvider returns the old provider value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldProvider(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldProvider is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldProvider requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProvider: %w", err)
	}
	return oldValue.Provider, nil
}

// ResetProvider reset all changes of the "provider" field.
func (m *UserMutation) ResetProvider() {
	m.provider = nil
}

// SetEmail sets the email field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the email value in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old email value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmail is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail reset all changes of the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the password field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the password value in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old password value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPassword is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword reset all changes of the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetAPIKey sets the api_key field.
func (m *UserMutation) SetAPIKey(s string) {
	m.api_key = &s
}

// APIKey returns the api_key value in the mutation.
func (m *UserMutation) APIKey() (r string, exists bool) {
	v := m.api_key
	if v == nil {
		return
	}
	return *v, true
}

// OldAPIKey returns the old api_key value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldAPIKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAPIKey is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAPIKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAPIKey: %w", err)
	}
	return oldValue.APIKey, nil
}

// ClearAPIKey clears the value of api_key.
func (m *UserMutation) ClearAPIKey() {
	m.api_key = nil
	m.clearedFields[user.FieldAPIKey] = struct{}{}
}

// APIKeyCleared returns if the field api_key was cleared in this mutation.
func (m *UserMutation) APIKeyCleared() bool {
	_, ok := m.clearedFields[user.FieldAPIKey]
	return ok
}

// ResetAPIKey reset all changes of the "api_key" field.
func (m *UserMutation) ResetAPIKey() {
	m.api_key = nil
	delete(m.clearedFields, user.FieldAPIKey)
}

// SetConfirmed sets the confirmed field.
func (m *UserMutation) SetConfirmed(b bool) {
	m.confirmed = &b
}

// Confirmed returns the confirmed value in the mutation.
func (m *UserMutation) Confirmed() (r bool, exists bool) {
	v := m.confirmed
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmed returns the old confirmed value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldConfirmed(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldConfirmed is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldConfirmed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmed: %w", err)
	}
	return oldValue.Confirmed, nil
}

// ClearConfirmed clears the value of confirmed.
func (m *UserMutation) ClearConfirmed() {
	m.confirmed = nil
	m.clearedFields[user.FieldConfirmed] = struct{}{}
}

// ConfirmedCleared returns if the field confirmed was cleared in this mutation.
func (m *UserMutation) ConfirmedCleared() bool {
	_, ok := m.clearedFields[user.FieldConfirmed]
	return ok
}

// ResetConfirmed reset all changes of the "confirmed" field.
func (m *UserMutation) ResetConfirmed() {
	m.confirmed = nil
	delete(m.clearedFields, user.FieldConfirmed)
}

// SetConfirmationSentAt sets the confirmation_sent_at field.
func (m *UserMutation) SetConfirmationSentAt(t time.Time) {
	m.confirmation_sent_at = &t
}

// ConfirmationSentAt returns the confirmation_sent_at value in the mutation.
func (m *UserMutation) ConfirmationSentAt() (r time.Time, exists bool) {
	v := m.confirmation_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmationSentAt returns the old confirmation_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldConfirmationSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldConfirmationSentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldConfirmationSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmationSentAt: %w", err)
	}
	return oldValue.ConfirmationSentAt, nil
}

// ClearConfirmationSentAt clears the value of confirmation_sent_at.
func (m *UserMutation) ClearConfirmationSentAt() {
	m.confirmation_sent_at = nil
	m.clearedFields[user.FieldConfirmationSentAt] = struct{}{}
}

// ConfirmationSentAtCleared returns if the field confirmation_sent_at was cleared in this mutation.
func (m *UserMutation) ConfirmationSentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldConfirmationSentAt]
	return ok
}

// ResetConfirmationSentAt reset all changes of the "confirmation_sent_at" field.
func (m *UserMutation) ResetConfirmationSentAt() {
	m.confirmation_sent_at = nil
	delete(m.clearedFields, user.FieldConfirmationSentAt)
}

// SetConfirmationToken sets the confirmation_token field.
func (m *UserMutation) SetConfirmationToken(s string) {
	m.confirmation_token = &s
}

// ConfirmationToken returns the confirmation_token value in the mutation.
func (m *UserMutation) ConfirmationToken() (r string, exists bool) {
	v := m.confirmation_token
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmationToken returns the old confirmation_token value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldConfirmationToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldConfirmationToken is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldConfirmationToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmationToken: %w", err)
	}
	return oldValue.ConfirmationToken, nil
}

// ClearConfirmationToken clears the value of confirmation_token.
func (m *UserMutation) ClearConfirmationToken() {
	m.confirmation_token = nil
	m.clearedFields[user.FieldConfirmationToken] = struct{}{}
}

// ConfirmationTokenCleared returns if the field confirmation_token was cleared in this mutation.
func (m *UserMutation) ConfirmationTokenCleared() bool {
	_, ok := m.clearedFields[user.FieldConfirmationToken]
	return ok
}

// ResetConfirmationToken reset all changes of the "confirmation_token" field.
func (m *UserMutation) ResetConfirmationToken() {
	m.confirmation_token = nil
	delete(m.clearedFields, user.FieldConfirmationToken)
}

// SetRecoverySentAt sets the recovery_sent_at field.
func (m *UserMutation) SetRecoverySentAt(t time.Time) {
	m.recovery_sent_at = &t
}

// RecoverySentAt returns the recovery_sent_at value in the mutation.
func (m *UserMutation) RecoverySentAt() (r time.Time, exists bool) {
	v := m.recovery_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRecoverySentAt returns the old recovery_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldRecoverySentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRecoverySentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRecoverySentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecoverySentAt: %w", err)
	}
	return oldValue.RecoverySentAt, nil
}

// ClearRecoverySentAt clears the value of recovery_sent_at.
func (m *UserMutation) ClearRecoverySentAt() {
	m.recovery_sent_at = nil
	m.clearedFields[user.FieldRecoverySentAt] = struct{}{}
}

// RecoverySentAtCleared returns if the field recovery_sent_at was cleared in this mutation.
func (m *UserMutation) RecoverySentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldRecoverySentAt]
	return ok
}

// ResetRecoverySentAt reset all changes of the "recovery_sent_at" field.
func (m *UserMutation) ResetRecoverySentAt() {
	m.recovery_sent_at = nil
	delete(m.clearedFields, user.FieldRecoverySentAt)
}

// SetRecoveryToken sets the recovery_token field.
func (m *UserMutation) SetRecoveryToken(s string) {
	m.recovery_token = &s
}

// RecoveryToken returns the recovery_token value in the mutation.
func (m *UserMutation) RecoveryToken() (r string, exists bool) {
	v := m.recovery_token
	if v == nil {
		return
	}
	return *v, true
}

// OldRecoveryToken returns the old recovery_token value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldRecoveryToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRecoveryToken is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRecoveryToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecoveryToken: %w", err)
	}
	return oldValue.RecoveryToken, nil
}

// ClearRecoveryToken clears the value of recovery_token.
func (m *UserMutation) ClearRecoveryToken() {
	m.recovery_token = nil
	m.clearedFields[user.FieldRecoveryToken] = struct{}{}
}

// RecoveryTokenCleared returns if the field recovery_token was cleared in this mutation.
func (m *UserMutation) RecoveryTokenCleared() bool {
	_, ok := m.clearedFields[user.FieldRecoveryToken]
	return ok
}

// ResetRecoveryToken reset all changes of the "recovery_token" field.
func (m *UserMutation) ResetRecoveryToken() {
	m.recovery_token = nil
	delete(m.clearedFields, user.FieldRecoveryToken)
}

// SetOtpSentAt sets the otp_sent_at field.
func (m *UserMutation) SetOtpSentAt(t time.Time) {
	m.otp_sent_at = &t
}

// OtpSentAt returns the otp_sent_at value in the mutation.
func (m *UserMutation) OtpSentAt() (r time.Time, exists bool) {
	v := m.otp_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldOtpSentAt returns the old otp_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldOtpSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOtpSentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOtpSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOtpSentAt: %w", err)
	}
	return oldValue.OtpSentAt, nil
}

// ClearOtpSentAt clears the value of otp_sent_at.
func (m *UserMutation) ClearOtpSentAt() {
	m.otp_sent_at = nil
	m.clearedFields[user.FieldOtpSentAt] = struct{}{}
}

// OtpSentAtCleared returns if the field otp_sent_at was cleared in this mutation.
func (m *UserMutation) OtpSentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldOtpSentAt]
	return ok
}

// ResetOtpSentAt reset all changes of the "otp_sent_at" field.
func (m *UserMutation) ResetOtpSentAt() {
	m.otp_sent_at = nil
	delete(m.clearedFields, user.FieldOtpSentAt)
}

// SetOtp sets the otp field.
func (m *UserMutation) SetOtp(s string) {
	m.otp = &s
}

// Otp returns the otp value in the mutation.
func (m *UserMutation) Otp() (r string, exists bool) {
	v := m.otp
	if v == nil {
		return
	}
	return *v, true
}

// OldOtp returns the old otp value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldOtp(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOtp is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOtp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOtp: %w", err)
	}
	return oldValue.Otp, nil
}

// ClearOtp clears the value of otp.
func (m *UserMutation) ClearOtp() {
	m.otp = nil
	m.clearedFields[user.FieldOtp] = struct{}{}
}

// OtpCleared returns if the field otp was cleared in this mutation.
func (m *UserMutation) OtpCleared() bool {
	_, ok := m.clearedFields[user.FieldOtp]
	return ok
}

// ResetOtp reset all changes of the "otp" field.
func (m *UserMutation) ResetOtp() {
	m.otp = nil
	delete(m.clearedFields, user.FieldOtp)
}

// SetEmailChange sets the email_change field.
func (m *UserMutation) SetEmailChange(s string) {
	m.email_change = &s
}

// EmailChange returns the email_change value in the mutation.
func (m *UserMutation) EmailChange() (r string, exists bool) {
	v := m.email_change
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChange returns the old email_change value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmailChange(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmailChange is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmailChange requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChange: %w", err)
	}
	return oldValue.EmailChange, nil
}

// ClearEmailChange clears the value of email_change.
func (m *UserMutation) ClearEmailChange() {
	m.email_change = nil
	m.clearedFields[user.FieldEmailChange] = struct{}{}
}

// EmailChangeCleared returns if the field email_change was cleared in this mutation.
func (m *UserMutation) EmailChangeCleared() bool {
	_, ok := m.clearedFields[user.FieldEmailChange]
	return ok
}

// ResetEmailChange reset all changes of the "email_change" field.
func (m *UserMutation) ResetEmailChange() {
	m.email_change = nil
	delete(m.clearedFields, user.FieldEmailChange)
}

// SetEmailChangeSentAt sets the email_change_sent_at field.
func (m *UserMutation) SetEmailChangeSentAt(t time.Time) {
	m.email_change_sent_at = &t
}

// EmailChangeSentAt returns the email_change_sent_at value in the mutation.
func (m *UserMutation) EmailChangeSentAt() (r time.Time, exists bool) {
	v := m.email_change_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChangeSentAt returns the old email_change_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmailChangeSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmailChangeSentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmailChangeSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChangeSentAt: %w", err)
	}
	return oldValue.EmailChangeSentAt, nil
}

// ClearEmailChangeSentAt clears the value of email_change_sent_at.
func (m *UserMutation) ClearEmailChangeSentAt() {
	m.email_change_sent_at = nil
	m.clearedFields[user.FieldEmailChangeSentAt] = struct{}{}
}

// EmailChangeSentAtCleared returns if the field email_change_sent_at was cleared in this mutation.
func (m *UserMutation) EmailChangeSentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldEmailChangeSentAt]
	return ok
}

// ResetEmailChangeSentAt reset all changes of the "email_change_sent_at" field.
func (m *UserMutation) ResetEmailChangeSentAt() {
	m.email_change_sent_at = nil
	delete(m.clearedFields, user.FieldEmailChangeSentAt)
}

// SetEmailChangeToken sets the email_change_token field.
func (m *UserMutation) SetEmailChangeToken(s string) {
	m.email_change_token = &s
}

// EmailChangeToken returns the email_change_token value in the mutation.
func (m *UserMutation) EmailChangeToken() (r string, exists bool) {
	v := m.email_change_token
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChangeToken returns the old email_change_token value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmailChangeToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmailChangeToken is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmailChangeToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChangeToken: %w", err)
	}
	return oldValue.EmailChangeToken, nil
}

// ClearEmailChangeToken clears the value of email_change_token.
func (m *UserMutation) ClearEmailChangeToken() {
	m.email_change_token = nil
	m.clearedFields[user.FieldEmailChangeToken] = struct{}{}
}

// EmailChangeTokenCleared returns if the field email_change_token was cleared in this mutation.
func (m *UserMutation) EmailChangeTokenCleared() bool {
	_, ok := m.clearedFields[user.FieldEmailChangeToken]
	return ok
}

// ResetEmailChangeToken reset all changes of the "email_change_token" field.
func (m *UserMutation) ResetEmailChangeToken() {
	m.email_change_token = nil
	delete(m.clearedFields, user.FieldEmailChangeToken)
}

// SetMetadata sets the metadata field.
func (m *UserMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *UserMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old metadata value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldMetadata(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMetadata is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ResetMetadata reset all changes of the "metadata" field.
func (m *UserMutation) ResetMetadata() {
	m.metadata = nil
}

// SetRoles sets the roles field.
func (m *UserMutation) SetRoles(s []string) {
	m.roles = &s
}

// Roles returns the roles value in the mutation.
func (m *UserMutation) Roles() (r []string, exists bool) {
	v := m.roles
	if v == nil {
		return
	}
	return *v, true
}

// OldRoles returns the old roles value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldRoles(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRoles is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRoles requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRoles: %w", err)
	}
	return oldValue.Roles, nil
}

// ClearRoles clears the value of roles.
func (m *UserMutation) ClearRoles() {
	m.roles = nil
	m.clearedFields[user.FieldRoles] = struct{}{}
}

// RolesCleared returns if the field roles was cleared in this mutation.
func (m *UserMutation) RolesCleared() bool {
	_, ok := m.clearedFields[user.FieldRoles]
	return ok
}

// ResetRoles reset all changes of the "roles" field.
func (m *UserMutation) ResetRoles() {
	m.roles = nil
	delete(m.clearedFields, user.FieldRoles)
}

// SetTeams sets the teams field.
func (m *UserMutation) SetTeams(value map[string]string) {
	m.teams = &value
}

// Teams returns the teams value in the mutation.
func (m *UserMutation) Teams() (r map[string]string, exists bool) {
	v := m.teams
	if v == nil {
		return
	}
	return *v, true
}

// OldTeams returns the old teams value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldTeams(ctx context.Context) (v map[string]string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTeams is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTeams requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTeams: %w", err)
	}
	return oldValue.Teams, nil
}

// ClearTeams clears the value of teams.
func (m *UserMutation) ClearTeams() {
	m.teams = nil
	m.clearedFields[user.FieldTeams] = struct{}{}
}

// TeamsCleared returns if the field teams was cleared in this mutation.
func (m *UserMutation) TeamsCleared() bool {
	_, ok := m.clearedFields[user.FieldTeams]
	return ok
}

// ResetTeams reset all changes of the "teams" field.
func (m *UserMutation) ResetTeams() {
	m.teams = nil
	delete(m.clearedFields, user.FieldTeams)
}

// SetCreatedAt sets the created_at field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetLastSigninAt sets the last_signin_at field.
func (m *UserMutation) SetLastSigninAt(t time.Time) {
	m.last_signin_at = &t
}

// LastSigninAt returns the last_signin_at value in the mutation.
func (m *UserMutation) LastSigninAt() (r time.Time, exists bool) {
	v := m.last_signin_at
	if v == nil {
		return
	}
	return *v, true
}

// OldLastSigninAt returns the old last_signin_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldLastSigninAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastSigninAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastSigninAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastSigninAt: %w", err)
	}
	return oldValue.LastSigninAt, nil
}

// ClearLastSigninAt clears the value of last_signin_at.
func (m *UserMutation) ClearLastSigninAt() {
	m.last_signin_at = nil
	m.clearedFields[user.FieldLastSigninAt] = struct{}{}
}

// LastSigninAtCleared returns if the field last_signin_at was cleared in this mutation.
func (m *UserMutation) LastSigninAtCleared() bool {
	_, ok := m.clearedFields[user.FieldLastSigninAt]
	return ok
}

// ResetLastSigninAt reset all changes of the "last_signin_at" field.
func (m *UserMutation) ResetLastSigninAt() {
	m.last_signin_at = nil
	delete(m.clearedFields, user.FieldLastSigninAt)
}

// SetWorkspaceID sets the workspace edge to Workspace by id.
func (m *UserMutation) SetWorkspaceID(id uuid.UUID) {
	m.workspace = &id
}

// ClearWorkspace clears the workspace edge to Workspace.
func (m *UserMutation) ClearWorkspace() {
	m.clearedworkspace = true
}

// WorkspaceCleared returns if the edge workspace was cleared.
func (m *UserMutation) WorkspaceCleared() bool {
	return m.clearedworkspace
}

// WorkspaceID returns the workspace id in the mutation.
func (m *UserMutation) WorkspaceID() (id uuid.UUID, exists bool) {
	if m.workspace != nil {
		return *m.workspace, true
	}
	return
}

// WorkspaceIDs returns the workspace ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// WorkspaceID instead. It exists only for internal usage by the builders.
func (m *UserMutation) WorkspaceIDs() (ids []uuid.UUID) {
	if id := m.workspace; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetWorkspace reset all changes of the "workspace" edge.
func (m *UserMutation) ResetWorkspace() {
	m.workspace = nil
	m.clearedworkspace = false
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (m *UserMutation) AddWorkspaceRoleIDs(ids ...uuid.UUID) {
	if m.workspace_roles == nil {
		m.workspace_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.workspace_roles[ids[i]] = struct{}{}
	}
}

// ClearWorkspaceRoles clears the workspace_roles edge to WorkspaceRole.
func (m *UserMutation) ClearWorkspaceRoles() {
	m.clearedworkspace_roles = true
}

// WorkspaceRolesCleared returns if the edge workspace_roles was cleared.
func (m *UserMutation) WorkspaceRolesCleared() bool {
	return m.clearedworkspace_roles
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (m *UserMutation) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) {
	if m.removedworkspace_roles == nil {
		m.removedworkspace_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedworkspace_roles[ids[i]] = struct{}{}
	}
}

// RemovedWorkspaceRoles returns the removed ids of workspace_roles.
func (m *UserMutation) RemovedWorkspaceRolesIDs() (ids []uuid.UUID) {
	for id := range m.removedworkspace_roles {
		ids = append(ids, id)
	}
	return
}

// WorkspaceRolesIDs returns the workspace_roles ids in the mutation.
func (m *UserMutation) WorkspaceRolesIDs() (ids []uuid.UUID) {
	for id := range m.workspace_roles {
		ids = append(ids, id)
	}
	return
}

// ResetWorkspaceRoles reset all changes of the "workspace_roles" edge.
func (m *UserMutation) ResetWorkspaceRoles() {
	m.workspace_roles = nil
	m.clearedworkspace_roles = false
	m.removedworkspace_roles = nil
}

// AddGroupRoleIDs adds the group_roles edge to GroupRole by ids.
func (m *UserMutation) AddGroupRoleIDs(ids ...uuid.UUID) {
	if m.group_roles == nil {
		m.group_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.group_roles[ids[i]] = struct{}{}
	}
}

// ClearGroupRoles clears the group_roles edge to GroupRole.
func (m *UserMutation) ClearGroupRoles() {
	m.clearedgroup_roles = true
}

// GroupRolesCleared returns if the edge group_roles was cleared.
func (m *UserMutation) GroupRolesCleared() bool {
	return m.clearedgroup_roles
}

// RemoveGroupRoleIDs removes the group_roles edge to GroupRole by ids.
func (m *UserMutation) RemoveGroupRoleIDs(ids ...uuid.UUID) {
	if m.removedgroup_roles == nil {
		m.removedgroup_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedgroup_roles[ids[i]] = struct{}{}
	}
}

// RemovedGroupRoles returns the removed ids of group_roles.
func (m *UserMutation) RemovedGroupRolesIDs() (ids []uuid.UUID) {
	for id := range m.removedgroup_roles {
		ids = append(ids, id)
	}
	return
}

// GroupRolesIDs returns the group_roles ids in the mutation.
func (m *UserMutation) GroupRolesIDs() (ids []uuid.UUID) {
	for id := range m.group_roles {
		ids = append(ids, id)
	}
	return
}

// ResetGroupRoles reset all changes of the "group_roles" edge.
func (m *UserMutation) ResetGroupRoles() {
	m.group_roles = nil
	m.clearedgroup_roles = false
	m.removedgroup_roles = nil
}

// AddUserRoleIDs adds the user_roles edge to UserRole by ids.
func (m *UserMutation) AddUserRoleIDs(ids ...uuid.UUID) {
	if m.user_roles == nil {
		m.user_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.user_roles[ids[i]] = struct{}{}
	}
}

// ClearUserRoles clears the user_roles edge to UserRole.
func (m *UserMutation) ClearUserRoles() {
	m.cleareduser_roles = true
}

// UserRolesCleared returns if the edge user_roles was cleared.
func (m *UserMutation) UserRolesCleared() bool {
	return m.cleareduser_roles
}

// RemoveUserRoleIDs removes the user_roles edge to UserRole by ids.
func (m *UserMutation) RemoveUserRoleIDs(ids ...uuid.UUID) {
	if m.removeduser_roles == nil {
		m.removeduser_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removeduser_roles[ids[i]] = struct{}{}
	}
}

// RemovedUserRoles returns the removed ids of user_roles.
func (m *UserMutation) RemovedUserRolesIDs() (ids []uuid.UUID) {
	for id := range m.removeduser_roles {
		ids = append(ids, id)
	}
	return
}

// UserRolesIDs returns the user_roles ids in the mutation.
func (m *UserMutation) UserRolesIDs() (ids []uuid.UUID) {
	for id := range m.user_roles {
		ids = append(ids, id)
	}
	return
}

// ResetUserRoles reset all changes of the "user_roles" edge.
func (m *UserMutation) ResetUserRoles() {
	m.user_roles = nil
	m.cleareduser_roles = false
	m.removeduser_roles = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 21)
	if m.billing_id != nil {
		fields = append(fields, user.FieldBillingID)
	}
	if m.provider != nil {
		fields = append(fields, user.FieldProvider)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.api_key != nil {
		fields = append(fields, user.FieldAPIKey)
	}
	if m.confirmed != nil {
		fields = append(fields, user.FieldConfirmed)
	}
	if m.confirmation_sent_at != nil {
		fields = append(fields, user.FieldConfirmationSentAt)
	}
	if m.confirmation_token != nil {
		fields = append(fields, user.FieldConfirmationToken)
	}
	if m.recovery_sent_at != nil {
		fields = append(fields, user.FieldRecoverySentAt)
	}
	if m.recovery_token != nil {
		fields = append(fields, user.FieldRecoveryToken)
	}
	if m.otp_sent_at != nil {
		fields = append(fields, user.FieldOtpSentAt)
	}
	if m.otp != nil {
		fields = append(fields, user.FieldOtp)
	}
	if m.email_change != nil {
		fields = append(fields, user.FieldEmailChange)
	}
	if m.email_change_sent_at != nil {
		fields = append(fields, user.FieldEmailChangeSentAt)
	}
	if m.email_change_token != nil {
		fields = append(fields, user.FieldEmailChangeToken)
	}
	if m.metadata != nil {
		fields = append(fields, user.FieldMetadata)
	}
	if m.roles != nil {
		fields = append(fields, user.FieldRoles)
	}
	if m.teams != nil {
		fields = append(fields, user.FieldTeams)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.last_signin_at != nil {
		fields = append(fields, user.FieldLastSigninAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldBillingID:
		return m.BillingID()
	case user.FieldProvider:
		return m.Provider()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldAPIKey:
		return m.APIKey()
	case user.FieldConfirmed:
		return m.Confirmed()
	case user.FieldConfirmationSentAt:
		return m.ConfirmationSentAt()
	case user.FieldConfirmationToken:
		return m.ConfirmationToken()
	case user.FieldRecoverySentAt:
		return m.RecoverySentAt()
	case user.FieldRecoveryToken:
		return m.RecoveryToken()
	case user.FieldOtpSentAt:
		return m.OtpSentAt()
	case user.FieldOtp:
		return m.Otp()
	case user.FieldEmailChange:
		return m.EmailChange()
	case user.FieldEmailChangeSentAt:
		return m.EmailChangeSentAt()
	case user.FieldEmailChangeToken:
		return m.EmailChangeToken()
	case user.FieldMetadata:
		return m.Metadata()
	case user.FieldRoles:
		return m.Roles()
	case user.FieldTeams:
		return m.Teams()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldLastSigninAt:
		return m.LastSigninAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldBillingID:
		return m.OldBillingID(ctx)
	case user.FieldProvider:
		return m.OldProvider(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldAPIKey:
		return m.OldAPIKey(ctx)
	case user.FieldConfirmed:
		return m.OldConfirmed(ctx)
	case user.FieldConfirmationSentAt:
		return m.OldConfirmationSentAt(ctx)
	case user.FieldConfirmationToken:
		return m.OldConfirmationToken(ctx)
	case user.FieldRecoverySentAt:
		return m.OldRecoverySentAt(ctx)
	case user.FieldRecoveryToken:
		return m.OldRecoveryToken(ctx)
	case user.FieldOtpSentAt:
		return m.OldOtpSentAt(ctx)
	case user.FieldOtp:
		return m.OldOtp(ctx)
	case user.FieldEmailChange:
		return m.OldEmailChange(ctx)
	case user.FieldEmailChangeSentAt:
		return m.OldEmailChangeSentAt(ctx)
	case user.FieldEmailChangeToken:
		return m.OldEmailChangeToken(ctx)
	case user.FieldMetadata:
		return m.OldMetadata(ctx)
	case user.FieldRoles:
		return m.OldRoles(ctx)
	case user.FieldTeams:
		return m.OldTeams(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldLastSigninAt:
		return m.OldLastSigninAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldBillingID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBillingID(v)
		return nil
	case user.FieldProvider:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProvider(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldAPIKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAPIKey(v)
		return nil
	case user.FieldConfirmed:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmed(v)
		return nil
	case user.FieldConfirmationSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmationSentAt(v)
		return nil
	case user.FieldConfirmationToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmationToken(v)
		return nil
	case user.FieldRecoverySentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecoverySentAt(v)
		return nil
	case user.FieldRecoveryToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecoveryToken(v)
		return nil
	case user.FieldOtpSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOtpSentAt(v)
		return nil
	case user.FieldOtp:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOtp(v)
		return nil
	case user.FieldEmailChange:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChange(v)
		return nil
	case user.FieldEmailChangeSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChangeSentAt(v)
		return nil
	case user.FieldEmailChangeToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChangeToken(v)
		return nil
	case user.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case user.FieldRoles:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRoles(v)
		return nil
	case user.FieldTeams:
		v, ok := value.(map[string]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTeams(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldLastSigninAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastSigninAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldBillingID) {
		fields = append(fields, user.FieldBillingID)
	}
	if m.FieldCleared(user.FieldAPIKey) {
		fields = append(fields, user.FieldAPIKey)
	}
	if m.FieldCleared(user.FieldConfirmed) {
		fields = append(fields, user.FieldConfirmed)
	}
	if m.FieldCleared(user.FieldConfirmationSentAt) {
		fields = append(fields, user.FieldConfirmationSentAt)
	}
	if m.FieldCleared(user.FieldConfirmationToken) {
		fields = append(fields, user.FieldConfirmationToken)
	}
	if m.FieldCleared(user.FieldRecoverySentAt) {
		fields = append(fields, user.FieldRecoverySentAt)
	}
	if m.FieldCleared(user.FieldRecoveryToken) {
		fields = append(fields, user.FieldRecoveryToken)
	}
	if m.FieldCleared(user.FieldOtpSentAt) {
		fields = append(fields, user.FieldOtpSentAt)
	}
	if m.FieldCleared(user.FieldOtp) {
		fields = append(fields, user.FieldOtp)
	}
	if m.FieldCleared(user.FieldEmailChange) {
		fields = append(fields, user.FieldEmailChange)
	}
	if m.FieldCleared(user.FieldEmailChangeSentAt) {
		fields = append(fields, user.FieldEmailChangeSentAt)
	}
	if m.FieldCleared(user.FieldEmailChangeToken) {
		fields = append(fields, user.FieldEmailChangeToken)
	}
	if m.FieldCleared(user.FieldRoles) {
		fields = append(fields, user.FieldRoles)
	}
	if m.FieldCleared(user.FieldTeams) {
		fields = append(fields, user.FieldTeams)
	}
	if m.FieldCleared(user.FieldLastSigninAt) {
		fields = append(fields, user.FieldLastSigninAt)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldBillingID:
		m.ClearBillingID()
		return nil
	case user.FieldAPIKey:
		m.ClearAPIKey()
		return nil
	case user.FieldConfirmed:
		m.ClearConfirmed()
		return nil
	case user.FieldConfirmationSentAt:
		m.ClearConfirmationSentAt()
		return nil
	case user.FieldConfirmationToken:
		m.ClearConfirmationToken()
		return nil
	case user.FieldRecoverySentAt:
		m.ClearRecoverySentAt()
		return nil
	case user.FieldRecoveryToken:
		m.ClearRecoveryToken()
		return nil
	case user.FieldOtpSentAt:
		m.ClearOtpSentAt()
		return nil
	case user.FieldOtp:
		m.ClearOtp()
		return nil
	case user.FieldEmailChange:
		m.ClearEmailChange()
		return nil
	case user.FieldEmailChangeSentAt:
		m.ClearEmailChangeSentAt()
		return nil
	case user.FieldEmailChangeToken:
		m.ClearEmailChangeToken()
		return nil
	case user.FieldRoles:
		m.ClearRoles()
		return nil
	case user.FieldTeams:
		m.ClearTeams()
		return nil
	case user.FieldLastSigninAt:
		m.ClearLastSigninAt()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldBillingID:
		m.ResetBillingID()
		return nil
	case user.FieldProvider:
		m.ResetProvider()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldAPIKey:
		m.ResetAPIKey()
		return nil
	case user.FieldConfirmed:
		m.ResetConfirmed()
		return nil
	case user.FieldConfirmationSentAt:
		m.ResetConfirmationSentAt()
		return nil
	case user.FieldConfirmationToken:
		m.ResetConfirmationToken()
		return nil
	case user.FieldRecoverySentAt:
		m.ResetRecoverySentAt()
		return nil
	case user.FieldRecoveryToken:
		m.ResetRecoveryToken()
		return nil
	case user.FieldOtpSentAt:
		m.ResetOtpSentAt()
		return nil
	case user.FieldOtp:
		m.ResetOtp()
		return nil
	case user.FieldEmailChange:
		m.ResetEmailChange()
		return nil
	case user.FieldEmailChangeSentAt:
		m.ResetEmailChangeSentAt()
		return nil
	case user.FieldEmailChangeToken:
		m.ResetEmailChangeToken()
		return nil
	case user.FieldMetadata:
		m.ResetMetadata()
		return nil
	case user.FieldRoles:
		m.ResetRoles()
		return nil
	case user.FieldTeams:
		m.ResetTeams()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldLastSigninAt:
		m.ResetLastSigninAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 4)
	if m.workspace != nil {
		edges = append(edges, user.EdgeWorkspace)
	}
	if m.workspace_roles != nil {
		edges = append(edges, user.EdgeWorkspaceRoles)
	}
	if m.group_roles != nil {
		edges = append(edges, user.EdgeGroupRoles)
	}
	if m.user_roles != nil {
		edges = append(edges, user.EdgeUserRoles)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeWorkspace:
		if id := m.workspace; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeWorkspaceRoles:
		ids := make([]ent.Value, 0, len(m.workspace_roles))
		for id := range m.workspace_roles {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeGroupRoles:
		ids := make([]ent.Value, 0, len(m.group_roles))
		for id := range m.group_roles {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeUserRoles:
		ids := make([]ent.Value, 0, len(m.user_roles))
		for id := range m.user_roles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 4)
	if m.removedworkspace_roles != nil {
		edges = append(edges, user.EdgeWorkspaceRoles)
	}
	if m.removedgroup_roles != nil {
		edges = append(edges, user.EdgeGroupRoles)
	}
	if m.removeduser_roles != nil {
		edges = append(edges, user.EdgeUserRoles)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeWorkspaceRoles:
		ids := make([]ent.Value, 0, len(m.removedworkspace_roles))
		for id := range m.removedworkspace_roles {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeGroupRoles:
		ids := make([]ent.Value, 0, len(m.removedgroup_roles))
		for id := range m.removedgroup_roles {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeUserRoles:
		ids := make([]ent.Value, 0, len(m.removeduser_roles))
		for id := range m.removeduser_roles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 4)
	if m.clearedworkspace {
		edges = append(edges, user.EdgeWorkspace)
	}
	if m.clearedworkspace_roles {
		edges = append(edges, user.EdgeWorkspaceRoles)
	}
	if m.clearedgroup_roles {
		edges = append(edges, user.EdgeGroupRoles)
	}
	if m.cleareduser_roles {
		edges = append(edges, user.EdgeUserRoles)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeWorkspace:
		return m.clearedworkspace
	case user.EdgeWorkspaceRoles:
		return m.clearedworkspace_roles
	case user.EdgeGroupRoles:
		return m.clearedgroup_roles
	case user.EdgeUserRoles:
		return m.cleareduser_roles
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeWorkspace:
		m.ClearWorkspace()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeWorkspace:
		m.ResetWorkspace()
		return nil
	case user.EdgeWorkspaceRoles:
		m.ResetWorkspaceRoles()
		return nil
	case user.EdgeGroupRoles:
		m.ResetGroupRoles()
		return nil
	case user.EdgeUserRoles:
		m.ResetUserRoles()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}

// UserRoleMutation represents an operation that mutate the UserRoles
// nodes in the graph.
type UserRoleMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	user_id       *uuid.UUID
	metadata      *map[string]interface{}
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	users         *uuid.UUID
	clearedusers  bool
	done          bool
	oldValue      func(context.Context) (*UserRole, error)
	predicates    []predicate.UserRole
}

var _ ent.Mutation = (*UserRoleMutation)(nil)

// userroleOption allows to manage the mutation configuration using functional options.
type userroleOption func(*UserRoleMutation)

// newUserRoleMutation creates new mutation for UserRole.
func newUserRoleMutation(c config, op Op, opts ...userroleOption) *UserRoleMutation {
	m := &UserRoleMutation{
		config:        c,
		op:            op,
		typ:           TypeUserRole,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserRoleID sets the id field of the mutation.
func withUserRoleID(id uuid.UUID) userroleOption {
	return func(m *UserRoleMutation) {
		var (
			err   error
			once  sync.Once
			value *UserRole
		)
		m.oldValue = func(ctx context.Context) (*UserRole, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserRole.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserRole sets the old UserRole of the mutation.
func withUserRole(node *UserRole) userroleOption {
	return func(m *UserRoleMutation) {
		m.oldValue = func(context.Context) (*UserRole, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserRoleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserRoleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on UserRole creation.
func (m *UserRoleMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserRoleMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *UserRoleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserRoleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the UserRole.
// If the UserRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserRoleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *UserRoleMutation) ResetName() {
	m.name = nil
}

// SetUserID sets the user_id field.
func (m *UserRoleMutation) SetUserID(u uuid.UUID) {
	m.user_id = &u
}

// UserID returns the user_id value in the mutation.
func (m *UserRoleMutation) UserID() (r uuid.UUID, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old user_id value of the UserRole.
// If the UserRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserRoleMutation) OldUserID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID reset all changes of the "user_id" field.
func (m *UserRoleMutation) ResetUserID() {
	m.user_id = nil
}

// SetMetadata sets the metadata field.
func (m *UserRoleMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *UserRoleMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old metadata value of the UserRole.
// If the UserRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserRoleMutation) OldMetadata(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMetadata is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ClearMetadata clears the value of metadata.
func (m *UserRoleMutation) ClearMetadata() {
	m.metadata = nil
	m.clearedFields[userrole.FieldMetadata] = struct{}{}
}

// MetadataCleared returns if the field metadata was cleared in this mutation.
func (m *UserRoleMutation) MetadataCleared() bool {
	_, ok := m.clearedFields[userrole.FieldMetadata]
	return ok
}

// ResetMetadata reset all changes of the "metadata" field.
func (m *UserRoleMutation) ResetMetadata() {
	m.metadata = nil
	delete(m.clearedFields, userrole.FieldMetadata)
}

// SetCreatedAt sets the created_at field.
func (m *UserRoleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *UserRoleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the UserRole.
// If the UserRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserRoleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *UserRoleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *UserRoleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *UserRoleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the UserRole.
// If the UserRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserRoleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *UserRoleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetUsersID sets the users edge to User by id.
func (m *UserRoleMutation) SetUsersID(id uuid.UUID) {
	m.users = &id
}

// ClearUsers clears the users edge to User.
func (m *UserRoleMutation) ClearUsers() {
	m.clearedusers = true
}

// UsersCleared returns if the edge users was cleared.
func (m *UserRoleMutation) UsersCleared() bool {
	return m.clearedusers
}

// UsersID returns the users id in the mutation.
func (m *UserRoleMutation) UsersID() (id uuid.UUID, exists bool) {
	if m.users != nil {
		return *m.users, true
	}
	return
}

// UsersIDs returns the users ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// UsersID instead. It exists only for internal usage by the builders.
func (m *UserRoleMutation) UsersIDs() (ids []uuid.UUID) {
	if id := m.users; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUsers reset all changes of the "users" edge.
func (m *UserRoleMutation) ResetUsers() {
	m.users = nil
	m.clearedusers = false
}

// Op returns the operation name.
func (m *UserRoleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (UserRole).
func (m *UserRoleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserRoleMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, userrole.FieldName)
	}
	if m.user_id != nil {
		fields = append(fields, userrole.FieldUserID)
	}
	if m.metadata != nil {
		fields = append(fields, userrole.FieldMetadata)
	}
	if m.created_at != nil {
		fields = append(fields, userrole.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, userrole.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserRoleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case userrole.FieldName:
		return m.Name()
	case userrole.FieldUserID:
		return m.UserID()
	case userrole.FieldMetadata:
		return m.Metadata()
	case userrole.FieldCreatedAt:
		return m.CreatedAt()
	case userrole.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserRoleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case userrole.FieldName:
		return m.OldName(ctx)
	case userrole.FieldUserID:
		return m.OldUserID(ctx)
	case userrole.FieldMetadata:
		return m.OldMetadata(ctx)
	case userrole.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case userrole.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown UserRole field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserRoleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case userrole.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case userrole.FieldUserID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case userrole.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case userrole.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case userrole.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown UserRole field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserRoleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserRoleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserRoleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown UserRole numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserRoleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(userrole.FieldMetadata) {
		fields = append(fields, userrole.FieldMetadata)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserRoleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserRoleMutation) ClearField(name string) error {
	switch name {
	case userrole.FieldMetadata:
		m.ClearMetadata()
		return nil
	}
	return fmt.Errorf("unknown UserRole nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserRoleMutation) ResetField(name string) error {
	switch name {
	case userrole.FieldName:
		m.ResetName()
		return nil
	case userrole.FieldUserID:
		m.ResetUserID()
		return nil
	case userrole.FieldMetadata:
		m.ResetMetadata()
		return nil
	case userrole.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case userrole.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown UserRole field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserRoleMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.users != nil {
		edges = append(edges, userrole.EdgeUsers)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserRoleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case userrole.EdgeUsers:
		if id := m.users; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserRoleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserRoleMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserRoleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedusers {
		edges = append(edges, userrole.EdgeUsers)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserRoleMutation) EdgeCleared(name string) bool {
	switch name {
	case userrole.EdgeUsers:
		return m.clearedusers
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserRoleMutation) ClearEdge(name string) error {
	switch name {
	case userrole.EdgeUsers:
		m.ClearUsers()
		return nil
	}
	return fmt.Errorf("unknown UserRole unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserRoleMutation) ResetEdge(name string) error {
	switch name {
	case userrole.EdgeUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown UserRole edge %s", name)
}

// WorkspaceMutation represents an operation that mutate the Workspaces
// nodes in the graph.
type WorkspaceMutation struct {
	config
	op                     Op
	typ                    string
	id                     *uuid.UUID
	name                   *string
	plan                   *string
	description            *string
	metadata               *map[string]interface{}
	created_at             *time.Time
	updated_at             *time.Time
	clearedFields          map[string]struct{}
	owner                  *uuid.UUID
	clearedowner           bool
	workspace_roles        map[uuid.UUID]struct{}
	removedworkspace_roles map[uuid.UUID]struct{}
	clearedworkspace_roles bool
	groups                 map[uuid.UUID]struct{}
	removedgroups          map[uuid.UUID]struct{}
	clearedgroups          bool
	done                   bool
	oldValue               func(context.Context) (*Workspace, error)
	predicates             []predicate.Workspace
}

var _ ent.Mutation = (*WorkspaceMutation)(nil)

// workspaceOption allows to manage the mutation configuration using functional options.
type workspaceOption func(*WorkspaceMutation)

// newWorkspaceMutation creates new mutation for Workspace.
func newWorkspaceMutation(c config, op Op, opts ...workspaceOption) *WorkspaceMutation {
	m := &WorkspaceMutation{
		config:        c,
		op:            op,
		typ:           TypeWorkspace,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWorkspaceID sets the id field of the mutation.
func withWorkspaceID(id uuid.UUID) workspaceOption {
	return func(m *WorkspaceMutation) {
		var (
			err   error
			once  sync.Once
			value *Workspace
		)
		m.oldValue = func(ctx context.Context) (*Workspace, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Workspace.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWorkspace sets the old Workspace of the mutation.
func withWorkspace(node *Workspace) workspaceOption {
	return func(m *WorkspaceMutation) {
		m.oldValue = func(context.Context) (*Workspace, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WorkspaceMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WorkspaceMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Workspace creation.
func (m *WorkspaceMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *WorkspaceMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *WorkspaceMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *WorkspaceMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Workspace.
// If the Workspace object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *WorkspaceMutation) ResetName() {
	m.name = nil
}

// SetPlan sets the plan field.
func (m *WorkspaceMutation) SetPlan(s string) {
	m.plan = &s
}

// Plan returns the plan value in the mutation.
func (m *WorkspaceMutation) Plan() (r string, exists bool) {
	v := m.plan
	if v == nil {
		return
	}
	return *v, true
}

// OldPlan returns the old plan value of the Workspace.
// If the Workspace object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceMutation) OldPlan(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPlan is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPlan requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPlan: %w", err)
	}
	return oldValue.Plan, nil
}

// ResetPlan reset all changes of the "plan" field.
func (m *WorkspaceMutation) ResetPlan() {
	m.plan = nil
}

// SetDescription sets the description field.
func (m *WorkspaceMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the description value in the mutation.
func (m *WorkspaceMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old description value of the Workspace.
// If the Workspace object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDescription is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of description.
func (m *WorkspaceMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[workspace.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the field description was cleared in this mutation.
func (m *WorkspaceMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[workspace.FieldDescription]
	return ok
}

// ResetDescription reset all changes of the "description" field.
func (m *WorkspaceMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, workspace.FieldDescription)
}

// SetMetadata sets the metadata field.
func (m *WorkspaceMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *WorkspaceMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old metadata value of the Workspace.
// If the Workspace object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceMutation) OldMetadata(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMetadata is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ClearMetadata clears the value of metadata.
func (m *WorkspaceMutation) ClearMetadata() {
	m.metadata = nil
	m.clearedFields[workspace.FieldMetadata] = struct{}{}
}

// MetadataCleared returns if the field metadata was cleared in this mutation.
func (m *WorkspaceMutation) MetadataCleared() bool {
	_, ok := m.clearedFields[workspace.FieldMetadata]
	return ok
}

// ResetMetadata reset all changes of the "metadata" field.
func (m *WorkspaceMutation) ResetMetadata() {
	m.metadata = nil
	delete(m.clearedFields, workspace.FieldMetadata)
}

// SetCreatedAt sets the created_at field.
func (m *WorkspaceMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *WorkspaceMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Workspace.
// If the Workspace object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *WorkspaceMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *WorkspaceMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *WorkspaceMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Workspace.
// If the Workspace object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *WorkspaceMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetOwnerID sets the owner edge to User by id.
func (m *WorkspaceMutation) SetOwnerID(id uuid.UUID) {
	m.owner = &id
}

// ClearOwner clears the owner edge to User.
func (m *WorkspaceMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared returns if the edge owner was cleared.
func (m *WorkspaceMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the owner id in the mutation.
func (m *WorkspaceMutation) OwnerID() (id uuid.UUID, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the owner ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *WorkspaceMutation) OwnerIDs() (ids []uuid.UUID) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner reset all changes of the "owner" edge.
func (m *WorkspaceMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (m *WorkspaceMutation) AddWorkspaceRoleIDs(ids ...uuid.UUID) {
	if m.workspace_roles == nil {
		m.workspace_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.workspace_roles[ids[i]] = struct{}{}
	}
}

// ClearWorkspaceRoles clears the workspace_roles edge to WorkspaceRole.
func (m *WorkspaceMutation) ClearWorkspaceRoles() {
	m.clearedworkspace_roles = true
}

// WorkspaceRolesCleared returns if the edge workspace_roles was cleared.
func (m *WorkspaceMutation) WorkspaceRolesCleared() bool {
	return m.clearedworkspace_roles
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (m *WorkspaceMutation) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) {
	if m.removedworkspace_roles == nil {
		m.removedworkspace_roles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedworkspace_roles[ids[i]] = struct{}{}
	}
}

// RemovedWorkspaceRoles returns the removed ids of workspace_roles.
func (m *WorkspaceMutation) RemovedWorkspaceRolesIDs() (ids []uuid.UUID) {
	for id := range m.removedworkspace_roles {
		ids = append(ids, id)
	}
	return
}

// WorkspaceRolesIDs returns the workspace_roles ids in the mutation.
func (m *WorkspaceMutation) WorkspaceRolesIDs() (ids []uuid.UUID) {
	for id := range m.workspace_roles {
		ids = append(ids, id)
	}
	return
}

// ResetWorkspaceRoles reset all changes of the "workspace_roles" edge.
func (m *WorkspaceMutation) ResetWorkspaceRoles() {
	m.workspace_roles = nil
	m.clearedworkspace_roles = false
	m.removedworkspace_roles = nil
}

// AddGroupIDs adds the groups edge to Group by ids.
func (m *WorkspaceMutation) AddGroupIDs(ids ...uuid.UUID) {
	if m.groups == nil {
		m.groups = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.groups[ids[i]] = struct{}{}
	}
}

// ClearGroups clears the groups edge to Group.
func (m *WorkspaceMutation) ClearGroups() {
	m.clearedgroups = true
}

// GroupsCleared returns if the edge groups was cleared.
func (m *WorkspaceMutation) GroupsCleared() bool {
	return m.clearedgroups
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (m *WorkspaceMutation) RemoveGroupIDs(ids ...uuid.UUID) {
	if m.removedgroups == nil {
		m.removedgroups = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedgroups[ids[i]] = struct{}{}
	}
}

// RemovedGroups returns the removed ids of groups.
func (m *WorkspaceMutation) RemovedGroupsIDs() (ids []uuid.UUID) {
	for id := range m.removedgroups {
		ids = append(ids, id)
	}
	return
}

// GroupsIDs returns the groups ids in the mutation.
func (m *WorkspaceMutation) GroupsIDs() (ids []uuid.UUID) {
	for id := range m.groups {
		ids = append(ids, id)
	}
	return
}

// ResetGroups reset all changes of the "groups" edge.
func (m *WorkspaceMutation) ResetGroups() {
	m.groups = nil
	m.clearedgroups = false
	m.removedgroups = nil
}

// Op returns the operation name.
func (m *WorkspaceMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Workspace).
func (m *WorkspaceMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *WorkspaceMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.name != nil {
		fields = append(fields, workspace.FieldName)
	}
	if m.plan != nil {
		fields = append(fields, workspace.FieldPlan)
	}
	if m.description != nil {
		fields = append(fields, workspace.FieldDescription)
	}
	if m.metadata != nil {
		fields = append(fields, workspace.FieldMetadata)
	}
	if m.created_at != nil {
		fields = append(fields, workspace.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, workspace.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *WorkspaceMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case workspace.FieldName:
		return m.Name()
	case workspace.FieldPlan:
		return m.Plan()
	case workspace.FieldDescription:
		return m.Description()
	case workspace.FieldMetadata:
		return m.Metadata()
	case workspace.FieldCreatedAt:
		return m.CreatedAt()
	case workspace.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *WorkspaceMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case workspace.FieldName:
		return m.OldName(ctx)
	case workspace.FieldPlan:
		return m.OldPlan(ctx)
	case workspace.FieldDescription:
		return m.OldDescription(ctx)
	case workspace.FieldMetadata:
		return m.OldMetadata(ctx)
	case workspace.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case workspace.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Workspace field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *WorkspaceMutation) SetField(name string, value ent.Value) error {
	switch name {
	case workspace.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case workspace.FieldPlan:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPlan(v)
		return nil
	case workspace.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case workspace.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case workspace.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case workspace.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Workspace field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *WorkspaceMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *WorkspaceMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *WorkspaceMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Workspace numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *WorkspaceMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(workspace.FieldDescription) {
		fields = append(fields, workspace.FieldDescription)
	}
	if m.FieldCleared(workspace.FieldMetadata) {
		fields = append(fields, workspace.FieldMetadata)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *WorkspaceMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *WorkspaceMutation) ClearField(name string) error {
	switch name {
	case workspace.FieldDescription:
		m.ClearDescription()
		return nil
	case workspace.FieldMetadata:
		m.ClearMetadata()
		return nil
	}
	return fmt.Errorf("unknown Workspace nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *WorkspaceMutation) ResetField(name string) error {
	switch name {
	case workspace.FieldName:
		m.ResetName()
		return nil
	case workspace.FieldPlan:
		m.ResetPlan()
		return nil
	case workspace.FieldDescription:
		m.ResetDescription()
		return nil
	case workspace.FieldMetadata:
		m.ResetMetadata()
		return nil
	case workspace.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case workspace.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Workspace field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *WorkspaceMutation) AddedEdges() []string {
	edges := make([]string, 0, 3)
	if m.owner != nil {
		edges = append(edges, workspace.EdgeOwner)
	}
	if m.workspace_roles != nil {
		edges = append(edges, workspace.EdgeWorkspaceRoles)
	}
	if m.groups != nil {
		edges = append(edges, workspace.EdgeGroups)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *WorkspaceMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case workspace.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	case workspace.EdgeWorkspaceRoles:
		ids := make([]ent.Value, 0, len(m.workspace_roles))
		for id := range m.workspace_roles {
			ids = append(ids, id)
		}
		return ids
	case workspace.EdgeGroups:
		ids := make([]ent.Value, 0, len(m.groups))
		for id := range m.groups {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *WorkspaceMutation) RemovedEdges() []string {
	edges := make([]string, 0, 3)
	if m.removedworkspace_roles != nil {
		edges = append(edges, workspace.EdgeWorkspaceRoles)
	}
	if m.removedgroups != nil {
		edges = append(edges, workspace.EdgeGroups)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *WorkspaceMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case workspace.EdgeWorkspaceRoles:
		ids := make([]ent.Value, 0, len(m.removedworkspace_roles))
		for id := range m.removedworkspace_roles {
			ids = append(ids, id)
		}
		return ids
	case workspace.EdgeGroups:
		ids := make([]ent.Value, 0, len(m.removedgroups))
		for id := range m.removedgroups {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *WorkspaceMutation) ClearedEdges() []string {
	edges := make([]string, 0, 3)
	if m.clearedowner {
		edges = append(edges, workspace.EdgeOwner)
	}
	if m.clearedworkspace_roles {
		edges = append(edges, workspace.EdgeWorkspaceRoles)
	}
	if m.clearedgroups {
		edges = append(edges, workspace.EdgeGroups)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *WorkspaceMutation) EdgeCleared(name string) bool {
	switch name {
	case workspace.EdgeOwner:
		return m.clearedowner
	case workspace.EdgeWorkspaceRoles:
		return m.clearedworkspace_roles
	case workspace.EdgeGroups:
		return m.clearedgroups
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *WorkspaceMutation) ClearEdge(name string) error {
	switch name {
	case workspace.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Workspace unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *WorkspaceMutation) ResetEdge(name string) error {
	switch name {
	case workspace.EdgeOwner:
		m.ResetOwner()
		return nil
	case workspace.EdgeWorkspaceRoles:
		m.ResetWorkspaceRoles()
		return nil
	case workspace.EdgeGroups:
		m.ResetGroups()
		return nil
	}
	return fmt.Errorf("unknown Workspace edge %s", name)
}

// WorkspaceRoleMutation represents an operation that mutate the WorkspaceRoles
// nodes in the graph.
type WorkspaceRoleMutation struct {
	config
	op                Op
	typ               string
	id                *uuid.UUID
	name              *string
	workspace_id      *uuid.UUID
	user_id           *uuid.UUID
	metadata          *map[string]interface{}
	created_at        *time.Time
	updated_at        *time.Time
	clearedFields     map[string]struct{}
	workspaces        *uuid.UUID
	clearedworkspaces bool
	users             *uuid.UUID
	clearedusers      bool
	done              bool
	oldValue          func(context.Context) (*WorkspaceRole, error)
	predicates        []predicate.WorkspaceRole
}

var _ ent.Mutation = (*WorkspaceRoleMutation)(nil)

// workspaceroleOption allows to manage the mutation configuration using functional options.
type workspaceroleOption func(*WorkspaceRoleMutation)

// newWorkspaceRoleMutation creates new mutation for WorkspaceRole.
func newWorkspaceRoleMutation(c config, op Op, opts ...workspaceroleOption) *WorkspaceRoleMutation {
	m := &WorkspaceRoleMutation{
		config:        c,
		op:            op,
		typ:           TypeWorkspaceRole,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWorkspaceRoleID sets the id field of the mutation.
func withWorkspaceRoleID(id uuid.UUID) workspaceroleOption {
	return func(m *WorkspaceRoleMutation) {
		var (
			err   error
			once  sync.Once
			value *WorkspaceRole
		)
		m.oldValue = func(ctx context.Context) (*WorkspaceRole, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().WorkspaceRole.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWorkspaceRole sets the old WorkspaceRole of the mutation.
func withWorkspaceRole(node *WorkspaceRole) workspaceroleOption {
	return func(m *WorkspaceRoleMutation) {
		m.oldValue = func(context.Context) (*WorkspaceRole, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WorkspaceRoleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WorkspaceRoleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on WorkspaceRole creation.
func (m *WorkspaceRoleMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *WorkspaceRoleMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *WorkspaceRoleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *WorkspaceRoleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the WorkspaceRole.
// If the WorkspaceRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceRoleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *WorkspaceRoleMutation) ResetName() {
	m.name = nil
}

// SetWorkspaceID sets the workspace_id field.
func (m *WorkspaceRoleMutation) SetWorkspaceID(u uuid.UUID) {
	m.workspace_id = &u
}

// WorkspaceID returns the workspace_id value in the mutation.
func (m *WorkspaceRoleMutation) WorkspaceID() (r uuid.UUID, exists bool) {
	v := m.workspace_id
	if v == nil {
		return
	}
	return *v, true
}

// OldWorkspaceID returns the old workspace_id value of the WorkspaceRole.
// If the WorkspaceRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceRoleMutation) OldWorkspaceID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldWorkspaceID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldWorkspaceID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWorkspaceID: %w", err)
	}
	return oldValue.WorkspaceID, nil
}

// ResetWorkspaceID reset all changes of the "workspace_id" field.
func (m *WorkspaceRoleMutation) ResetWorkspaceID() {
	m.workspace_id = nil
}

// SetUserID sets the user_id field.
func (m *WorkspaceRoleMutation) SetUserID(u uuid.UUID) {
	m.user_id = &u
}

// UserID returns the user_id value in the mutation.
func (m *WorkspaceRoleMutation) UserID() (r uuid.UUID, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old user_id value of the WorkspaceRole.
// If the WorkspaceRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceRoleMutation) OldUserID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID reset all changes of the "user_id" field.
func (m *WorkspaceRoleMutation) ResetUserID() {
	m.user_id = nil
}

// SetMetadata sets the metadata field.
func (m *WorkspaceRoleMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *WorkspaceRoleMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old metadata value of the WorkspaceRole.
// If the WorkspaceRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceRoleMutation) OldMetadata(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMetadata is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ClearMetadata clears the value of metadata.
func (m *WorkspaceRoleMutation) ClearMetadata() {
	m.metadata = nil
	m.clearedFields[workspacerole.FieldMetadata] = struct{}{}
}

// MetadataCleared returns if the field metadata was cleared in this mutation.
func (m *WorkspaceRoleMutation) MetadataCleared() bool {
	_, ok := m.clearedFields[workspacerole.FieldMetadata]
	return ok
}

// ResetMetadata reset all changes of the "metadata" field.
func (m *WorkspaceRoleMutation) ResetMetadata() {
	m.metadata = nil
	delete(m.clearedFields, workspacerole.FieldMetadata)
}

// SetCreatedAt sets the created_at field.
func (m *WorkspaceRoleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *WorkspaceRoleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the WorkspaceRole.
// If the WorkspaceRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceRoleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *WorkspaceRoleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *WorkspaceRoleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *WorkspaceRoleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the WorkspaceRole.
// If the WorkspaceRole object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *WorkspaceRoleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *WorkspaceRoleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetWorkspacesID sets the workspaces edge to Workspace by id.
func (m *WorkspaceRoleMutation) SetWorkspacesID(id uuid.UUID) {
	m.workspaces = &id
}

// ClearWorkspaces clears the workspaces edge to Workspace.
func (m *WorkspaceRoleMutation) ClearWorkspaces() {
	m.clearedworkspaces = true
}

// WorkspacesCleared returns if the edge workspaces was cleared.
func (m *WorkspaceRoleMutation) WorkspacesCleared() bool {
	return m.clearedworkspaces
}

// WorkspacesID returns the workspaces id in the mutation.
func (m *WorkspaceRoleMutation) WorkspacesID() (id uuid.UUID, exists bool) {
	if m.workspaces != nil {
		return *m.workspaces, true
	}
	return
}

// WorkspacesIDs returns the workspaces ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// WorkspacesID instead. It exists only for internal usage by the builders.
func (m *WorkspaceRoleMutation) WorkspacesIDs() (ids []uuid.UUID) {
	if id := m.workspaces; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetWorkspaces reset all changes of the "workspaces" edge.
func (m *WorkspaceRoleMutation) ResetWorkspaces() {
	m.workspaces = nil
	m.clearedworkspaces = false
}

// SetUsersID sets the users edge to User by id.
func (m *WorkspaceRoleMutation) SetUsersID(id uuid.UUID) {
	m.users = &id
}

// ClearUsers clears the users edge to User.
func (m *WorkspaceRoleMutation) ClearUsers() {
	m.clearedusers = true
}

// UsersCleared returns if the edge users was cleared.
func (m *WorkspaceRoleMutation) UsersCleared() bool {
	return m.clearedusers
}

// UsersID returns the users id in the mutation.
func (m *WorkspaceRoleMutation) UsersID() (id uuid.UUID, exists bool) {
	if m.users != nil {
		return *m.users, true
	}
	return
}

// UsersIDs returns the users ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// UsersID instead. It exists only for internal usage by the builders.
func (m *WorkspaceRoleMutation) UsersIDs() (ids []uuid.UUID) {
	if id := m.users; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUsers reset all changes of the "users" edge.
func (m *WorkspaceRoleMutation) ResetUsers() {
	m.users = nil
	m.clearedusers = false
}

// Op returns the operation name.
func (m *WorkspaceRoleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (WorkspaceRole).
func (m *WorkspaceRoleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *WorkspaceRoleMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.name != nil {
		fields = append(fields, workspacerole.FieldName)
	}
	if m.workspace_id != nil {
		fields = append(fields, workspacerole.FieldWorkspaceID)
	}
	if m.user_id != nil {
		fields = append(fields, workspacerole.FieldUserID)
	}
	if m.metadata != nil {
		fields = append(fields, workspacerole.FieldMetadata)
	}
	if m.created_at != nil {
		fields = append(fields, workspacerole.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, workspacerole.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *WorkspaceRoleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case workspacerole.FieldName:
		return m.Name()
	case workspacerole.FieldWorkspaceID:
		return m.WorkspaceID()
	case workspacerole.FieldUserID:
		return m.UserID()
	case workspacerole.FieldMetadata:
		return m.Metadata()
	case workspacerole.FieldCreatedAt:
		return m.CreatedAt()
	case workspacerole.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *WorkspaceRoleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case workspacerole.FieldName:
		return m.OldName(ctx)
	case workspacerole.FieldWorkspaceID:
		return m.OldWorkspaceID(ctx)
	case workspacerole.FieldUserID:
		return m.OldUserID(ctx)
	case workspacerole.FieldMetadata:
		return m.OldMetadata(ctx)
	case workspacerole.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case workspacerole.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown WorkspaceRole field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *WorkspaceRoleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case workspacerole.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case workspacerole.FieldWorkspaceID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWorkspaceID(v)
		return nil
	case workspacerole.FieldUserID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case workspacerole.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case workspacerole.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case workspacerole.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown WorkspaceRole field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *WorkspaceRoleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *WorkspaceRoleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *WorkspaceRoleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown WorkspaceRole numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *WorkspaceRoleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(workspacerole.FieldMetadata) {
		fields = append(fields, workspacerole.FieldMetadata)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *WorkspaceRoleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *WorkspaceRoleMutation) ClearField(name string) error {
	switch name {
	case workspacerole.FieldMetadata:
		m.ClearMetadata()
		return nil
	}
	return fmt.Errorf("unknown WorkspaceRole nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *WorkspaceRoleMutation) ResetField(name string) error {
	switch name {
	case workspacerole.FieldName:
		m.ResetName()
		return nil
	case workspacerole.FieldWorkspaceID:
		m.ResetWorkspaceID()
		return nil
	case workspacerole.FieldUserID:
		m.ResetUserID()
		return nil
	case workspacerole.FieldMetadata:
		m.ResetMetadata()
		return nil
	case workspacerole.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case workspacerole.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown WorkspaceRole field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *WorkspaceRoleMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.workspaces != nil {
		edges = append(edges, workspacerole.EdgeWorkspaces)
	}
	if m.users != nil {
		edges = append(edges, workspacerole.EdgeUsers)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *WorkspaceRoleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case workspacerole.EdgeWorkspaces:
		if id := m.workspaces; id != nil {
			return []ent.Value{*id}
		}
	case workspacerole.EdgeUsers:
		if id := m.users; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *WorkspaceRoleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *WorkspaceRoleMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *WorkspaceRoleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedworkspaces {
		edges = append(edges, workspacerole.EdgeWorkspaces)
	}
	if m.clearedusers {
		edges = append(edges, workspacerole.EdgeUsers)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *WorkspaceRoleMutation) EdgeCleared(name string) bool {
	switch name {
	case workspacerole.EdgeWorkspaces:
		return m.clearedworkspaces
	case workspacerole.EdgeUsers:
		return m.clearedusers
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *WorkspaceRoleMutation) ClearEdge(name string) error {
	switch name {
	case workspacerole.EdgeWorkspaces:
		m.ClearWorkspaces()
		return nil
	case workspacerole.EdgeUsers:
		m.ClearUsers()
		return nil
	}
	return fmt.Errorf("unknown WorkspaceRole unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *WorkspaceRoleMutation) ResetEdge(name string) error {
	switch name {
	case workspacerole.EdgeWorkspaces:
		m.ResetWorkspaces()
		return nil
	case workspacerole.EdgeUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown WorkspaceRole edge %s", name)
}
