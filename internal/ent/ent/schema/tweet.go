package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("content"),
		field.String("userid"),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return nil
}
