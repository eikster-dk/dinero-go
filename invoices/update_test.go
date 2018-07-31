package invoices

import (
	"testing"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestUpdateInvoice_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	id := "0869dd82-7cac-445a-b5db-1fa2037054f4"
	invoice, err := Get(c, id)
	if err != nil {
		t.Error("we could not setup the test, issues getting the invoice: ", err)
	}

	params := UpdateInvoice{
		PaymentConditions: dinero.PaymentConditions{
			PaymentConditionNumberOfDays: 8,
			PaymentConditionType:         dinero.Netto,
		},
		Timestamp:         invoice.Timestamp,
		ContactID:         "3e389a20-d206-4c4b-acff-3cff102db328",
		ShowLinesInclVat:  false,
		Currency:          "DKK",
		Language:          "da-DK",
		ExternalReference: "golangSDK",
		Description:       "Update from Golang SDK",
		Comment:           "This is a comment",
		Date:              dinero.DateNow(),
		ProductLines: []InvoiceLine{
			InvoiceLine{
				BaseAmountValue: 20000,
				Quantity:        1,
				AccountNumber:   1000,
				Description:     "Update from Golang SDK",
				LineType:        "Product",
				Unit:            "hours",
			},
		},
		Address: "A secret place on earth!",
	}

	if _, err := Update(c, id, params); err != nil {
		t.Error("Can't update invoice, err: ", err)
	}
}
