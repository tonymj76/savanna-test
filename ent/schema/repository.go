package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").NotEmpty(),
		field.String("URL").NotEmpty(),
		field.String("language").NotEmpty(),
		field.Int("open_issues_count"),
		field.Int("watchers_count"),
		field.Int("star_count"),
		field.Int("forks_count"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("commits", Commit.Type),
	}
}