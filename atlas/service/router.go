package service

type Router interface {
	GetPreference(paymentMode string, metadata map[string]interface{}) (string, error)
}
