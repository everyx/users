// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"time"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
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

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
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

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
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
