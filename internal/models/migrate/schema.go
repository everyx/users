// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// SessionsxColumns holds the columns for the "Sessionsx" table.
	SessionsxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "data", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
	}
	// SessionsxTable holds the schema information for the "Sessionsx" table.
	SessionsxTable = &schema.Table{
		Name:        "Sessionsx",
		Columns:     SessionsxColumns,
		PrimaryKey:  []*schema.Column{SessionsxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersxColumns holds the columns for the "Usersx" table.
	UsersxColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "confirmed", Type: field.TypeBool, Nullable: true},
		{Name: "confirmation_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "confirmation_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "recovery_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "recovery_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email_change", Type: field.TypeString, Nullable: true},
		{Name: "email_change_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "email_change_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_signin_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersxTable holds the schema information for the "Usersx" table.
	UsersxTable = &schema.Table{
		Name:        "Usersx",
		Columns:     UsersxColumns,
		PrimaryKey:  []*schema.Column{UsersxColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SessionsxTable,
		UsersxTable,
	}
)

func init() {
}
