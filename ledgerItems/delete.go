package ledgeritems

import (
	"net/http"

	dinero "github.com/eikc/dinero-go"
)

// Delete is used to delete ledgeritems from an existing ledger in dinero.
func Delete(api dinero.API, ledgerItems []LedgerItemModel) error {
	route := "v1/{organizationID}/ledgeritems/delete"

	return api.Call(http.MethodDelete, route, ledgerItems, nil)
}
