package entries

import (
	"testing"
	"time"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestGetEntries(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	now := time.Now()
	fromDate := dinero.NewDate(now.Year(), 1, 1)

	params := GetRequestParams{
		FromDate:     fromDate,
		ToDate:       dinero.DateNow(),
		IncludePrimo: false,
	}

	_, err := Get(c, params)
	if err != nil {
		t.Fatal(err)
	}
}
