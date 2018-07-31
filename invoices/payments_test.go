package invoices

import (
	"testing"

	"github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestCreatePaymentOnInvoice_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	invoiceID := "f25ed5df-3fb4-49b5-98ce-31711238fc10"
	invoice, err := Get(c, invoiceID)
	if err != nil {
		t.Error("error occured while setting up the test: ", err)
	}

	params := CreatePaymentParams{
		Timestamp:            invoice.Timestamp,
		DepositAccountNumber: 55000,
		ExternalReference:    "GolangSDK",
		PaymentDate:          dinero.DateNow(),
		Description:          "Paid with the awesome GolangSDK",
		Amount:               1000,
	}

	if _, err := CreatePayment(c, invoiceID, params); err != nil {
		t.Error("Error occured while creating payment: ", err)
	}
}

func TestGetAllPaymentsOnInvoice_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	invoiceID := "f25ed5df-3fb4-49b5-98ce-31711238fc10"
	if _, err := GetPayments(c, invoiceID); err != nil {
		t.Error("We could not fetch the payments of the invoice: ", err)
	}
}

func TestDeletePaymentOnInvoice_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	invoiceID := "f25ed5df-3fb4-49b5-98ce-31711238fc10"
	invoice, err := Get(c, invoiceID)
	if err != nil {
		t.Error("error occured while setting up the test: ", err)
	}

	params := CreatePaymentParams{
		Timestamp:            invoice.Timestamp,
		DepositAccountNumber: 55000,
		ExternalReference:    "GolangSDK",
		PaymentDate:          dinero.DateNow(),
		Description:          "Paid with the awesome GolangSDK",
		Amount:               1000,
	}

	payment, err := CreatePayment(c, invoiceID, params)
	if err != nil {
		t.Error("error occured while setting up the test: ", err)
	}

	deleteParams := DeletePaymentParams{
		InvoiceID: invoice.ID,
		PaymentID: payment.ID,
		Timestamp: payment.Timestamp,
	}

	if _, err = DeletePayment(c, deleteParams); err != nil {
		t.Error("We could not delete the newly created payment: ", err)
	}
}
