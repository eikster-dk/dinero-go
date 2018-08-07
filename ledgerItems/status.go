package ledgeritems

import (
	"net/http"

	dinero "github.com/eikc/dinero-go"
)

// Status is used to get the status of a given set of ledger items
func Status(api dinero.API, items []LedgerItemStatusModel) ([]LedgerItemStatus, error) {
	route := "v1/{organizationID}/ledgeritems/status"

	var statusItems []LedgerItemStatus
	if err := api.Call(http.MethodPost, route, items, &statusItems); err != nil {
		return nil, err
	}

	return statusItems, nil
}
