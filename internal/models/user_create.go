// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/users/internal/models/grouprole"
	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/adnaan/users/internal/models/workspace"
	"github.com/adnaan/users/internal/models/workspacerole"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetBillingID sets the billing_id field.
func (uc *UserCreate) SetBillingID(s string) *UserCreate {
	uc.mutation.SetBillingID(s)
	return uc
}

// SetNillableBillingID sets the billing_id field if the given value is not nil.
func (uc *UserCreate) SetNillableBillingID(s *string) *UserCreate {
	if s != nil {
		uc.SetBillingID(*s)
	}
	return uc
}

// SetProvider sets the provider field.
func (uc *UserCreate) SetProvider(s string) *UserCreate {
	uc.mutation.SetProvider(s)
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPassword sets the password field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetAPIKey sets the api_key field.
func (uc *UserCreate) SetAPIKey(s string) *UserCreate {
	uc.mutation.SetAPIKey(s)
	return uc
}

// SetNillableAPIKey sets the api_key field if the given value is not nil.
func (uc *UserCreate) SetNillableAPIKey(s *string) *UserCreate {
	if s != nil {
		uc.SetAPIKey(*s)
	}
	return uc
}

// SetConfirmed sets the confirmed field.
func (uc *UserCreate) SetConfirmed(b bool) *UserCreate {
	uc.mutation.SetConfirmed(b)
	return uc
}

// SetNillableConfirmed sets the confirmed field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmed(b *bool) *UserCreate {
	if b != nil {
		uc.SetConfirmed(*b)
	}
	return uc
}

// SetConfirmationSentAt sets the confirmation_sent_at field.
func (uc *UserCreate) SetConfirmationSentAt(t time.Time) *UserCreate {
	uc.mutation.SetConfirmationSentAt(t)
	return uc
}

// SetNillableConfirmationSentAt sets the confirmation_sent_at field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmationSentAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetConfirmationSentAt(*t)
	}
	return uc
}

// SetConfirmationToken sets the confirmation_token field.
func (uc *UserCreate) SetConfirmationToken(s string) *UserCreate {
	uc.mutation.SetConfirmationToken(s)
	return uc
}

// SetNillableConfirmationToken sets the confirmation_token field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmationToken(s *string) *UserCreate {
	if s != nil {
		uc.SetConfirmationToken(*s)
	}
	return uc
}

// SetRecoverySentAt sets the recovery_sent_at field.
func (uc *UserCreate) SetRecoverySentAt(t time.Time) *UserCreate {
	uc.mutation.SetRecoverySentAt(t)
	return uc
}

// SetNillableRecoverySentAt sets the recovery_sent_at field if the given value is not nil.
func (uc *UserCreate) SetNillableRecoverySentAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetRecoverySentAt(*t)
	}
	return uc
}

// SetRecoveryToken sets the recovery_token field.
func (uc *UserCreate) SetRecoveryToken(s string) *UserCreate {
	uc.mutation.SetRecoveryToken(s)
	return uc
}

// SetNillableRecoveryToken sets the recovery_token field if the given value is not nil.
func (uc *UserCreate) SetNillableRecoveryToken(s *string) *UserCreate {
	if s != nil {
		uc.SetRecoveryToken(*s)
	}
	return uc
}

// SetOtpSentAt sets the otp_sent_at field.
func (uc *UserCreate) SetOtpSentAt(t time.Time) *UserCreate {
	uc.mutation.SetOtpSentAt(t)
	return uc
}

// SetNillableOtpSentAt sets the otp_sent_at field if the given value is not nil.
func (uc *UserCreate) SetNillableOtpSentAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetOtpSentAt(*t)
	}
	return uc
}

// SetOtp sets the otp field.
func (uc *UserCreate) SetOtp(s string) *UserCreate {
	uc.mutation.SetOtp(s)
	return uc
}

// SetNillableOtp sets the otp field if the given value is not nil.
func (uc *UserCreate) SetNillableOtp(s *string) *UserCreate {
	if s != nil {
		uc.SetOtp(*s)
	}
	return uc
}

// SetEmailChange sets the email_change field.
func (uc *UserCreate) SetEmailChange(s string) *UserCreate {
	uc.mutation.SetEmailChange(s)
	return uc
}

// SetNillableEmailChange sets the email_change field if the given value is not nil.
func (uc *UserCreate) SetNillableEmailChange(s *string) *UserCreate {
	if s != nil {
		uc.SetEmailChange(*s)
	}
	return uc
}

// SetEmailChangeSentAt sets the email_change_sent_at field.
func (uc *UserCreate) SetEmailChangeSentAt(t time.Time) *UserCreate {
	uc.mutation.SetEmailChangeSentAt(t)
	return uc
}

// SetNillableEmailChangeSentAt sets the email_change_sent_at field if the given value is not nil.
func (uc *UserCreate) SetNillableEmailChangeSentAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetEmailChangeSentAt(*t)
	}
	return uc
}

// SetEmailChangeToken sets the email_change_token field.
func (uc *UserCreate) SetEmailChangeToken(s string) *UserCreate {
	uc.mutation.SetEmailChangeToken(s)
	return uc
}

// SetNillableEmailChangeToken sets the email_change_token field if the given value is not nil.
func (uc *UserCreate) SetNillableEmailChangeToken(s *string) *UserCreate {
	if s != nil {
		uc.SetEmailChangeToken(*s)
	}
	return uc
}

// SetMetadata sets the metadata field.
func (uc *UserCreate) SetMetadata(m map[string]interface{}) *UserCreate {
	uc.mutation.SetMetadata(m)
	return uc
}

// SetRoles sets the roles field.
func (uc *UserCreate) SetRoles(s []string) *UserCreate {
	uc.mutation.SetRoles(s)
	return uc
}

// SetTeams sets the teams field.
func (uc *UserCreate) SetTeams(m map[string]string) *UserCreate {
	uc.mutation.SetTeams(m)
	return uc
}

// SetCreatedAt sets the created_at field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the updated_at field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetLastSigninAt sets the last_signin_at field.
func (uc *UserCreate) SetLastSigninAt(t time.Time) *UserCreate {
	uc.mutation.SetLastSigninAt(t)
	return uc
}

// SetNillableLastSigninAt sets the last_signin_at field if the given value is not nil.
func (uc *UserCreate) SetNillableLastSigninAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastSigninAt(*t)
	}
	return uc
}

// SetID sets the id field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetWorkspaceID sets the workspace edge to Workspace by id.
func (uc *UserCreate) SetWorkspaceID(id uuid.UUID) *UserCreate {
	uc.mutation.SetWorkspaceID(id)
	return uc
}

// SetNillableWorkspaceID sets the workspace edge to Workspace by id if the given value is not nil.
func (uc *UserCreate) SetNillableWorkspaceID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetWorkspaceID(*id)
	}
	return uc
}

// SetWorkspace sets the workspace edge to Workspace.
func (uc *UserCreate) SetWorkspace(w *Workspace) *UserCreate {
	return uc.SetWorkspaceID(w.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (uc *UserCreate) AddWorkspaceRoleIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddWorkspaceRoleIDs(ids...)
	return uc
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (uc *UserCreate) AddWorkspaceRoles(w ...*WorkspaceRole) *UserCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddWorkspaceRoleIDs(ids...)
}

// AddGroupRoleIDs adds the group_roles edge to GroupRole by ids.
func (uc *UserCreate) AddGroupRoleIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddGroupRoleIDs(ids...)
	return uc
}

// AddGroupRoles adds the group_roles edges to GroupRole.
func (uc *UserCreate) AddGroupRoles(g ...*GroupRole) *UserCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupRoleIDs(ids...)
}

// AddUserRoleIDs adds the user_roles edge to UserRole by ids.
func (uc *UserCreate) AddUserRoleIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddUserRoleIDs(ids...)
	return uc
}

// AddUserRoles adds the user_roles edges to UserRole.
func (uc *UserCreate) AddUserRoles(u ...*UserRole) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Confirmed(); !ok {
		v := user.DefaultConfirmed
		uc.mutation.SetConfirmed(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if v, ok := uc.mutation.BillingID(); ok {
		if err := user.BillingIDValidator(v); err != nil {
			return &ValidationError{Name: "billing_id", err: fmt.Errorf("models: validator failed for field \"billing_id\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New("models: missing required field \"provider\"")}
	}
	if v, ok := uc.mutation.Provider(); ok {
		if err := user.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("models: validator failed for field \"provider\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("models: missing required field \"email\"")}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("models: missing required field \"password\"")}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("models: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := uc.mutation.APIKey(); ok {
		if err := user.APIKeyValidator(v); err != nil {
			return &ValidationError{Name: "api_key", err: fmt.Errorf("models: validator failed for field \"api_key\": %w", err)}
		}
	}
	if v, ok := uc.mutation.ConfirmationToken(); ok {
		if err := user.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf("models: validator failed for field \"confirmation_token\": %w", err)}
		}
	}
	if v, ok := uc.mutation.RecoveryToken(); ok {
		if err := user.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf("models: validator failed for field \"recovery_token\": %w", err)}
		}
	}
	if v, ok := uc.mutation.Otp(); ok {
		if err := user.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf("models: validator failed for field \"otp\": %w", err)}
		}
	}
	if v, ok := uc.mutation.EmailChange(); ok {
		if err := user.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf("models: validator failed for field \"email_change\": %w", err)}
		}
	}
	if v, ok := uc.mutation.EmailChangeToken(); ok {
		if err := user.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf("models: validator failed for field \"email_change_token\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "metadata", err: errors.New("models: missing required field \"metadata\"")}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("models: missing required field \"created_at\"")}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("models: missing required field \"updated_at\"")}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.BillingID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBillingID,
		})
		_node.BillingID = value
	}
	if value, ok := uc.mutation.Provider(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldProvider,
		})
		_node.Provider = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.APIKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAPIKey,
		})
		_node.APIKey = value
	}
	if value, ok := uc.mutation.Confirmed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldConfirmed,
		})
		_node.Confirmed = value
	}
	if value, ok := uc.mutation.ConfirmationSentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldConfirmationSentAt,
		})
		_node.ConfirmationSentAt = &value
	}
	if value, ok := uc.mutation.ConfirmationToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldConfirmationToken,
		})
		_node.ConfirmationToken = &value
	}
	if value, ok := uc.mutation.RecoverySentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldRecoverySentAt,
		})
		_node.RecoverySentAt = &value
	}
	if value, ok := uc.mutation.RecoveryToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRecoveryToken,
		})
		_node.RecoveryToken = &value
	}
	if value, ok := uc.mutation.OtpSentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldOtpSentAt,
		})
		_node.OtpSentAt = &value
	}
	if value, ok := uc.mutation.Otp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOtp,
		})
		_node.Otp = &value
	}
	if value, ok := uc.mutation.EmailChange(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmailChange,
		})
		_node.EmailChange = value
	}
	if value, ok := uc.mutation.EmailChangeSentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldEmailChangeSentAt,
		})
		_node.EmailChangeSentAt = &value
	}
	if value, ok := uc.mutation.EmailChangeToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmailChangeToken,
		})
		_node.EmailChangeToken = &value
	}
	if value, ok := uc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := uc.mutation.Roles(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldRoles,
		})
		_node.Roles = value
	}
	if value, ok := uc.mutation.Teams(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldTeams,
		})
		_node.Teams = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.LastSigninAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastSigninAt,
		})
		_node.LastSigninAt = &value
	}
	if nodes := uc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.WorkspaceTable,
			Columns: []string{user.WorkspaceColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WorkspaceRolesTable,
			Columns: []string{user.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.GroupRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GroupRolesTable,
			Columns: []string{user.GroupRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: grouprole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserRolesTable,
			Columns: []string{user.UserRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: userrole.FieldID,
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

// UserCreateBulk is the builder for creating a bulk of User entities.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
