// Code generated by ent, DO NOT EDIT.

package gitcommit

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the gitcommit type in the database.
	Label = "git_commit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGitcommit holds the string denoting the gitcommit field in the database.
	FieldGitcommit = "gitcommit"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// Table holds the table name of the gitcommit in the database.
	Table = "git_commits"
)

// Columns holds all SQL columns for gitcommit fields.
var Columns = []string{
	FieldID,
	FieldGitcommit,
	FieldURL,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "git_commits"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"repository_git_commits",
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

// OrderOption defines the ordering options for the GitCommit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}
