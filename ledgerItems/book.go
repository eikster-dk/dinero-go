package ledgeritems

import (
	"net/http"

	"github.com/eikc/dinero-go"
)

// Book function is used to book a set of ledgeritems which are given as paramenter
func Book(api dinero.API, items []LedgerItemModel) error {
	route := "v1/{organizationID}/ledgeritems/book"

	return api.Call(http.MethodPost, route, items, nil)
}
