package contacts

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/eikc/dinero-go"
)

// Fields Params
const (
	ContactGUID                  = "ContactGuid"
	CreatedAt                    = "CreatedAt"
	UpdatedAt                    = "UpdatedAt"
	DeletedAt                    = "DeletedAt"
	IsDebitor                    = "IsDebitor"
	IsCreditor                   = "IsCreditor"
	ExternalReference            = "ExternalReference"
	Name                         = "Name"
	Street                       = "Street"
	ZipCode                      = "Zipcode"
	City                         = "City"
	CountryKey                   = "CountryKey"
	Phone                        = "Phone"
	Email                        = "Email"
	Webpage                      = "Webpage"
	AttPerson                    = "AttPerson"
	VatNumber                    = "VatNumber"
	EanNumber                    = "EanNumber"
	PaymentConditionType         = "PaymentConditionType"
	PaymentConditionNumberOfDays = "PaymentConditionNumberOfDays"
	IsPerson                     = "IsPerson"
)

// Contact represents a contact from the dinero api
type Contact struct {
	ContactGUID                  string    `json:"ContactGuid"`
	CreatedAt                    time.Time `json:"CreatedAt"`
	UpdatedAt                    time.Time `json:"UpdatedAt"`
	DeletedAt                    time.Time `json:"DeletedAt"`
	IsDebitor                    bool      `json:"IsDebitor"`
	IsCreditor                   bool      `json:"IsCreditor"`
	ExternalReference            string    `json:"ExternalReference"`
	Name                         string    `json:"Name"`
	Street                       string    `json:"Street"`
	ZipCode                      string    `json:"ZipCode"`
	City                         string    `json:"City"`
	CountryKey                   string    `json:"CountryKey"`
	Phone                        string    `json:"Phone"`
	Email                        string    `json:"Email"`
	Webpage                      string    `json:"Webpage"`
	AttPerson                    string    `json:"AttPerson"`
	VatNumber                    string    `json:"VatNumber"`
	EanNumber                    string    `json:"EanNumber"`
	PaymentConditionType         string    `json:"PaymentConditionType"`
	PaymentConditionNumberOfDays int       `json:"PaymentConditionNumberOfDays"`
	IsPerson                     bool      `json:"IsPerson"`
}

// ContactList returns the paginated result of the contacts
type ContactList struct {
	Collection []Contact               `json:"Collection"`
	Pagination dinero.PaginationResult `json:"Pagination"`
}

// ListParams are parameter options for getting a list of contacts
// See more: https://api.dinero.dk/v1/docs/Api/GET-v1-organizationId-contacts_fields_queryFilter_changesSince_deletedOnly_page_pageSize
type ListParams struct {
	Fields      []string
	QueryFilter []string
	ChangeSince time.Time
	DeletedOnly bool
	Page        int
	PageSize    int
}

// Restore a deleted contact from the given organization
func Restore(api dinero.API) {
}

// List retrieves a list of contacts for the organization order by UpdatedAt
// If fields are not specified then it defaults to name,contactGuid
func List(api dinero.API, params ListParams) (*ContactList, error) {
	route := "v1/{organizationID}/contacts"
	q := url.Values{}

	if params.Fields != nil {
		fields := dinero.BuildFieldsQuery(params.Fields...)
		q.Add("fields", fields)
	}

	if params.QueryFilter != nil {
		filter := dinero.BuildFieldsQuery(params.QueryFilter...)
		q.Add("queryFilter", filter)
	}

	var defaultTime time.Time
	if params.ChangeSince != defaultTime {
		q.Add("changesSince", params.ChangeSince.Format(time.RFC3339))
	}

	q.Add("deletedOnly", strconv.FormatBool(params.DeletedOnly))
	q.Add("page", strconv.FormatInt(int64(params.Page), 10))
	q.Add("pageSize", strconv.FormatInt(int64(params.PageSize), 10))

	encodedQueryString := q.Encode()
	route = fmt.Sprintf("%v?%v", route, encodedQueryString)

	var contacts ContactList
	if err := api.Call(http.MethodGet, route, nil, &contacts); err != nil {
		return nil, err
	}

	return &contacts, nil
}

// Get retrieves contact information for the contact with the given id
func Get() {
}

// Add a new contact to the organization
func Add() {
}

// Update an existing contact
func Update() {}

// Delete a contact from the given organization
func Delete() {}
