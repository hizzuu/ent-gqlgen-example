package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			Unique().
			Immutable(),
		field.String("nickname"),
		field.String("bio"),
		field.Enum("role").
			NamedValues(
				"general", "GENERAL",
				"official", "OFFICIAL",
			).Annotations(
			entgql.OrderField("ROLE"),
		),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.From("avatar", Image.Type).
			Ref("users").
			Unique(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("nickname"),
		index.Fields("uid").
			Unique(),
	}
}
