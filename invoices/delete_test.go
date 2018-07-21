package invoices

import (
	"testing"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/dinerotest"
)

func TestInvoiceCanBeDeleted_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

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
		Description:       "deleting from golangSDK",
		Comment:           "This is a comment",
		Date:              dinero.DateNow(),
		ProductLines: []InvoiceLine{
			InvoiceLine{
				BaseAmountValue: 10000,
				Quantity:        1,
				AccountNumber:   1000,
				Description:     "deleting from golangSDK",
				LineType:        "Product",
				Unit:            "hours",
			},
		},
		Address: "A secret place on earth!",
	}

	invoice, err := Save(c, params)
	if err != nil {
		t.Error("Saving invoice introduced an error: ", err)
	}

	deleteParams := DeleteInvoiceParams{
		InvoiceID: invoice.ID,
		Timestamp: invoice.Timestamp,
	}

	if err := Delete(c, deleteParams); err != nil {
		t.Error("failed to delete draft invoice: ", err)
	}
}
