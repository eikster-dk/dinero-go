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
	Date              dinero.Date   `json:"date,omitempty"`
	ProductLines      []InvoiceLine `json:"productLines,omitempty"`
	Address           string        `json:"address,omitempty"`
}

// InvoiceCreated returns the response from dinero when a invoice is created successfully
type InvoiceCreated struct {
	ID        string `json:"Guid"`
	Timestamp string `json:"timeStamp"`
}

// Save saves an invoice as draft.
func Save(api dinero.API, params CreateInvoice) (*InvoiceCreated, error) {
	route := "v1/{organizationID}/invoices"

	var invoiceCreated InvoiceCreated
	if err := api.Call(http.MethodPost, route, &params, &invoiceCreated); err != nil {
		return nil, err
	}

	return &invoiceCreated, nil
}
