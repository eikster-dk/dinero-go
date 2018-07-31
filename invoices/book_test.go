package invoices

import (
	"testing"

	"github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestBookInvoice_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	paramsForInvoiceToBook := CreateInvoice{
		PaymentConditions: dinero.PaymentConditions{
			PaymentConditionNumberOfDays: 8,
			PaymentConditionType:         dinero.Netto,
		},
		ContactID:         "3e389a20-d206-4c4b-acff-3cff102db328",
		ShowLinesInclVat:  false,
		Currency:          "DKK",
		Language:          "da-DK",
		ExternalReference: "golangSDK",
		Description:       "booking from GolangSDK",
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

	resp, err := Save(c, paramsForInvoiceToBook)
	if err != nil {
		t.Error("we could not add an invoice to book, err: ", err)
	}

	if _, err := Book(c, resp.ID, resp.Timestamp); err != nil {
		t.Error("Booking invoice failed, err: ", err)
	}
}
