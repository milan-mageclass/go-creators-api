package creatorsapi

import (
	"fmt"
	"strings"

	"github.com/milan-mageclass/go-creators-api/api"
)

// APIError describes a non-2xx API response while preserving the decoded error payload.
type APIError struct {
	StatusCode int
	Errors     []api.Error
	Body       []byte
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}

	if len(e.Errors) == 0 {
		return fmt.Sprintf("request failed with status %d: %s", e.StatusCode, strings.TrimSpace(string(e.Body)))
	}

	first := e.Errors[0]
	if first.Code != "" && first.Message != "" {
		return fmt.Sprintf("request failed with status %d: %s: %s", e.StatusCode, first.Code, first.Message)
	}
	if first.Message != "" {
		return fmt.Sprintf("request failed with status %d: %s", e.StatusCode, first.Message)
	}
	if first.Code != "" {
		return fmt.Sprintf("request failed with status %d: %s", e.StatusCode, first.Code)
	}

	return fmt.Sprintf("request failed with status %d", e.StatusCode)
}

func apiErrorsFromResponse(v any) []api.Error {
	switch resp := v.(type) {
	case *api.GetItemsResponse:
		return append([]api.Error(nil), resp.Errors...)
	case *api.GetVariationsResponse:
		return append([]api.Error(nil), resp.Errors...)
	default:
		return nil
	}
}
