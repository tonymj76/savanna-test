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
)

// GitCommitUpdate is the builder for updating GitCommit entities.
type GitCommitUpdate struct {
	config
	hooks    []Hook
	mutation *GitCommitMutation
}

// Where appends a list predicates to the GitCommitUpdate builder.
func (gcu *GitCommitUpdate) Where(ps ...predicate.GitCommit) *GitCommitUpdate {
	gcu.mutation.Where(ps...)
	return gcu
}

// SetGitcommit sets the "gitcommit" field.
func (gcu *GitCommitUpdate) SetGitcommit(a any) *GitCommitUpdate {
	gcu.mutation.SetGitcommit(a)
	return gcu
}

// SetURL sets the "url" field.
func (gcu *GitCommitUpdate) SetURL(s string) *GitCommitUpdate {
	gcu.mutation.SetURL(s)
	return gcu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (gcu *GitCommitUpdate) SetNillableURL(s *string) *GitCommitUpdate {
	if s != nil {
		gcu.SetURL(*s)
	}
	return gcu
}

// SetDate sets the "date" field.
func (gcu *GitCommitUpdate) SetDate(t time.Time) *GitCommitUpdate {
	gcu.mutation.SetDate(t)
	return gcu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (gcu *GitCommitUpdate) SetNillableDate(t *time.Time) *GitCommitUpdate {
	if t != nil {
		gcu.SetDate(*t)
	}
	return gcu
}

// Mutation returns the GitCommitMutation object of the builder.
func (gcu *GitCommitUpdate) Mutation() *GitCommitMutation {
	return gcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gcu *GitCommitUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gcu.sqlSave, gcu.mutation, gcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gcu *GitCommitUpdate) SaveX(ctx context.Context) int {
	affected, err := gcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gcu *GitCommitUpdate) Exec(ctx context.Context) error {
	_, err := gcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcu *GitCommitUpdate) ExecX(ctx context.Context) {
	if err := gcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gcu *GitCommitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(gitcommit.Table, gitcommit.Columns, sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt))
	if ps := gcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcu.mutation.Gitcommit(); ok {
		_spec.SetField(gitcommit.FieldGitcommit, field.TypeJSON, value)
	}
	if value, ok := gcu.mutation.URL(); ok {
		_spec.SetField(gitcommit.FieldURL, field.TypeString, value)
	}
	if value, ok := gcu.mutation.Date(); ok {
		_spec.SetField(gitcommit.FieldDate, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gitcommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gcu.mutation.done = true
	return n, nil
}

// GitCommitUpdateOne is the builder for updating a single GitCommit entity.
type GitCommitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GitCommitMutation
}

// SetGitcommit sets the "gitcommit" field.
func (gcuo *GitCommitUpdateOne) SetGitcommit(a any) *GitCommitUpdateOne {
	gcuo.mutation.SetGitcommit(a)
	return gcuo
}

// SetURL sets the "url" field.
func (gcuo *GitCommitUpdateOne) SetURL(s string) *GitCommitUpdateOne {
	gcuo.mutation.SetURL(s)
	return gcuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (gcuo *GitCommitUpdateOne) SetNillableURL(s *string) *GitCommitUpdateOne {
	if s != nil {
		gcuo.SetURL(*s)
	}
	return gcuo
}

// SetDate sets the "date" field.
func (gcuo *GitCommitUpdateOne) SetDate(t time.Time) *GitCommitUpdateOne {
	gcuo.mutation.SetDate(t)
	return gcuo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (gcuo *GitCommitUpdateOne) SetNillableDate(t *time.Time) *GitCommitUpdateOne {
	if t != nil {
		gcuo.SetDate(*t)
	}
	return gcuo
}

// Mutation returns the GitCommitMutation object of the builder.
func (gcuo *GitCommitUpdateOne) Mutation() *GitCommitMutation {
	return gcuo.mutation
}

// Where appends a list predicates to the GitCommitUpdate builder.
func (gcuo *GitCommitUpdateOne) Where(ps ...predicate.GitCommit) *GitCommitUpdateOne {
	gcuo.mutation.Where(ps...)
	return gcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gcuo *GitCommitUpdateOne) Select(field string, fields ...string) *GitCommitUpdateOne {
	gcuo.fields = append([]string{field}, fields...)
	return gcuo
}

// Save executes the query and returns the updated GitCommit entity.
func (gcuo *GitCommitUpdateOne) Save(ctx context.Context) (*GitCommit, error) {
	return withHooks(ctx, gcuo.sqlSave, gcuo.mutation, gcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gcuo *GitCommitUpdateOne) SaveX(ctx context.Context) *GitCommit {
	node, err := gcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gcuo *GitCommitUpdateOne) Exec(ctx context.Context) error {
	_, err := gcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcuo *GitCommitUpdateOne) ExecX(ctx context.Context) {
	if err := gcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gcuo *GitCommitUpdateOne) sqlSave(ctx context.Context) (_node *GitCommit, err error) {
	_spec := sqlgraph.NewUpdateSpec(gitcommit.Table, gitcommit.Columns, sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt))
	id, ok := gcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GitCommit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gitcommit.FieldID)
		for _, f := range fields {
			if !gitcommit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gitcommit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcuo.mutation.Gitcommit(); ok {
		_spec.SetField(gitcommit.FieldGitcommit, field.TypeJSON, value)
	}
	if value, ok := gcuo.mutation.URL(); ok {
		_spec.SetField(gitcommit.FieldURL, field.TypeString, value)
	}
	if value, ok := gcuo.mutation.Date(); ok {
		_spec.SetField(gitcommit.FieldDate, field.TypeTime, value)
	}
	_node = &GitCommit{config: gcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gitcommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gcuo.mutation.done = true
	return _node, nil
}
