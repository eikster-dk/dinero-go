package accountingYears

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestAccountingYear_GET_integration(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	_, err := Get(c)
	if err != nil {
		t.Errorf("Error getting accounting years: %v", err)
	}
}
