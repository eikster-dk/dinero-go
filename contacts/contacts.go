package contacts

import (
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

// ListParams are parameter options for getting a list of contacts
// Fields: A comma separated list of fields to include in the response.
// Possible values are Name, ContactGuid, ExternalReference, IsPerson, Street, Zipcode, City, CountryKey, Phone, Email,
// Webpage, AttPerson, VatNumber, EanNumber, PaymentConditionType, PaymentConditionNumberOfDays, CreatedAt, UpdatedAt and DeletedAt.
// Notice that it's not case sensitive, the property name will be returned the way you request it.
// If left empty it defaults to name and contactGuid (notice small start letter).
// QueryFilter: Filter specific for contacts. Filtering can be applied to following fields:
// ExternalReference,Name, Email, VatNumber, EanNumber, IsPerson. See API documentation for filtering format.
// If left empty no filtering is applied.
// ChangeSince: [Generic Filter Option] Only return contacts that was created, deleted or updated at or after given time.
// If left empty, this filter will not be applied, and contacts will be returned regardless of change history.
// The time must be UTC and in the format 'YYYY-MM-DDTHH:mm:ssZ' example: '2015-08-18T06:36:22Z'.
// DeletedOnly: [Generic Filter Option] Only select deleted contacts. If left empty, will default to false.
// Page: The 0-based page number
// PageSize: The maximum number of items to include in a page. Maximum 1000.
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
func List(api dinero.API, params ListParams) {

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
