package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("order_date").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.String("shipping_address"),
		field.String("phone"),
		field.Enum("payment_method").Values("Cash", "Banking"),
		field.Int32("total_price"),
		field.String("tracking_number"),
		field.Enum("order_status").Values("Pending", "Processed", "Shipped", "Cancelled", "PickedUp"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
