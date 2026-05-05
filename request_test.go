package creatorsapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/milan-mageclass/go-creators-api/api"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestGetItemsReturnsAPIError(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		CredentialID:      "id",
		CredentialSecret:  "secret",
		CredentialVersion: "2.1",
		PartnerTag:        "tag-20",
		Marketplace:       "www.amazon.com",
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				switch {
				case strings.Contains(req.URL.Host, "amazoncognito.com"):
					return jsonResponse(t, req, http.StatusOK, tokenResponse{
						AccessToken: "test-token",
						ExpiresIn:   3600,
						TokenType:   "Bearer",
					}), nil
				case req.URL.Host == "creatorsapi.amazon" && req.URL.Path == "/catalog/v1/getItems":
					return jsonResponse(t, req, http.StatusBadRequest, api.GetItemsResponse{
						Errors: []api.Error{
							{
								Type:    "ValidationException",
								Code:    "InvalidParameterValue",
								Message: "itemIds must not be empty",
							},
						},
					}), nil
				default:
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL.String())
					return nil, nil
				}
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	_, err = client.GetItems(context.Background(), &api.GetItemsParams{
		ItemIDs: []string{"bad-id"},
	})
	if err == nil {
		t.Fatal("GetItems() error = nil, want APIError")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("errors.As(%T, *APIError) = false", err)
	}

	if apiErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("APIError.StatusCode = %d, want %d", apiErr.StatusCode, http.StatusBadRequest)
	}
	if len(apiErr.Errors) != 1 {
		t.Fatalf("len(APIError.Errors) = %d, want 1", len(apiErr.Errors))
	}
	if apiErr.Errors[0].Code != "InvalidParameterValue" {
		t.Fatalf("APIError.Errors[0].Code = %q, want %q", apiErr.Errors[0].Code, "InvalidParameterValue")
	}
	if !strings.Contains(apiErr.Error(), "InvalidParameterValue") {
		t.Fatalf("APIError.Error() = %q, want code in string", apiErr.Error())
	}
	if !strings.Contains(string(apiErr.Body), "\"errors\"") {
		t.Fatalf("APIError.Body = %q, want raw response body", string(apiErr.Body))
	}
}

func TestGetItemsReturnsAPIErrorForNonJSONErrorBody(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		CredentialID:      "id",
		CredentialSecret:  "secret",
		CredentialVersion: "2.1",
		PartnerTag:        "tag-20",
		Marketplace:       "www.amazon.com",
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				switch {
				case strings.Contains(req.URL.Host, "amazoncognito.com"):
					return jsonResponse(t, req, http.StatusOK, tokenResponse{
						AccessToken: "test-token",
						ExpiresIn:   3600,
						TokenType:   "Bearer",
					}), nil
				case req.URL.Host == "creatorsapi.amazon" && req.URL.Path == "/catalog/v1/getItems":
					return textResponse(req, http.StatusBadGateway, "upstream proxy error"), nil
				default:
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL.String())
					return nil, nil
				}
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	_, err = client.GetItems(context.Background(), &api.GetItemsParams{
		ItemIDs: []string{"B00MNV8E0C"},
	})
	if err == nil {
		t.Fatal("GetItems() error = nil, want APIError")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("errors.As(%T, *APIError) = false", err)
	}
	if apiErr.StatusCode != http.StatusBadGateway {
		t.Fatalf("APIError.StatusCode = %d, want %d", apiErr.StatusCode, http.StatusBadGateway)
	}
	if string(apiErr.Body) != "upstream proxy error" {
		t.Fatalf("APIError.Body = %q, want %q", string(apiErr.Body), "upstream proxy error")
	}
	if len(apiErr.Errors) != 0 {
		t.Fatalf("len(APIError.Errors) = %d, want 0", len(apiErr.Errors))
	}
}

func TestGetItemsReturnsAPIErrorForEmptyErrorBody(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		CredentialID:      "id",
		CredentialSecret:  "secret",
		CredentialVersion: "2.1",
		PartnerTag:        "tag-20",
		Marketplace:       "www.amazon.com",
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				switch {
				case strings.Contains(req.URL.Host, "amazoncognito.com"):
					return jsonResponse(t, req, http.StatusOK, tokenResponse{
						AccessToken: "test-token",
						ExpiresIn:   3600,
						TokenType:   "Bearer",
					}), nil
				case req.URL.Host == "creatorsapi.amazon" && req.URL.Path == "/catalog/v1/getItems":
					return textResponse(req, http.StatusServiceUnavailable, ""), nil
				default:
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL.String())
					return nil, nil
				}
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	_, err = client.GetItems(context.Background(), &api.GetItemsParams{
		ItemIDs: []string{"B00MNV8E0C"},
	})
	if err == nil {
		t.Fatal("GetItems() error = nil, want APIError")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("errors.As(%T, *APIError) = false", err)
	}
	if apiErr.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("APIError.StatusCode = %d, want %d", apiErr.StatusCode, http.StatusServiceUnavailable)
	}
	if len(apiErr.Body) != 0 {
		t.Fatalf("len(APIError.Body) = %d, want 0", len(apiErr.Body))
	}
}

func TestGetItemsReturnsSuccessResponse(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		CredentialID:      "id",
		CredentialSecret:  "secret",
		CredentialVersion: "2.1",
		PartnerTag:        "tag-20",
		Marketplace:       "www.amazon.com",
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				switch {
				case strings.Contains(req.URL.Host, "amazoncognito.com"):
					return jsonResponse(t, req, http.StatusOK, tokenResponse{
						AccessToken: "test-token",
						ExpiresIn:   3600,
						TokenType:   "Bearer",
					}), nil
				case req.URL.Host == "creatorsapi.amazon" && req.URL.Path == "/catalog/v1/getItems":
					return jsonResponse(t, req, http.StatusOK, api.GetItemsResponse{
						ItemsResult: api.ItemsResult{
							Items: []api.Item{{ASIN: "B00MNV8E0C"}},
						},
					}), nil
				default:
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL.String())
					return nil, nil
				}
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	resp, err := client.GetItems(context.Background(), &api.GetItemsParams{
		ItemIDs: []string{"B00MNV8E0C"},
	})
	if err != nil {
		t.Fatalf("GetItems() error = %v", err)
	}
	if resp == nil || len(resp.ItemsResult.Items) != 1 || resp.ItemsResult.Items[0].ASIN != "B00MNV8E0C" {
		t.Fatalf("GetItems() response = %#v", resp)
	}
}

func TestGetItemsReturnsHTTPErrorForTokenNonJSONErrorBody(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		CredentialID:      "id",
		CredentialSecret:  "secret",
		CredentialVersion: "2.1",
		PartnerTag:        "tag-20",
		Marketplace:       "www.amazon.com",
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				switch {
				case strings.Contains(req.URL.Host, "amazoncognito.com"):
					return textResponse(req, http.StatusBadGateway, "bad gateway"), nil
				default:
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL.String())
					return nil, nil
				}
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	_, err = client.GetItems(context.Background(), &api.GetItemsParams{
		ItemIDs: []string{"B00MNV8E0C"},
	})
	if err == nil {
		t.Fatal("GetItems() error = nil, want HTTPError")
	}

	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("errors.As(%T, *HTTPError) = false", err)
	}
	if httpErr.StatusCode != http.StatusBadGateway {
		t.Fatalf("HTTPError.StatusCode = %d, want %d", httpErr.StatusCode, http.StatusBadGateway)
	}
	if string(httpErr.Body) != "bad gateway" {
		t.Fatalf("HTTPError.Body = %q, want %q", string(httpErr.Body), "bad gateway")
	}
}

func TestAPIErrorsFromResponseSupportsCommonResponses(t *testing.T) {
	t.Parallel()

	itemErrors := apiErrorsFromResponse(&api.GetItemsResponse{
		Errors: []api.Error{{Code: "items"}},
	})
	if len(itemErrors) != 1 || itemErrors[0].Code != "items" {
		t.Fatalf("apiErrorsFromResponse(GetItemsResponse) = %#v", itemErrors)
	}

	variationErrors := apiErrorsFromResponse(&api.GetVariationsResponse{
		Errors: []api.Error{{Code: "variations"}},
	})
	if len(variationErrors) != 1 || variationErrors[0].Code != "variations" {
		t.Fatalf("apiErrorsFromResponse(GetVariationsResponse) = %#v", variationErrors)
	}
}

func jsonResponse(t *testing.T, req *http.Request, status int, body any) *http.Response {
	t.Helper()

	payload, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(string(payload))),
		Request:    req,
	}
}

func textResponse(req *http.Request, status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    req,
	}
}
