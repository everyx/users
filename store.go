package users

import (
	"context"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/adnaan/users/internal"

	"github.com/adnaan/users/internal/models"
)

const defaultMaxAge = 60 * 60 * 24 * 30 // 30 days
const defaultPath = "/"

// UserStore represents the data store for the User model
type UserStore interface {
	// Create, Delete
	New(email, password string, meta map[string]interface{}) (string, error)
	GetUser(id string) (map[string]interface{}, error)
	GetUserByEmail(email string) (map[string]interface{}, error)
	GetUserByConfirmationToken(token string) (map[string]interface{}, error)
	GetUserByRecoveryToken(token string) (map[string]interface{}, error)
	GetUserByEmailChangeToken(token string) (map[string]interface{}, error)
	DeleteUser(id string) error

	// Update Password
	UpdatePassword(id, password string) error
	// Update Email
	UpdateEmail(id, email string) error

	// Confirm Email
	SaveConfirmationToken(id, token string) error
	SaveConfirmationTokenSentAt(id string, tokenSentAt time.Time) error
	MarkConfirmed(id string, confirmed bool) error
	DeleteConfirmToken(id string) error

	// Recover Password
	SaveRecoveryToken(id, token string) error
	SaveRecoveryTokenSentAt(id string, tokenSentAt time.Time) error
	DeleteRecoveryToken(id string) error

	// Change Email
	SaveEmailChangeToken(id, email, token string) error
	SaveEmailChangeTokenSentAt(id string, tokenSentAt time.Time) error
	DeleteEmailChangeToken(id string) error

	// Metadata
	UpsertMetaData(id string, metaData map[string]interface{}) error
	DeleteKeysMetaData(id string, keys []string) error
	DeleteAllMetadata(id string) error

	// Timestamps
	SetUpdatedAt(id string, time time.Time) error
	SetLastSignInAt(id string, time time.Time) error
}

func NewDefaultUserStore(ctx context.Context, driver, dataSource string) (UserStore, error) {
	client, err := models.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	return &internal.DefaultUserStore{Ctx: ctx, Client: client, Driver: driver, DataSource: dataSource}, nil
}

func NewDefaultSessionStore(ctx context.Context, driver, dataSource string, keyPairs ...[]byte) (sessions.Store, error) {
	client, err := models.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	return &internal.DefaultSessionStore{Ctx: ctx, Client: client, Driver: driver, DataSource: dataSource,
		Codecs: securecookie.CodecsFromPairs(keyPairs...),
		SessionOpts: &sessions.Options{
			Path:   defaultPath,
			MaxAge: defaultMaxAge,
		}}, nil
}
