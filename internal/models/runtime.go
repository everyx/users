// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"time"

	"github.com/adnaan/users/internal/models/schema"
	"github.com/adnaan/users/internal/models/session"
	"github.com/adnaan/users/internal/models/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionFields[2].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionFields[3].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescProvider is the schema descriptor for provider field.
	userDescProvider := userFields[1].Descriptor()
	// user.ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	user.ProviderValidator = userDescProvider.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAPIKey is the schema descriptor for api_key field.
	userDescAPIKey := userFields[4].Descriptor()
	// user.APIKeyValidator is a validator for the "api_key" field. It is called by the builders before save.
	user.APIKeyValidator = userDescAPIKey.Validators[0].(func(string) error)
	// userDescConfirmed is the schema descriptor for confirmed field.
	userDescConfirmed := userFields[5].Descriptor()
	// user.DefaultConfirmed holds the default value on creation for the confirmed field.
	user.DefaultConfirmed = userDescConfirmed.Default.(bool)
	// userDescConfirmationToken is the schema descriptor for confirmation_token field.
	userDescConfirmationToken := userFields[7].Descriptor()
	// user.ConfirmationTokenValidator is a validator for the "confirmation_token" field. It is called by the builders before save.
	user.ConfirmationTokenValidator = userDescConfirmationToken.Validators[0].(func(string) error)
	// userDescRecoveryToken is the schema descriptor for recovery_token field.
	userDescRecoveryToken := userFields[9].Descriptor()
	// user.RecoveryTokenValidator is a validator for the "recovery_token" field. It is called by the builders before save.
	user.RecoveryTokenValidator = userDescRecoveryToken.Validators[0].(func(string) error)
	// userDescOtp is the schema descriptor for otp field.
	userDescOtp := userFields[11].Descriptor()
	// user.OtpValidator is a validator for the "otp" field. It is called by the builders before save.
	user.OtpValidator = userDescOtp.Validators[0].(func(string) error)
	// userDescEmailChange is the schema descriptor for email_change field.
	userDescEmailChange := userFields[12].Descriptor()
	// user.EmailChangeValidator is a validator for the "email_change" field. It is called by the builders before save.
	user.EmailChangeValidator = userDescEmailChange.Validators[0].(func(string) error)
	// userDescEmailChangeToken is the schema descriptor for email_change_token field.
	userDescEmailChangeToken := userFields[14].Descriptor()
	// user.EmailChangeTokenValidator is a validator for the "email_change_token" field. It is called by the builders before save.
	user.EmailChangeTokenValidator = userDescEmailChangeToken.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[16].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[17].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
