// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// GroupsxColumns holds the columns for the "groupsx" table.
	GroupsxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "workspace_groups", Type: field.TypeUUID, Nullable: true},
	}
	// GroupsxTable holds the schema information for the "groupsx" table.
	GroupsxTable = &schema.Table{
		Name:       "groupsx",
		Columns:    GroupsxColumns,
		PrimaryKey: []*schema.Column{GroupsxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "groupsx_workspacesx_groups",
				Columns: []*schema.Column{GroupsxColumns[6]},

				RefColumns: []*schema.Column{WorkspacesxColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "groupsx"},
	}
	// GrouprolesxColumns holds the columns for the "grouprolesx" table.
	GrouprolesxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "group_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "group_group_roles", Type: field.TypeUUID, Nullable: true},
		{Name: "user_group_roles", Type: field.TypeUUID, Nullable: true},
	}
	// GrouprolesxTable holds the schema information for the "grouprolesx" table.
	GrouprolesxTable = &schema.Table{
		Name:       "grouprolesx",
		Columns:    GrouprolesxColumns,
		PrimaryKey: []*schema.Column{GrouprolesxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "grouprolesx_groupsx_group_roles",
				Columns: []*schema.Column{GrouprolesxColumns[7]},

				RefColumns: []*schema.Column{GroupsxColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "grouprolesx_usersx_group_roles",
				Columns: []*schema.Column{GrouprolesxColumns[8]},

				RefColumns: []*schema.Column{UsersxColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "grouprolesx"},
	}
	// PermissionsxColumns holds the columns for the "permissionsx" table.
	PermissionsxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "action", Type: field.TypeString},
		{Name: "target", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PermissionsxTable holds the schema information for the "permissionsx" table.
	PermissionsxTable = &schema.Table{
		Name:        "permissionsx",
		Columns:     PermissionsxColumns,
		PrimaryKey:  []*schema.Column{PermissionsxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "permissionsx"},
	}
	// SessionsxColumns holds the columns for the "sessionsx" table.
	SessionsxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "data", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
	}
	// SessionsxTable holds the schema information for the "sessionsx" table.
	SessionsxTable = &schema.Table{
		Name:        "sessionsx",
		Columns:     SessionsxColumns,
		PrimaryKey:  []*schema.Column{SessionsxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "sessionsx"},
	}
	// UsersxColumns holds the columns for the "usersx" table.
	UsersxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "billing_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "provider", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "api_key", Type: field.TypeString, Nullable: true},
		{Name: "confirmed", Type: field.TypeBool, Nullable: true},
		{Name: "confirmation_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "confirmation_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "recovery_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "recovery_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "otp_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "otp", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email_change", Type: field.TypeString, Nullable: true},
		{Name: "email_change_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "email_change_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "roles", Type: field.TypeJSON, Nullable: true},
		{Name: "teams", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_signin_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersxTable holds the schema information for the "usersx" table.
	UsersxTable = &schema.Table{
		Name:        "usersx",
		Columns:     UsersxColumns,
		PrimaryKey:  []*schema.Column{UsersxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "usersx"},
	}
	// UserrolesxColumns holds the columns for the "userrolesx" table.
	UserrolesxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_user_roles", Type: field.TypeUUID, Nullable: true},
	}
	// UserrolesxTable holds the schema information for the "userrolesx" table.
	UserrolesxTable = &schema.Table{
		Name:       "userrolesx",
		Columns:    UserrolesxColumns,
		PrimaryKey: []*schema.Column{UserrolesxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "userrolesx_usersx_user_roles",
				Columns: []*schema.Column{UserrolesxColumns[6]},

				RefColumns: []*schema.Column{UsersxColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "userrolesx"},
	}
	// WorkspacesxColumns holds the columns for the "workspacesx" table.
	WorkspacesxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "plan", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_workspace", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// WorkspacesxTable holds the schema information for the "workspacesx" table.
	WorkspacesxTable = &schema.Table{
		Name:       "workspacesx",
		Columns:    WorkspacesxColumns,
		PrimaryKey: []*schema.Column{WorkspacesxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workspacesx_usersx_workspace",
				Columns: []*schema.Column{WorkspacesxColumns[7]},

				RefColumns: []*schema.Column{UsersxColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "workspacesx"},
	}
	// WorkspacerolesxColumns holds the columns for the "workspacerolesx" table.
	WorkspacerolesxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "workspace_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_workspace_roles", Type: field.TypeUUID, Nullable: true},
		{Name: "workspace_workspace_roles", Type: field.TypeUUID, Nullable: true},
	}
	// WorkspacerolesxTable holds the schema information for the "workspacerolesx" table.
	WorkspacerolesxTable = &schema.Table{
		Name:       "workspacerolesx",
		Columns:    WorkspacerolesxColumns,
		PrimaryKey: []*schema.Column{WorkspacerolesxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workspacerolesx_usersx_workspace_roles",
				Columns: []*schema.Column{WorkspacerolesxColumns[7]},

				RefColumns: []*schema.Column{UsersxColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "workspacerolesx_workspacesx_workspace_roles",
				Columns: []*schema.Column{WorkspacerolesxColumns[8]},

				RefColumns: []*schema.Column{WorkspacesxColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "workspacerolesx"},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsxTable,
		GrouprolesxTable,
		PermissionsxTable,
		SessionsxTable,
		UsersxTable,
		UserrolesxTable,
		WorkspacesxTable,
		WorkspacerolesxTable,
	}
)

func init() {
	GroupsxTable.ForeignKeys[0].RefTable = WorkspacesxTable
	GrouprolesxTable.ForeignKeys[0].RefTable = GroupsxTable
	GrouprolesxTable.ForeignKeys[1].RefTable = UsersxTable
	UserrolesxTable.ForeignKeys[0].RefTable = UsersxTable
	WorkspacesxTable.ForeignKeys[0].RefTable = UsersxTable
	WorkspacerolesxTable.ForeignKeys[0].RefTable = UsersxTable
	WorkspacerolesxTable.ForeignKeys[1].RefTable = WorkspacesxTable
}
