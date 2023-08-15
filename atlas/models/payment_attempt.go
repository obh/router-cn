package models

import (
	"atlas/ent"
	"time"
)

type PaymentAttemptReq struct {
	PaymentIntentId      int               `json:"payment_intent_id"`
	PaymentMethod        string            `json:"payment_method"`
	PaymentMethodDetails map[string]string `json:"payment_method_details"`
}

type PaymentAttempt struct {
	Id                   int               `json:"id"`
	PaymentIntentId      int               `json:"payment_intent_id"`
	PaymentMethod        string            `json:"payment_method"`
	PaymentMethodDetails map[string]string `json:"payment_method_details"`
	PaymentProcessor     string            `json:"processor"`
	PaymentProcessorRef  string            `json:"processor_reference"`
	PaymentHash          string            `json:"payment_hash"`
	Metadata             string            `json:"metadata"`
	Amount               float32           `json:"amount"`
	Currency             string            `json:"currency"`
	Status               string            `json:"status"`
	AddedOn              time.Time         `json:"added_on"`
	UpdatedOn            time.Time         `json:"updated_on"`
}

func ConvertEntToPA(p ent.PaymentAttempt) *PaymentAttempt {
	return &PaymentAttempt{
		Id:               p.ID,
		PaymentIntentId:  p.PaymentIntentID,
		PaymentProcessor: p.Processor,
		PaymentMethod:    p.PaymentMethod,
		Amount:           float32(p.Amount),
		Currency:         p.Currency,
		Status:           p.Status,
		AddedOn:          p.AddedOn,
		UpdatedOn:        p.UpdatedOn,

		PaymentHash: p.PaymentHash,
		Metadata:    p.Metadata,
	}
}
