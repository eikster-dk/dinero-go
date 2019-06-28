package invoices

import (
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
	"testing"
)

func TestGetInvoiceTemplates(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	if _ ,err := GetInvoiceTemplates(c); err != nil {
		t.Error("error occured while getting invoice templates: ", err)
	}
}