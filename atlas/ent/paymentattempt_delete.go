// Code generated by ent, DO NOT EDIT.

package ent

import (
	"atlas/ent/paymentattempt"
	"atlas/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentAttemptDelete is the builder for deleting a PaymentAttempt entity.
type PaymentAttemptDelete struct {
	config
	hooks    []Hook
	mutation *PaymentAttemptMutation
}

// Where appends a list predicates to the PaymentAttemptDelete builder.
func (pad *PaymentAttemptDelete) Where(ps ...predicate.PaymentAttempt) *PaymentAttemptDelete {
	pad.mutation.Where(ps...)
	return pad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pad *PaymentAttemptDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pad.sqlExec, pad.mutation, pad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pad *PaymentAttemptDelete) ExecX(ctx context.Context) int {
	n, err := pad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pad *PaymentAttemptDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(paymentattempt.Table, sqlgraph.NewFieldSpec(paymentattempt.FieldID, field.TypeInt))
	if ps := pad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pad.mutation.done = true
	return affected, err
}

// PaymentAttemptDeleteOne is the builder for deleting a single PaymentAttempt entity.
type PaymentAttemptDeleteOne struct {
	pad *PaymentAttemptDelete
}

// Where appends a list predicates to the PaymentAttemptDelete builder.
func (pado *PaymentAttemptDeleteOne) Where(ps ...predicate.PaymentAttempt) *PaymentAttemptDeleteOne {
	pado.pad.mutation.Where(ps...)
	return pado
}

// Exec executes the deletion query.
func (pado *PaymentAttemptDeleteOne) Exec(ctx context.Context) error {
	n, err := pado.pad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{paymentattempt.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pado *PaymentAttemptDeleteOne) ExecX(ctx context.Context) {
	if err := pado.Exec(ctx); err != nil {
		panic(err)
	}
}