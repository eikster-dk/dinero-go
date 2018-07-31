package invoices

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestSendingInvoiceAsEmail_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

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
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

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
