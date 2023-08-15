package razorpay

type Order struct {
	ID         string  `json:"id"`
	Entity     string  `json:"entity"`
	Amount     int     `json:"amount"`
	AmountPaid int     `json:"amount_paid"`
	AmountDue  int     `json:"amount_due"`
	Currency   string  `json:"currency"`
	Receipt    string  `json:"receipt"`
	OfferID    *string `json:"offer_id"`
	Status     string  `json:"status"`
	Attempts   int     `json:"attempts"`
	CreatedAt  int     `json:"created_at"`
}
