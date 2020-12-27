// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/workspace"
	"github.com/adnaan/users/internal/models/workspacerole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceRoleUpdate is the builder for updating WorkspaceRole entities.
type WorkspaceRoleUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceRoleMutation
}

// Where adds a new predicate for the builder.
func (wru *WorkspaceRoleUpdate) Where(ps ...predicate.WorkspaceRole) *WorkspaceRoleUpdate {
	wru.mutation.predicates = append(wru.mutation.predicates, ps...)
	return wru
}

// SetName sets the name field.
func (wru *WorkspaceRoleUpdate) SetName(s string) *WorkspaceRoleUpdate {
	wru.mutation.SetName(s)
	return wru
}

// SetWorkspaceID sets the workspace_id field.
func (wru *WorkspaceRoleUpdate) SetWorkspaceID(u uuid.UUID) *WorkspaceRoleUpdate {
	wru.mutation.SetWorkspaceID(u)
	return wru
}

// SetUserID sets the user_id field.
func (wru *WorkspaceRoleUpdate) SetUserID(u uuid.UUID) *WorkspaceRoleUpdate {
	wru.mutation.SetUserID(u)
	return wru
}

// SetMetadata sets the metadata field.
func (wru *WorkspaceRoleUpdate) SetMetadata(m map[string]interface{}) *WorkspaceRoleUpdate {
	wru.mutation.SetMetadata(m)
	return wru
}

// ClearMetadata clears the value of metadata.
func (wru *WorkspaceRoleUpdate) ClearMetadata() *WorkspaceRoleUpdate {
	wru.mutation.ClearMetadata()
	return wru
}

// SetUpdatedAt sets the updated_at field.
func (wru *WorkspaceRoleUpdate) SetUpdatedAt(t time.Time) *WorkspaceRoleUpdate {
	wru.mutation.SetUpdatedAt(t)
	return wru
}

// SetWorkspacesID sets the workspaces edge to Workspace by id.
func (wru *WorkspaceRoleUpdate) SetWorkspacesID(id uuid.UUID) *WorkspaceRoleUpdate {
	wru.mutation.SetWorkspacesID(id)
	return wru
}

// SetWorkspaces sets the workspaces edge to Workspace.
func (wru *WorkspaceRoleUpdate) SetWorkspaces(w *Workspace) *WorkspaceRoleUpdate {
	return wru.SetWorkspacesID(w.ID)
}

// SetUsersID sets the users edge to User by id.
func (wru *WorkspaceRoleUpdate) SetUsersID(id uuid.UUID) *WorkspaceRoleUpdate {
	wru.mutation.SetUsersID(id)
	return wru
}

// SetUsers sets the users edge to User.
func (wru *WorkspaceRoleUpdate) SetUsers(u *User) *WorkspaceRoleUpdate {
	return wru.SetUsersID(u.ID)
}

// Mutation returns the WorkspaceRoleMutation object of the builder.
func (wru *WorkspaceRoleUpdate) Mutation() *WorkspaceRoleMutation {
	return wru.mutation
}

// ClearWorkspaces clears the "workspaces" edge to type Workspace.
func (wru *WorkspaceRoleUpdate) ClearWorkspaces() *WorkspaceRoleUpdate {
	wru.mutation.ClearWorkspaces()
	return wru
}

// ClearUsers clears the "users" edge to type User.
func (wru *WorkspaceRoleUpdate) ClearUsers() *WorkspaceRoleUpdate {
	wru.mutation.ClearUsers()
	return wru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wru *WorkspaceRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wru.defaults()
	if len(wru.hooks) == 0 {
		if err = wru.check(); err != nil {
			return 0, err
		}
		affected, err = wru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wru.check(); err != nil {
				return 0, err
			}
			wru.mutation = mutation
			affected, err = wru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wru.hooks) - 1; i >= 0; i-- {
			mut = wru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wru *WorkspaceRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := wru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wru *WorkspaceRoleUpdate) Exec(ctx context.Context) error {
	_, err := wru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wru *WorkspaceRoleUpdate) ExecX(ctx context.Context) {
	if err := wru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wru *WorkspaceRoleUpdate) defaults() {
	if _, ok := wru.mutation.UpdatedAt(); !ok {
		v := workspacerole.UpdateDefaultUpdatedAt()
		wru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wru *WorkspaceRoleUpdate) check() error {
	if v, ok := wru.mutation.Name(); ok {
		if err := workspacerole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := wru.mutation.WorkspacesID(); wru.mutation.WorkspacesCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"workspaces\"")
	}
	if _, ok := wru.mutation.UsersID(); wru.mutation.UsersCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"users\"")
	}
	return nil
}

func (wru *WorkspaceRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspacerole.Table,
			Columns: workspacerole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspacerole.FieldID,
			},
		},
	}
	if ps := wru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspacerole.FieldName,
		})
	}
	if value, ok := wru.mutation.WorkspaceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspacerole.FieldWorkspaceID,
		})
	}
	if value, ok := wru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspacerole.FieldUserID,
		})
	}
	if value, ok := wru.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspacerole.FieldMetadata,
		})
	}
	if wru.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: workspacerole.FieldMetadata,
		})
	}
	if value, ok := wru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspacerole.FieldUpdatedAt,
		})
	}
	if wru.mutation.WorkspacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.WorkspacesTable,
			Columns: []string{workspacerole.WorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wru.mutation.WorkspacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.WorkspacesTable,
			Columns: []string{workspacerole.WorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.UsersTable,
			Columns: []string{workspacerole.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.UsersTable,
			Columns: []string{workspacerole.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspacerole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkspaceRoleUpdateOne is the builder for updating a single WorkspaceRole entity.
type WorkspaceRoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *WorkspaceRoleMutation
}

// SetName sets the name field.
func (wruo *WorkspaceRoleUpdateOne) SetName(s string) *WorkspaceRoleUpdateOne {
	wruo.mutation.SetName(s)
	return wruo
}

// SetWorkspaceID sets the workspace_id field.
func (wruo *WorkspaceRoleUpdateOne) SetWorkspaceID(u uuid.UUID) *WorkspaceRoleUpdateOne {
	wruo.mutation.SetWorkspaceID(u)
	return wruo
}

// SetUserID sets the user_id field.
func (wruo *WorkspaceRoleUpdateOne) SetUserID(u uuid.UUID) *WorkspaceRoleUpdateOne {
	wruo.mutation.SetUserID(u)
	return wruo
}

// SetMetadata sets the metadata field.
func (wruo *WorkspaceRoleUpdateOne) SetMetadata(m map[string]interface{}) *WorkspaceRoleUpdateOne {
	wruo.mutation.SetMetadata(m)
	return wruo
}

// ClearMetadata clears the value of metadata.
func (wruo *WorkspaceRoleUpdateOne) ClearMetadata() *WorkspaceRoleUpdateOne {
	wruo.mutation.ClearMetadata()
	return wruo
}

// SetUpdatedAt sets the updated_at field.
func (wruo *WorkspaceRoleUpdateOne) SetUpdatedAt(t time.Time) *WorkspaceRoleUpdateOne {
	wruo.mutation.SetUpdatedAt(t)
	return wruo
}

// SetWorkspacesID sets the workspaces edge to Workspace by id.
func (wruo *WorkspaceRoleUpdateOne) SetWorkspacesID(id uuid.UUID) *WorkspaceRoleUpdateOne {
	wruo.mutation.SetWorkspacesID(id)
	return wruo
}

// SetWorkspaces sets the workspaces edge to Workspace.
func (wruo *WorkspaceRoleUpdateOne) SetWorkspaces(w *Workspace) *WorkspaceRoleUpdateOne {
	return wruo.SetWorkspacesID(w.ID)
}

// SetUsersID sets the users edge to User by id.
func (wruo *WorkspaceRoleUpdateOne) SetUsersID(id uuid.UUID) *WorkspaceRoleUpdateOne {
	wruo.mutation.SetUsersID(id)
	return wruo
}

// SetUsers sets the users edge to User.
func (wruo *WorkspaceRoleUpdateOne) SetUsers(u *User) *WorkspaceRoleUpdateOne {
	return wruo.SetUsersID(u.ID)
}

// Mutation returns the WorkspaceRoleMutation object of the builder.
func (wruo *WorkspaceRoleUpdateOne) Mutation() *WorkspaceRoleMutation {
	return wruo.mutation
}

// ClearWorkspaces clears the "workspaces" edge to type Workspace.
func (wruo *WorkspaceRoleUpdateOne) ClearWorkspaces() *WorkspaceRoleUpdateOne {
	wruo.mutation.ClearWorkspaces()
	return wruo
}

// ClearUsers clears the "users" edge to type User.
func (wruo *WorkspaceRoleUpdateOne) ClearUsers() *WorkspaceRoleUpdateOne {
	wruo.mutation.ClearUsers()
	return wruo
}

// Save executes the query and returns the updated entity.
func (wruo *WorkspaceRoleUpdateOne) Save(ctx context.Context) (*WorkspaceRole, error) {
	var (
		err  error
		node *WorkspaceRole
	)
	wruo.defaults()
	if len(wruo.hooks) == 0 {
		if err = wruo.check(); err != nil {
			return nil, err
		}
		node, err = wruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wruo.check(); err != nil {
				return nil, err
			}
			wruo.mutation = mutation
			node, err = wruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wruo.hooks) - 1; i >= 0; i-- {
			mut = wruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wruo *WorkspaceRoleUpdateOne) SaveX(ctx context.Context) *WorkspaceRole {
	node, err := wruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wruo *WorkspaceRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := wruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wruo *WorkspaceRoleUpdateOne) ExecX(ctx context.Context) {
	if err := wruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wruo *WorkspaceRoleUpdateOne) defaults() {
	if _, ok := wruo.mutation.UpdatedAt(); !ok {
		v := workspacerole.UpdateDefaultUpdatedAt()
		wruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wruo *WorkspaceRoleUpdateOne) check() error {
	if v, ok := wruo.mutation.Name(); ok {
		if err := workspacerole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := wruo.mutation.WorkspacesID(); wruo.mutation.WorkspacesCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"workspaces\"")
	}
	if _, ok := wruo.mutation.UsersID(); wruo.mutation.UsersCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"users\"")
	}
	return nil
}

func (wruo *WorkspaceRoleUpdateOne) sqlSave(ctx context.Context) (_node *WorkspaceRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspacerole.Table,
			Columns: workspacerole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspacerole.FieldID,
			},
		},
	}
	id, ok := wruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WorkspaceRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := wruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspacerole.FieldName,
		})
	}
	if value, ok := wruo.mutation.WorkspaceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspacerole.FieldWorkspaceID,
		})
	}
	if value, ok := wruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspacerole.FieldUserID,
		})
	}
	if value, ok := wruo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspacerole.FieldMetadata,
		})
	}
	if wruo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: workspacerole.FieldMetadata,
		})
	}
	if value, ok := wruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspacerole.FieldUpdatedAt,
		})
	}
	if wruo.mutation.WorkspacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.WorkspacesTable,
			Columns: []string{workspacerole.WorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wruo.mutation.WorkspacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.WorkspacesTable,
			Columns: []string{workspacerole.WorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.UsersTable,
			Columns: []string{workspacerole.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.UsersTable,
			Columns: []string{workspacerole.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkspaceRole{config: wruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, wruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspacerole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
