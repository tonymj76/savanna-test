// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tonymj76/savannah/ent/gitcommit"
	"github.com/tonymj76/savannah/ent/predicate"
	"github.com/tonymj76/savannah/ent/repository"
)

// RepositoryUpdate is the builder for updating Repository entities.
type RepositoryUpdate struct {
	config
	hooks    []Hook
	mutation *RepositoryMutation
}

// Where appends a list predicates to the RepositoryUpdate builder.
func (ru *RepositoryUpdate) Where(ps ...predicate.Repository) *RepositoryUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RepositoryUpdate) SetName(s string) *RepositoryUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableName(s *string) *RepositoryUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetDescription sets the "description" field.
func (ru *RepositoryUpdate) SetDescription(s string) *RepositoryUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableDescription(s *string) *RepositoryUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// SetURL sets the "URL" field.
func (ru *RepositoryUpdate) SetURL(s string) *RepositoryUpdate {
	ru.mutation.SetURL(s)
	return ru
}

// SetNillableURL sets the "URL" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableURL(s *string) *RepositoryUpdate {
	if s != nil {
		ru.SetURL(*s)
	}
	return ru
}

// SetLanguage sets the "language" field.
func (ru *RepositoryUpdate) SetLanguage(s string) *RepositoryUpdate {
	ru.mutation.SetLanguage(s)
	return ru
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableLanguage(s *string) *RepositoryUpdate {
	if s != nil {
		ru.SetLanguage(*s)
	}
	return ru
}

// SetOpenIssuesCount sets the "open_issues_count" field.
func (ru *RepositoryUpdate) SetOpenIssuesCount(i int) *RepositoryUpdate {
	ru.mutation.ResetOpenIssuesCount()
	ru.mutation.SetOpenIssuesCount(i)
	return ru
}

// SetNillableOpenIssuesCount sets the "open_issues_count" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableOpenIssuesCount(i *int) *RepositoryUpdate {
	if i != nil {
		ru.SetOpenIssuesCount(*i)
	}
	return ru
}

// AddOpenIssuesCount adds i to the "open_issues_count" field.
func (ru *RepositoryUpdate) AddOpenIssuesCount(i int) *RepositoryUpdate {
	ru.mutation.AddOpenIssuesCount(i)
	return ru
}

// SetWatchersCount sets the "watchers_count" field.
func (ru *RepositoryUpdate) SetWatchersCount(i int) *RepositoryUpdate {
	ru.mutation.ResetWatchersCount()
	ru.mutation.SetWatchersCount(i)
	return ru
}

// SetNillableWatchersCount sets the "watchers_count" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableWatchersCount(i *int) *RepositoryUpdate {
	if i != nil {
		ru.SetWatchersCount(*i)
	}
	return ru
}

// AddWatchersCount adds i to the "watchers_count" field.
func (ru *RepositoryUpdate) AddWatchersCount(i int) *RepositoryUpdate {
	ru.mutation.AddWatchersCount(i)
	return ru
}

// SetStarCount sets the "star_count" field.
func (ru *RepositoryUpdate) SetStarCount(i int) *RepositoryUpdate {
	ru.mutation.ResetStarCount()
	ru.mutation.SetStarCount(i)
	return ru
}

// SetNillableStarCount sets the "star_count" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableStarCount(i *int) *RepositoryUpdate {
	if i != nil {
		ru.SetStarCount(*i)
	}
	return ru
}

// AddStarCount adds i to the "star_count" field.
func (ru *RepositoryUpdate) AddStarCount(i int) *RepositoryUpdate {
	ru.mutation.AddStarCount(i)
	return ru
}

// SetForksCount sets the "forks_count" field.
func (ru *RepositoryUpdate) SetForksCount(i int) *RepositoryUpdate {
	ru.mutation.ResetForksCount()
	ru.mutation.SetForksCount(i)
	return ru
}

// SetNillableForksCount sets the "forks_count" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableForksCount(i *int) *RepositoryUpdate {
	if i != nil {
		ru.SetForksCount(*i)
	}
	return ru
}

// AddForksCount adds i to the "forks_count" field.
func (ru *RepositoryUpdate) AddForksCount(i int) *RepositoryUpdate {
	ru.mutation.AddForksCount(i)
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *RepositoryUpdate) SetCreatedAt(t time.Time) *RepositoryUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableCreatedAt(t *time.Time) *RepositoryUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RepositoryUpdate) SetUpdatedAt(t time.Time) *RepositoryUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableUpdatedAt(t *time.Time) *RepositoryUpdate {
	if t != nil {
		ru.SetUpdatedAt(*t)
	}
	return ru
}

// AddGitCommitIDs adds the "gitCommits" edge to the GitCommit entity by IDs.
func (ru *RepositoryUpdate) AddGitCommitIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.AddGitCommitIDs(ids...)
	return ru
}

// AddGitCommits adds the "gitCommits" edges to the GitCommit entity.
func (ru *RepositoryUpdate) AddGitCommits(g ...*GitCommit) *RepositoryUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ru.AddGitCommitIDs(ids...)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ru *RepositoryUpdate) Mutation() *RepositoryMutation {
	return ru.mutation
}

// ClearGitCommits clears all "gitCommits" edges to the GitCommit entity.
func (ru *RepositoryUpdate) ClearGitCommits() *RepositoryUpdate {
	ru.mutation.ClearGitCommits()
	return ru
}

// RemoveGitCommitIDs removes the "gitCommits" edge to GitCommit entities by IDs.
func (ru *RepositoryUpdate) RemoveGitCommitIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.RemoveGitCommitIDs(ids...)
	return ru
}

// RemoveGitCommits removes "gitCommits" edges to GitCommit entities.
func (ru *RepositoryUpdate) RemoveGitCommits(g ...*GitCommit) *RepositoryUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ru.RemoveGitCommitIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RepositoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepositoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepositoryUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepositoryUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RepositoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(repository.Table, repository.Columns, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(repository.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(repository.FieldDescription, field.TypeString, value)
	}
	if value, ok := ru.mutation.URL(); ok {
		_spec.SetField(repository.FieldURL, field.TypeString, value)
	}
	if value, ok := ru.mutation.Language(); ok {
		_spec.SetField(repository.FieldLanguage, field.TypeString, value)
	}
	if value, ok := ru.mutation.OpenIssuesCount(); ok {
		_spec.SetField(repository.FieldOpenIssuesCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedOpenIssuesCount(); ok {
		_spec.AddField(repository.FieldOpenIssuesCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.WatchersCount(); ok {
		_spec.SetField(repository.FieldWatchersCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedWatchersCount(); ok {
		_spec.AddField(repository.FieldWatchersCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.StarCount(); ok {
		_spec.SetField(repository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedStarCount(); ok {
		_spec.AddField(repository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.ForksCount(); ok {
		_spec.SetField(repository.FieldForksCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedForksCount(); ok {
		_spec.AddField(repository.FieldForksCount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(repository.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(repository.FieldUpdatedAt, field.TypeTime, value)
	}
	if ru.mutation.GitCommitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.GitCommitsTable,
			Columns: []string{repository.GitCommitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedGitCommitsIDs(); len(nodes) > 0 && !ru.mutation.GitCommitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.GitCommitsTable,
			Columns: []string{repository.GitCommitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.GitCommitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.GitCommitsTable,
			Columns: []string{repository.GitCommitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RepositoryUpdateOne is the builder for updating a single Repository entity.
type RepositoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepositoryMutation
}

// SetName sets the "name" field.
func (ruo *RepositoryUpdateOne) SetName(s string) *RepositoryUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableName(s *string) *RepositoryUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RepositoryUpdateOne) SetDescription(s string) *RepositoryUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableDescription(s *string) *RepositoryUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// SetURL sets the "URL" field.
func (ruo *RepositoryUpdateOne) SetURL(s string) *RepositoryUpdateOne {
	ruo.mutation.SetURL(s)
	return ruo
}

// SetNillableURL sets the "URL" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableURL(s *string) *RepositoryUpdateOne {
	if s != nil {
		ruo.SetURL(*s)
	}
	return ruo
}

// SetLanguage sets the "language" field.
func (ruo *RepositoryUpdateOne) SetLanguage(s string) *RepositoryUpdateOne {
	ruo.mutation.SetLanguage(s)
	return ruo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableLanguage(s *string) *RepositoryUpdateOne {
	if s != nil {
		ruo.SetLanguage(*s)
	}
	return ruo
}

// SetOpenIssuesCount sets the "open_issues_count" field.
func (ruo *RepositoryUpdateOne) SetOpenIssuesCount(i int) *RepositoryUpdateOne {
	ruo.mutation.ResetOpenIssuesCount()
	ruo.mutation.SetOpenIssuesCount(i)
	return ruo
}

// SetNillableOpenIssuesCount sets the "open_issues_count" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableOpenIssuesCount(i *int) *RepositoryUpdateOne {
	if i != nil {
		ruo.SetOpenIssuesCount(*i)
	}
	return ruo
}

// AddOpenIssuesCount adds i to the "open_issues_count" field.
func (ruo *RepositoryUpdateOne) AddOpenIssuesCount(i int) *RepositoryUpdateOne {
	ruo.mutation.AddOpenIssuesCount(i)
	return ruo
}

// SetWatchersCount sets the "watchers_count" field.
func (ruo *RepositoryUpdateOne) SetWatchersCount(i int) *RepositoryUpdateOne {
	ruo.mutation.ResetWatchersCount()
	ruo.mutation.SetWatchersCount(i)
	return ruo
}

// SetNillableWatchersCount sets the "watchers_count" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableWatchersCount(i *int) *RepositoryUpdateOne {
	if i != nil {
		ruo.SetWatchersCount(*i)
	}
	return ruo
}

// AddWatchersCount adds i to the "watchers_count" field.
func (ruo *RepositoryUpdateOne) AddWatchersCount(i int) *RepositoryUpdateOne {
	ruo.mutation.AddWatchersCount(i)
	return ruo
}

// SetStarCount sets the "star_count" field.
func (ruo *RepositoryUpdateOne) SetStarCount(i int) *RepositoryUpdateOne {
	ruo.mutation.ResetStarCount()
	ruo.mutation.SetStarCount(i)
	return ruo
}

// SetNillableStarCount sets the "star_count" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableStarCount(i *int) *RepositoryUpdateOne {
	if i != nil {
		ruo.SetStarCount(*i)
	}
	return ruo
}

// AddStarCount adds i to the "star_count" field.
func (ruo *RepositoryUpdateOne) AddStarCount(i int) *RepositoryUpdateOne {
	ruo.mutation.AddStarCount(i)
	return ruo
}

// SetForksCount sets the "forks_count" field.
func (ruo *RepositoryUpdateOne) SetForksCount(i int) *RepositoryUpdateOne {
	ruo.mutation.ResetForksCount()
	ruo.mutation.SetForksCount(i)
	return ruo
}

// SetNillableForksCount sets the "forks_count" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableForksCount(i *int) *RepositoryUpdateOne {
	if i != nil {
		ruo.SetForksCount(*i)
	}
	return ruo
}

// AddForksCount adds i to the "forks_count" field.
func (ruo *RepositoryUpdateOne) AddForksCount(i int) *RepositoryUpdateOne {
	ruo.mutation.AddForksCount(i)
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *RepositoryUpdateOne) SetCreatedAt(t time.Time) *RepositoryUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableCreatedAt(t *time.Time) *RepositoryUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RepositoryUpdateOne) SetUpdatedAt(t time.Time) *RepositoryUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *RepositoryUpdateOne {
	if t != nil {
		ruo.SetUpdatedAt(*t)
	}
	return ruo
}

// AddGitCommitIDs adds the "gitCommits" edge to the GitCommit entity by IDs.
func (ruo *RepositoryUpdateOne) AddGitCommitIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.AddGitCommitIDs(ids...)
	return ruo
}

// AddGitCommits adds the "gitCommits" edges to the GitCommit entity.
func (ruo *RepositoryUpdateOne) AddGitCommits(g ...*GitCommit) *RepositoryUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ruo.AddGitCommitIDs(ids...)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ruo *RepositoryUpdateOne) Mutation() *RepositoryMutation {
	return ruo.mutation
}

// ClearGitCommits clears all "gitCommits" edges to the GitCommit entity.
func (ruo *RepositoryUpdateOne) ClearGitCommits() *RepositoryUpdateOne {
	ruo.mutation.ClearGitCommits()
	return ruo
}

// RemoveGitCommitIDs removes the "gitCommits" edge to GitCommit entities by IDs.
func (ruo *RepositoryUpdateOne) RemoveGitCommitIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.RemoveGitCommitIDs(ids...)
	return ruo
}

// RemoveGitCommits removes "gitCommits" edges to GitCommit entities.
func (ruo *RepositoryUpdateOne) RemoveGitCommits(g ...*GitCommit) *RepositoryUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ruo.RemoveGitCommitIDs(ids...)
}

// Where appends a list predicates to the RepositoryUpdate builder.
func (ruo *RepositoryUpdateOne) Where(ps ...predicate.Repository) *RepositoryUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RepositoryUpdateOne) Select(field string, fields ...string) *RepositoryUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Repository entity.
func (ruo *RepositoryUpdateOne) Save(ctx context.Context) (*Repository, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) SaveX(ctx context.Context) *Repository {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RepositoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RepositoryUpdateOne) sqlSave(ctx context.Context) (_node *Repository, err error) {
	_spec := sqlgraph.NewUpdateSpec(repository.Table, repository.Columns, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Repository.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repository.FieldID)
		for _, f := range fields {
			if !repository.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != repository.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(repository.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(repository.FieldDescription, field.TypeString, value)
	}
	if value, ok := ruo.mutation.URL(); ok {
		_spec.SetField(repository.FieldURL, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Language(); ok {
		_spec.SetField(repository.FieldLanguage, field.TypeString, value)
	}
	if value, ok := ruo.mutation.OpenIssuesCount(); ok {
		_spec.SetField(repository.FieldOpenIssuesCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedOpenIssuesCount(); ok {
		_spec.AddField(repository.FieldOpenIssuesCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.WatchersCount(); ok {
		_spec.SetField(repository.FieldWatchersCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedWatchersCount(); ok {
		_spec.AddField(repository.FieldWatchersCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.StarCount(); ok {
		_spec.SetField(repository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedStarCount(); ok {
		_spec.AddField(repository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.ForksCount(); ok {
		_spec.SetField(repository.FieldForksCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedForksCount(); ok {
		_spec.AddField(repository.FieldForksCount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(repository.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(repository.FieldUpdatedAt, field.TypeTime, value)
	}
	if ruo.mutation.GitCommitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.GitCommitsTable,
			Columns: []string{repository.GitCommitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedGitCommitsIDs(); len(nodes) > 0 && !ruo.mutation.GitCommitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.GitCommitsTable,
			Columns: []string{repository.GitCommitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.GitCommitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.GitCommitsTable,
			Columns: []string{repository.GitCommitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Repository{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
