package invoices

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// SendInvoice are the model of sending an invoice to a given user
type SendInvoice struct {
	Timestamp              string `json:"timestamp,omitempty"`
	Sender                 string `json:"sender,omitempty"`
	CCToEmail              string `json:"ccToEmail,omitempty"`
	Receiver               string `json:"receiver,omitempty"`
	Subject                string `json:"subject,omitempty"`
	Message                string `json:"message,omitempty"`
	AddVoucherAsAttachment bool   `json:"addVoucherAsAttachment,omitempty"`
}

// EmailSent is the result of the recipients of the email that was sent to
type EmailSent struct {
	Recipients []struct {
		Email string `json:"email"`
	}
}

type timestampModel struct {
	Timestamp string `json:"timestamp"`
}

// SendEmail sends an email with link to a public version of the invoice where it can be printed or downloaded as a pdf.
func SendEmail(api dinero.API, id string, params SendInvoice) (*EmailSent, error) {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/email", id)

	var emailSent EmailSent
	if err := api.Call(http.MethodPost, route, &params, &emailSent); err != nil {
		return nil, err
	}

	return &emailSent, nil
}

// SendPreReminder sends a pre reminder email with link to a public version of the invoice where it
// can be printed or downloaded as a pdf. The invoice needs to be overdue to send the reminder.
// A pre-reminder is a mail reminding the customer, that the invoice is overdue.
// This will not cause a reminder to be created in Dinero, this is only a mailout.
func SendPreReminder(api dinero.API, id string, params SendInvoice) error {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/email/pre-reminder", id)

	return api.Call(http.MethodPost, route, nil, nil)
}

// SendEInvoice sends an e-invoice to an EAN customer
func SendEInvoice(api dinero.API, id, timestamp string) error {
	route := fmt.Sprintf("v1/{organizationID}/invoices/%v/e-invoice", id)

	body := timestampModel{timestamp}

	return api.Call(http.MethodPost, route, body, nil)
}
