package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

type bookInvoice struct {
	Timestamp string `json:"timestamp"`
}

// Book books a given invoice
func Book(api dinero.API, id, timestamp string) (*dinero.TimestampResponse, error) {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/book", id)
	body := bookInvoice{timestamp}

	var invoiceBooked dinero.TimestampResponse
	if err := api.Call(http.MethodPost, route, body, &invoiceBooked); err != nil {
		return nil, err
	}

	return &invoiceBooked, nil
}
