// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SessionsxTable,
		UsersxTable,
	}
)

func init() {
}
