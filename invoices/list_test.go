package invoices

import (
	"testing"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/dinerotest"
)

func TestInvoicesList_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	var params ListParams
	if _, err := List(c, params); err != nil {
		t.Error("error occured while getting a list of invoices: ", err)
	}
}
