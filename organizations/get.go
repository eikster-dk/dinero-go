package organizations

import (
	"fmt"
	"net/http"

	"github.com/eikc/dinero-go"
)

// OrganizationProperty is property that is apart of the
// organization data model and can be used to return only
// the needed properties of the organization model.
type OrganizationProperty string

// List of values for field filtering
const (
	ID        OrganizationProperty = "Id"
	Name      OrganizationProperty = "Name"
	IsPro     OrganizationProperty = "IsPro"
	IsVatFree OrganizationProperty = "IsVatFree"
)

// Organization represents the data model from the api.
type Organization struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	IsPro       bool   `json:"isPro,omitempty"`
	IsPayingPro bool   `json:"isPayingPro,omitempty"`
	IsVatFree   bool   `json:"isVatFree,omitempty"`
}

// Get is used to get organizations related to the authorized API key.
func Get(api dinero.API, fields ...OrganizationProperty) ([]Organization, error) {
	route := "v1/organizations"

	if len(fields) > 0 {
		fieldStrings := toString(fields)
		queryFields := dinero.BuildFieldsQuery(fieldStrings...)
		route = fmt.Sprintf("%s?fields=%s", route, queryFields)
	}

	var organizations []Organization

	if err := api.Call(http.MethodGet, route, nil, &organizations); err != nil {
		return nil, err
	}

	return organizations, nil
}

func toString(f []OrganizationProperty) []string {
	s := make([]string, len(f))

	for i, f := range f {
		s[i] = string(f)
	}

	return s
}
