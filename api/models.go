package api

type Error struct {
	Type    string `json:"__type,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ItemsResult struct {
	Items []Item `json:"items,omitempty"`
}

type VariationsResult struct {
	Items            []Item           `json:"items,omitempty"`
	VariationSummary VariationSummary `json:"variationSummary,omitempty"`
}

type Item struct {
	ASIN                string               `json:"asin,omitempty"`
	BrowseNodeInfo      BrowseNodeInfo       `json:"browseNodeInfo,omitempty"`
	DetailPageURL       string               `json:"detailPageUrl,omitempty"`
	Images              Images               `json:"images,omitempty"`
	ItemInfo            ItemInfo             `json:"itemInfo,omitempty"`
	OffersV2            OffersV2             `json:"offersV2,omitempty"`
	ParentASIN          string               `json:"parentASIN,omitempty"`
	Score               float32              `json:"score,omitempty"`
	CustomerReviews     *CustomerReviews     `json:"customerReviews,omitempty"`
	VariationAttributes []VariationAttribute `json:"variationAttributes,omitempty"`
}

type VariationSummary struct {
	PageCount           int                  `json:"pageCount,omitempty"`
	Price               Price                `json:"price,omitempty"`
	VariationCount      int                  `json:"variationCount,omitempty"`
	VariationDimensions []VariationDimension `json:"variationDimensions,omitempty"`
}

type Price struct {
	HighestPrice OfferPrice `json:"highestPrice,omitempty"`
	LowestPrice  OfferPrice `json:"lowestPrice,omitempty"`
}

type OfferPrice struct {
	Amount        float32      `json:"amount,omitempty"`
	Currency      string       `json:"currency,omitempty"`
	DisplayAmount string       `json:"displayAmount,omitempty"`
	PricePerUnit  float32      `json:"pricePerUnit,omitempty"`
	Savings       OfferSavings `json:"savings,omitempty"`
}

type OfferSavings struct {
	Amount        float32 `json:"amount,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	DisplayAmount string  `json:"displayAmount,omitempty"`
	Percentage    int     `json:"percentage,omitempty"`
	PricePerUnit  float32 `json:"pricePerUnit,omitempty"`
}

type VariationDimension struct {
	DisplayName string   `json:"displayName,omitempty"`
	Locale      string   `json:"locale,omitempty"`
	Name        string   `json:"name,omitempty"`
	Values      []string `json:"values,omitempty"`
}

type VariationAttribute struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type BrowseNodeInfo struct {
	BrowseNodes      []BrowseNode     `json:"browseNodes,omitempty"`
	WebsiteSalesRank WebsiteSalesRank `json:"websiteSalesRank,omitempty"`
}

type BrowseNode struct {
	Ancestor        *BrowseNodeAncestor `json:"ancestor,omitempty"`
	Children        []BrowseNodeChild   `json:"children,omitempty"`
	ContextFreeName string              `json:"contextFreeName,omitempty"`
	DisplayName     string              `json:"displayName,omitempty"`
	ID              string              `json:"id,omitempty"`
	IsRoot          bool                `json:"isRoot,omitempty"`
	SalesRank       int                 `json:"salesRank,omitempty"`
}

type BrowseNodeAncestor struct {
	Ancestor        *BrowseNodeAncestor `json:"ancestor,omitempty"`
	ContextFreeName string              `json:"contextFreeName,omitempty"`
	DisplayName     string              `json:"displayName,omitempty"`
	ID              string              `json:"id,omitempty"`
}

type BrowseNodeChild struct {
	ContextFreeName string `json:"contextFreeName,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	ID              string `json:"id,omitempty"`
}

type WebsiteSalesRank struct {
	DisplayName     string `json:"displayName,omitempty"`
	ContextFreeName string `json:"contextFreeName,omitempty"`
	ID              string `json:"id,omitempty"`
	SalesRank       int    `json:"salesRank,omitempty"`
}

type CustomerReviews struct {
	Count      *int            `json:"count,omitempty"`
	StarRating *CustomerRating `json:"starRating,omitempty"`
}

type CustomerRating struct {
	Value *float64 `json:"value,omitempty"`
}

type Images struct {
	Primary  ImageType   `json:"primary,omitempty"`
	Variants []ImageType `json:"variants,omitempty"`
}

type ImageType struct {
	Small   ImageSize `json:"small,omitempty"`
	Medium  ImageSize `json:"medium,omitempty"`
	Large   ImageSize `json:"large,omitempty"`
	HighRes ImageSize `json:"highRes,omitempty"`
}

type ImageSize struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type ItemInfo struct {
	ByLineInfo      ByLineInfo      `json:"byLineInfo,omitempty"`
	ContentInfo     ContentInfo     `json:"contentInfo,omitempty"`
	ContentRating   ContentRating   `json:"contentRating,omitempty"`
	Classifications Classifications `json:"classifications,omitempty"`
	ExternalIDs     ExternalIDs     `json:"externalIds,omitempty"`
	Features        MultiValue      `json:"features,omitempty"`
	ManufactureInfo ManufactureInfo `json:"manufactureInfo,omitempty"`
	ProductInfo     ProductInfo     `json:"productInfo,omitempty"`
	TechnicalInfo   TechnicalInfo   `json:"technicalInfo,omitempty"`
	Title           StringValue     `json:"title,omitempty"`
	TradeInInfo     TradeInInfo     `json:"tradeInInfo,omitempty"`
}

type ByLineInfo struct {
	Brand        StringValue   `json:"brand,omitempty"`
	Contributors []Contributor `json:"contributors,omitempty"`
	Manufacturer StringValue   `json:"manufacturer,omitempty"`
}

type Contributor struct {
	Locale string `json:"locale,omitempty"`
	Name   string `json:"name,omitempty"`
	Role   string `json:"role,omitempty"`
}

type ContentInfo struct {
	Edition         StringValue  `json:"edition,omitempty"`
	Languages       Languages    `json:"languages,omitempty"`
	PagesCount      IntegerValue `json:"pagesCount,omitempty"`
	PublicationDate StringValue  `json:"publicationDate,omitempty"`
}

type Languages struct {
	DisplayValues []LanguageType `json:"displayValues,omitempty"`
	Label         string         `json:"label,omitempty"`
	Locale        string         `json:"locale,omitempty"`
}

type LanguageType struct {
	DisplayValue string `json:"displayValue,omitempty"`
	Type         string `json:"type,omitempty"`
}

type ContentRating struct {
	AudienceRating StringValue `json:"audienceRating,omitempty"`
}

type Classifications struct {
	Binding      StringValue `json:"binding,omitempty"`
	ProductGroup StringValue `json:"productGroup,omitempty"`
}

type ExternalIDs struct {
	EANs  MultiValue `json:"eans,omitempty"`
	ISBNs MultiValue `json:"isbns,omitempty"`
	UPCs  MultiValue `json:"upcs,omitempty"`
}

type ManufactureInfo struct {
	ItemPartNumber StringValue `json:"itemPartNumber,omitempty"`
	Model          StringValue `json:"model,omitempty"`
	Warranty       StringValue `json:"warranty,omitempty"`
}

type ProductInfo struct {
	Color          StringValue     `json:"color,omitempty"`
	IsAdultProduct BoolValue       `json:"isAdultProduct,omitempty"`
	ItemDimensions DimensionValues `json:"itemDimensions,omitempty"`
	ReleaseDate    StringValue     `json:"releaseDate,omitempty"`
	Size           StringValue     `json:"size,omitempty"`
	UnitCount      IntegerValue    `json:"unitCount,omitempty"`
}

type DimensionValues struct {
	Height UnitBasedAttribute `json:"height,omitempty"`
	Length UnitBasedAttribute `json:"length,omitempty"`
	Weight UnitBasedAttribute `json:"weight,omitempty"`
	Width  UnitBasedAttribute `json:"width,omitempty"`
}

type TechnicalInfo struct {
	Formats MultiValue `json:"formats,omitempty"`
}

type TradeInInfo struct {
	IsEligibleForTradeIn bool `json:"isEligibleForTradeIn,omitempty"`
}

type StringValue struct {
	DisplayValue string `json:"displayValue,omitempty"`
	Label        string `json:"label,omitempty"`
	Locale       string `json:"locale,omitempty"`
}

type BoolValue struct {
	DisplayValue bool   `json:"displayValue,omitempty"`
	Label        string `json:"label,omitempty"`
	Locale       string `json:"locale,omitempty"`
}

type IntegerValue struct {
	DisplayValue int    `json:"displayValue,omitempty"`
	Label        string `json:"label,omitempty"`
	Locale       string `json:"locale,omitempty"`
}

type UnitBasedAttribute struct {
	DisplayValue float32 `json:"displayValue,omitempty"`
	Label        string  `json:"label,omitempty"`
	Locale       string  `json:"locale,omitempty"`
	Unit         string  `json:"unit,omitempty"`
}

type MultiValue struct {
	DisplayValues []string `json:"displayValues,omitempty"`
	Label         string   `json:"label,omitempty"`
	Locale        string   `json:"locale,omitempty"`
}

type OffersV2 struct {
	Listings []OfferV2Listing `json:"listings,omitempty"`
}

type OfferV2Listing struct {
	Availability   OfferV2Availability  `json:"availability,omitempty"`
	Condition      OfferV2Condition     `json:"condition,omitempty"`
	DealDetails    OfferV2DealDetails   `json:"dealDetails,omitempty"`
	IsBuyBoxWinner bool                 `json:"isBuyBoxWinner,omitempty"`
	LoyaltyPoints  OfferV2LoyaltyPoints `json:"loyaltyPoints,omitempty"`
	MerchantInfo   OfferV2MerchantInfo  `json:"merchantInfo,omitempty"`
	Price          OfferV2Price         `json:"price,omitempty"`
	Type           string               `json:"type,omitempty"`
	ViolatesMAP    bool                 `json:"violatesMAP,omitempty"`
}

type OfferV2Availability struct {
	MaxOrderQuantity int    `json:"maxOrderQuantity,omitempty"`
	Message          string `json:"message,omitempty"`
	MinOrderQuantity int    `json:"minOrderQuantity,omitempty"`
	Type             string `json:"type,omitempty"`
}

type OfferV2Condition struct {
	ConditionNote string `json:"conditionNote,omitempty"`
	SubCondition  string `json:"subCondition,omitempty"`
	Value         string `json:"value,omitempty"`
}

type OfferV2DealDetails struct {
	AccessType                        string `json:"accessType,omitempty"`
	Badge                             string `json:"badge,omitempty"`
	EarlyAccessDurationInMilliseconds int64  `json:"earlyAccessDurationInMilliseconds,omitempty"`
	EndTime                           string `json:"endTime,omitempty"`
	PercentClaimed                    int    `json:"percentClaimed,omitempty"`
	StartTime                         string `json:"startTime,omitempty"`
}

type OfferV2LoyaltyPoints struct {
	Points int `json:"points,omitempty"`
}

type OfferV2MerchantInfo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OfferV2Price struct {
	Money        OfferV2Money        `json:"money,omitempty"`
	PricePerUnit OfferV2Money        `json:"pricePerUnit,omitempty"`
	SavingBasis  *OfferV2SavingBasis `json:"savingBasis,omitempty"`
	Savings      *OfferV2Savings     `json:"savings,omitempty"`
}

type OfferV2Money struct {
	Amount        float64 `json:"amount,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	DisplayAmount string  `json:"displayAmount,omitempty"`
}

type OfferV2SavingBasis struct {
	Money                OfferV2Money `json:"money,omitempty"`
	SavingBasisType      string       `json:"savingBasisType,omitempty"`
	SavingBasisTypeLabel string       `json:"savingBasisTypeLabel,omitempty"`
}

type OfferV2Savings struct {
	Money      OfferV2Money `json:"money,omitempty"`
	Percentage int          `json:"percentage,omitempty"`
}
