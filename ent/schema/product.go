package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("imageUrl"),
		field.String("summary"),
		field.Float("price"),
		field.Int("quantity"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sub", SubCategory.Type).Ref("product").Required(),
		edge.To("cart", Cart.Type),
	}
}
