// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BillingID applies equality check predicate on the "billing_id" field. It's identical to BillingIDEQ.
func BillingID(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBillingID), v))
	})
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvider), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// APIKey applies equality check predicate on the "api_key" field. It's identical to APIKeyEQ.
func APIKey(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIKey), v))
	})
}

// Confirmed applies equality check predicate on the "confirmed" field. It's identical to ConfirmedEQ.
func Confirmed(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfirmed), v))
	})
}

// ConfirmationSentAt applies equality check predicate on the "confirmation_sent_at" field. It's identical to ConfirmationSentAtEQ.
func ConfirmationSentAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfirmationSentAt), v))
	})
}

// ConfirmationToken applies equality check predicate on the "confirmation_token" field. It's identical to ConfirmationTokenEQ.
func ConfirmationToken(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfirmationToken), v))
	})
}

// RecoverySentAt applies equality check predicate on the "recovery_sent_at" field. It's identical to RecoverySentAtEQ.
func RecoverySentAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecoverySentAt), v))
	})
}

// RecoveryToken applies equality check predicate on the "recovery_token" field. It's identical to RecoveryTokenEQ.
func RecoveryToken(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecoveryToken), v))
	})
}

// OtpSentAt applies equality check predicate on the "otp_sent_at" field. It's identical to OtpSentAtEQ.
func OtpSentAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtpSentAt), v))
	})
}

// Otp applies equality check predicate on the "otp" field. It's identical to OtpEQ.
func Otp(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtp), v))
	})
}

// EmailChange applies equality check predicate on the "email_change" field. It's identical to EmailChangeEQ.
func EmailChange(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailChange), v))
	})
}

// EmailChangeSentAt applies equality check predicate on the "email_change_sent_at" field. It's identical to EmailChangeSentAtEQ.
func EmailChangeSentAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailChangeSentAt), v))
	})
}

// EmailChangeToken applies equality check predicate on the "email_change_token" field. It's identical to EmailChangeTokenEQ.
func EmailChangeToken(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailChangeToken), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// LastSigninAt applies equality check predicate on the "last_signin_at" field. It's identical to LastSigninAtEQ.
func LastSigninAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSigninAt), v))
	})
}

// BillingIDEQ applies the EQ predicate on the "billing_id" field.
func BillingIDEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBillingID), v))
	})
}

// BillingIDNEQ applies the NEQ predicate on the "billing_id" field.
func BillingIDNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBillingID), v))
	})
}

// BillingIDIn applies the In predicate on the "billing_id" field.
func BillingIDIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBillingID), v...))
	})
}

// BillingIDNotIn applies the NotIn predicate on the "billing_id" field.
func BillingIDNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBillingID), v...))
	})
}

// BillingIDGT applies the GT predicate on the "billing_id" field.
func BillingIDGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBillingID), v))
	})
}

// BillingIDGTE applies the GTE predicate on the "billing_id" field.
func BillingIDGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBillingID), v))
	})
}

// BillingIDLT applies the LT predicate on the "billing_id" field.
func BillingIDLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBillingID), v))
	})
}

// BillingIDLTE applies the LTE predicate on the "billing_id" field.
func BillingIDLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBillingID), v))
	})
}

// BillingIDContains applies the Contains predicate on the "billing_id" field.
func BillingIDContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBillingID), v))
	})
}

// BillingIDHasPrefix applies the HasPrefix predicate on the "billing_id" field.
func BillingIDHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBillingID), v))
	})
}

// BillingIDHasSuffix applies the HasSuffix predicate on the "billing_id" field.
func BillingIDHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBillingID), v))
	})
}

// BillingIDIsNil applies the IsNil predicate on the "billing_id" field.
func BillingIDIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBillingID)))
	})
}

// BillingIDNotNil applies the NotNil predicate on the "billing_id" field.
func BillingIDNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBillingID)))
	})
}

// BillingIDEqualFold applies the EqualFold predicate on the "billing_id" field.
func BillingIDEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBillingID), v))
	})
}

// BillingIDContainsFold applies the ContainsFold predicate on the "billing_id" field.
func BillingIDContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBillingID), v))
	})
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvider), v))
	})
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProvider), v))
	})
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProvider), v...))
	})
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProvider), v...))
	})
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProvider), v))
	})
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProvider), v))
	})
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProvider), v))
	})
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProvider), v))
	})
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProvider), v))
	})
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProvider), v))
	})
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProvider), v))
	})
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProvider), v))
	})
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProvider), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// APIKeyEQ applies the EQ predicate on the "api_key" field.
func APIKeyEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIKey), v))
	})
}

// APIKeyNEQ applies the NEQ predicate on the "api_key" field.
func APIKeyNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAPIKey), v))
	})
}

// APIKeyIn applies the In predicate on the "api_key" field.
func APIKeyIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAPIKey), v...))
	})
}

// APIKeyNotIn applies the NotIn predicate on the "api_key" field.
func APIKeyNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAPIKey), v...))
	})
}

// APIKeyGT applies the GT predicate on the "api_key" field.
func APIKeyGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAPIKey), v))
	})
}

// APIKeyGTE applies the GTE predicate on the "api_key" field.
func APIKeyGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAPIKey), v))
	})
}

// APIKeyLT applies the LT predicate on the "api_key" field.
func APIKeyLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAPIKey), v))
	})
}

// APIKeyLTE applies the LTE predicate on the "api_key" field.
func APIKeyLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAPIKey), v))
	})
}

// APIKeyContains applies the Contains predicate on the "api_key" field.
func APIKeyContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAPIKey), v))
	})
}

// APIKeyHasPrefix applies the HasPrefix predicate on the "api_key" field.
func APIKeyHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAPIKey), v))
	})
}

// APIKeyHasSuffix applies the HasSuffix predicate on the "api_key" field.
func APIKeyHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAPIKey), v))
	})
}

// APIKeyIsNil applies the IsNil predicate on the "api_key" field.
func APIKeyIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAPIKey)))
	})
}

// APIKeyNotNil applies the NotNil predicate on the "api_key" field.
func APIKeyNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAPIKey)))
	})
}

// APIKeyEqualFold applies the EqualFold predicate on the "api_key" field.
func APIKeyEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAPIKey), v))
	})
}

// APIKeyContainsFold applies the ContainsFold predicate on the "api_key" field.
func APIKeyContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAPIKey), v))
	})
}

// ConfirmedEQ applies the EQ predicate on the "confirmed" field.
func ConfirmedEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfirmed), v))
	})
}

// ConfirmedNEQ applies the NEQ predicate on the "confirmed" field.
func ConfirmedNEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfirmed), v))
	})
}

// ConfirmedIsNil applies the IsNil predicate on the "confirmed" field.
func ConfirmedIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldConfirmed)))
	})
}

// ConfirmedNotNil applies the NotNil predicate on the "confirmed" field.
func ConfirmedNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldConfirmed)))
	})
}

// ConfirmationSentAtEQ applies the EQ predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfirmationSentAt), v))
	})
}

// ConfirmationSentAtNEQ applies the NEQ predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfirmationSentAt), v))
	})
}

// ConfirmationSentAtIn applies the In predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfirmationSentAt), v...))
	})
}

// ConfirmationSentAtNotIn applies the NotIn predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfirmationSentAt), v...))
	})
}

// ConfirmationSentAtGT applies the GT predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfirmationSentAt), v))
	})
}

// ConfirmationSentAtGTE applies the GTE predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfirmationSentAt), v))
	})
}

// ConfirmationSentAtLT applies the LT predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfirmationSentAt), v))
	})
}

// ConfirmationSentAtLTE applies the LTE predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfirmationSentAt), v))
	})
}

// ConfirmationSentAtIsNil applies the IsNil predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldConfirmationSentAt)))
	})
}

// ConfirmationSentAtNotNil applies the NotNil predicate on the "confirmation_sent_at" field.
func ConfirmationSentAtNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldConfirmationSentAt)))
	})
}

// ConfirmationTokenEQ applies the EQ predicate on the "confirmation_token" field.
func ConfirmationTokenEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenNEQ applies the NEQ predicate on the "confirmation_token" field.
func ConfirmationTokenNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenIn applies the In predicate on the "confirmation_token" field.
func ConfirmationTokenIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfirmationToken), v...))
	})
}

// ConfirmationTokenNotIn applies the NotIn predicate on the "confirmation_token" field.
func ConfirmationTokenNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfirmationToken), v...))
	})
}

// ConfirmationTokenGT applies the GT predicate on the "confirmation_token" field.
func ConfirmationTokenGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenGTE applies the GTE predicate on the "confirmation_token" field.
func ConfirmationTokenGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenLT applies the LT predicate on the "confirmation_token" field.
func ConfirmationTokenLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenLTE applies the LTE predicate on the "confirmation_token" field.
func ConfirmationTokenLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenContains applies the Contains predicate on the "confirmation_token" field.
func ConfirmationTokenContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenHasPrefix applies the HasPrefix predicate on the "confirmation_token" field.
func ConfirmationTokenHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenHasSuffix applies the HasSuffix predicate on the "confirmation_token" field.
func ConfirmationTokenHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenIsNil applies the IsNil predicate on the "confirmation_token" field.
func ConfirmationTokenIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldConfirmationToken)))
	})
}

// ConfirmationTokenNotNil applies the NotNil predicate on the "confirmation_token" field.
func ConfirmationTokenNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldConfirmationToken)))
	})
}

// ConfirmationTokenEqualFold applies the EqualFold predicate on the "confirmation_token" field.
func ConfirmationTokenEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConfirmationToken), v))
	})
}

// ConfirmationTokenContainsFold applies the ContainsFold predicate on the "confirmation_token" field.
func ConfirmationTokenContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConfirmationToken), v))
	})
}

// RecoverySentAtEQ applies the EQ predicate on the "recovery_sent_at" field.
func RecoverySentAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecoverySentAt), v))
	})
}

// RecoverySentAtNEQ applies the NEQ predicate on the "recovery_sent_at" field.
func RecoverySentAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecoverySentAt), v))
	})
}

// RecoverySentAtIn applies the In predicate on the "recovery_sent_at" field.
func RecoverySentAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecoverySentAt), v...))
	})
}

// RecoverySentAtNotIn applies the NotIn predicate on the "recovery_sent_at" field.
func RecoverySentAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecoverySentAt), v...))
	})
}

// RecoverySentAtGT applies the GT predicate on the "recovery_sent_at" field.
func RecoverySentAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecoverySentAt), v))
	})
}

// RecoverySentAtGTE applies the GTE predicate on the "recovery_sent_at" field.
func RecoverySentAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecoverySentAt), v))
	})
}

// RecoverySentAtLT applies the LT predicate on the "recovery_sent_at" field.
func RecoverySentAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecoverySentAt), v))
	})
}

// RecoverySentAtLTE applies the LTE predicate on the "recovery_sent_at" field.
func RecoverySentAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecoverySentAt), v))
	})
}

// RecoverySentAtIsNil applies the IsNil predicate on the "recovery_sent_at" field.
func RecoverySentAtIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecoverySentAt)))
	})
}

// RecoverySentAtNotNil applies the NotNil predicate on the "recovery_sent_at" field.
func RecoverySentAtNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecoverySentAt)))
	})
}

// RecoveryTokenEQ applies the EQ predicate on the "recovery_token" field.
func RecoveryTokenEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenNEQ applies the NEQ predicate on the "recovery_token" field.
func RecoveryTokenNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenIn applies the In predicate on the "recovery_token" field.
func RecoveryTokenIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecoveryToken), v...))
	})
}

// RecoveryTokenNotIn applies the NotIn predicate on the "recovery_token" field.
func RecoveryTokenNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecoveryToken), v...))
	})
}

// RecoveryTokenGT applies the GT predicate on the "recovery_token" field.
func RecoveryTokenGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenGTE applies the GTE predicate on the "recovery_token" field.
func RecoveryTokenGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenLT applies the LT predicate on the "recovery_token" field.
func RecoveryTokenLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenLTE applies the LTE predicate on the "recovery_token" field.
func RecoveryTokenLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenContains applies the Contains predicate on the "recovery_token" field.
func RecoveryTokenContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenHasPrefix applies the HasPrefix predicate on the "recovery_token" field.
func RecoveryTokenHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenHasSuffix applies the HasSuffix predicate on the "recovery_token" field.
func RecoveryTokenHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenIsNil applies the IsNil predicate on the "recovery_token" field.
func RecoveryTokenIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecoveryToken)))
	})
}

// RecoveryTokenNotNil applies the NotNil predicate on the "recovery_token" field.
func RecoveryTokenNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecoveryToken)))
	})
}

// RecoveryTokenEqualFold applies the EqualFold predicate on the "recovery_token" field.
func RecoveryTokenEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecoveryToken), v))
	})
}

// RecoveryTokenContainsFold applies the ContainsFold predicate on the "recovery_token" field.
func RecoveryTokenContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecoveryToken), v))
	})
}

// OtpSentAtEQ applies the EQ predicate on the "otp_sent_at" field.
func OtpSentAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtpSentAt), v))
	})
}

// OtpSentAtNEQ applies the NEQ predicate on the "otp_sent_at" field.
func OtpSentAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtpSentAt), v))
	})
}

// OtpSentAtIn applies the In predicate on the "otp_sent_at" field.
func OtpSentAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtpSentAt), v...))
	})
}

// OtpSentAtNotIn applies the NotIn predicate on the "otp_sent_at" field.
func OtpSentAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtpSentAt), v...))
	})
}

// OtpSentAtGT applies the GT predicate on the "otp_sent_at" field.
func OtpSentAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtpSentAt), v))
	})
}

// OtpSentAtGTE applies the GTE predicate on the "otp_sent_at" field.
func OtpSentAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtpSentAt), v))
	})
}

// OtpSentAtLT applies the LT predicate on the "otp_sent_at" field.
func OtpSentAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtpSentAt), v))
	})
}

// OtpSentAtLTE applies the LTE predicate on the "otp_sent_at" field.
func OtpSentAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtpSentAt), v))
	})
}

// OtpSentAtIsNil applies the IsNil predicate on the "otp_sent_at" field.
func OtpSentAtIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOtpSentAt)))
	})
}

// OtpSentAtNotNil applies the NotNil predicate on the "otp_sent_at" field.
func OtpSentAtNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOtpSentAt)))
	})
}

// OtpEQ applies the EQ predicate on the "otp" field.
func OtpEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOtp), v))
	})
}

// OtpNEQ applies the NEQ predicate on the "otp" field.
func OtpNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOtp), v))
	})
}

// OtpIn applies the In predicate on the "otp" field.
func OtpIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOtp), v...))
	})
}

// OtpNotIn applies the NotIn predicate on the "otp" field.
func OtpNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOtp), v...))
	})
}

// OtpGT applies the GT predicate on the "otp" field.
func OtpGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOtp), v))
	})
}

// OtpGTE applies the GTE predicate on the "otp" field.
func OtpGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOtp), v))
	})
}

// OtpLT applies the LT predicate on the "otp" field.
func OtpLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOtp), v))
	})
}

// OtpLTE applies the LTE predicate on the "otp" field.
func OtpLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOtp), v))
	})
}

// OtpContains applies the Contains predicate on the "otp" field.
func OtpContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOtp), v))
	})
}

// OtpHasPrefix applies the HasPrefix predicate on the "otp" field.
func OtpHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOtp), v))
	})
}

// OtpHasSuffix applies the HasSuffix predicate on the "otp" field.
func OtpHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOtp), v))
	})
}

// OtpIsNil applies the IsNil predicate on the "otp" field.
func OtpIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOtp)))
	})
}

// OtpNotNil applies the NotNil predicate on the "otp" field.
func OtpNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOtp)))
	})
}

// OtpEqualFold applies the EqualFold predicate on the "otp" field.
func OtpEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOtp), v))
	})
}

// OtpContainsFold applies the ContainsFold predicate on the "otp" field.
func OtpContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOtp), v))
	})
}

// EmailChangeEQ applies the EQ predicate on the "email_change" field.
func EmailChangeEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailChange), v))
	})
}

// EmailChangeNEQ applies the NEQ predicate on the "email_change" field.
func EmailChangeNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmailChange), v))
	})
}

// EmailChangeIn applies the In predicate on the "email_change" field.
func EmailChangeIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmailChange), v...))
	})
}

// EmailChangeNotIn applies the NotIn predicate on the "email_change" field.
func EmailChangeNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmailChange), v...))
	})
}

// EmailChangeGT applies the GT predicate on the "email_change" field.
func EmailChangeGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmailChange), v))
	})
}

// EmailChangeGTE applies the GTE predicate on the "email_change" field.
func EmailChangeGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmailChange), v))
	})
}

// EmailChangeLT applies the LT predicate on the "email_change" field.
func EmailChangeLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmailChange), v))
	})
}

// EmailChangeLTE applies the LTE predicate on the "email_change" field.
func EmailChangeLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmailChange), v))
	})
}

// EmailChangeContains applies the Contains predicate on the "email_change" field.
func EmailChangeContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmailChange), v))
	})
}

// EmailChangeHasPrefix applies the HasPrefix predicate on the "email_change" field.
func EmailChangeHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmailChange), v))
	})
}

// EmailChangeHasSuffix applies the HasSuffix predicate on the "email_change" field.
func EmailChangeHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmailChange), v))
	})
}

// EmailChangeIsNil applies the IsNil predicate on the "email_change" field.
func EmailChangeIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEmailChange)))
	})
}

// EmailChangeNotNil applies the NotNil predicate on the "email_change" field.
func EmailChangeNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEmailChange)))
	})
}

// EmailChangeEqualFold applies the EqualFold predicate on the "email_change" field.
func EmailChangeEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmailChange), v))
	})
}

// EmailChangeContainsFold applies the ContainsFold predicate on the "email_change" field.
func EmailChangeContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmailChange), v))
	})
}

// EmailChangeSentAtEQ applies the EQ predicate on the "email_change_sent_at" field.
func EmailChangeSentAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailChangeSentAt), v))
	})
}

// EmailChangeSentAtNEQ applies the NEQ predicate on the "email_change_sent_at" field.
func EmailChangeSentAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmailChangeSentAt), v))
	})
}

// EmailChangeSentAtIn applies the In predicate on the "email_change_sent_at" field.
func EmailChangeSentAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmailChangeSentAt), v...))
	})
}

// EmailChangeSentAtNotIn applies the NotIn predicate on the "email_change_sent_at" field.
func EmailChangeSentAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmailChangeSentAt), v...))
	})
}

// EmailChangeSentAtGT applies the GT predicate on the "email_change_sent_at" field.
func EmailChangeSentAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmailChangeSentAt), v))
	})
}

// EmailChangeSentAtGTE applies the GTE predicate on the "email_change_sent_at" field.
func EmailChangeSentAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmailChangeSentAt), v))
	})
}

// EmailChangeSentAtLT applies the LT predicate on the "email_change_sent_at" field.
func EmailChangeSentAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmailChangeSentAt), v))
	})
}

// EmailChangeSentAtLTE applies the LTE predicate on the "email_change_sent_at" field.
func EmailChangeSentAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmailChangeSentAt), v))
	})
}

// EmailChangeSentAtIsNil applies the IsNil predicate on the "email_change_sent_at" field.
func EmailChangeSentAtIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEmailChangeSentAt)))
	})
}

// EmailChangeSentAtNotNil applies the NotNil predicate on the "email_change_sent_at" field.
func EmailChangeSentAtNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEmailChangeSentAt)))
	})
}

// EmailChangeTokenEQ applies the EQ predicate on the "email_change_token" field.
func EmailChangeTokenEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenNEQ applies the NEQ predicate on the "email_change_token" field.
func EmailChangeTokenNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenIn applies the In predicate on the "email_change_token" field.
func EmailChangeTokenIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmailChangeToken), v...))
	})
}

// EmailChangeTokenNotIn applies the NotIn predicate on the "email_change_token" field.
func EmailChangeTokenNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmailChangeToken), v...))
	})
}

// EmailChangeTokenGT applies the GT predicate on the "email_change_token" field.
func EmailChangeTokenGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenGTE applies the GTE predicate on the "email_change_token" field.
func EmailChangeTokenGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenLT applies the LT predicate on the "email_change_token" field.
func EmailChangeTokenLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenLTE applies the LTE predicate on the "email_change_token" field.
func EmailChangeTokenLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenContains applies the Contains predicate on the "email_change_token" field.
func EmailChangeTokenContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenHasPrefix applies the HasPrefix predicate on the "email_change_token" field.
func EmailChangeTokenHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenHasSuffix applies the HasSuffix predicate on the "email_change_token" field.
func EmailChangeTokenHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenIsNil applies the IsNil predicate on the "email_change_token" field.
func EmailChangeTokenIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEmailChangeToken)))
	})
}

// EmailChangeTokenNotNil applies the NotNil predicate on the "email_change_token" field.
func EmailChangeTokenNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEmailChangeToken)))
	})
}

// EmailChangeTokenEqualFold applies the EqualFold predicate on the "email_change_token" field.
func EmailChangeTokenEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmailChangeToken), v))
	})
}

// EmailChangeTokenContainsFold applies the ContainsFold predicate on the "email_change_token" field.
func EmailChangeTokenContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmailChangeToken), v))
	})
}

// RolesIsNil applies the IsNil predicate on the "roles" field.
func RolesIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRoles)))
	})
}

// RolesNotNil applies the NotNil predicate on the "roles" field.
func RolesNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRoles)))
	})
}

// TeamsIsNil applies the IsNil predicate on the "teams" field.
func TeamsIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTeams)))
	})
}

// TeamsNotNil applies the NotNil predicate on the "teams" field.
func TeamsNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTeams)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// LastSigninAtEQ applies the EQ predicate on the "last_signin_at" field.
func LastSigninAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSigninAt), v))
	})
}

// LastSigninAtNEQ applies the NEQ predicate on the "last_signin_at" field.
func LastSigninAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSigninAt), v))
	})
}

// LastSigninAtIn applies the In predicate on the "last_signin_at" field.
func LastSigninAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastSigninAt), v...))
	})
}

// LastSigninAtNotIn applies the NotIn predicate on the "last_signin_at" field.
func LastSigninAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastSigninAt), v...))
	})
}

// LastSigninAtGT applies the GT predicate on the "last_signin_at" field.
func LastSigninAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSigninAt), v))
	})
}

// LastSigninAtGTE applies the GTE predicate on the "last_signin_at" field.
func LastSigninAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSigninAt), v))
	})
}

// LastSigninAtLT applies the LT predicate on the "last_signin_at" field.
func LastSigninAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSigninAt), v))
	})
}

// LastSigninAtLTE applies the LTE predicate on the "last_signin_at" field.
func LastSigninAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSigninAt), v))
	})
}

// LastSigninAtIsNil applies the IsNil predicate on the "last_signin_at" field.
func LastSigninAtIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastSigninAt)))
	})
}

// LastSigninAtNotNil applies the NotNil predicate on the "last_signin_at" field.
func LastSigninAtNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastSigninAt)))
	})
}

// HasWorkspace applies the HasEdge predicate on the "workspace" edge.
func HasWorkspace() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, WorkspaceTable, WorkspaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspaceWith applies the HasEdge predicate on the "workspace" edge with a given conditions (other predicates).
func HasWorkspaceWith(preds ...predicate.Workspace) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, WorkspaceTable, WorkspaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkspaceRoles applies the HasEdge predicate on the "workspace_roles" edge.
func HasWorkspaceRoles() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceRolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkspaceRolesTable, WorkspaceRolesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspaceRolesWith applies the HasEdge predicate on the "workspace_roles" edge with a given conditions (other predicates).
func HasWorkspaceRolesWith(preds ...predicate.WorkspaceRole) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceRolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkspaceRolesTable, WorkspaceRolesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroupRoles applies the HasEdge predicate on the "group_roles" edge.
func HasGroupRoles() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupRolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GroupRolesTable, GroupRolesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupRolesWith applies the HasEdge predicate on the "group_roles" edge with a given conditions (other predicates).
func HasGroupRolesWith(preds ...predicate.GroupRole) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupRolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GroupRolesTable, GroupRolesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserRoles applies the HasEdge predicate on the "user_roles" edge.
func HasUserRoles() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserRolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserRolesTable, UserRolesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserRolesWith applies the HasEdge predicate on the "user_roles" edge with a given conditions (other predicates).
func HasUserRolesWith(preds ...predicate.UserRole) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserRolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserRolesTable, UserRolesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
