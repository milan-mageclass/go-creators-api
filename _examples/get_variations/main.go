package main

import (
	"context"
	"fmt"
	"os"

	creatorsapi "github.com/milan-mageclass/go-creators-api"
	"github.com/milan-mageclass/go-creators-api/api"
)

func main() {
	client, err := creatorsapi.NewClient(creatorsapi.Config{
		CredentialID:      os.Getenv("CREATORS_API_CLIENT_ID"),
		CredentialSecret:  os.Getenv("CREATORS_API_CLIENT_SECRET"),
		CredentialVersion: os.Getenv("CREATORS_API_CLIENT_VERSION"),
		PartnerTag:        os.Getenv("CREATORS_API_PARTNER_TAG"),
		Marketplace:       "www.amazon.com",
	})
	if err != nil {
		panic(err)
	}

	response, err := client.GetVariations(context.Background(), &api.GetVariationsParams{
		ASIN: "B0CQTM51XC",
		Resources: []api.Resource{
			api.BrowseNodeInfoBrowseNodes,
			api.BrowseNodeInfoBrowseNodesAncestor,
			api.BrowseNodeInfoBrowseNodesSalesRank,
			api.BrowseNodeInfoWebsiteSalesRank,
			api.CustomerReviewsCount,
			api.CustomerReviewsStarRating,
			api.ImagesPrimarySmall,
			api.ImagesPrimaryMedium,
			api.ImagesPrimaryLarge,
			api.ImagesPrimaryHighRes,
			api.ImagesVariantsSmall,
			api.ImagesVariantsMedium,
			api.ImagesVariantsLarge,
			api.ImagesVariantsHighRes,
			api.ItemInfoByLineInfo,
			api.ItemInfoContentInfo,
			api.ItemInfoContentRating,
			api.ItemInfoClassifications,
			api.ItemInfoExternalIDs,
			api.ItemInfoFeatures,
			api.ItemInfoManufactureInfo,
			api.ItemInfoProductInfo,
			api.ItemInfoTechnicalInfo,
			api.ItemInfoTitle,
			api.ItemInfoTradeInInfo,
			api.ParentASIN,
			api.VariationSummaryPriceHighestPrice,
			api.VariationSummaryPriceLowestPrice,
			api.VariationSummaryVariationDimension,
			api.OffersV2ListingsAvailability,
			api.OffersV2ListingsCondition,
			api.OffersV2ListingsDealDetails,
			api.OffersV2ListingsIsBuyBoxWinner,
			api.OffersV2ListingsLoyaltyPoints,
			api.OffersV2ListingsMerchantInfo,
			api.OffersV2ListingsPrice,
			api.OffersV2ListingsType,
		},
	})
	if err != nil {
		panic(err)
	}

	for _, item := range response.VariationsResult.Items {
		fmt.Printf("%s\n%s\n\n", item.ItemInfo.Title.DisplayValue, item.DetailPageURL)
	}
}
