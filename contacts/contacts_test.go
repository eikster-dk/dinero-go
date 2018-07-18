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

func TestAddPrivateContact_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	params := ContactParams{
		Name:                         "Hello awesome",
		ExternalReference:            "external",
		AttPerson:                    "",
		City:                         "city",
		EanNumber:                    "",
		Email:                        "test@test.dk",
		PaymentConditionType:         Netto,
		Phone:                        "88 88 88 88",
		Street:                       "street",
		VatNumber:                    "",
		Webpage:                      "http://awesome.dk",
		ZipCode:                      "2700",
		CountryKey:                   "DK",
		IsPerson:                     true,
		PaymentConditionNumberOfDays: 10,
	}
	_, err := Add(c, params)

	if err != nil {
		t.Errorf("We can not create a contact: %v", err)
	}
}

func TestUpdateContact_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	id := "3e389a20-d206-4c4b-acff-3cff102db328"

	params := ContactParams{
		Name:       "testing awesome update",
		CountryKey: "DK",
		IsPerson:   false,
	}

	if err := Update(c, id, params); err != nil {
		t.Errorf("We could not update a contact: %v", err)
	}
}

func TestDeleteAndRestoreContact_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.NewClient(key, secret)
	c.Authorize(apiKey, organizationID)

	id := "3617dafe-c066-450c-8cbc-8d082a9f3907"

	if err := Delete(c, id); err != nil {
		t.Errorf("We could not delete the contact correctly: %v", err)
		t.Fail()
	}

	if err := Restore(c, id); err != nil {
		t.Errorf("We could not restore the contact correctly: %v", err)
		t.Fail()
	}
}
