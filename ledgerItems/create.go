package ledgerItems

import (
	"net/http"

	"github.com/eikc/dinero-go"
)

func Create(api dinero.API, items []LedgerItem) ([]LedgerItemModel, error) {
	route := "v1.1/{organizationID}/ledgeritems"

	var ledgerItems []LedgerItemModel
	if err := api.Call(http.MethodPost, route, items, &ledgerItems); err != nil {
		return nil, err
	}

	return ledgerItems, nil
}
