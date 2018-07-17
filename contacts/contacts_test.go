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

func TestGetContact_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	id := "3e389a20-d206-4c4b-acff-3cff102db328"

	_, err := Get(c, id)
	if err != nil {
		t.Errorf("Failed getting contact by ID: %v", err)
	}
}
