package creatorsapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func tokenEndpoint(version string) (string, error) {
	switch version {
	case "2.1":
		return "https://creatorsapi.auth.us-east-1.amazoncognito.com/oauth2/token", nil
	case "2.2":
		return "https://creatorsapi.auth.eu-south-2.amazoncognito.com/oauth2/token", nil
	case "2.3":
		return "https://creatorsapi.auth.us-west-2.amazoncognito.com/oauth2/token", nil
	case "3.1":
		return "https://api.amazon.com/auth/o2/token", nil
	case "3.2":
		return "https://api.amazon.co.uk/auth/o2/token", nil
	case "3.3":
		return "https://api.amazon.co.jp/auth/o2/token", nil
	default:
		return "", fmt.Errorf("unsupported credential version %q", version)
	}
}

func tokenScope(version string) (string, error) {
	switch {
	case strings.HasPrefix(version, "2."):
		return creatorsScopeV2, nil
	case strings.HasPrefix(version, "3."):
		return creatorsScopeV3, nil
	default:
		return "", fmt.Errorf("unsupported credential version %q", version)
	}
}

func (c *Client) fetchAccessToken(ctx context.Context) (string, time.Time, error) {
	endpoint, err := tokenEndpoint(c.credentialVersion)
	if err != nil {
		return "", time.Time{}, err
	}
	scope, err := tokenScope(c.credentialVersion)
	if err != nil {
		return "", time.Time{}, err
	}

	var requestBody io.Reader
	contentType := "application/x-www-form-urlencoded"
	if strings.HasPrefix(c.credentialVersion, "3.") {
		bodyBytes, err := json.Marshal(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     c.credentialID,
			"client_secret": c.credentialSecret,
			"scope":         scope,
		})
		if err != nil {
			return "", time.Time{}, err
		}
		requestBody = bytes.NewReader(bodyBytes)
		contentType = "application/json"
	} else {
		form := url.Values{}
		form.Set("grant_type", "client_credentials")
		form.Set("client_id", c.credentialID)
		form.Set("client_secret", c.credentialSecret)
		form.Set("scope", scope)
		requestBody = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, requestBody)
	if err != nil {
		return "", time.Time{}, err
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := c.httpClientSnapshot().Do(req)
	if err != nil {
		return "", time.Time{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", time.Time{}, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return "", time.Time{}, &HTTPError{
			StatusCode: resp.StatusCode,
			Body:       append([]byte(nil), body...),
		}
	}

	var parsed tokenResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", time.Time{}, fmt.Errorf("invalid token response: %w", err)
	}
	if parsed.AccessToken == "" {
		return "", time.Time{}, fmt.Errorf("token response missing access token")
	}

	expiresIn := time.Hour
	if parsed.ExpiresIn > 0 {
		expiresIn = time.Duration(parsed.ExpiresIn) * time.Second
	}
	expiry := time.Now().Add(expiresIn)
	if expiresIn > tokenExpirySkew {
		expiry = expiry.Add(-tokenExpirySkew)
	}

	return parsed.AccessToken, expiry, nil
}
