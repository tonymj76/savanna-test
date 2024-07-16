// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommitsColumns holds the columns for the "commits" table.
	CommitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "author", Type: field.TypeJSON},
		{Name: "url", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "repository_commits", Type: field.TypeInt, Nullable: true},
	}
	// CommitsTable holds the schema information for the "commits" table.
	CommitsTable = &schema.Table{
		Name:       "commits",
		Columns:    CommitsColumns,
		PrimaryKey: []*schema.Column{CommitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "commits_repositories_commits",
				Columns:    []*schema.Column{CommitsColumns[4]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RepositoriesColumns holds the columns for the "repositories" table.
	RepositoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "language", Type: field.TypeString},
		{Name: "open_issues_count", Type: field.TypeInt},
		{Name: "watchers_count", Type: field.TypeInt},
		{Name: "star_count", Type: field.TypeInt},
		{Name: "forks_count", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// RepositoriesTable holds the schema information for the "repositories" table.
	RepositoriesTable = &schema.Table{
		Name:       "repositories",
		Columns:    RepositoriesColumns,
		PrimaryKey: []*schema.Column{RepositoriesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommitsTable,
		RepositoriesTable,
	}
)

func init() {
	CommitsTable.ForeignKeys[0].RefTable = RepositoriesTable
}
