package creatorsapi

import (
	"context"

	"github.com/milan-mageclass/go-creators-api/api"
)

func (c *Client) GetItems(ctx context.Context, params *api.GetItemsParams) (*api.GetItemsResponse, error) {
	response := api.GetItemsResponse{}
	if err := c.executeOperation(ctx, api.OperationGetItems, params, &response); err != nil {
		return &response, err
	}
	return &response, nil
}

func (c *Client) GetVariations(ctx context.Context, params *api.GetVariationsParams) (*api.GetVariationsResponse, error) {
	response := api.GetVariationsResponse{}
	if err := c.executeOperation(ctx, api.OperationGetVariations, params, &response); err != nil {
		return &response, err
	}
	return &response, nil
}
