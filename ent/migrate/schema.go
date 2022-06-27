// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TenantsColumns holds the columns for the "tenants" table.
	TenantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// TenantsTable holds the schema information for the "tenants" table.
	TenantsTable = &schema.Table{
		Name:       "tenants",
		Columns:    TenantsColumns,
		PrimaryKey: []*schema.Column{TenantsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_tenants_users",
				Columns:    []*schema.Column{UsersColumns[3]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TenantsTable,
		UsersTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = TenantsTable
}
