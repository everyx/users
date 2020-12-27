// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"time"

	"github.com/adnaan/users/internal/models/grouprole"
	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/user"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/adnaan/users/internal/models/workspace"
	"github.com/adnaan/users/internal/models/workspacerole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetBillingID sets the billing_id field.
func (uu *UserUpdate) SetBillingID(s string) *UserUpdate {
	uu.mutation.SetBillingID(s)
	return uu
}

// SetNillableBillingID sets the billing_id field if the given value is not nil.
func (uu *UserUpdate) SetNillableBillingID(s *string) *UserUpdate {
	if s != nil {
		uu.SetBillingID(*s)
	}
	return uu
}

// ClearBillingID clears the value of billing_id.
func (uu *UserUpdate) ClearBillingID() *UserUpdate {
	uu.mutation.ClearBillingID()
	return uu
}

// SetProvider sets the provider field.
func (uu *UserUpdate) SetProvider(s string) *UserUpdate {
	uu.mutation.SetProvider(s)
	return uu
}

// SetEmail sets the email field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the password field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetAPIKey sets the api_key field.
func (uu *UserUpdate) SetAPIKey(s string) *UserUpdate {
	uu.mutation.SetAPIKey(s)
	return uu
}

// SetNillableAPIKey sets the api_key field if the given value is not nil.
func (uu *UserUpdate) SetNillableAPIKey(s *string) *UserUpdate {
	if s != nil {
		uu.SetAPIKey(*s)
	}
	return uu
}

// ClearAPIKey clears the value of api_key.
func (uu *UserUpdate) ClearAPIKey() *UserUpdate {
	uu.mutation.ClearAPIKey()
	return uu
}

// SetConfirmed sets the confirmed field.
func (uu *UserUpdate) SetConfirmed(b bool) *UserUpdate {
	uu.mutation.SetConfirmed(b)
	return uu
}

// SetNillableConfirmed sets the confirmed field if the given value is not nil.
func (uu *UserUpdate) SetNillableConfirmed(b *bool) *UserUpdate {
	if b != nil {
		uu.SetConfirmed(*b)
	}
	return uu
}

// ClearConfirmed clears the value of confirmed.
func (uu *UserUpdate) ClearConfirmed() *UserUpdate {
	uu.mutation.ClearConfirmed()
	return uu
}

// SetConfirmationSentAt sets the confirmation_sent_at field.
func (uu *UserUpdate) SetConfirmationSentAt(t time.Time) *UserUpdate {
	uu.mutation.SetConfirmationSentAt(t)
	return uu
}

// SetNillableConfirmationSentAt sets the confirmation_sent_at field if the given value is not nil.
func (uu *UserUpdate) SetNillableConfirmationSentAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetConfirmationSentAt(*t)
	}
	return uu
}

// ClearConfirmationSentAt clears the value of confirmation_sent_at.
func (uu *UserUpdate) ClearConfirmationSentAt() *UserUpdate {
	uu.mutation.ClearConfirmationSentAt()
	return uu
}

// SetConfirmationToken sets the confirmation_token field.
func (uu *UserUpdate) SetConfirmationToken(s string) *UserUpdate {
	uu.mutation.SetConfirmationToken(s)
	return uu
}

// SetNillableConfirmationToken sets the confirmation_token field if the given value is not nil.
func (uu *UserUpdate) SetNillableConfirmationToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetConfirmationToken(*s)
	}
	return uu
}

// ClearConfirmationToken clears the value of confirmation_token.
func (uu *UserUpdate) ClearConfirmationToken() *UserUpdate {
	uu.mutation.ClearConfirmationToken()
	return uu
}

// SetRecoverySentAt sets the recovery_sent_at field.
func (uu *UserUpdate) SetRecoverySentAt(t time.Time) *UserUpdate {
	uu.mutation.SetRecoverySentAt(t)
	return uu
}

// SetNillableRecoverySentAt sets the recovery_sent_at field if the given value is not nil.
func (uu *UserUpdate) SetNillableRecoverySentAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetRecoverySentAt(*t)
	}
	return uu
}

// ClearRecoverySentAt clears the value of recovery_sent_at.
func (uu *UserUpdate) ClearRecoverySentAt() *UserUpdate {
	uu.mutation.ClearRecoverySentAt()
	return uu
}

// SetRecoveryToken sets the recovery_token field.
func (uu *UserUpdate) SetRecoveryToken(s string) *UserUpdate {
	uu.mutation.SetRecoveryToken(s)
	return uu
}

// SetNillableRecoveryToken sets the recovery_token field if the given value is not nil.
func (uu *UserUpdate) SetNillableRecoveryToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetRecoveryToken(*s)
	}
	return uu
}

// ClearRecoveryToken clears the value of recovery_token.
func (uu *UserUpdate) ClearRecoveryToken() *UserUpdate {
	uu.mutation.ClearRecoveryToken()
	return uu
}

// SetOtpSentAt sets the otp_sent_at field.
func (uu *UserUpdate) SetOtpSentAt(t time.Time) *UserUpdate {
	uu.mutation.SetOtpSentAt(t)
	return uu
}

// SetNillableOtpSentAt sets the otp_sent_at field if the given value is not nil.
func (uu *UserUpdate) SetNillableOtpSentAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetOtpSentAt(*t)
	}
	return uu
}

// ClearOtpSentAt clears the value of otp_sent_at.
func (uu *UserUpdate) ClearOtpSentAt() *UserUpdate {
	uu.mutation.ClearOtpSentAt()
	return uu
}

// SetOtp sets the otp field.
func (uu *UserUpdate) SetOtp(s string) *UserUpdate {
	uu.mutation.SetOtp(s)
	return uu
}

// SetNillableOtp sets the otp field if the given value is not nil.
func (uu *UserUpdate) SetNillableOtp(s *string) *UserUpdate {
	if s != nil {
		uu.SetOtp(*s)
	}
	return uu
}

// ClearOtp clears the value of otp.
func (uu *UserUpdate) ClearOtp() *UserUpdate {
	uu.mutation.ClearOtp()
	return uu
}

// SetEmailChange sets the email_change field.
func (uu *UserUpdate) SetEmailChange(s string) *UserUpdate {
	uu.mutation.SetEmailChange(s)
	return uu
}

// SetNillableEmailChange sets the email_change field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmailChange(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmailChange(*s)
	}
	return uu
}

// ClearEmailChange clears the value of email_change.
func (uu *UserUpdate) ClearEmailChange() *UserUpdate {
	uu.mutation.ClearEmailChange()
	return uu
}

// SetEmailChangeSentAt sets the email_change_sent_at field.
func (uu *UserUpdate) SetEmailChangeSentAt(t time.Time) *UserUpdate {
	uu.mutation.SetEmailChangeSentAt(t)
	return uu
}

// SetNillableEmailChangeSentAt sets the email_change_sent_at field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmailChangeSentAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetEmailChangeSentAt(*t)
	}
	return uu
}

// ClearEmailChangeSentAt clears the value of email_change_sent_at.
func (uu *UserUpdate) ClearEmailChangeSentAt() *UserUpdate {
	uu.mutation.ClearEmailChangeSentAt()
	return uu
}

// SetEmailChangeToken sets the email_change_token field.
func (uu *UserUpdate) SetEmailChangeToken(s string) *UserUpdate {
	uu.mutation.SetEmailChangeToken(s)
	return uu
}

// SetNillableEmailChangeToken sets the email_change_token field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmailChangeToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmailChangeToken(*s)
	}
	return uu
}

// ClearEmailChangeToken clears the value of email_change_token.
func (uu *UserUpdate) ClearEmailChangeToken() *UserUpdate {
	uu.mutation.ClearEmailChangeToken()
	return uu
}

// SetMetadata sets the metadata field.
func (uu *UserUpdate) SetMetadata(m map[string]interface{}) *UserUpdate {
	uu.mutation.SetMetadata(m)
	return uu
}

// SetRoles sets the roles field.
func (uu *UserUpdate) SetRoles(s []string) *UserUpdate {
	uu.mutation.SetRoles(s)
	return uu
}

// ClearRoles clears the value of roles.
func (uu *UserUpdate) ClearRoles() *UserUpdate {
	uu.mutation.ClearRoles()
	return uu
}

// SetTeams sets the teams field.
func (uu *UserUpdate) SetTeams(m map[string]string) *UserUpdate {
	uu.mutation.SetTeams(m)
	return uu
}

// ClearTeams clears the value of teams.
func (uu *UserUpdate) ClearTeams() *UserUpdate {
	uu.mutation.ClearTeams()
	return uu
}

// SetUpdatedAt sets the updated_at field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetLastSigninAt sets the last_signin_at field.
func (uu *UserUpdate) SetLastSigninAt(t time.Time) *UserUpdate {
	uu.mutation.SetLastSigninAt(t)
	return uu
}

// SetNillableLastSigninAt sets the last_signin_at field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastSigninAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastSigninAt(*t)
	}
	return uu
}

// ClearLastSigninAt clears the value of last_signin_at.
func (uu *UserUpdate) ClearLastSigninAt() *UserUpdate {
	uu.mutation.ClearLastSigninAt()
	return uu
}

// SetWorkspaceID sets the workspace edge to Workspace by id.
func (uu *UserUpdate) SetWorkspaceID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetWorkspaceID(id)
	return uu
}

// SetNillableWorkspaceID sets the workspace edge to Workspace by id if the given value is not nil.
func (uu *UserUpdate) SetNillableWorkspaceID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetWorkspaceID(*id)
	}
	return uu
}

// SetWorkspace sets the workspace edge to Workspace.
func (uu *UserUpdate) SetWorkspace(w *Workspace) *UserUpdate {
	return uu.SetWorkspaceID(w.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (uu *UserUpdate) AddWorkspaceRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddWorkspaceRoleIDs(ids...)
	return uu
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (uu *UserUpdate) AddWorkspaceRoles(w ...*WorkspaceRole) *UserUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.AddWorkspaceRoleIDs(ids...)
}

// AddGroupRoleIDs adds the group_roles edge to GroupRole by ids.
func (uu *UserUpdate) AddGroupRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddGroupRoleIDs(ids...)
	return uu
}

// AddGroupRoles adds the group_roles edges to GroupRole.
func (uu *UserUpdate) AddGroupRoles(g ...*GroupRole) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupRoleIDs(ids...)
}

// AddUserRoleIDs adds the user_roles edge to UserRole by ids.
func (uu *UserUpdate) AddUserRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddUserRoleIDs(ids...)
	return uu
}

// AddUserRoles adds the user_roles edges to UserRole.
func (uu *UserUpdate) AddUserRoles(u ...*UserRole) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUserRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearWorkspace clears the "workspace" edge to type Workspace.
func (uu *UserUpdate) ClearWorkspace() *UserUpdate {
	uu.mutation.ClearWorkspace()
	return uu
}

// ClearWorkspaceRoles clears all "workspace_roles" edges to type WorkspaceRole.
func (uu *UserUpdate) ClearWorkspaceRoles() *UserUpdate {
	uu.mutation.ClearWorkspaceRoles()
	return uu
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (uu *UserUpdate) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveWorkspaceRoleIDs(ids...)
	return uu
}

// RemoveWorkspaceRoles removes workspace_roles edges to WorkspaceRole.
func (uu *UserUpdate) RemoveWorkspaceRoles(w ...*WorkspaceRole) *UserUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.RemoveWorkspaceRoleIDs(ids...)
}

// ClearGroupRoles clears all "group_roles" edges to type GroupRole.
func (uu *UserUpdate) ClearGroupRoles() *UserUpdate {
	uu.mutation.ClearGroupRoles()
	return uu
}

// RemoveGroupRoleIDs removes the group_roles edge to GroupRole by ids.
func (uu *UserUpdate) RemoveGroupRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveGroupRoleIDs(ids...)
	return uu
}

// RemoveGroupRoles removes group_roles edges to GroupRole.
func (uu *UserUpdate) RemoveGroupRoles(g ...*GroupRole) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupRoleIDs(ids...)
}

// ClearUserRoles clears all "user_roles" edges to type UserRole.
func (uu *UserUpdate) ClearUserRoles() *UserUpdate {
	uu.mutation.ClearUserRoles()
	return uu
}

// RemoveUserRoleIDs removes the user_roles edge to UserRole by ids.
func (uu *UserUpdate) RemoveUserRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveUserRoleIDs(ids...)
	return uu
}

// RemoveUserRoles removes user_roles edges to UserRole.
func (uu *UserUpdate) RemoveUserRoles(u ...*UserRole) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUserRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.BillingID(); ok {
		if err := user.BillingIDValidator(v); err != nil {
			return &ValidationError{Name: "billing_id", err: fmt.Errorf("models: validator failed for field \"billing_id\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Provider(); ok {
		if err := user.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("models: validator failed for field \"provider\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("models: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := uu.mutation.APIKey(); ok {
		if err := user.APIKeyValidator(v); err != nil {
			return &ValidationError{Name: "api_key", err: fmt.Errorf("models: validator failed for field \"api_key\": %w", err)}
		}
	}
	if v, ok := uu.mutation.ConfirmationToken(); ok {
		if err := user.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf("models: validator failed for field \"confirmation_token\": %w", err)}
		}
	}
	if v, ok := uu.mutation.RecoveryToken(); ok {
		if err := user.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf("models: validator failed for field \"recovery_token\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Otp(); ok {
		if err := user.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf("models: validator failed for field \"otp\": %w", err)}
		}
	}
	if v, ok := uu.mutation.EmailChange(); ok {
		if err := user.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf("models: validator failed for field \"email_change\": %w", err)}
		}
	}
	if v, ok := uu.mutation.EmailChangeToken(); ok {
		if err := user.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf("models: validator failed for field \"email_change_token\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.BillingID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBillingID,
		})
	}
	if uu.mutation.BillingIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldBillingID,
		})
	}
	if value, ok := uu.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldProvider,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uu.mutation.APIKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAPIKey,
		})
	}
	if uu.mutation.APIKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAPIKey,
		})
	}
	if value, ok := uu.mutation.Confirmed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldConfirmed,
		})
	}
	if uu.mutation.ConfirmedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldConfirmed,
		})
	}
	if value, ok := uu.mutation.ConfirmationSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldConfirmationSentAt,
		})
	}
	if uu.mutation.ConfirmationSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldConfirmationSentAt,
		})
	}
	if value, ok := uu.mutation.ConfirmationToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldConfirmationToken,
		})
	}
	if uu.mutation.ConfirmationTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldConfirmationToken,
		})
	}
	if value, ok := uu.mutation.RecoverySentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldRecoverySentAt,
		})
	}
	if uu.mutation.RecoverySentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldRecoverySentAt,
		})
	}
	if value, ok := uu.mutation.RecoveryToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRecoveryToken,
		})
	}
	if uu.mutation.RecoveryTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldRecoveryToken,
		})
	}
	if value, ok := uu.mutation.OtpSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldOtpSentAt,
		})
	}
	if uu.mutation.OtpSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldOtpSentAt,
		})
	}
	if value, ok := uu.mutation.Otp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOtp,
		})
	}
	if uu.mutation.OtpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldOtp,
		})
	}
	if value, ok := uu.mutation.EmailChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmailChange,
		})
	}
	if uu.mutation.EmailChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmailChange,
		})
	}
	if value, ok := uu.mutation.EmailChangeSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldEmailChangeSentAt,
		})
	}
	if uu.mutation.EmailChangeSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldEmailChangeSentAt,
		})
	}
	if value, ok := uu.mutation.EmailChangeToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmailChangeToken,
		})
	}
	if uu.mutation.EmailChangeTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmailChangeToken,
		})
	}
	if value, ok := uu.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldMetadata,
		})
	}
	if value, ok := uu.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldRoles,
		})
	}
	if uu.mutation.RolesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldRoles,
		})
	}
	if value, ok := uu.mutation.Teams(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldTeams,
		})
	}
	if uu.mutation.TeamsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldTeams,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uu.mutation.LastSigninAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastSigninAt,
		})
	}
	if uu.mutation.LastSigninAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldLastSigninAt,
		})
	}
	if uu.mutation.WorkspaceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.WorkspaceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedWorkspaceRolesIDs(); len(nodes) > 0 && !uu.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GroupRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupRolesIDs(); len(nodes) > 0 && !uu.mutation.GroupRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UserRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserRolesIDs(); len(nodes) > 0 && !uu.mutation.UserRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetBillingID sets the billing_id field.
func (uuo *UserUpdateOne) SetBillingID(s string) *UserUpdateOne {
	uuo.mutation.SetBillingID(s)
	return uuo
}

// SetNillableBillingID sets the billing_id field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBillingID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBillingID(*s)
	}
	return uuo
}

// ClearBillingID clears the value of billing_id.
func (uuo *UserUpdateOne) ClearBillingID() *UserUpdateOne {
	uuo.mutation.ClearBillingID()
	return uuo
}

// SetProvider sets the provider field.
func (uuo *UserUpdateOne) SetProvider(s string) *UserUpdateOne {
	uuo.mutation.SetProvider(s)
	return uuo
}

// SetEmail sets the email field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the password field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetAPIKey sets the api_key field.
func (uuo *UserUpdateOne) SetAPIKey(s string) *UserUpdateOne {
	uuo.mutation.SetAPIKey(s)
	return uuo
}

// SetNillableAPIKey sets the api_key field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAPIKey(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAPIKey(*s)
	}
	return uuo
}

// ClearAPIKey clears the value of api_key.
func (uuo *UserUpdateOne) ClearAPIKey() *UserUpdateOne {
	uuo.mutation.ClearAPIKey()
	return uuo
}

// SetConfirmed sets the confirmed field.
func (uuo *UserUpdateOne) SetConfirmed(b bool) *UserUpdateOne {
	uuo.mutation.SetConfirmed(b)
	return uuo
}

// SetNillableConfirmed sets the confirmed field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableConfirmed(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetConfirmed(*b)
	}
	return uuo
}

// ClearConfirmed clears the value of confirmed.
func (uuo *UserUpdateOne) ClearConfirmed() *UserUpdateOne {
	uuo.mutation.ClearConfirmed()
	return uuo
}

// SetConfirmationSentAt sets the confirmation_sent_at field.
func (uuo *UserUpdateOne) SetConfirmationSentAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetConfirmationSentAt(t)
	return uuo
}

// SetNillableConfirmationSentAt sets the confirmation_sent_at field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableConfirmationSentAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetConfirmationSentAt(*t)
	}
	return uuo
}

// ClearConfirmationSentAt clears the value of confirmation_sent_at.
func (uuo *UserUpdateOne) ClearConfirmationSentAt() *UserUpdateOne {
	uuo.mutation.ClearConfirmationSentAt()
	return uuo
}

// SetConfirmationToken sets the confirmation_token field.
func (uuo *UserUpdateOne) SetConfirmationToken(s string) *UserUpdateOne {
	uuo.mutation.SetConfirmationToken(s)
	return uuo
}

// SetNillableConfirmationToken sets the confirmation_token field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableConfirmationToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetConfirmationToken(*s)
	}
	return uuo
}

// ClearConfirmationToken clears the value of confirmation_token.
func (uuo *UserUpdateOne) ClearConfirmationToken() *UserUpdateOne {
	uuo.mutation.ClearConfirmationToken()
	return uuo
}

// SetRecoverySentAt sets the recovery_sent_at field.
func (uuo *UserUpdateOne) SetRecoverySentAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetRecoverySentAt(t)
	return uuo
}

// SetNillableRecoverySentAt sets the recovery_sent_at field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRecoverySentAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetRecoverySentAt(*t)
	}
	return uuo
}

// ClearRecoverySentAt clears the value of recovery_sent_at.
func (uuo *UserUpdateOne) ClearRecoverySentAt() *UserUpdateOne {
	uuo.mutation.ClearRecoverySentAt()
	return uuo
}

// SetRecoveryToken sets the recovery_token field.
func (uuo *UserUpdateOne) SetRecoveryToken(s string) *UserUpdateOne {
	uuo.mutation.SetRecoveryToken(s)
	return uuo
}

// SetNillableRecoveryToken sets the recovery_token field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRecoveryToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRecoveryToken(*s)
	}
	return uuo
}

// ClearRecoveryToken clears the value of recovery_token.
func (uuo *UserUpdateOne) ClearRecoveryToken() *UserUpdateOne {
	uuo.mutation.ClearRecoveryToken()
	return uuo
}

// SetOtpSentAt sets the otp_sent_at field.
func (uuo *UserUpdateOne) SetOtpSentAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetOtpSentAt(t)
	return uuo
}

// SetNillableOtpSentAt sets the otp_sent_at field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOtpSentAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetOtpSentAt(*t)
	}
	return uuo
}

// ClearOtpSentAt clears the value of otp_sent_at.
func (uuo *UserUpdateOne) ClearOtpSentAt() *UserUpdateOne {
	uuo.mutation.ClearOtpSentAt()
	return uuo
}

// SetOtp sets the otp field.
func (uuo *UserUpdateOne) SetOtp(s string) *UserUpdateOne {
	uuo.mutation.SetOtp(s)
	return uuo
}

// SetNillableOtp sets the otp field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOtp(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetOtp(*s)
	}
	return uuo
}

// ClearOtp clears the value of otp.
func (uuo *UserUpdateOne) ClearOtp() *UserUpdateOne {
	uuo.mutation.ClearOtp()
	return uuo
}

// SetEmailChange sets the email_change field.
func (uuo *UserUpdateOne) SetEmailChange(s string) *UserUpdateOne {
	uuo.mutation.SetEmailChange(s)
	return uuo
}

// SetNillableEmailChange sets the email_change field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmailChange(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmailChange(*s)
	}
	return uuo
}

// ClearEmailChange clears the value of email_change.
func (uuo *UserUpdateOne) ClearEmailChange() *UserUpdateOne {
	uuo.mutation.ClearEmailChange()
	return uuo
}

// SetEmailChangeSentAt sets the email_change_sent_at field.
func (uuo *UserUpdateOne) SetEmailChangeSentAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetEmailChangeSentAt(t)
	return uuo
}

// SetNillableEmailChangeSentAt sets the email_change_sent_at field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmailChangeSentAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetEmailChangeSentAt(*t)
	}
	return uuo
}

// ClearEmailChangeSentAt clears the value of email_change_sent_at.
func (uuo *UserUpdateOne) ClearEmailChangeSentAt() *UserUpdateOne {
	uuo.mutation.ClearEmailChangeSentAt()
	return uuo
}

// SetEmailChangeToken sets the email_change_token field.
func (uuo *UserUpdateOne) SetEmailChangeToken(s string) *UserUpdateOne {
	uuo.mutation.SetEmailChangeToken(s)
	return uuo
}

// SetNillableEmailChangeToken sets the email_change_token field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmailChangeToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmailChangeToken(*s)
	}
	return uuo
}

// ClearEmailChangeToken clears the value of email_change_token.
func (uuo *UserUpdateOne) ClearEmailChangeToken() *UserUpdateOne {
	uuo.mutation.ClearEmailChangeToken()
	return uuo
}

// SetMetadata sets the metadata field.
func (uuo *UserUpdateOne) SetMetadata(m map[string]interface{}) *UserUpdateOne {
	uuo.mutation.SetMetadata(m)
	return uuo
}

// SetRoles sets the roles field.
func (uuo *UserUpdateOne) SetRoles(s []string) *UserUpdateOne {
	uuo.mutation.SetRoles(s)
	return uuo
}

// ClearRoles clears the value of roles.
func (uuo *UserUpdateOne) ClearRoles() *UserUpdateOne {
	uuo.mutation.ClearRoles()
	return uuo
}

// SetTeams sets the teams field.
func (uuo *UserUpdateOne) SetTeams(m map[string]string) *UserUpdateOne {
	uuo.mutation.SetTeams(m)
	return uuo
}

// ClearTeams clears the value of teams.
func (uuo *UserUpdateOne) ClearTeams() *UserUpdateOne {
	uuo.mutation.ClearTeams()
	return uuo
}

// SetUpdatedAt sets the updated_at field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetLastSigninAt sets the last_signin_at field.
func (uuo *UserUpdateOne) SetLastSigninAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastSigninAt(t)
	return uuo
}

// SetNillableLastSigninAt sets the last_signin_at field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastSigninAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastSigninAt(*t)
	}
	return uuo
}

// ClearLastSigninAt clears the value of last_signin_at.
func (uuo *UserUpdateOne) ClearLastSigninAt() *UserUpdateOne {
	uuo.mutation.ClearLastSigninAt()
	return uuo
}

// SetWorkspaceID sets the workspace edge to Workspace by id.
func (uuo *UserUpdateOne) SetWorkspaceID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetWorkspaceID(id)
	return uuo
}

// SetNillableWorkspaceID sets the workspace edge to Workspace by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWorkspaceID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetWorkspaceID(*id)
	}
	return uuo
}

// SetWorkspace sets the workspace edge to Workspace.
func (uuo *UserUpdateOne) SetWorkspace(w *Workspace) *UserUpdateOne {
	return uuo.SetWorkspaceID(w.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (uuo *UserUpdateOne) AddWorkspaceRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddWorkspaceRoleIDs(ids...)
	return uuo
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (uuo *UserUpdateOne) AddWorkspaceRoles(w ...*WorkspaceRole) *UserUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.AddWorkspaceRoleIDs(ids...)
}

// AddGroupRoleIDs adds the group_roles edge to GroupRole by ids.
func (uuo *UserUpdateOne) AddGroupRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddGroupRoleIDs(ids...)
	return uuo
}

// AddGroupRoles adds the group_roles edges to GroupRole.
func (uuo *UserUpdateOne) AddGroupRoles(g ...*GroupRole) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupRoleIDs(ids...)
}

// AddUserRoleIDs adds the user_roles edge to UserRole by ids.
func (uuo *UserUpdateOne) AddUserRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddUserRoleIDs(ids...)
	return uuo
}

// AddUserRoles adds the user_roles edges to UserRole.
func (uuo *UserUpdateOne) AddUserRoles(u ...*UserRole) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUserRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearWorkspace clears the "workspace" edge to type Workspace.
func (uuo *UserUpdateOne) ClearWorkspace() *UserUpdateOne {
	uuo.mutation.ClearWorkspace()
	return uuo
}

// ClearWorkspaceRoles clears all "workspace_roles" edges to type WorkspaceRole.
func (uuo *UserUpdateOne) ClearWorkspaceRoles() *UserUpdateOne {
	uuo.mutation.ClearWorkspaceRoles()
	return uuo
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (uuo *UserUpdateOne) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveWorkspaceRoleIDs(ids...)
	return uuo
}

// RemoveWorkspaceRoles removes workspace_roles edges to WorkspaceRole.
func (uuo *UserUpdateOne) RemoveWorkspaceRoles(w ...*WorkspaceRole) *UserUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.RemoveWorkspaceRoleIDs(ids...)
}

// ClearGroupRoles clears all "group_roles" edges to type GroupRole.
func (uuo *UserUpdateOne) ClearGroupRoles() *UserUpdateOne {
	uuo.mutation.ClearGroupRoles()
	return uuo
}

// RemoveGroupRoleIDs removes the group_roles edge to GroupRole by ids.
func (uuo *UserUpdateOne) RemoveGroupRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveGroupRoleIDs(ids...)
	return uuo
}

// RemoveGroupRoles removes group_roles edges to GroupRole.
func (uuo *UserUpdateOne) RemoveGroupRoles(g ...*GroupRole) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupRoleIDs(ids...)
}

// ClearUserRoles clears all "user_roles" edges to type UserRole.
func (uuo *UserUpdateOne) ClearUserRoles() *UserUpdateOne {
	uuo.mutation.ClearUserRoles()
	return uuo
}

// RemoveUserRoleIDs removes the user_roles edge to UserRole by ids.
func (uuo *UserUpdateOne) RemoveUserRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveUserRoleIDs(ids...)
	return uuo
}

// RemoveUserRoles removes user_roles edges to UserRole.
func (uuo *UserUpdateOne) RemoveUserRoles(u ...*UserRole) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUserRoleIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.BillingID(); ok {
		if err := user.BillingIDValidator(v); err != nil {
			return &ValidationError{Name: "billing_id", err: fmt.Errorf("models: validator failed for field \"billing_id\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Provider(); ok {
		if err := user.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("models: validator failed for field \"provider\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("models: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.APIKey(); ok {
		if err := user.APIKeyValidator(v); err != nil {
			return &ValidationError{Name: "api_key", err: fmt.Errorf("models: validator failed for field \"api_key\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.ConfirmationToken(); ok {
		if err := user.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf("models: validator failed for field \"confirmation_token\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.RecoveryToken(); ok {
		if err := user.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf("models: validator failed for field \"recovery_token\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Otp(); ok {
		if err := user.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf("models: validator failed for field \"otp\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.EmailChange(); ok {
		if err := user.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf("models: validator failed for field \"email_change\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.EmailChangeToken(); ok {
		if err := user.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf("models: validator failed for field \"email_change_token\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.BillingID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBillingID,
		})
	}
	if uuo.mutation.BillingIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldBillingID,
		})
	}
	if value, ok := uuo.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldProvider,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.APIKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAPIKey,
		})
	}
	if uuo.mutation.APIKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAPIKey,
		})
	}
	if value, ok := uuo.mutation.Confirmed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldConfirmed,
		})
	}
	if uuo.mutation.ConfirmedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldConfirmed,
		})
	}
	if value, ok := uuo.mutation.ConfirmationSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldConfirmationSentAt,
		})
	}
	if uuo.mutation.ConfirmationSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldConfirmationSentAt,
		})
	}
	if value, ok := uuo.mutation.ConfirmationToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldConfirmationToken,
		})
	}
	if uuo.mutation.ConfirmationTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldConfirmationToken,
		})
	}
	if value, ok := uuo.mutation.RecoverySentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldRecoverySentAt,
		})
	}
	if uuo.mutation.RecoverySentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldRecoverySentAt,
		})
	}
	if value, ok := uuo.mutation.RecoveryToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRecoveryToken,
		})
	}
	if uuo.mutation.RecoveryTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldRecoveryToken,
		})
	}
	if value, ok := uuo.mutation.OtpSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldOtpSentAt,
		})
	}
	if uuo.mutation.OtpSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldOtpSentAt,
		})
	}
	if value, ok := uuo.mutation.Otp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOtp,
		})
	}
	if uuo.mutation.OtpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldOtp,
		})
	}
	if value, ok := uuo.mutation.EmailChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmailChange,
		})
	}
	if uuo.mutation.EmailChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmailChange,
		})
	}
	if value, ok := uuo.mutation.EmailChangeSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldEmailChangeSentAt,
		})
	}
	if uuo.mutation.EmailChangeSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldEmailChangeSentAt,
		})
	}
	if value, ok := uuo.mutation.EmailChangeToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmailChangeToken,
		})
	}
	if uuo.mutation.EmailChangeTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmailChangeToken,
		})
	}
	if value, ok := uuo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldMetadata,
		})
	}
	if value, ok := uuo.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldRoles,
		})
	}
	if uuo.mutation.RolesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldRoles,
		})
	}
	if value, ok := uuo.mutation.Teams(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldTeams,
		})
	}
	if uuo.mutation.TeamsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldTeams,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uuo.mutation.LastSigninAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastSigninAt,
		})
	}
	if uuo.mutation.LastSigninAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldLastSigninAt,
		})
	}
	if uuo.mutation.WorkspaceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.WorkspaceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedWorkspaceRolesIDs(); len(nodes) > 0 && !uuo.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GroupRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupRolesIDs(); len(nodes) > 0 && !uuo.mutation.GroupRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UserRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserRolesIDs(); len(nodes) > 0 && !uuo.mutation.UserRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
