package provider

import (
	"atlas/models"
	"context"
)

type Provider interface {
	GetName() string
	CreatePayment(c context.Context, req *models.PaymentIntent) (string, error)
	// GetStatus(c context.Context, req *models.PaymentAttempt) (*models.PaymentAttempt, error)
}
