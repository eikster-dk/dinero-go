package accountingYears

import (
	"testing"

	"github.com/eikc/dinero-go/dinerotest"

	"github.com/eikc/dinero-go"
)

func TestAccountingYear_GET_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	_, err := Get(c)
	if err != nil {
		t.Errorf("Error getting accounting years: %v", err)
	}
}
