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
	ContactGUID                  string      `json:"ContactGuid,omitempty"`
	CreatedAt                    dinero.Time `json:"CreatedAt,omitempty"`
	UpdatedAt                    dinero.Time `json:"UpdatedAt,omitempty"`
	DeletedAt                    dinero.Time `json:"DeletedAt,omitempty"`
	IsDebitor                    bool        `json:"IsDebitor,omitempty"`
	IsCreditor                   bool        `json:"IsCreditor,omitempty"`
	ExternalReference            string      `json:"ExternalReference,omitempty"`
	Name                         string      `json:"Name,omitempty"`
	Street                       string      `json:"Street,omitempty"`
	ZipCode                      string      `json:"ZipCode,omitempty"`
	City                         string      `json:"City,omitempty"`
	CountryKey                   string      `json:"CountryKey,omitempty"`
	Phone                        string      `json:"Phone,omitempty"`
	Email                        string      `json:"Email,omitempty"`
	Webpage                      string      `json:"Webpage,omitempty"`
	AttPerson                    string      `json:"AttPerson,omitempty"`
	VatNumber                    string      `json:"VatNumber,omitempty"`
	EanNumber                    string      `json:"EanNumber,omitempty"`
	PaymentConditionType         string      `json:"PaymentConditionType,omitempty"`
	PaymentConditionNumberOfDays int         `json:"PaymentConditionNumberOfDays,omitempty"`
	IsPerson                     bool        `json:"IsPerson,omitempty"`
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
func Restore(api dinero.API, ID string) error {
	route := fmt.Sprintf("v1/{organizationID}/contacts/%v/restore", ID)

	return api.Call(http.MethodPost, route, nil, nil)
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
func Get(c dinero.API, ID string) (*Contact, error) {
	route := "v1/{organizationID}/contacts/" + ID

	var contact Contact
	if err := c.Call(http.MethodGet, route, nil, &contact); err != nil {
		return nil, err
	}

	return &contact, nil
}

// ContactParams is the parameters you provide to create a contact in dinero's api
type ContactParams struct {
	ExternalReference            string `json:"externalReference,omitempty"`
	Name                         string `json:"name,omitempty"`
	Street                       string `json:"street,omitempty"`
	ZipCode                      string `json:"zipcode,omitempty"`
	City                         string `json:"city,omitempty"`
	CountryKey                   string `json:"countryKey,omitempty"`
	Phone                        string `json:"phone,omitempty"`
	Email                        string `json:"email,omitempty"`
	Webpage                      string `json:"webpage,omitempty"`
	AttPerson                    string `json:"attPerson,omitempty"`
	VatNumber                    string `json:"vatNumber,omitempty"`
	EanNumber                    string `json:"eanNumber,omitempty"`
	PaymentConditionType         string `json:"paymentConditionType,omitempty"`
	PaymentConditionNumberOfDays int    `json:"paymentConditionNumberOfDays,omitempty"`
	IsPerson                     bool   `json:"isPerson"`
}

// ContactCreated is the returned type when a contact is created successfully
type ContactCreated struct {
	ID string `json:"ContactGUID"`
}

// Add a new contact to the organization
func Add(api dinero.API, params ContactParams) (*ContactCreated, error) {
	route := "v1/{organizationID}/contacts"

	var created ContactCreated
	if err := api.Call(http.MethodPost, route, &params, &created); err != nil {
		return nil, err
	}

	return &created, nil
}

// Update an existing contact
func Update(api dinero.API, id string, params ContactParams) error {
	route := fmt.Sprintf("v1/{organizationID}/contacts/%v", id)

	return api.Call(http.MethodPut, route, params, nil)
}

// Delete a contact from the given organization
func Delete(api dinero.API, ID string) error {
	route := fmt.Sprintf("v1/{organizationID}/contacts/%v", ID)

	return api.Call(http.MethodDelete, route, nil, nil)
}
