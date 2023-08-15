package service

import (
	"atlas/models"
	"context"
)

type PaymentIntent interface {
	Create(c context.Context, req models.PaymentIntentReq) (*models.PaymentIntent, error)
	Get(c context.Context, id int) (*models.PaymentIntent, error)
}
