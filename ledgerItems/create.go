package ledgeritems

import (
	"net/http"

	"github.com/eikc/dinero-go"
)

// Create will create ledgeritems in dinero
func Create(api dinero.API, items []LedgerItem) ([]LedgerItemModel, error) {
	route := "v1.1/{organizationID}/ledgeritems"

	var ledgerItems []LedgerItemModel
	if err := api.Call(http.MethodPost, route, items, &ledgerItems); err != nil {
		return nil, err
	}

	return ledgerItems, nil
}
