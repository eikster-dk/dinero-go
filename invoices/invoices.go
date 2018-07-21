package invoices

import (
	"github.com/eikc/dinero-go"
)

// InvoiceLine represents a line in an invoice
type InvoiceLine struct {
	BaseAmountValue float64 `json:"baseAmountValue,omitempty"`
	ProductID       string  `json:"productGuid,omitempty"`
	Description     string  `json:"description,omitempty"`
	Comments        string  `json:"comments,omitempty"`
	Quantity        float64 `json:"quantity,omitempty"`
	AccountNumber   int     `json:"accountNumber,omitempty"`
	Unit            string  `json:"unit,omitempty"`
	Discount        float64 `json:"discount,omitempty"`
	LineType        string  `json:"lineType,omitempty"`
}

// TotalLine represents a line total
type TotalLine struct {
	Type        string  `json:"type,omitempty"`
	TotalAmount float64 `json:"totalAmount,omitempty"`
	Position    int     `json:"position,omitempty"`
	Label       string  `json:"label,omitempty"`
}

// Invoice represents an invoice json object
type Invoice struct {
	PaymentDate                  string      `json:"paymentDate"`
	PaymentStatus                string      `json:"paymentStatus"`
	PaymentConditionNumberOfDays int         `json:"paymentConditionNumberOfDays"`
	PaymentConditionType         string      `json:"paymentConditionType"`
	Status                       string      `json:"status"`
	ContactID                    string      `json:"contactGuid"`
	ID                           string      `json:"guid"`
	Timestamp                    string      `json:"timestamp"`
	Created                      dinero.Time `json:"createdAt"`
	UpdatedAt                    dinero.Time `json:"updatedAt"`
	DeletedAt                    dinero.Time `json:"deletedAt"`
	Number                       int         `json:"number"`
	ContactName                  string      `json:"contactName"`
	ShowLinesInclVat             bool        `json:"showLinesInclVat"`
	TotalExclVat                 float64     `json:"totalExclVat"`
	TotalVatableAmount           float64     `json:"totalVatableAmount"`
	TotalInclVat                 float64     `json:"totalInclVat"`
	TotalNonVatableAmount        float64     `json:"totalNonVatableAmount"`
	TotalVat                     float64     `json:"totalVat"`
	TotalLines                   []TotalLine `json:"totalLines"`
	Currency                     string      `json:"currency"`
	Language                     string      `json:"language"`
	ExternalReference            string      `json:"externalReference"`
	Description                  string      `json:"description"`
}

// Delete invoice. The invoice cannot be deleted if booked.
func Delete() {}
