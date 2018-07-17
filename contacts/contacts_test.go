package contacts

import (
	"testing"

	"github.com/eikc/dinero-go"

	"github.com/eikc/dinero-go/dinerotest"
)

func TestListContacts_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	params := ListParams{

		Page:     0,
		PageSize: 100,
	}

	_, err := List(c, params)

	if err != nil {
		t.Errorf("Failed getting the contact list: %v", err)
	}
}
