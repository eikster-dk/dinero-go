package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// UpdateInvoice is the model for updating an existing invoice
type UpdateInvoice struct {
	dinero.PaymentConditions
	Timestamp         string        `json:"timestamp"`
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

// Update an existing invoice. The invoice cannot be updated if booked.
// endpoint used is version: 1.2
func Update(api dinero.API, id string, params UpdateInvoice) (*dinero.TimestampResponse, error) {
	route := fmt.Sprint("v1.2/{organizationID}/invoices/", id)

	var invoiceUpdated dinero.TimestampResponse
	if err := api.Call(http.MethodPut, route, &params, &invoiceUpdated); err != nil {
		return nil, err
	}

	return &invoiceUpdated, nil
}
