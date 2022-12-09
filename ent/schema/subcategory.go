package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubCategory holds the schema definition for the SubCategory entity.
type SubCategory struct {
	ent.Schema
}

// Fields of the SubCategory.
func (SubCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique(),
		field.String("description"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the SubCategory.
func (SubCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("subs").Unique(),
		edge.To("product", Product.Type),
	}
}
