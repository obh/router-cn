package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PaymentAttempt struct {
	ent.Schema
}

func (PaymentAttempt) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("payment_intent_id"),
		field.String("payment_method"),
		field.String("processor").Optional(),
		field.String("processor_ref").Optional(),
		field.Text("metadata").Optional(),
		field.String("payment_hash").Optional(),
		field.String("status"),
		field.Float("amount"),
		field.String("currency"),
		field.Time("added_on").Optional(),
		field.Time("updated_on").Optional(),
	}
}

func (PaymentAttempt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment_intent", PaymentIntent.Type).
			Ref("payment_attempts").
			Field("payment_intent_id").
			Unique().Required(),
	}
}

func (PaymentAttempt) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "paymentattempt"},
	}
}
