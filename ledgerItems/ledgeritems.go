package ledgeritems

import (
	"github.com/eikc/dinero-go"
)

// LedgerItem is used as a create new ledger items in dinero
type LedgerItem struct {
	ID                      string      `json:"id,omitempty"`
	VoucherNumber           int         `json:"voucherNumber,omitempty"`
	AccountNumber           int         `json:"accountNumber,omitempty"`
	AccountVatCode          string      `json:"accountVatCode,omitempty"`
	Amount                  float64     `json:"amount,omitempty"`
	BalancingAccountNumber  int         `json:"balancingAccountNumber,omitempty"`
	BalancingAccountVatCode string      `json:"balancingAccountVatCode,omitempty"`
	Description             string      `json:"description,omitempty"`
	VoucherDate             dinero.Date `json:"voucherDate,omitempty"`
	FileGUID                string      `json:"fileGuid,omitempty"`
}

// LedgerItemModel is used to book a ledger item
type LedgerItemModel struct {
	ID      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
}

// LedgerItemStatusModel is the request item sent to dinero
// to receive the ledgeritem status
type LedgerItemStatusModel struct {
	ID string `json:"id,omitempty"`
}

// LedgerItemStatus is the response from dinero
// It contains the status of a booked ledger item.
type LedgerItemStatus struct {
	ID            string `json:"id,omitempty"`
	Status        string `json:"status,omitempty"`
	VoucherNumber int    `json:"voucherNumber,omitempty"`
	Version       string `json:"version,omitempty"`
}
