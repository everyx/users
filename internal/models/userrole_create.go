// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserRoleCreate is the builder for creating a UserRole entity.
type UserRoleCreate struct {
	config
	mutation *UserRoleMutation
	hooks    []Hook
}

// SetName sets the name field.
func (urc *UserRoleCreate) SetName(s string) *UserRoleCreate {
	urc.mutation.SetName(s)
	return urc
}

// SetUserID sets the user_id field.
func (urc *UserRoleCreate) SetUserID(u uuid.UUID) *UserRoleCreate {
	urc.mutation.SetUserID(u)
	return urc
}

// SetMetadata sets the metadata field.
func (urc *UserRoleCreate) SetMetadata(m map[string]interface{}) *UserRoleCreate {
	urc.mutation.SetMetadata(m)
	return urc
}

// SetCreatedAt sets the created_at field.
func (urc *UserRoleCreate) SetCreatedAt(t time.Time) *UserRoleCreate {
	urc.mutation.SetCreatedAt(t)
	return urc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (urc *UserRoleCreate) SetNillableCreatedAt(t *time.Time) *UserRoleCreate {
	if t != nil {
		urc.SetCreatedAt(*t)
	}
	return urc
}

// SetUpdatedAt sets the updated_at field.
func (urc *UserRoleCreate) SetUpdatedAt(t time.Time) *UserRoleCreate {
	urc.mutation.SetUpdatedAt(t)
	return urc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (urc *UserRoleCreate) SetNillableUpdatedAt(t *time.Time) *UserRoleCreate {
	if t != nil {
		urc.SetUpdatedAt(*t)
	}
	return urc
}

// SetID sets the id field.
func (urc *UserRoleCreate) SetID(u uuid.UUID) *UserRoleCreate {
	urc.mutation.SetID(u)
	return urc
}

// SetUsersID sets the users edge to User by id.
func (urc *UserRoleCreate) SetUsersID(id uuid.UUID) *UserRoleCreate {
	urc.mutation.SetUsersID(id)
	return urc
}

// SetUsers sets the users edge to User.
func (urc *UserRoleCreate) SetUsers(u *User) *UserRoleCreate {
	return urc.SetUsersID(u.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (urc *UserRoleCreate) Mutation() *UserRoleMutation {
	return urc.mutation
}

// Save creates the UserRole in the database.
func (urc *UserRoleCreate) Save(ctx context.Context) (*UserRole, error) {
	var (
		err  error
		node *UserRole
	)
	urc.defaults()
	if len(urc.hooks) == 0 {
		if err = urc.check(); err != nil {
			return nil, err
		}
		node, err = urc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = urc.check(); err != nil {
				return nil, err
			}
			urc.mutation = mutation
			node, err = urc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(urc.hooks) - 1; i >= 0; i-- {
			mut = urc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, urc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRoleCreate) SaveX(ctx context.Context) *UserRole {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (urc *UserRoleCreate) defaults() {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		v := userrole.DefaultCreatedAt()
		urc.mutation.SetCreatedAt(v)
	}
	if _, ok := urc.mutation.UpdatedAt(); !ok {
		v := userrole.DefaultUpdatedAt()
		urc.mutation.SetUpdatedAt(v)
	}
	if _, ok := urc.mutation.ID(); !ok {
		v := userrole.DefaultID()
		urc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRoleCreate) check() error {
	if _, ok := urc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("models: missing required field \"name\"")}
	}
	if v, ok := urc.mutation.Name(); ok {
		if err := userrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := urc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New("models: missing required field \"user_id\"")}
	}
	if _, ok := urc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("models: missing required field \"created_at\"")}
	}
	if _, ok := urc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("models: missing required field \"updated_at\"")}
	}
	if _, ok := urc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New("models: missing required edge \"users\"")}
	}
	return nil
}

func (urc *UserRoleCreate) sqlSave(ctx context.Context) (*UserRole, error) {
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (urc *UserRoleCreate) createSpec() (*UserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRole{config: urc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		}
	)
	if id, ok := urc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := urc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userrole.FieldName,
		})
		_node.Name = value
	}
	if value, ok := urc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userrole.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := urc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: userrole.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := urc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := urc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := urc.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserRoleCreateBulk is the builder for creating a bulk of UserRole entities.
type UserRoleCreateBulk struct {
	config
	builders []*UserRoleCreate
}

// Save creates the UserRole entities in the database.
func (urcb *UserRoleCreateBulk) Save(ctx context.Context) ([]*UserRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRole, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (urcb *UserRoleCreateBulk) SaveX(ctx context.Context) []*UserRole {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
