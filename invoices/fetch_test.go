package invoices

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestFetchInvoiceTotal_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	params := FetchParams{
		PaymentConditionNumberOfDays: 10,
		PaymentConditionType:         "Netto",
		ContactID:                    "3e389a20-d206-4c4b-acff-3cff102db328",
		ShowLinesInclVat:             false,
		Currency:                     "DKK",
		ProductLines: []InvoiceLine{
			{
				BaseAmountValue: 1000,
				Description:     "AWesome golang work",
				Quantity:        1,
				Unit:            "hours",
				LineType:        "product",
				AccountNumber:   1000,
			},
		},
		Date: "2018-02-28",
	}
	_, err := Fetch(c, params)
	if err != nil {
		t.Errorf("We could not fetch invoice totals: %v", err)
	}
}
