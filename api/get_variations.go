package api

import "errors"

type GetVariationsResponse struct {
	Errors           []Error          `json:"errors,omitempty"`
	VariationsResult VariationsResult `json:"variationsResult,omitempty"`
}

type GetVariationsParams struct {
	ASIN                  string
	Condition             string
	CurrencyOfPreference  string
	LanguagesOfPreference []string
	Merchant              string
	OfferCount            int
	Resources             []Resource
	VariationCount        int
	VariationPage         int
}

func (p GetVariationsParams) ResourceList() []Resource {
	return p.Resources
}

func (p GetVariationsParams) Payload() (map[string]any, error) {
	if p.ASIN == "" {
		return nil, errors.New("asin is required")
	}

	payload := map[string]any{
		"asin": p.ASIN,
	}

	if p.Condition != "" {
		payload["condition"] = p.Condition
	}
	if p.CurrencyOfPreference != "" {
		payload["currencyOfPreference"] = p.CurrencyOfPreference
	}
	if len(p.LanguagesOfPreference) > 0 {
		payload["languagesOfPreference"] = p.LanguagesOfPreference
	}
	if p.Merchant != "" {
		payload["merchant"] = p.Merchant
	}
	if p.OfferCount > 0 {
		payload["offerCount"] = p.OfferCount
	}
	if len(p.Resources) > 0 {
		payload["resources"] = p.Resources
	}
	if p.VariationCount > 0 {
		payload["variationCount"] = p.VariationCount
	}
	if p.VariationPage > 0 {
		payload["variationPage"] = p.VariationPage
	}

	return payload, nil
}
