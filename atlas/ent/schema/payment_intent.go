package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PaymentIntent struct {
	ent.Schema
}

func (PaymentIntent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("customer_email"),
		field.String("customer_phone"),
		field.String("customer_address"),
		field.String("status"),
		field.Float("amount"),
		field.String("currency"),
		field.Time("added_on").Optional(),
		field.Time("updated_on").Optional(),
	}
}

func (PaymentIntent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payment_attempts", PaymentAttempt.Type),
	}
}

func (PaymentIntent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "paymentintent"},
	}
}
