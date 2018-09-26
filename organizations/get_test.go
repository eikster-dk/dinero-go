package organizations

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestGetOrganizations(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	if _, err := Get(c); err != nil {
		t.Error("We could not get organizations, error: ", err)
	}
}

func TestGetOrganizationWithFields(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	if _, err := Get(c, Name, IsPro, IsVatFree); err != nil {
		t.Error("We could not get organizations, error: ", err)
	}
}
