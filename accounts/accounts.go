package accounts

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// Params for filtering what is returned for both deposits and accounts
const (
	Name   = "Name"
	Number = "AccountNumber"
)

// Deposit represents a deposit in dinero regnskab
type Deposit struct {
	Number int    `json:"AccountName"`
	Name   string `json:"Name"`
}

// Account represents an account in dinero regnskab
type Account struct {
	Number int    `json:"AccountName"`
	Name   string `json:"Name"`
}

// GetAccounts Gets the list of entry accounts for the organization
func GetAccounts(api dinero.API, wantedFields ...string) ([]Account, error) {
	route := getRouteWithFields("v1/{organizationID}/accounts/entry", wantedFields)

	var accounts []Account
	if err := api.Call(http.MethodGet, route, nil, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

// GetDeposits gets the list of deposit accounts for the organization
func GetDeposits(api dinero.API, wantedFields ...string) ([]Deposit, error) {
	route := getRouteWithFields("v1/{organizationID}/accounts/deposit", wantedFields)

	var deposits []Deposit
	if err := api.Call(http.MethodGet, route, nil, &deposits); err != nil {
		return nil, err
	}

	return deposits, nil
}

func getRouteWithFields(route string, wantedFields []string) string {
	if wantedFields != nil {
		fieldQuery := dinero.BuildFieldsQuery(wantedFields...)
		route = fmt.Sprintf("%v?fields=%v", route, fieldQuery)
	}

	return route
}
