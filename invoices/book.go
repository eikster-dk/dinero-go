package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

type bookInvoice struct {
	Timestamp string `json:"timestamp"`
}

// InvoiceBooked is the response containing the timestamp of the invoice when a booking is succesfull
type InvoiceBooked struct {
	ID        string `json:"guid"`
	Timestamp string `json:"timeStamp"`
}

// Book books a given invoice
func Book(api dinero.API, id, timestamp string) (*InvoiceBooked, error) {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/book", id)
	body := bookInvoice{timestamp}

	var invoiceBooked InvoiceBooked
	if err := api.Call(http.MethodPost, route, body, &invoiceBooked); err != nil {
		return nil, err
	}

	return &invoiceBooked, nil
}
