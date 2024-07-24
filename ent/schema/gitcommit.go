package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GitCommit holds the schema definition for the GitCommit entity.
type GitCommit struct {
	ent.Schema
}

// Fields of the GitCommit.
func (GitCommit) Fields() []ent.Field {
	return []ent.Field{
		field.Any("author"),
		field.String("url"),
		field.Time("date"),
	}
}

// Edges of the GitCommit.
func (GitCommit) Edges() []ent.Edge {
	return nil
}
