package entries

import (
	"net/http"
	"net/url"
	"strconv"

	dinero "github.com/eikc/dinero-go"
)

type Entry struct {
	AccountNumber int         `json:"accountNumber,omitempty"`
	AccountName   string      `json:"accountName,omitempty"`
	Date          dinero.Date `json:"date,omitempty"`
	VoucherNumber int         `json:"voucherNumber,omitempty"`
	VoucherType   string      `json:"voucherType,omitempty"`
	Description   string      `json:"description,omitempty"`
	VatType       string      `json:"vatType,omitempty"`
	VatCode       string      `json:"vatCode,omitempty"`
	Amount        float64     `json:"amount,omitempty"`
	EntryGUID     string      `json:"entryGuid,omitempty"`
}

type GetRequestParams struct {
	FromDate     dinero.Date
	ToDate       dinero.Date
	IncludePrimo bool
}

func Get(api dinero.API, params GetRequestParams) ([]Entry, error) {
	query := url.Values{}
	query.Add("fromDate", string(params.FromDate))
	query.Add("toDate", string(params.ToDate))
	query.Add("includePrimo", strconv.FormatBool(params.IncludePrimo))

	queryString := query.Encode()
	route := "v1/{organizationID}/entries?" + queryString

	var entries []Entry
	if err := api.Call(http.MethodGet, route, nil, &entries); err != nil {
		return nil, err
	}

	return entries, nil

}
