package vattypes

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestGetVatTypes(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	if _, err := Get(c); err != nil {
		t.Error("We could not get vattypes, error: ", err)
	}
}
