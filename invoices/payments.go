package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// CreatePaymentParams is the model needed to create a payment
type CreatePaymentParams struct {
	Timestamp               string      `json:"timestamp,omitempty"`
	DepositAccountNumber    int         `json:"depositAccountNumber,omitempty"`
	ExternalReference       string      `json:"externalReference,omitempty"`
	PaymentDate             dinero.Date `json:"paymentDate,omitempty"`
	Description             string      `json:"description,omitempty"`
	Amount                  float64     `json:"amount,omitempty"`
	AmountInForeignCurrency float64     `json:"amountInForeignCurrency,omitempty"`
}

// InvoicePayments presents a collection of payments and the remaining,paid, and total of the invoice
type InvoicePayments struct {
	Payments                         []InvoicePayment `json:"payments"`
	RemainingAmount                  float64          `json:"remainingAmount"`
	PaidAmount                       float64          `json:"paidAmount"`
	InvoiceTotalInclReminderExpenses float64          `json:"invoiceTotalIncludingReminderExpenses"`
}

// InvoicePayment represents a payment in dinero
type InvoicePayment struct {
	ID                      string      `json:"guid"`
	DepositAccountNumber    int         `json:"depositAccountNumber"`
	ExternalReference       string      `json:"externalReference"`
	PaymentDate             dinero.Date `json:"paymentDate"`
	Description             string      `json:"description"`
	Amount                  float64     `json:"amount"`
	AmountInForeignCurrency float64     `json:"amountInForeignCurrency"`
}

// DeletePaymentParams are the params needed to delete a payment from an invoice
type DeletePaymentParams struct {
	InvoiceID string `json:"-"`
	PaymentID string `json:"-"`
	Timestamp string `json:"timestamp,omitempty"`
}

// CreatePayment creates a payment for an invoice. Payments can only be added to a booked invoice.
func CreatePayment(api dinero.API, id string, params CreatePaymentParams) (*dinero.TimestampResponse, error) {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/payments", id)

	var timestampResp dinero.TimestampResponse
	if err := api.Call(http.MethodPost, route, &params, &timestampResp); err != nil {
		return nil, err
	}

	return &timestampResp, nil
}

// GetPayments gets the payments for an invoice
func GetPayments(api dinero.API, id string) (*InvoicePayments, error) {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/payments", id)

	var paymentResp InvoicePayments
	if err := api.Call(http.MethodGet, route, nil, &paymentResp); err != nil {
		return nil, err
	}

	return &paymentResp, nil
}

// DeletePayment deletes a payment from an invoice. Only booked invoices can have payments.
func DeletePayment(api dinero.API, params DeletePaymentParams) (*dinero.TimestampResponse, error) {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/payments/%v", params.InvoiceID, params.PaymentID)

	var resp dinero.TimestampResponse
	if err := api.Call(http.MethodDelete, route, &params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
