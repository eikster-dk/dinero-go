package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// Get invoice as json or pdf. Define the Accept header of your request to either
// 'application/json' or 'application/octet-stream'. PDF's can only be generated from booked invoices.
func Get(api dinero.API, id string) (*Invoice, error) {
	route := fmt.Sprint("v1/{organizationID}/invoices/", id)

	var invoice Invoice
	if err := api.Call(http.MethodGet, route, nil, &invoice); err != nil {
		return nil, err
	}

	return &invoice, nil
}
