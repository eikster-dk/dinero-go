package ledgeritems

import (
	"testing"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestStatusOfLedgeritems(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	ledgerItems := []LedgerItem{
		{
			VoucherNumber:           1,
			AccountNumber:           1000,
			AccountVatCode:          "I25",
			Amount:                  200,
			BalancingAccountNumber:  55000,
			BalancingAccountVatCode: "I25",
			Description:             "Hello from golang!",
			VoucherDate:             dinero.DateNow(),
		},
	}

	item, err := Create(c, ledgerItems)
	if err != nil {
		t.Errorf("Could not setup test in dinero: %v", err)
	}

	if err := Book(c, item); err != nil {
		t.Errorf("Could not setup test, booking of ledger failed: %v", err)
	}

	statusItems := []LedgerItemStatusModel{
		{
			ID: item[0].ID,
		},
	}

	_, err = Status(c, statusItems)
	if err != nil {
		t.Fatal(err)
	}
}
