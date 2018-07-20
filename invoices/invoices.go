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
	Type        string
	TotalAmount float64
	Position    int
	Label       string
}

// Invoice represents an invoice json object
type Invoice struct {
	PaymentDate                  string      `json:"paymentDate"`
	PaymentStatus                string      `json:paymentStatus"`
	PaymentConditionNumberOfDays int         `json:"paymentConditionNumberOfDays"`
	PaymentConditionType         string      `json"paymentConditionType"`
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

// Book books a given invoice
func Book() {}

// Send an email with link to a public version of the invoice where it can be printed or downloaded as a pdf.
func Send() {}

// SendPreReminder sends a pre reminder email with link to a public version of the invoice where it
// can be printed or downloaded as a pdf. The invoice needs to be overdue to send the reminder.
// A pre-reminder is a mail reminding the customer, that the invoice is overdue.
// This will not cause a reminder to be created in Dinero, this is only a mailout.
func SendPreReminder() {}

// SendEInvoice sends an e-invoice to an EAN customer
func SendEInvoice() {}

// Update an existing invoice. The invoice cannot be updated if booked.
// endpoint used is version: 1.2
func Update() {}

// CreatePayment creates a payment for an invoice. Payments can only be added to a booked invoice.
func CreatePayment() {}

// DeletePayment deletes a payment from an invoice. Only booked invoices can have payments.
func DeletePayment() {}

// GetPayments gets the payments for an invoice
func GetPayments() {}

// GenerateCreditNote generates and saves a credit note draft of a given booked invoice.
func GenerateCreditNote() {}

// Save saves an invoice as draft.
func Save() {}

// Delete invoice. The invoice cannot be deleted if booked.
func Delete() {}

// List receives a list of invoices for the organization.
func List() {}
