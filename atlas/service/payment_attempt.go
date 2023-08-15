package service

import (
	"atlas/models"
	"context"
)

type PaymentAttempt interface {
	CreateAttempt(c context.Context, pi *models.PaymentIntent, req *models.PaymentAttemptReq) (*models.PaymentAttempt, error)
}
