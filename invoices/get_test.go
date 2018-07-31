package invoices

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestGetInvoice_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	if _, err := Get(c, "f25ed5df-3fb4-49b5-98ce-31711238fc10"); err != nil {
		t.Error("We could not get an invoice, error: ", err)
	}
}
