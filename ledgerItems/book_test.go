package ledgeritems

import (
	"testing"

	dinero "github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func TestLedgerBooking(t *testing.T) {
	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	ledgerItems := []LedgerItem{
		{
			VoucherNumber:           1,
			AccountNumber:           1000,
			AccountVatCode:          "I25",
			Amount:                  -200,
			BalancingAccountNumber:  55000,
			BalancingAccountVatCode: "I25",
			Description:             "Hello from golang!",
			VoucherDate:             dinero.DateNow(),
		},
		{
			VoucherNumber:           1,
			AccountNumber:           1000,
			AccountVatCode:          "I25",
			Amount:                  -200,
			BalancingAccountNumber:  55000,
			BalancingAccountVatCode: "I25",
			Description:             "Hello from golang line 2",
			VoucherDate:             dinero.DateNow(),
		},
	}

	items, err := Create(c, ledgerItems)
	if err != nil {
		t.Errorf("We can't setup test correctly: %v", err)
	}

	if err := Book(c, items); err != nil {
		t.Errorf("Booking of ledger items failed: %v", err)
	}
}
