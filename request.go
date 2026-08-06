package creatorsapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/milan-mageclass/go-creators-api/api"
)

const catalogEndpoint = "https://creatorsapi.amazon/catalog/v1"

type payloadResourceLister interface {
	Payload() (map[string]any, error)
	ResourceList() []api.Resource
}

func (c *Client) executeOperation(ctx context.Context, operation api.Operation, params payloadResourceLister, v any) error {
	if params == nil {
		return fmt.Errorf("nil parameters")
	}

	if err := operation.Validate(params.ResourceList()); err != nil {
		return err
	}

	payload, err := params.Payload()
	if err != nil {
		return err
	}
	payload["partnerTag"] = c.partnerTag
	payload["marketplace"] = c.marketplace

	token, err := c.token(ctx)
	if err != nil {
		return err
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s", catalogEndpoint, operation.Path()), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	authorization := fmt.Sprintf("Bearer %s", token)
	if !strings.HasPrefix(c.credentialVersion, "3.") {
		authorization += fmt.Sprintf(", Version %s", c.credentialVersion)
	}
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-marketplace", c.marketplace)

	resp, err := c.httpClientSnapshot().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		_ = json.Unmarshal(respBody, v)
		return &APIError{
			StatusCode: resp.StatusCode,
			Errors:     apiErrorsFromResponse(v),
			Body:       append([]byte(nil), respBody...),
		}
	}

	if err := json.Unmarshal(respBody, v); err != nil {
		return fmt.Errorf("invalid response: %w", err)
	}

	return nil
}
