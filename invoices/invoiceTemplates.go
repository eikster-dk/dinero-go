package invoices

import (
	"github.com/eikc/dinero-go"
	"net/http"
)

type InvoiceTemplate struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	IsDefault             bool   `json:"isDefault"`
	Theme                 int    `json:"theme"`
	PrimaryColor          string `json:"primaryColor"`
	SecondaryColor        string `json:"secondaryColor"`
	AddressPlacement      int    `json:"addressPlacement"`
	Font                  string `json:"font"`
	LogoType              int    `json:"logoType"`
	LogoText              string `json:"logoText"`
	LogoFileGuid          string `json:"logoFileGuid"`
	MaxImageWidthMm       int    `json:"maxImageWidthMm"`
	ShowCompanyInfo       bool   `json:"showCompanyInfo"`
	ShowPaymentConditions bool   `json:"showPaymentConditions"`
	ShowLogo              bool   `json:"showLogo"`
	ShowLineQuantity      bool   `json:"showLineQuantity"`
	ShowLineUnit          bool   `json:"showLineUnit"`
	ImageFileUploadGuid   string `json:"imageFileUploadGuid"`
	ImageTextColor        string `json:"imageTextColor"`
}

// GetInvoiceTemplates will get the list of invoice templates available to the user
func GetInvoiceTemplates(api dinero.API) ([]InvoiceTemplate, error) {
	route := "v1/{organizationID}/invoices/templates"
	var invoiceTemplate []InvoiceTemplate

	if err := api.Call(http.MethodGet, route, nil, &invoiceTemplate); err != nil {
		return nil, err
	}

	return invoiceTemplate, nil
}
