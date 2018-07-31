package accounts

import (
	"testing"

	"github.com/eikc/dinero-go/internal"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

func Test_getRouteWithFields(t *testing.T) {
	t.Parallel()

	type args struct {
		route        string
		wantedFields []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test route for accounts with nil fields",
			args{"v1/{organizationId}/accounts/entry", nil},
			"v1/{organizationId}/accounts/entry",
		},
		{
			"test routes for account with name field",
			args{"v1/{organizationId}/accounts/entry", []string{Name}},
			"v1/{organizationId}/accounts/entry?fields=Name",
		},
		{
			"test routes for account with name and number field",
			args{"v1/{organizationId}/accounts/entry", []string{Name, Number}},
			"v1/{organizationId}/accounts/entry?fields=Name,AccountNumber",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRouteWithFields(tt.args.route, tt.args.wantedFields); got != tt.want {
				t.Errorf("getRouteWithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDeposits_integration(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	_, err := GetDeposits(c)
	if err != nil {
		t.Errorf("got an error while calling the deposits endpoint: %v", err)
	}
}

func TestGetDeposits_OnlyReturnNumberField_integration(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	deposits, err := GetDeposits(c, Number)
	if err != nil {
		t.Errorf("Got an error whil calling the deposits endpoint %v", err)
	}

	first := deposits[0]

	var defaultString string
	var defaultInt int
	if first.Name != defaultString && first.Number == defaultInt {
		t.Errorf("the returned accounts are including the Name, first is: %v", first)
	}
}

func TestGetAccounts_integration(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	_, err := GetAccounts(c)
	if err != nil {
		t.Errorf("Got an error while fetching accounts: %v", err)
	}
}

func TestGetAccountsWithOnlyName_Integration(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip(dinerotest.IntegrationTestText)
	}

	c := internal.GetClient()

	accounts, err := GetAccounts(c, Name)
	if err != nil {
		t.Errorf("Got an error while fetching accounts with name field only: %v", err)
	}

	first := accounts[0]

	var defaultString string
	var defaultNumber int

	if first.Name == defaultString && first.Number != defaultNumber {
		t.Errorf("The returned account is not what is expected: %v", first)
	}
}
