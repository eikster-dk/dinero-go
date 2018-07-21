package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// DeleteInvoiceParams are the needed params to delete the invoice
type DeleteInvoiceParams struct {
	InvoiceID string `json:"-"`
	Timestamp string `json:"timestamp"`
}

// Delete invoice. The invoice cannot be deleted if booked.
func Delete(api dinero.API, params DeleteInvoiceParams) error {
	route := fmt.Sprint("v1/{organizationID}/invoices/", params.InvoiceID)

	return api.Call(http.MethodDelete, route, nil, nil)
}
