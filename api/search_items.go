package api

import "errors"

type SearchItemsResponse struct {
	Errors       []Error      `json:"errors,omitempty"`
	SearchResult SearchResult `json:"searchResult,omitempty"`
}

type SearchItemsParams struct {
	Actor                 string
	Artist                string
	Author                string
	Availability          string
	Brand                 string
	BrowseNodeID          string
	Condition             string
	CurrencyOfPreference  string
	DeliveryFlags         []string
	ItemCount             int
	ItemPage              int
	Keywords              string
	LanguagesOfPreference []string
	MaxPrice              int
	MinPrice              int
	MinReviewsRating      int
	MinSavingPercent      int
	Properties            map[string]string
	Resources             []Resource
	SearchIndex           string
	SortBy                string
	Title                 string
}

func (p SearchItemsParams) ResourceList() []Resource {
	return p.Resources
}

func (p SearchItemsParams) Payload() (map[string]any, error) {
	if p.Actor == "" && p.Artist == "" && p.Author == "" && p.Brand == "" && p.Keywords == "" && p.Title == "" {
		return nil, errors.New("one or more search terms required")
	}

	payload := map[string]any{}
	if p.Actor != "" {
		payload["actor"] = p.Actor
	}
	if p.Artist != "" {
		payload["artist"] = p.Artist
	}
	if p.Author != "" {
		payload["author"] = p.Author
	}
	if p.Availability != "" {
		payload["availability"] = p.Availability
	}
	if p.Brand != "" {
		payload["brand"] = p.Brand
	}
	if p.BrowseNodeID != "" {
		payload["browseNodeId"] = p.BrowseNodeID
	}
	if p.Condition != "" {
		payload["condition"] = p.Condition
	}
	if p.CurrencyOfPreference != "" {
		payload["currencyOfPreference"] = p.CurrencyOfPreference
	}
	if len(p.DeliveryFlags) > 0 {
		payload["deliveryFlags"] = p.DeliveryFlags
	}
	if p.ItemCount > 0 {
		payload["itemCount"] = p.ItemCount
	}
	if p.ItemPage > 0 {
		payload["itemPage"] = p.ItemPage
	}
	if p.Keywords != "" {
		payload["keywords"] = p.Keywords
	}
	if len(p.LanguagesOfPreference) > 0 {
		payload["languagesOfPreference"] = p.LanguagesOfPreference
	}
	if p.MaxPrice > 0 {
		payload["maxPrice"] = p.MaxPrice
	}
	if p.MinPrice > 0 {
		payload["minPrice"] = p.MinPrice
	}
	if p.MinReviewsRating > 0 {
		payload["minReviewsRating"] = p.MinReviewsRating
	}
	if p.MinSavingPercent > 0 {
		payload["minSavingPercent"] = p.MinSavingPercent
	}
	if len(p.Properties) > 0 {
		payload["properties"] = p.Properties
	}
	if len(p.Resources) > 0 {
		payload["resources"] = p.Resources
	}
	if p.SearchIndex != "" {
		payload["searchIndex"] = p.SearchIndex
	}
	if p.SortBy != "" {
		payload["sortBy"] = p.SortBy
	}
	if p.Title != "" {
		payload["title"] = p.Title
	}

	return payload, nil
}
