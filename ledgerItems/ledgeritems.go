package ledgerItems

import (
	"github.com/eikc/dinero-go"
)

// LedgerItem is used as a create new ledger items in dinero
type LedgerItem struct {
	Id                      string      `json:"id,omitempty"`
	VoucherNumber           int         `json:"voucherNumber,omitempty"`
	AccountNumber           int         `json:"accountNumber,omitempty"`
	AccountVatCode          string      `json:"accountVatCode,omitempty"`
	Amount                  float64     `json:"amount,omitempty"`
	BalancingAccountNumber  int         `json:"balancingAccountNumber,omitempty"`
	BalancingAccountVatCode string      `json:"balancingAccountVatCode,omitempty"`
	Description             string      `json:"description,omitempty"`
	VoucherDate             dinero.Date `json:"voucherDate,omitempty"`
	FileGuid                string      `json:"fileGuid,omitempty"`
}

// LedgerItemModel is used to book a ledger item
type LedgerItemModel struct {
	ID      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
}
