package vattypes

import (
	"net/http"

	"github.com/eikc/dinero-go"
)

// VatType represents the vat type data model from dinero
type VatType struct {
	VatCode string  `json:"vatCode,omitempty"`
	Name    string  `json:"name,omitempty"`
	VatRate float64 `json:"vatRate,omitempty"`
}

// Get is used to get the vattypes from a given company in dinero
func Get(api dinero.API) ([]VatType, error) {
	route := "v1/{organizationID}/vatTypes"

	var vattypes []VatType
	if err := api.Call(http.MethodGet, route, nil, &vattypes); err != nil {
		return nil, err
	}

	return vattypes, nil
}
