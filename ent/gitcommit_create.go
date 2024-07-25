// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tonymj76/savannah/ent/gitcommit"
)

// GitCommitCreate is the builder for creating a GitCommit entity.
type GitCommitCreate struct {
	config
	mutation *GitCommitMutation
	hooks    []Hook
}

// SetGitcommit sets the "gitcommit" field.
func (gcc *GitCommitCreate) SetGitcommit(a any) *GitCommitCreate {
	gcc.mutation.SetGitcommit(a)
	return gcc
}

// SetURL sets the "url" field.
func (gcc *GitCommitCreate) SetURL(s string) *GitCommitCreate {
	gcc.mutation.SetURL(s)
	return gcc
}

// SetDate sets the "date" field.
func (gcc *GitCommitCreate) SetDate(t time.Time) *GitCommitCreate {
	gcc.mutation.SetDate(t)
	return gcc
}

// Mutation returns the GitCommitMutation object of the builder.
func (gcc *GitCommitCreate) Mutation() *GitCommitMutation {
	return gcc.mutation
}

// Save creates the GitCommit in the database.
func (gcc *GitCommitCreate) Save(ctx context.Context) (*GitCommit, error) {
	return withHooks(ctx, gcc.sqlSave, gcc.mutation, gcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gcc *GitCommitCreate) SaveX(ctx context.Context) *GitCommit {
	v, err := gcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcc *GitCommitCreate) Exec(ctx context.Context) error {
	_, err := gcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcc *GitCommitCreate) ExecX(ctx context.Context) {
	if err := gcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gcc *GitCommitCreate) check() error {
	if _, ok := gcc.mutation.Gitcommit(); !ok {
		return &ValidationError{Name: "gitcommit", err: errors.New(`ent: missing required field "GitCommit.gitcommit"`)}
	}
	if _, ok := gcc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "GitCommit.url"`)}
	}
	if _, ok := gcc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "GitCommit.date"`)}
	}
	return nil
}

func (gcc *GitCommitCreate) sqlSave(ctx context.Context) (*GitCommit, error) {
	if err := gcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gcc.mutation.id = &_node.ID
	gcc.mutation.done = true
	return _node, nil
}

func (gcc *GitCommitCreate) createSpec() (*GitCommit, *sqlgraph.CreateSpec) {
	var (
		_node = &GitCommit{config: gcc.config}
		_spec = sqlgraph.NewCreateSpec(gitcommit.Table, sqlgraph.NewFieldSpec(gitcommit.FieldID, field.TypeInt))
	)
	if value, ok := gcc.mutation.Gitcommit(); ok {
		_spec.SetField(gitcommit.FieldGitcommit, field.TypeJSON, value)
		_node.Gitcommit = value
	}
	if value, ok := gcc.mutation.URL(); ok {
		_spec.SetField(gitcommit.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := gcc.mutation.Date(); ok {
		_spec.SetField(gitcommit.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	return _node, _spec
}

// GitCommitCreateBulk is the builder for creating many GitCommit entities in bulk.
type GitCommitCreateBulk struct {
	config
	err      error
	builders []*GitCommitCreate
}

// Save creates the GitCommit entities in the database.
func (gccb *GitCommitCreateBulk) Save(ctx context.Context) ([]*GitCommit, error) {
	if gccb.err != nil {
		return nil, gccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gccb.builders))
	nodes := make([]*GitCommit, len(gccb.builders))
	mutators := make([]Mutator, len(gccb.builders))
	for i := range gccb.builders {
		func(i int, root context.Context) {
			builder := gccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GitCommitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gccb *GitCommitCreateBulk) SaveX(ctx context.Context) []*GitCommit {
	v, err := gccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gccb *GitCommitCreateBulk) Exec(ctx context.Context) error {
	_, err := gccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gccb *GitCommitCreateBulk) ExecX(ctx context.Context) {
	if err := gccb.Exec(ctx); err != nil {
		panic(err)
	}
}
