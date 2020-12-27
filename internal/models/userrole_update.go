// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserRoleUpdate is the builder for updating UserRole entities.
type UserRoleUpdate struct {
	config
	hooks    []Hook
	mutation *UserRoleMutation
}

// Where adds a new predicate for the builder.
func (uru *UserRoleUpdate) Where(ps ...predicate.UserRole) *UserRoleUpdate {
	uru.mutation.predicates = append(uru.mutation.predicates, ps...)
	return uru
}

// SetName sets the name field.
func (uru *UserRoleUpdate) SetName(s string) *UserRoleUpdate {
	uru.mutation.SetName(s)
	return uru
}

// SetUserID sets the user_id field.
func (uru *UserRoleUpdate) SetUserID(u uuid.UUID) *UserRoleUpdate {
	uru.mutation.SetUserID(u)
	return uru
}

// SetMetadata sets the metadata field.
func (uru *UserRoleUpdate) SetMetadata(m map[string]interface{}) *UserRoleUpdate {
	uru.mutation.SetMetadata(m)
	return uru
}

// ClearMetadata clears the value of metadata.
func (uru *UserRoleUpdate) ClearMetadata() *UserRoleUpdate {
	uru.mutation.ClearMetadata()
	return uru
}

// SetUpdatedAt sets the updated_at field.
func (uru *UserRoleUpdate) SetUpdatedAt(t time.Time) *UserRoleUpdate {
	uru.mutation.SetUpdatedAt(t)
	return uru
}

// SetUsersID sets the users edge to User by id.
func (uru *UserRoleUpdate) SetUsersID(id uuid.UUID) *UserRoleUpdate {
	uru.mutation.SetUsersID(id)
	return uru
}

// SetUsers sets the users edge to User.
func (uru *UserRoleUpdate) SetUsers(u *User) *UserRoleUpdate {
	return uru.SetUsersID(u.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (uru *UserRoleUpdate) Mutation() *UserRoleMutation {
	return uru.mutation
}

// ClearUsers clears the "users" edge to type User.
func (uru *UserRoleUpdate) ClearUsers() *UserRoleUpdate {
	uru.mutation.ClearUsers()
	return uru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uru.defaults()
	if len(uru.hooks) == 0 {
		if err = uru.check(); err != nil {
			return 0, err
		}
		affected, err = uru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uru.check(); err != nil {
				return 0, err
			}
			uru.mutation = mutation
			affected, err = uru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uru.hooks) - 1; i >= 0; i-- {
			mut = uru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserRoleUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserRoleUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uru *UserRoleUpdate) defaults() {
	if _, ok := uru.mutation.UpdatedAt(); !ok {
		v := userrole.UpdateDefaultUpdatedAt()
		uru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uru *UserRoleUpdate) check() error {
	if v, ok := uru.mutation.Name(); ok {
		if err := userrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := uru.mutation.UsersID(); uru.mutation.UsersCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"users\"")
	}
	return nil
}

func (uru *UserRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userrole.Table,
			Columns: userrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		},
	}
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userrole.FieldName,
		})
	}
	if value, ok := uru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userrole.FieldUserID,
		})
	}
	if value, ok := uru.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: userrole.FieldMetadata,
		})
	}
	if uru.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: userrole.FieldMetadata,
		})
	}
	if value, ok := uru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userrole.FieldUpdatedAt,
		})
	}
	if uru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UsersTable,
			Columns: []string{userrole.UsersColumn},
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
	if nodes := uru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UsersTable,
			Columns: []string{userrole.UsersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserRoleUpdateOne is the builder for updating a single UserRole entity.
type UserRoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserRoleMutation
}

// SetName sets the name field.
func (uruo *UserRoleUpdateOne) SetName(s string) *UserRoleUpdateOne {
	uruo.mutation.SetName(s)
	return uruo
}

// SetUserID sets the user_id field.
func (uruo *UserRoleUpdateOne) SetUserID(u uuid.UUID) *UserRoleUpdateOne {
	uruo.mutation.SetUserID(u)
	return uruo
}

// SetMetadata sets the metadata field.
func (uruo *UserRoleUpdateOne) SetMetadata(m map[string]interface{}) *UserRoleUpdateOne {
	uruo.mutation.SetMetadata(m)
	return uruo
}

// ClearMetadata clears the value of metadata.
func (uruo *UserRoleUpdateOne) ClearMetadata() *UserRoleUpdateOne {
	uruo.mutation.ClearMetadata()
	return uruo
}

// SetUpdatedAt sets the updated_at field.
func (uruo *UserRoleUpdateOne) SetUpdatedAt(t time.Time) *UserRoleUpdateOne {
	uruo.mutation.SetUpdatedAt(t)
	return uruo
}

// SetUsersID sets the users edge to User by id.
func (uruo *UserRoleUpdateOne) SetUsersID(id uuid.UUID) *UserRoleUpdateOne {
	uruo.mutation.SetUsersID(id)
	return uruo
}

// SetUsers sets the users edge to User.
func (uruo *UserRoleUpdateOne) SetUsers(u *User) *UserRoleUpdateOne {
	return uruo.SetUsersID(u.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (uruo *UserRoleUpdateOne) Mutation() *UserRoleMutation {
	return uruo.mutation
}

// ClearUsers clears the "users" edge to type User.
func (uruo *UserRoleUpdateOne) ClearUsers() *UserRoleUpdateOne {
	uruo.mutation.ClearUsers()
	return uruo
}

// Save executes the query and returns the updated entity.
func (uruo *UserRoleUpdateOne) Save(ctx context.Context) (*UserRole, error) {
	var (
		err  error
		node *UserRole
	)
	uruo.defaults()
	if len(uruo.hooks) == 0 {
		if err = uruo.check(); err != nil {
			return nil, err
		}
		node, err = uruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uruo.check(); err != nil {
				return nil, err
			}
			uruo.mutation = mutation
			node, err = uruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uruo.hooks) - 1; i >= 0; i-- {
			mut = uruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserRoleUpdateOne) SaveX(ctx context.Context) *UserRole {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserRoleUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uruo *UserRoleUpdateOne) defaults() {
	if _, ok := uruo.mutation.UpdatedAt(); !ok {
		v := userrole.UpdateDefaultUpdatedAt()
		uruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uruo *UserRoleUpdateOne) check() error {
	if v, ok := uruo.mutation.Name(); ok {
		if err := userrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := uruo.mutation.UsersID(); uruo.mutation.UsersCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"users\"")
	}
	return nil
}

func (uruo *UserRoleUpdateOne) sqlSave(ctx context.Context) (_node *UserRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userrole.Table,
			Columns: userrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		},
	}
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userrole.FieldName,
		})
	}
	if value, ok := uruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userrole.FieldUserID,
		})
	}
	if value, ok := uruo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: userrole.FieldMetadata,
		})
	}
	if uruo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: userrole.FieldMetadata,
		})
	}
	if value, ok := uruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userrole.FieldUpdatedAt,
		})
	}
	if uruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UsersTable,
			Columns: []string{userrole.UsersColumn},
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
	if nodes := uruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UsersTable,
			Columns: []string{userrole.UsersColumn},
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
	_node = &UserRole{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
