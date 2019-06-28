package invoices

import (
	"net/http"

	"github.com/eikc/dinero-go"
)

// CreateInvoice represents the model for creating an invoice
type CreateInvoice struct {
	dinero.PaymentConditions
	ContactID         string        `json:"contactGuid,omitempty"`
	ShowLinesInclVat  bool          `json:"showLinesInclVat,omitempty"`
	Currency          string        `json:"currency,omitempty"`
	Language          string        `json:"language,omitempty"`
	ExternalReference string        `json:"externalReference,omitempty"`
	Description       string        `json:"description,omitempty"`
	Comment           string        `json:"comment,omitempty"`
	InvoiceTemplateId string        `json:"invoiceTemplateId,omitempty"`
	Date              dinero.Date   `json:"date,omitempty"`
	ProductLines      []InvoiceLine `json:"productLines,omitempty"`
	Address           string        `json:"address,omitempty"`
}

// Save saves an invoice as draft.
func Save(api dinero.API, params CreateInvoice) (*dinero.TimestampResponse, error) {
	route := "v1/{organizationID}/invoices"

	var invoiceCreated dinero.TimestampResponse
	if err := api.Call(http.MethodPost, route, &params, &invoiceCreated); err != nil {
		return nil, err
	}

	return &invoiceCreated, nil
}
