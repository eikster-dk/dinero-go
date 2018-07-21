package invoices

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/eikc/dinero-go"
)

const (
	timeFormat = "2006-01-02"
)

// const for ascending and descending sorting params
const (
	ASC  = "ascending"
	DESC = "descending"
)

// Fields for what to be returned of properties on the list
// if null, the list is defaulted to: ID,ContactName,Date,Description
// Possible values are:
const (
	Number            = "Number"
	ID                = "Guid"
	ContactName       = "ContactName"
	Date              = "Date"
	PaymentDate       = "PaymentDate"
	Description       = "Description"
	Currency          = "Currency"
	Status            = "Status"
	MailOutStatus     = "MailOutStatus"
	TotalExclVatInDKK = "TotalExclVatInDkk"
	TotalInclVatInDKK = "TotalInclVatInDkk"
	TotalExclVat      = "TotalExclVat"
	TotalInclVat      = "TotalInclVat"
	CreatedAt         = "CreatedAt"
	UpdatedAt         = "UpdatedAt"
	DeletedAt         = "DeletedAt"
)

// Status filter types
const (
	Draft    = "Draft"
	Booked   = "Booked"
	Paid     = "Paid"
	OverPaid = "OverPaid"
	Overdue  = "Overdue"
)

// Sort Options
const (
	VoucherNumber = "VoucherNumber"
	VoucherDate   = "VoucherDate"
)

// ListParams are all the params that can be used to return a list of invoices
type ListParams struct {
	StartDate      *time.Time
	EndDate        *time.Time
	Fields         []string
	FreeTextSearch *string
	StatusFilter   *string
	ChangesSince   *dinero.Time
	DeletedOnly    *bool
	Page           *int
	PageSize       *int
	Sort           []string
	SortOrder      *string
}

func (params *ListParams) getQueryString() string {
	query := url.Values{}
	if params.StartDate != nil {
		formatted := params.StartDate.Format(timeFormat)
		query.Add("startDate", formatted)
	}

	if params.EndDate != nil {
		formatted := params.EndDate.Format(timeFormat)
		query.Add("endDate", formatted)
	}

	if params.Fields != nil {
		fields := dinero.BuildFieldsQuery(params.Fields...)
		query.Add("fields", fields)
	}

	if params.FreeTextSearch != nil {
		query.Add("freeTextSearch", *params.FreeTextSearch)
	}

	if params.ChangesSince != nil {
		marshalled, _ := params.ChangesSince.MarshalJSON()
		query.Add("changesSince", string(marshalled))
	}

	if params.DeletedOnly != nil {
		query.Add("deletedOnly", strconv.FormatBool(*params.DeletedOnly))
	}

	if params.Page != nil {
		query.Add("page", strconv.FormatInt(int64(*params.Page), 10))
	}

	if params.PageSize != nil {
		query.Add("pageSize", strconv.FormatInt(int64(*params.PageSize), 10))
	}

	if params.Sort != nil {
		sortFields := dinero.BuildFieldsQuery(params.Sort...)
		query.Add("sort", sortFields)
	}

	if params.SortOrder != nil {
		query.Add("sortOrder", *params.SortOrder)
	}

	return query.Encode()
}

// InvoiceList is the returned response from dinero when getting a list of invoices
type InvoiceList struct {
	Collection []Invoice
	Pagination dinero.PaginationResult
}

// List receives a list of invoices for the organization.
func List(api dinero.API, params ListParams) (*InvoiceList, error) {
	route := fmt.Sprint("v1/{organizationID}/invoices?", params.getQueryString())

	var invoiceList InvoiceList
	if err := api.Call(http.MethodGet, route, nil, &invoiceList); err != nil {
		return nil, err
	}

	return &invoiceList, nil
}
