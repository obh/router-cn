package cashfree

import (
	"time"
)

type CustomerDetails struct {
	CustomerID            string `json:"customer_id"`
	CustomerName          string `json:"customer_name"`
	CustomerEmail         string `json:"customer_email"`
	CustomerPhone         string `json:"customer_phone"`
	CustomerAccountNumber string `json:"customer_account_number"`
	CustomerIFSC          string `json:"customer_ifsc"`
	CustomerBankCode      int    `json:"customer_bank_code"`
}

type OrderMeta struct {
	ReturnURL      string `json:"return_url"`
	NotifyURL      string `json:"notify_url"`
	PaymentMethods string `json:"payment_methods"`
}

type Payments struct {
	URL string `json:"url"`
}

type Refunds struct {
	URL string `json:"url"`
}

type Settlements struct {
	URL string `json:"url"`
}

type Order struct {
	CfOrderID        int             `json:"cf_order_id"`
	CreatedAt        time.Time       `json:"created_at"`
	CustomerDetails  CustomerDetails `json:"customer_details"`
	Entity           string          `json:"entity"`
	OrderAmount      float64         `json:"order_amount"`
	OrderCurrency    string          `json:"order_currency"`
	OrderExpiryTime  time.Time       `json:"order_expiry_time"`
	OrderID          string          `json:"order_id"`
	OrderMeta        OrderMeta       `json:"order_meta"`
	OrderNote        interface{}     `json:"order_note"`
	OrderSplits      []interface{}   `json:"order_splits"`
	OrderStatus      string          `json:"order_status"`
	OrderTags        interface{}     `json:"order_tags"`
	PaymentSessionID string          `json:"payment_session_id"`
	Payments         Payments        `json:"payments"`
	Refunds          Refunds         `json:"refunds"`
	Settlements      Settlements     `json:"settlements"`
	TerminalData     interface{}     `json:"terminal_data"`
}

type PaymentData struct {
	URL         string      `json:"url"`
	Payload     interface{} `json:"payload"`
	ContentType interface{} `json:"content_type"`
	Method      interface{} `json:"method"`
}

type Payment struct {
	Action        string      `json:"action"`
	CfPaymentID   int         `json:"cf_payment_id"`
	Channel       string      `json:"channel"`
	Data          PaymentData `json:"data"`
	PaymentAmount float64     `json:"payment_amount"`
	PaymentMethod string      `json:"payment_method"`
}
