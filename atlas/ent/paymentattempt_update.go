// Code generated by ent, DO NOT EDIT.

package ent

import (
	"atlas/ent/paymentattempt"
	"atlas/ent/paymentintent"
	"atlas/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentAttemptUpdate is the builder for updating PaymentAttempt entities.
type PaymentAttemptUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentAttemptMutation
}

// Where appends a list predicates to the PaymentAttemptUpdate builder.
func (pau *PaymentAttemptUpdate) Where(ps ...predicate.PaymentAttempt) *PaymentAttemptUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetPaymentIntentID sets the "payment_intent_id" field.
func (pau *PaymentAttemptUpdate) SetPaymentIntentID(i int) *PaymentAttemptUpdate {
	pau.mutation.SetPaymentIntentID(i)
	return pau
}

// SetPaymentMethod sets the "payment_method" field.
func (pau *PaymentAttemptUpdate) SetPaymentMethod(s string) *PaymentAttemptUpdate {
	pau.mutation.SetPaymentMethod(s)
	return pau
}

// SetProcessor sets the "processor" field.
func (pau *PaymentAttemptUpdate) SetProcessor(s string) *PaymentAttemptUpdate {
	pau.mutation.SetProcessor(s)
	return pau
}

// SetNillableProcessor sets the "processor" field if the given value is not nil.
func (pau *PaymentAttemptUpdate) SetNillableProcessor(s *string) *PaymentAttemptUpdate {
	if s != nil {
		pau.SetProcessor(*s)
	}
	return pau
}

// ClearProcessor clears the value of the "processor" field.
func (pau *PaymentAttemptUpdate) ClearProcessor() *PaymentAttemptUpdate {
	pau.mutation.ClearProcessor()
	return pau
}

// SetProcessorRef sets the "processor_ref" field.
func (pau *PaymentAttemptUpdate) SetProcessorRef(s string) *PaymentAttemptUpdate {
	pau.mutation.SetProcessorRef(s)
	return pau
}

// SetNillableProcessorRef sets the "processor_ref" field if the given value is not nil.
func (pau *PaymentAttemptUpdate) SetNillableProcessorRef(s *string) *PaymentAttemptUpdate {
	if s != nil {
		pau.SetProcessorRef(*s)
	}
	return pau
}

// ClearProcessorRef clears the value of the "processor_ref" field.
func (pau *PaymentAttemptUpdate) ClearProcessorRef() *PaymentAttemptUpdate {
	pau.mutation.ClearProcessorRef()
	return pau
}

// SetMetadata sets the "metadata" field.
func (pau *PaymentAttemptUpdate) SetMetadata(s string) *PaymentAttemptUpdate {
	pau.mutation.SetMetadata(s)
	return pau
}

// SetNillableMetadata sets the "metadata" field if the given value is not nil.
func (pau *PaymentAttemptUpdate) SetNillableMetadata(s *string) *PaymentAttemptUpdate {
	if s != nil {
		pau.SetMetadata(*s)
	}
	return pau
}

// ClearMetadata clears the value of the "metadata" field.
func (pau *PaymentAttemptUpdate) ClearMetadata() *PaymentAttemptUpdate {
	pau.mutation.ClearMetadata()
	return pau
}

// SetPaymentHash sets the "payment_hash" field.
func (pau *PaymentAttemptUpdate) SetPaymentHash(s string) *PaymentAttemptUpdate {
	pau.mutation.SetPaymentHash(s)
	return pau
}

// SetNillablePaymentHash sets the "payment_hash" field if the given value is not nil.
func (pau *PaymentAttemptUpdate) SetNillablePaymentHash(s *string) *PaymentAttemptUpdate {
	if s != nil {
		pau.SetPaymentHash(*s)
	}
	return pau
}

// ClearPaymentHash clears the value of the "payment_hash" field.
func (pau *PaymentAttemptUpdate) ClearPaymentHash() *PaymentAttemptUpdate {
	pau.mutation.ClearPaymentHash()
	return pau
}

// SetStatus sets the "status" field.
func (pau *PaymentAttemptUpdate) SetStatus(s string) *PaymentAttemptUpdate {
	pau.mutation.SetStatus(s)
	return pau
}

// SetAmount sets the "amount" field.
func (pau *PaymentAttemptUpdate) SetAmount(f float64) *PaymentAttemptUpdate {
	pau.mutation.ResetAmount()
	pau.mutation.SetAmount(f)
	return pau
}

// AddAmount adds f to the "amount" field.
func (pau *PaymentAttemptUpdate) AddAmount(f float64) *PaymentAttemptUpdate {
	pau.mutation.AddAmount(f)
	return pau
}

// SetCurrency sets the "currency" field.
func (pau *PaymentAttemptUpdate) SetCurrency(s string) *PaymentAttemptUpdate {
	pau.mutation.SetCurrency(s)
	return pau
}

// SetAddedOn sets the "added_on" field.
func (pau *PaymentAttemptUpdate) SetAddedOn(t time.Time) *PaymentAttemptUpdate {
	pau.mutation.SetAddedOn(t)
	return pau
}

// SetNillableAddedOn sets the "added_on" field if the given value is not nil.
func (pau *PaymentAttemptUpdate) SetNillableAddedOn(t *time.Time) *PaymentAttemptUpdate {
	if t != nil {
		pau.SetAddedOn(*t)
	}
	return pau
}

// ClearAddedOn clears the value of the "added_on" field.
func (pau *PaymentAttemptUpdate) ClearAddedOn() *PaymentAttemptUpdate {
	pau.mutation.ClearAddedOn()
	return pau
}

// SetUpdatedOn sets the "updated_on" field.
func (pau *PaymentAttemptUpdate) SetUpdatedOn(t time.Time) *PaymentAttemptUpdate {
	pau.mutation.SetUpdatedOn(t)
	return pau
}

// SetNillableUpdatedOn sets the "updated_on" field if the given value is not nil.
func (pau *PaymentAttemptUpdate) SetNillableUpdatedOn(t *time.Time) *PaymentAttemptUpdate {
	if t != nil {
		pau.SetUpdatedOn(*t)
	}
	return pau
}

// ClearUpdatedOn clears the value of the "updated_on" field.
func (pau *PaymentAttemptUpdate) ClearUpdatedOn() *PaymentAttemptUpdate {
	pau.mutation.ClearUpdatedOn()
	return pau
}

// SetPaymentIntent sets the "payment_intent" edge to the PaymentIntent entity.
func (pau *PaymentAttemptUpdate) SetPaymentIntent(p *PaymentIntent) *PaymentAttemptUpdate {
	return pau.SetPaymentIntentID(p.ID)
}

// Mutation returns the PaymentAttemptMutation object of the builder.
func (pau *PaymentAttemptUpdate) Mutation() *PaymentAttemptMutation {
	return pau.mutation
}

// ClearPaymentIntent clears the "payment_intent" edge to the PaymentIntent entity.
func (pau *PaymentAttemptUpdate) ClearPaymentIntent() *PaymentAttemptUpdate {
	pau.mutation.ClearPaymentIntent()
	return pau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *PaymentAttemptUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *PaymentAttemptUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *PaymentAttemptUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *PaymentAttemptUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *PaymentAttemptUpdate) check() error {
	if _, ok := pau.mutation.PaymentIntentID(); pau.mutation.PaymentIntentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PaymentAttempt.payment_intent"`)
	}
	return nil
}

func (pau *PaymentAttemptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(paymentattempt.Table, paymentattempt.Columns, sqlgraph.NewFieldSpec(paymentattempt.FieldID, field.TypeInt))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.PaymentMethod(); ok {
		_spec.SetField(paymentattempt.FieldPaymentMethod, field.TypeString, value)
	}
	if value, ok := pau.mutation.Processor(); ok {
		_spec.SetField(paymentattempt.FieldProcessor, field.TypeString, value)
	}
	if pau.mutation.ProcessorCleared() {
		_spec.ClearField(paymentattempt.FieldProcessor, field.TypeString)
	}
	if value, ok := pau.mutation.ProcessorRef(); ok {
		_spec.SetField(paymentattempt.FieldProcessorRef, field.TypeString, value)
	}
	if pau.mutation.ProcessorRefCleared() {
		_spec.ClearField(paymentattempt.FieldProcessorRef, field.TypeString)
	}
	if value, ok := pau.mutation.Metadata(); ok {
		_spec.SetField(paymentattempt.FieldMetadata, field.TypeString, value)
	}
	if pau.mutation.MetadataCleared() {
		_spec.ClearField(paymentattempt.FieldMetadata, field.TypeString)
	}
	if value, ok := pau.mutation.PaymentHash(); ok {
		_spec.SetField(paymentattempt.FieldPaymentHash, field.TypeString, value)
	}
	if pau.mutation.PaymentHashCleared() {
		_spec.ClearField(paymentattempt.FieldPaymentHash, field.TypeString)
	}
	if value, ok := pau.mutation.Status(); ok {
		_spec.SetField(paymentattempt.FieldStatus, field.TypeString, value)
	}
	if value, ok := pau.mutation.Amount(); ok {
		_spec.SetField(paymentattempt.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pau.mutation.AddedAmount(); ok {
		_spec.AddField(paymentattempt.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pau.mutation.Currency(); ok {
		_spec.SetField(paymentattempt.FieldCurrency, field.TypeString, value)
	}
	if value, ok := pau.mutation.AddedOn(); ok {
		_spec.SetField(paymentattempt.FieldAddedOn, field.TypeTime, value)
	}
	if pau.mutation.AddedOnCleared() {
		_spec.ClearField(paymentattempt.FieldAddedOn, field.TypeTime)
	}
	if value, ok := pau.mutation.UpdatedOn(); ok {
		_spec.SetField(paymentattempt.FieldUpdatedOn, field.TypeTime, value)
	}
	if pau.mutation.UpdatedOnCleared() {
		_spec.ClearField(paymentattempt.FieldUpdatedOn, field.TypeTime)
	}
	if pau.mutation.PaymentIntentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   paymentattempt.PaymentIntentTable,
			Columns: []string{paymentattempt.PaymentIntentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentintent.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.PaymentIntentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   paymentattempt.PaymentIntentTable,
			Columns: []string{paymentattempt.PaymentIntentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentintent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentattempt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// PaymentAttemptUpdateOne is the builder for updating a single PaymentAttempt entity.
type PaymentAttemptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentAttemptMutation
}

// SetPaymentIntentID sets the "payment_intent_id" field.
func (pauo *PaymentAttemptUpdateOne) SetPaymentIntentID(i int) *PaymentAttemptUpdateOne {
	pauo.mutation.SetPaymentIntentID(i)
	return pauo
}

// SetPaymentMethod sets the "payment_method" field.
func (pauo *PaymentAttemptUpdateOne) SetPaymentMethod(s string) *PaymentAttemptUpdateOne {
	pauo.mutation.SetPaymentMethod(s)
	return pauo
}

// SetProcessor sets the "processor" field.
func (pauo *PaymentAttemptUpdateOne) SetProcessor(s string) *PaymentAttemptUpdateOne {
	pauo.mutation.SetProcessor(s)
	return pauo
}

// SetNillableProcessor sets the "processor" field if the given value is not nil.
func (pauo *PaymentAttemptUpdateOne) SetNillableProcessor(s *string) *PaymentAttemptUpdateOne {
	if s != nil {
		pauo.SetProcessor(*s)
	}
	return pauo
}

// ClearProcessor clears the value of the "processor" field.
func (pauo *PaymentAttemptUpdateOne) ClearProcessor() *PaymentAttemptUpdateOne {
	pauo.mutation.ClearProcessor()
	return pauo
}

// SetProcessorRef sets the "processor_ref" field.
func (pauo *PaymentAttemptUpdateOne) SetProcessorRef(s string) *PaymentAttemptUpdateOne {
	pauo.mutation.SetProcessorRef(s)
	return pauo
}

// SetNillableProcessorRef sets the "processor_ref" field if the given value is not nil.
func (pauo *PaymentAttemptUpdateOne) SetNillableProcessorRef(s *string) *PaymentAttemptUpdateOne {
	if s != nil {
		pauo.SetProcessorRef(*s)
	}
	return pauo
}

// ClearProcessorRef clears the value of the "processor_ref" field.
func (pauo *PaymentAttemptUpdateOne) ClearProcessorRef() *PaymentAttemptUpdateOne {
	pauo.mutation.ClearProcessorRef()
	return pauo
}

// SetMetadata sets the "metadata" field.
func (pauo *PaymentAttemptUpdateOne) SetMetadata(s string) *PaymentAttemptUpdateOne {
	pauo.mutation.SetMetadata(s)
	return pauo
}

// SetNillableMetadata sets the "metadata" field if the given value is not nil.
func (pauo *PaymentAttemptUpdateOne) SetNillableMetadata(s *string) *PaymentAttemptUpdateOne {
	if s != nil {
		pauo.SetMetadata(*s)
	}
	return pauo
}

// ClearMetadata clears the value of the "metadata" field.
func (pauo *PaymentAttemptUpdateOne) ClearMetadata() *PaymentAttemptUpdateOne {
	pauo.mutation.ClearMetadata()
	return pauo
}

// SetPaymentHash sets the "payment_hash" field.
func (pauo *PaymentAttemptUpdateOne) SetPaymentHash(s string) *PaymentAttemptUpdateOne {
	pauo.mutation.SetPaymentHash(s)
	return pauo
}

// SetNillablePaymentHash sets the "payment_hash" field if the given value is not nil.
func (pauo *PaymentAttemptUpdateOne) SetNillablePaymentHash(s *string) *PaymentAttemptUpdateOne {
	if s != nil {
		pauo.SetPaymentHash(*s)
	}
	return pauo
}

// ClearPaymentHash clears the value of the "payment_hash" field.
func (pauo *PaymentAttemptUpdateOne) ClearPaymentHash() *PaymentAttemptUpdateOne {
	pauo.mutation.ClearPaymentHash()
	return pauo
}

// SetStatus sets the "status" field.
func (pauo *PaymentAttemptUpdateOne) SetStatus(s string) *PaymentAttemptUpdateOne {
	pauo.mutation.SetStatus(s)
	return pauo
}

// SetAmount sets the "amount" field.
func (pauo *PaymentAttemptUpdateOne) SetAmount(f float64) *PaymentAttemptUpdateOne {
	pauo.mutation.ResetAmount()
	pauo.mutation.SetAmount(f)
	return pauo
}

// AddAmount adds f to the "amount" field.
func (pauo *PaymentAttemptUpdateOne) AddAmount(f float64) *PaymentAttemptUpdateOne {
	pauo.mutation.AddAmount(f)
	return pauo
}

// SetCurrency sets the "currency" field.
func (pauo *PaymentAttemptUpdateOne) SetCurrency(s string) *PaymentAttemptUpdateOne {
	pauo.mutation.SetCurrency(s)
	return pauo
}

// SetAddedOn sets the "added_on" field.
func (pauo *PaymentAttemptUpdateOne) SetAddedOn(t time.Time) *PaymentAttemptUpdateOne {
	pauo.mutation.SetAddedOn(t)
	return pauo
}

// SetNillableAddedOn sets the "added_on" field if the given value is not nil.
func (pauo *PaymentAttemptUpdateOne) SetNillableAddedOn(t *time.Time) *PaymentAttemptUpdateOne {
	if t != nil {
		pauo.SetAddedOn(*t)
	}
	return pauo
}

// ClearAddedOn clears the value of the "added_on" field.
func (pauo *PaymentAttemptUpdateOne) ClearAddedOn() *PaymentAttemptUpdateOne {
	pauo.mutation.ClearAddedOn()
	return pauo
}

// SetUpdatedOn sets the "updated_on" field.
func (pauo *PaymentAttemptUpdateOne) SetUpdatedOn(t time.Time) *PaymentAttemptUpdateOne {
	pauo.mutation.SetUpdatedOn(t)
	return pauo
}

// SetNillableUpdatedOn sets the "updated_on" field if the given value is not nil.
func (pauo *PaymentAttemptUpdateOne) SetNillableUpdatedOn(t *time.Time) *PaymentAttemptUpdateOne {
	if t != nil {
		pauo.SetUpdatedOn(*t)
	}
	return pauo
}

// ClearUpdatedOn clears the value of the "updated_on" field.
func (pauo *PaymentAttemptUpdateOne) ClearUpdatedOn() *PaymentAttemptUpdateOne {
	pauo.mutation.ClearUpdatedOn()
	return pauo
}

// SetPaymentIntent sets the "payment_intent" edge to the PaymentIntent entity.
func (pauo *PaymentAttemptUpdateOne) SetPaymentIntent(p *PaymentIntent) *PaymentAttemptUpdateOne {
	return pauo.SetPaymentIntentID(p.ID)
}

// Mutation returns the PaymentAttemptMutation object of the builder.
func (pauo *PaymentAttemptUpdateOne) Mutation() *PaymentAttemptMutation {
	return pauo.mutation
}

// ClearPaymentIntent clears the "payment_intent" edge to the PaymentIntent entity.
func (pauo *PaymentAttemptUpdateOne) ClearPaymentIntent() *PaymentAttemptUpdateOne {
	pauo.mutation.ClearPaymentIntent()
	return pauo
}

// Where appends a list predicates to the PaymentAttemptUpdate builder.
func (pauo *PaymentAttemptUpdateOne) Where(ps ...predicate.PaymentAttempt) *PaymentAttemptUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *PaymentAttemptUpdateOne) Select(field string, fields ...string) *PaymentAttemptUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated PaymentAttempt entity.
func (pauo *PaymentAttemptUpdateOne) Save(ctx context.Context) (*PaymentAttempt, error) {
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *PaymentAttemptUpdateOne) SaveX(ctx context.Context) *PaymentAttempt {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *PaymentAttemptUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *PaymentAttemptUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *PaymentAttemptUpdateOne) check() error {
	if _, ok := pauo.mutation.PaymentIntentID(); pauo.mutation.PaymentIntentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PaymentAttempt.payment_intent"`)
	}
	return nil
}

func (pauo *PaymentAttemptUpdateOne) sqlSave(ctx context.Context) (_node *PaymentAttempt, err error) {
	if err := pauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(paymentattempt.Table, paymentattempt.Columns, sqlgraph.NewFieldSpec(paymentattempt.FieldID, field.TypeInt))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PaymentAttempt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentattempt.FieldID)
		for _, f := range fields {
			if !paymentattempt.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != paymentattempt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.PaymentMethod(); ok {
		_spec.SetField(paymentattempt.FieldPaymentMethod, field.TypeString, value)
	}
	if value, ok := pauo.mutation.Processor(); ok {
		_spec.SetField(paymentattempt.FieldProcessor, field.TypeString, value)
	}
	if pauo.mutation.ProcessorCleared() {
		_spec.ClearField(paymentattempt.FieldProcessor, field.TypeString)
	}
	if value, ok := pauo.mutation.ProcessorRef(); ok {
		_spec.SetField(paymentattempt.FieldProcessorRef, field.TypeString, value)
	}
	if pauo.mutation.ProcessorRefCleared() {
		_spec.ClearField(paymentattempt.FieldProcessorRef, field.TypeString)
	}
	if value, ok := pauo.mutation.Metadata(); ok {
		_spec.SetField(paymentattempt.FieldMetadata, field.TypeString, value)
	}
	if pauo.mutation.MetadataCleared() {
		_spec.ClearField(paymentattempt.FieldMetadata, field.TypeString)
	}
	if value, ok := pauo.mutation.PaymentHash(); ok {
		_spec.SetField(paymentattempt.FieldPaymentHash, field.TypeString, value)
	}
	if pauo.mutation.PaymentHashCleared() {
		_spec.ClearField(paymentattempt.FieldPaymentHash, field.TypeString)
	}
	if value, ok := pauo.mutation.Status(); ok {
		_spec.SetField(paymentattempt.FieldStatus, field.TypeString, value)
	}
	if value, ok := pauo.mutation.Amount(); ok {
		_spec.SetField(paymentattempt.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pauo.mutation.AddedAmount(); ok {
		_spec.AddField(paymentattempt.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pauo.mutation.Currency(); ok {
		_spec.SetField(paymentattempt.FieldCurrency, field.TypeString, value)
	}
	if value, ok := pauo.mutation.AddedOn(); ok {
		_spec.SetField(paymentattempt.FieldAddedOn, field.TypeTime, value)
	}
	if pauo.mutation.AddedOnCleared() {
		_spec.ClearField(paymentattempt.FieldAddedOn, field.TypeTime)
	}
	if value, ok := pauo.mutation.UpdatedOn(); ok {
		_spec.SetField(paymentattempt.FieldUpdatedOn, field.TypeTime, value)
	}
	if pauo.mutation.UpdatedOnCleared() {
		_spec.ClearField(paymentattempt.FieldUpdatedOn, field.TypeTime)
	}
	if pauo.mutation.PaymentIntentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   paymentattempt.PaymentIntentTable,
			Columns: []string{paymentattempt.PaymentIntentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentintent.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.PaymentIntentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   paymentattempt.PaymentIntentTable,
			Columns: []string{paymentattempt.PaymentIntentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentintent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PaymentAttempt{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentattempt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
