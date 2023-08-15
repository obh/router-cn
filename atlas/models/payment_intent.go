package models

import (
	"atlas/ent"
	"time"
)

type PaymentIntentReq struct {
	CartId        string  `json:"cart_id"`
	CustomerEmail string  `json:"customer_email"`
	CustomerPhone string  `json:"customer_phone"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
}

type PaymentIntent struct {
	Id            int       `json:"id"`
	CartId        string    `json:"cart_id"`
	CustomerEmail string    `json:"customer_email"`
	CustomerPhone string    `json:"customer_phone"`
	Amount        float32   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	AddedOn       time.Time `json:"added_on"`
	UpdatedOn     time.Time `json:"updated_on"`
}

func ConvertEntToModel(p ent.PaymentIntent) *PaymentIntent {
	return &PaymentIntent{
		Id:            p.ID,
		CustomerEmail: p.CustomerEmail,
		CustomerPhone: p.CustomerPhone,
		Amount:        float32(p.Amount),
		Currency:      p.Currency,
		Status:        p.Status,
	}
}
