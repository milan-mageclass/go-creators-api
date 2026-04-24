package api

import "errors"

type GetItemsResponse struct {
	Errors      []Error     `json:"errors,omitempty"`
	ItemsResult ItemsResult `json:"itemsResult,omitempty"`
}

type GetItemsParams struct {
	ItemIDType            string
	ItemIDs               []string
	Condition             string
	CurrencyOfPreference  string
	LanguagesOfPreference []string
	Merchant              string
	OfferCount            int
	Resources             []Resource
}

func (p GetItemsParams) ResourceList() []Resource {
	return p.Resources
}

func (p GetItemsParams) Payload() (map[string]any, error) {
	if len(p.ItemIDs) == 0 {
		return nil, errors.New("one or more item ids required")
	}

	payload := map[string]any{
		"itemIdType": "ASIN",
		"itemIds":    p.ItemIDs,
	}

	if p.ItemIDType != "" {
		payload["itemIdType"] = p.ItemIDType
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

	return payload, nil
}
