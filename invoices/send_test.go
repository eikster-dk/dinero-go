package invoices

import (
	"testing"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/dinerotest"
)

func TestSendingInvoiceAsEmail_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	invoiceID := "961532ac-be75-4d97-afeb-74ed724334dc"
	emailParams := SendInvoice{
		Sender:  "GolangSDK@gmail.com",
		Subject: "Golang SDK sends an invoice!",
	}

	if _, err := SendEmail(c, invoiceID, emailParams); err != nil {
		t.Error("We can't send an email from the SDK: ", err)
	}
}

func TestSendingPreReminder_integration(t *testing.T) {
	t.Skip("Skipped since we have no invoice in correct state to be tested on")

	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	invoiceID := "961532ac-be75-4d97-afeb-74ed724334dc"
	emailParams := SendInvoice{
		Sender:  "GolangSDK@gmail.com",
		Subject: "Golang SDK sends an invoice!",
	}

	if err := SendPreReminder(c, invoiceID, emailParams); err != nil {
		t.Error("We can't send an email from the SDK: ", err)
	}
}

func TestSendEInvoice_integration(t *testing.T) {
	t.Skip("We have no way to test the EAN invoices.. maybe it will come in the future")
}
