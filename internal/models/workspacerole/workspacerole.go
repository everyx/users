// Code generated (@generated) by entc, DO NOT EDIT.

package workspacerole

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workspacerole type in the database.
	Label = "workspace_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeWorkspaces holds the string denoting the workspaces edge name in mutations.
	EdgeWorkspaces = "workspaces"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"

	// Table holds the table name of the workspacerole in the database.
	Table = "workspacerolesx"
	// WorkspacesTable is the table the holds the workspaces relation/edge.
	WorkspacesTable = "workspacerolesx"
	// WorkspacesInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspacesInverseTable = "workspacesx"
	// WorkspacesColumn is the table column denoting the workspaces relation/edge.
	WorkspacesColumn = "workspace_workspace_roles"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "workspacerolesx"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "usersx"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_workspace_roles"
)

// Columns holds all SQL columns for workspacerole fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldWorkspaceID,
	FieldUserID,
	FieldMetadata,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the WorkspaceRole type.
var ForeignKeys = []string{
	"user_workspace_roles",
	"workspace_workspace_roles",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
