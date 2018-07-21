package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// GenerateCreditnoteParams is the parameters required to generate a creditnote
type GenerateCreditnoteParams struct {
	InvoiceID string `json:"-"`
	Timestamp string `json:"timestamp"`
}

// GenerateCreditNote generates and saves a credit note draft of a given booked invoice.
func GenerateCreditNote(api dinero.API, params GenerateCreditnoteParams) (*dinero.TimestampResponse, error) {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/generate-creditnote", params.InvoiceID)

	var resp dinero.TimestampResponse
	if err := api.Call(http.MethodPost, route, &params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
