package users

import (
	"context"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/adnaan/users/internal"

	"github.com/adnaan/users/internal/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const defaultMaxAge = 60 * 60 * 24 * 30 // 30 days
const defaultPath = "/"

// UserStore represents the data store for the User model
type UserStore interface {
	// Create, Delete
	New(email, password, role, provider string, meta map[string]interface{}, sendMailFunc func(string, string) error) (string, error)
	UserData(id string) (string, string, string, map[string]interface{}, error)
	UserIDByEmail(email string) (string, error)
	UserIDByConfirmationToken(token string) (string, error)
	UserIDByRecoveryToken(token string) (string, error)
	UserIDByEmailChangeToken(token string) (string, error)
	UserIDByOTP(token string) (string, error)
	UserIDByAPIKey(apiKey string) (string, error)
	DeleteUser(id string) error

	GetPassword(id string) (string, error)
	GetAPIKey(id string) (string, error)
	GetEmailChange(id string) (string, error)
	IsEmailConfirmed(id string) (bool, error)
	GetRoles(id string) ([]string, error)

	AddRole(id, role string) error
	DeleteRole(id, role string) error
	ClearRoles(id string) error
	// Update Password
	UpdatePassword(id, password string) error
	// Update API Key
	UpdateAPIKey(id, apiKey string) error
	// Update Email
	UpdateEmail(id, email string) error
	UpdateBillingID(id, billingID string) error
	UpdateProvider(id, provider string) error

	// Confirm Email
	SaveConfirmationToken(id, token string) error
	SaveConfirmationTokenSentAt(id string, tokenSentAt time.Time) error
	MarkConfirmed(id string, confirmed bool) error
	DeleteConfirmToken(id string) error

	// One time password
	SaveOTP(id, otp string) error
	SaveOTPSentAt(id string, otpSentAt time.Time) error
	DeleteOTP(id string) error

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
	Close() error
}

type SessionsStore interface {
	sessions.Store
	Close() error
}

type WorkspaceStore interface {
	GetWorkspace(id string) (string, string, map[string]interface{}, error)
	CreateWorkspace(userID, name, description, plan string, metadata map[string]interface{}) (string, error)
	UpdateWorkspace(workspaceID, name, description, plan string, metadata map[string]interface{}) error
	DeleteWorkspace(workspaceID string) error
	GetUserWorkspaces(userID string) (map[string][]string, error)

	WorkspaceUpsertUser(workspaceID, userID, role string) error
	WorkspaceRemoveUser(workspaceID, userID string) error
	GetWorkspaceGroups(workspaceID string) ([]string, error)

	GetGroup(id string) (string, string, map[string]interface{}, error)
	CreateGroup(userID, workspaceID, name, description string, metadata map[string]interface{}) (string, error)
	UpdateGroup(groupID, name, description string, metadata map[string]interface{}) error
	DeleteGroup(groupID string) error
	GetUserGroupRoles(userID string) (map[string]string, error)
	GroupUpsertUser(groupID, userID, role string) error
	GroupRemoveUser(groupID, userID string) error

	CreatePermission(roleID, action, target string) error
	UpdatePermission(roleID, action, target string) error
	DeletePermission(roleID, action string) error
	GetUserPermissions(userID string) (map[string]string, error)

	GetUserRoles(userID string) ([]string, error)
	CreateUserRole(userID, role string) error
	DeleteUserRole(userID, role string) error
}

func NewDefaultUserStore(ctx context.Context, driver, dataSource string) (UserStore, error) {
	client, err := models.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}
	return &internal.DefaultUserStore{Ctx: ctx, Client: client, Driver: driver, DataSource: dataSource}, nil
}

func NewDefaultWorkspaceStore(ctx context.Context, driver, dataSource string) (WorkspaceStore, error) {
	client, err := models.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}
	return &internal.DefaultWorkspaceStore{Ctx: ctx, Client: client, Driver: driver, DataSource: dataSource}, nil
}

func NewDefaultSessionStore(ctx context.Context, driver, dataSource string, keyPairs ...[]byte) (SessionsStore, error) {
	client, err := models.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}
	return &internal.DefaultSessionStore{Ctx: ctx, Client: client, Driver: driver, DataSource: dataSource,
		Codecs: securecookie.CodecsFromPairs(keyPairs...),
		SessionOpts: &sessions.Options{
			Path:   defaultPath,
			MaxAge: defaultMaxAge,
		}}, nil
}
