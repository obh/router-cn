// Code generated by ent, DO NOT EDIT.

package paymentintent

import (
	"atlas/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldID, id))
}

// CustomerEmail applies equality check predicate on the "customer_email" field. It's identical to CustomerEmailEQ.
func CustomerEmail(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCustomerEmail, v))
}

// CustomerPhone applies equality check predicate on the "customer_phone" field. It's identical to CustomerPhoneEQ.
func CustomerPhone(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCustomerPhone, v))
}

// CustomerAddress applies equality check predicate on the "customer_address" field. It's identical to CustomerAddressEQ.
func CustomerAddress(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCustomerAddress, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldStatus, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldAmount, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCurrency, v))
}

// AddedOn applies equality check predicate on the "added_on" field. It's identical to AddedOnEQ.
func AddedOn(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldAddedOn, v))
}

// UpdatedOn applies equality check predicate on the "updated_on" field. It's identical to UpdatedOnEQ.
func UpdatedOn(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldUpdatedOn, v))
}

// CustomerEmailEQ applies the EQ predicate on the "customer_email" field.
func CustomerEmailEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCustomerEmail, v))
}

// CustomerEmailNEQ applies the NEQ predicate on the "customer_email" field.
func CustomerEmailNEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldCustomerEmail, v))
}

// CustomerEmailIn applies the In predicate on the "customer_email" field.
func CustomerEmailIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldCustomerEmail, vs...))
}

// CustomerEmailNotIn applies the NotIn predicate on the "customer_email" field.
func CustomerEmailNotIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldCustomerEmail, vs...))
}

// CustomerEmailGT applies the GT predicate on the "customer_email" field.
func CustomerEmailGT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldCustomerEmail, v))
}

// CustomerEmailGTE applies the GTE predicate on the "customer_email" field.
func CustomerEmailGTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldCustomerEmail, v))
}

// CustomerEmailLT applies the LT predicate on the "customer_email" field.
func CustomerEmailLT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldCustomerEmail, v))
}

// CustomerEmailLTE applies the LTE predicate on the "customer_email" field.
func CustomerEmailLTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldCustomerEmail, v))
}

// CustomerEmailContains applies the Contains predicate on the "customer_email" field.
func CustomerEmailContains(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContains(FieldCustomerEmail, v))
}

// CustomerEmailHasPrefix applies the HasPrefix predicate on the "customer_email" field.
func CustomerEmailHasPrefix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasPrefix(FieldCustomerEmail, v))
}

// CustomerEmailHasSuffix applies the HasSuffix predicate on the "customer_email" field.
func CustomerEmailHasSuffix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasSuffix(FieldCustomerEmail, v))
}

// CustomerEmailEqualFold applies the EqualFold predicate on the "customer_email" field.
func CustomerEmailEqualFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEqualFold(FieldCustomerEmail, v))
}

// CustomerEmailContainsFold applies the ContainsFold predicate on the "customer_email" field.
func CustomerEmailContainsFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContainsFold(FieldCustomerEmail, v))
}

// CustomerPhoneEQ applies the EQ predicate on the "customer_phone" field.
func CustomerPhoneEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCustomerPhone, v))
}

// CustomerPhoneNEQ applies the NEQ predicate on the "customer_phone" field.
func CustomerPhoneNEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldCustomerPhone, v))
}

// CustomerPhoneIn applies the In predicate on the "customer_phone" field.
func CustomerPhoneIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldCustomerPhone, vs...))
}

// CustomerPhoneNotIn applies the NotIn predicate on the "customer_phone" field.
func CustomerPhoneNotIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldCustomerPhone, vs...))
}

// CustomerPhoneGT applies the GT predicate on the "customer_phone" field.
func CustomerPhoneGT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldCustomerPhone, v))
}

// CustomerPhoneGTE applies the GTE predicate on the "customer_phone" field.
func CustomerPhoneGTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldCustomerPhone, v))
}

// CustomerPhoneLT applies the LT predicate on the "customer_phone" field.
func CustomerPhoneLT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldCustomerPhone, v))
}

// CustomerPhoneLTE applies the LTE predicate on the "customer_phone" field.
func CustomerPhoneLTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldCustomerPhone, v))
}

// CustomerPhoneContains applies the Contains predicate on the "customer_phone" field.
func CustomerPhoneContains(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContains(FieldCustomerPhone, v))
}

// CustomerPhoneHasPrefix applies the HasPrefix predicate on the "customer_phone" field.
func CustomerPhoneHasPrefix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasPrefix(FieldCustomerPhone, v))
}

// CustomerPhoneHasSuffix applies the HasSuffix predicate on the "customer_phone" field.
func CustomerPhoneHasSuffix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasSuffix(FieldCustomerPhone, v))
}

// CustomerPhoneEqualFold applies the EqualFold predicate on the "customer_phone" field.
func CustomerPhoneEqualFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEqualFold(FieldCustomerPhone, v))
}

// CustomerPhoneContainsFold applies the ContainsFold predicate on the "customer_phone" field.
func CustomerPhoneContainsFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContainsFold(FieldCustomerPhone, v))
}

// CustomerAddressEQ applies the EQ predicate on the "customer_address" field.
func CustomerAddressEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCustomerAddress, v))
}

// CustomerAddressNEQ applies the NEQ predicate on the "customer_address" field.
func CustomerAddressNEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldCustomerAddress, v))
}

// CustomerAddressIn applies the In predicate on the "customer_address" field.
func CustomerAddressIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldCustomerAddress, vs...))
}

// CustomerAddressNotIn applies the NotIn predicate on the "customer_address" field.
func CustomerAddressNotIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldCustomerAddress, vs...))
}

// CustomerAddressGT applies the GT predicate on the "customer_address" field.
func CustomerAddressGT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldCustomerAddress, v))
}

// CustomerAddressGTE applies the GTE predicate on the "customer_address" field.
func CustomerAddressGTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldCustomerAddress, v))
}

// CustomerAddressLT applies the LT predicate on the "customer_address" field.
func CustomerAddressLT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldCustomerAddress, v))
}

// CustomerAddressLTE applies the LTE predicate on the "customer_address" field.
func CustomerAddressLTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldCustomerAddress, v))
}

// CustomerAddressContains applies the Contains predicate on the "customer_address" field.
func CustomerAddressContains(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContains(FieldCustomerAddress, v))
}

// CustomerAddressHasPrefix applies the HasPrefix predicate on the "customer_address" field.
func CustomerAddressHasPrefix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasPrefix(FieldCustomerAddress, v))
}

// CustomerAddressHasSuffix applies the HasSuffix predicate on the "customer_address" field.
func CustomerAddressHasSuffix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasSuffix(FieldCustomerAddress, v))
}

// CustomerAddressEqualFold applies the EqualFold predicate on the "customer_address" field.
func CustomerAddressEqualFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEqualFold(FieldCustomerAddress, v))
}

// CustomerAddressContainsFold applies the ContainsFold predicate on the "customer_address" field.
func CustomerAddressContainsFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContainsFold(FieldCustomerAddress, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContainsFold(FieldStatus, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldAmount, v))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyContains applies the Contains predicate on the "currency" field.
func CurrencyContains(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContains(FieldCurrency, v))
}

// CurrencyHasPrefix applies the HasPrefix predicate on the "currency" field.
func CurrencyHasPrefix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasPrefix(FieldCurrency, v))
}

// CurrencyHasSuffix applies the HasSuffix predicate on the "currency" field.
func CurrencyHasSuffix(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldHasSuffix(FieldCurrency, v))
}

// CurrencyEqualFold applies the EqualFold predicate on the "currency" field.
func CurrencyEqualFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEqualFold(FieldCurrency, v))
}

// CurrencyContainsFold applies the ContainsFold predicate on the "currency" field.
func CurrencyContainsFold(v string) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldContainsFold(FieldCurrency, v))
}

// AddedOnEQ applies the EQ predicate on the "added_on" field.
func AddedOnEQ(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldAddedOn, v))
}

// AddedOnNEQ applies the NEQ predicate on the "added_on" field.
func AddedOnNEQ(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldAddedOn, v))
}

// AddedOnIn applies the In predicate on the "added_on" field.
func AddedOnIn(vs ...time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldAddedOn, vs...))
}

// AddedOnNotIn applies the NotIn predicate on the "added_on" field.
func AddedOnNotIn(vs ...time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldAddedOn, vs...))
}

// AddedOnGT applies the GT predicate on the "added_on" field.
func AddedOnGT(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldAddedOn, v))
}

// AddedOnGTE applies the GTE predicate on the "added_on" field.
func AddedOnGTE(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldAddedOn, v))
}

// AddedOnLT applies the LT predicate on the "added_on" field.
func AddedOnLT(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldAddedOn, v))
}

// AddedOnLTE applies the LTE predicate on the "added_on" field.
func AddedOnLTE(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldAddedOn, v))
}

// AddedOnIsNil applies the IsNil predicate on the "added_on" field.
func AddedOnIsNil() predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIsNull(FieldAddedOn))
}

// AddedOnNotNil applies the NotNil predicate on the "added_on" field.
func AddedOnNotNil() predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotNull(FieldAddedOn))
}

// UpdatedOnEQ applies the EQ predicate on the "updated_on" field.
func UpdatedOnEQ(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldEQ(FieldUpdatedOn, v))
}

// UpdatedOnNEQ applies the NEQ predicate on the "updated_on" field.
func UpdatedOnNEQ(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNEQ(FieldUpdatedOn, v))
}

// UpdatedOnIn applies the In predicate on the "updated_on" field.
func UpdatedOnIn(vs ...time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIn(FieldUpdatedOn, vs...))
}

// UpdatedOnNotIn applies the NotIn predicate on the "updated_on" field.
func UpdatedOnNotIn(vs ...time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotIn(FieldUpdatedOn, vs...))
}

// UpdatedOnGT applies the GT predicate on the "updated_on" field.
func UpdatedOnGT(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGT(FieldUpdatedOn, v))
}

// UpdatedOnGTE applies the GTE predicate on the "updated_on" field.
func UpdatedOnGTE(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldGTE(FieldUpdatedOn, v))
}

// UpdatedOnLT applies the LT predicate on the "updated_on" field.
func UpdatedOnLT(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLT(FieldUpdatedOn, v))
}

// UpdatedOnLTE applies the LTE predicate on the "updated_on" field.
func UpdatedOnLTE(v time.Time) predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldLTE(FieldUpdatedOn, v))
}

// UpdatedOnIsNil applies the IsNil predicate on the "updated_on" field.
func UpdatedOnIsNil() predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldIsNull(FieldUpdatedOn))
}

// UpdatedOnNotNil applies the NotNil predicate on the "updated_on" field.
func UpdatedOnNotNil() predicate.PaymentIntent {
	return predicate.PaymentIntent(sql.FieldNotNull(FieldUpdatedOn))
}

// HasPaymentAttempts applies the HasEdge predicate on the "payment_attempts" edge.
func HasPaymentAttempts() predicate.PaymentIntent {
	return predicate.PaymentIntent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentAttemptsTable, PaymentAttemptsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentAttemptsWith applies the HasEdge predicate on the "payment_attempts" edge with a given conditions (other predicates).
func HasPaymentAttemptsWith(preds ...predicate.PaymentAttempt) predicate.PaymentIntent {
	return predicate.PaymentIntent(func(s *sql.Selector) {
		step := newPaymentAttemptsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PaymentIntent) predicate.PaymentIntent {
	return predicate.PaymentIntent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PaymentIntent) predicate.PaymentIntent {
	return predicate.PaymentIntent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PaymentIntent) predicate.PaymentIntent {
	return predicate.PaymentIntent(func(s *sql.Selector) {
		p(s.Not())
	})
}
