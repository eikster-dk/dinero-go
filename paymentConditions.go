package dinero

// Payment terms
const (
	Netto           = "Netto"
	NettoCash       = "NettoCash"
	CurrentMonthOut = "CurrentMonthNettoOut"
)

// PaymentConditions represents the properties for paymentConditions
type PaymentConditions struct {
	PaymentConditionNumberOfDays int    `json:"paymentConditionNumberOfDays,omitempty"`
	PaymentConditionType         string `json:"paymentConditionType,omitempty"`
}
