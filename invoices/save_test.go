package invoices

import (
	"testing"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestSaveInvoice_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	params := CreateInvoice{
		PaymentConditions: dinero.PaymentConditions{
			PaymentConditionNumberOfDays: 8,
			PaymentConditionType:         dinero.Netto,
		},
		ContactID:         "3e389a20-d206-4c4b-acff-3cff102db328",
		ShowLinesInclVat:  false,
		Currency:          "DKK",
		Language:          "da-DK",
		ExternalReference: "golangSDK",
		Description:       "Hello fro golangSDK",
		Comment:           "This is a comment",
		Date:              dinero.DateNow(),
		ProductLines: []InvoiceLine{
			InvoiceLine{
				BaseAmountValue: 10000,
				Quantity:        1,
				AccountNumber:   1000,
				Description:     "Awesomeness",
				LineType:        "Product",
				Unit:            "hours",
			},
		},
		Address: "A secret place on earth!",
	}
	if _, err := Save(c, params); err != nil {
		t.Error("Saving invoice introduced an error: ", err)
	}
}
