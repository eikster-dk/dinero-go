package accountingYears

import (
	"net/http"

	"github.com/eikc/dinero-go"
)

// AccountingYear represent the model from the dinero accounting year API
type AccountingYear struct {
	Name     string
	FromDate string
	ToDate   string
}

// Get a list of an organizations accounting years.
func Get(api dinero.API) ([]AccountingYear, error) {
	route := "v1/{organizationID}/accountingYears"
	var accountingYears []AccountingYear
	if err := api.Call(http.MethodGet, route, nil, &accountingYears); err != nil {
		return nil, err
	}

	return accountingYears, nil
}
