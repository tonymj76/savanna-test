// Code generated by ent, DO NOT EDIT.

package repository

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the repository type in the database.
	Label = "repository"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldOpenIssuesCount holds the string denoting the open_issues_count field in the database.
	FieldOpenIssuesCount = "open_issues_count"
	// FieldWatchersCount holds the string denoting the watchers_count field in the database.
	FieldWatchersCount = "watchers_count"
	// FieldStarCount holds the string denoting the star_count field in the database.
	FieldStarCount = "star_count"
	// FieldForksCount holds the string denoting the forks_count field in the database.
	FieldForksCount = "forks_count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeGitCommits holds the string denoting the gitcommits edge name in mutations.
	EdgeGitCommits = "gitCommits"
	// Table holds the table name of the repository in the database.
	Table = "repositories"
	// GitCommitsTable is the table that holds the gitCommits relation/edge.
	GitCommitsTable = "git_commits"
	// GitCommitsInverseTable is the table name for the GitCommit entity.
	// It exists in this package in order to avoid circular dependency with the "gitcommit" package.
	GitCommitsInverseTable = "git_commits"
	// GitCommitsColumn is the table column denoting the gitCommits relation/edge.
	GitCommitsColumn = "repository_git_commits"
)

// Columns holds all SQL columns for repository fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldURL,
	FieldLanguage,
	FieldOpenIssuesCount,
	FieldWatchersCount,
	FieldStarCount,
	FieldForksCount,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Repository queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByURL orders the results by the URL field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByOpenIssuesCount orders the results by the open_issues_count field.
func ByOpenIssuesCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOpenIssuesCount, opts...).ToFunc()
}

// ByWatchersCount orders the results by the watchers_count field.
func ByWatchersCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWatchersCount, opts...).ToFunc()
}

// ByStarCount orders the results by the star_count field.
func ByStarCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStarCount, opts...).ToFunc()
}

// ByForksCount orders the results by the forks_count field.
func ByForksCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForksCount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByGitCommitsCount orders the results by gitCommits count.
func ByGitCommitsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGitCommitsStep(), opts...)
	}
}

// ByGitCommits orders the results by gitCommits terms.
func ByGitCommits(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGitCommitsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newGitCommitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GitCommitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GitCommitsTable, GitCommitsColumn),
	)
}
