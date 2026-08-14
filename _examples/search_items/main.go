package main

import (
	"context"
	"fmt"
	"log"
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
		log.Fatal(err)
	}

	response, err := client.SearchItems(context.Background(), &api.SearchItemsParams{
		Keywords:  "756655338912",
		ItemCount: 5,
		Resources: []api.Resource{
			api.ItemInfoTitle,
			api.ImagesPrimaryMedium,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range response.SearchResult.Items {
		fmt.Printf("%s: %s\n", item.ASIN, item.ItemInfo.Title.DisplayValue)
	}
}
