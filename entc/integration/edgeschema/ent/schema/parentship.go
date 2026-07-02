// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Parentship holds the edge schema definition of the Parentship relationship.
type Parentship struct {
	ent.Schema
}

func (Parentship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("child_id", "parent_id"),
	}
}

// Fields of the Parentship.
func (Parentship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("weight").
			Default(1),
		field.Time("created_at").
			Default(time.Now),
		field.Int("parent_id").
			Immutable(),
		field.Int("child_id").
			Immutable(),
	}
}

// Edges of the Parentship.
func (Parentship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", User.Type).
			Unique().
			Required().
			Immutable().
			Field("parent_id"),
		edge.To("child", User.Type).
			Unique().
			Required().
			Immutable().
			Field("child_id"),
	}
}

// Indexes of the Parentship.
func (Parentship) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		// By default, Ent generates a unique index named <T>_<FK1>_<FK2>
		// for edge-schemas with an ID field to enforce the uniqueness of
		// the edges reside in the join table. However, in this case it is
		// skipped because we define it explicitly in the definition below.
		index.Fields("parent_id", "child_id").
			Unique().
			StorageKey("parentships_edge"),
	}
}

// Policy defines the privacy policy of the Parentship.
func (Parentship) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
