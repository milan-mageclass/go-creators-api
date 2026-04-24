package api

type Resource string

const (
	BrowseNodeInfoBrowseNodes          Resource = "browseNodeInfo.browseNodes"
	BrowseNodeInfoBrowseNodesAncestor  Resource = "browseNodeInfo.browseNodes.ancestor"
	BrowseNodeInfoBrowseNodesSalesRank Resource = "browseNodeInfo.browseNodes.salesRank"
	BrowseNodeInfoWebsiteSalesRank     Resource = "browseNodeInfo.websiteSalesRank"
	CustomerReviewsCount               Resource = "customerReviews.count"
	CustomerReviewsStarRating          Resource = "customerReviews.starRating"
	ImagesPrimarySmall                 Resource = "images.primary.small"
	ImagesPrimaryMedium                Resource = "images.primary.medium"
	ImagesPrimaryLarge                 Resource = "images.primary.large"
	ImagesPrimaryHighRes               Resource = "images.primary.highRes"
	ImagesVariantsSmall                Resource = "images.variants.small"
	ImagesVariantsMedium               Resource = "images.variants.medium"
	ImagesVariantsLarge                Resource = "images.variants.large"
	ImagesVariantsHighRes              Resource = "images.variants.highRes"
	ItemInfoByLineInfo                 Resource = "itemInfo.byLineInfo"
	ItemInfoContentInfo                Resource = "itemInfo.contentInfo"
	ItemInfoContentRating              Resource = "itemInfo.contentRating"
	ItemInfoClassifications            Resource = "itemInfo.classifications"
	ItemInfoExternalIDs                Resource = "itemInfo.externalIds"
	ItemInfoFeatures                   Resource = "itemInfo.features"
	ItemInfoManufactureInfo            Resource = "itemInfo.manufactureInfo"
	ItemInfoProductInfo                Resource = "itemInfo.productInfo"
	ItemInfoTechnicalInfo              Resource = "itemInfo.technicalInfo"
	ItemInfoTitle                      Resource = "itemInfo.title"
	ItemInfoTradeInInfo                Resource = "itemInfo.tradeInInfo"
	OffersV2ListingsAvailability       Resource = "offersV2.listings.availability"
	OffersV2ListingsCondition          Resource = "offersV2.listings.condition"
	OffersV2ListingsDealDetails        Resource = "offersV2.listings.dealDetails"
	OffersV2ListingsIsBuyBoxWinner     Resource = "offersV2.listings.isBuyBoxWinner"
	OffersV2ListingsLoyaltyPoints      Resource = "offersV2.listings.loyaltyPoints"
	OffersV2ListingsMerchantInfo       Resource = "offersV2.listings.merchantInfo"
	OffersV2ListingsPrice              Resource = "offersV2.listings.price"
	OffersV2ListingsType               Resource = "offersV2.listings.type"
	ParentASIN                         Resource = "parentASIN"
	VariationSummaryPriceHighestPrice  Resource = "variationSummary.price.highestPrice"
	VariationSummaryPriceLowestPrice   Resource = "variationSummary.price.lowestPrice"
	VariationSummaryVariationDimension Resource = "variationSummary.variationDimension"
	SearchRefinements                  Resource = "searchRefinements"
)

var resourceOperationsMap = map[Resource][]Operation{
	BrowseNodeInfoBrowseNodes:          {OperationGetItems, OperationGetVariations, OperationSearchItems},
	BrowseNodeInfoBrowseNodesAncestor:  {OperationGetItems, OperationGetVariations, OperationSearchItems},
	BrowseNodeInfoBrowseNodesSalesRank: {OperationGetItems, OperationGetVariations, OperationSearchItems},
	BrowseNodeInfoWebsiteSalesRank:     {OperationGetItems, OperationGetVariations, OperationSearchItems},
	CustomerReviewsCount:               {OperationGetItems, OperationGetVariations, OperationSearchItems},
	CustomerReviewsStarRating:          {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesPrimarySmall:                 {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesPrimaryMedium:                {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesPrimaryLarge:                 {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesPrimaryHighRes:               {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesVariantsSmall:                {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesVariantsMedium:               {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesVariantsLarge:                {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ImagesVariantsHighRes:              {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoByLineInfo:                 {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoContentInfo:                {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoContentRating:              {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoClassifications:            {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoExternalIDs:                {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoFeatures:                   {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoManufactureInfo:            {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoProductInfo:                {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoTechnicalInfo:              {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoTitle:                      {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ItemInfoTradeInInfo:                {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsAvailability:       {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsCondition:          {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsDealDetails:        {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsIsBuyBoxWinner:     {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsLoyaltyPoints:      {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsMerchantInfo:       {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsPrice:              {OperationGetItems, OperationGetVariations, OperationSearchItems},
	OffersV2ListingsType:               {OperationGetItems, OperationGetVariations, OperationSearchItems},
	ParentASIN:                         {OperationGetItems, OperationGetVariations, OperationSearchItems},
	VariationSummaryPriceHighestPrice:  {OperationGetVariations},
	VariationSummaryPriceLowestPrice:   {OperationGetVariations},
	VariationSummaryVariationDimension: {OperationGetVariations},
	SearchRefinements:                  {OperationSearchItems},
}
