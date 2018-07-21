package invoices

import (
	"net/http"

	"github.com/eikc/dinero-go"
)

// FetchParams are the values required to get invoice total, line sums etc.
type FetchParams struct {
	PaymentConditionNumberOfDays int           `json:"paymentConditionNumberOfDays"`
	PaymentConditionType         string        `json:"paymentConditionType"`
	ContactID                    string        `json:"contactGuid"`
	ShowLinesInclVat             bool          `json:"showLinesInclVat"`
	Currency                     string        `json:"currency"`
	Language                     string        `json:"language"`
	ExternalReference            string        `json:"externalReference"`
	Description                  string        `json:"description"`
	Comment                      string        `json:"comment"`
	Date                         string        `json:"date"`
	ProductLines                 []InvoiceLine `json:"productLines"`
	Address                      string        `json:"address"`
}

// FetchResponse is the response from the fetch call
type FetchResponse struct {
	PaymentConditionNumberOfDays int           `json:"paymentConditionNumberOfDays"`
	PaymentConditionType         string        `json:"paymentConditionType"`
	ContactID                    string        `json:"contactGuid"`
	Number                       int           `json:"number"`
	ContactName                  string        `json:"contactName"`
	ShowLinesInclVat             bool          `json:"showLinesInclVat"`
	TotalExclVat                 float64       `json:"totalExclVat"`
	TotalVatableAmount           float64       `json:"totalVatableAmount"`
	TotalInclVat                 float64       `json:"totalInclVat"`
	TotalNonVatableAmount        float64       `json:"totalNonVatableAmount"`
	TotalVat                     float64       `json:"totalVat"`
	TotalLines                   []TotalLine   `json:"totalLines"`
	Currency                     string        `json:"currency"`
	Language                     string        `json:"language"`
	ExternalReference            string        `json:"externalReference"`
	Description                  string        `json:"description"`
	Comment                      string        `json:"comment"`
	Date                         string        `json:"date"`
	ProductLines                 []InvoiceLine `json:"productLines"`
	Address                      string        `json:"address"`
}

// Fetch a invoice to get total, line sums and payment date calculations.
func Fetch(api dinero.API, params FetchParams) (*FetchResponse, error) {
	route := "v1/{organizationID}/invoices/fetch"

	var resp FetchResponse
	if err := api.Call(http.MethodPost, route, &params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
