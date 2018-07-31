package invoices

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestInvoicesList_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	var params ListParams
	if _, err := List(c, params); err != nil {
		t.Error("error occured while getting a list of invoices: ", err)
	}
}
