package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Commit holds the schema definition for the Commit entity.
type Commit struct {
	ent.Schema
}

// Fields of the Commit.
func (Commit) Fields() []ent.Field {
	return []ent.Field{
		field.Any("author"),
		field.String("url"),
		field.Time("date"),
	}
}

// Edges of the Commit.
func (Commit) Edges() []ent.Edge {
	return nil
}
