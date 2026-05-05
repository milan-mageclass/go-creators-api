package creatorsapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	default:
		return "", fmt.Errorf("unsupported credential version %q", version)
	}
}

func (c *Client) fetchAccessToken(ctx context.Context) (string, time.Time, error) {
	endpoint, err := tokenEndpoint(c.credentialVersion)
	if err != nil {
		return "", time.Time{}, err
	}

	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", c.credentialID)
	form.Set("client_secret", c.credentialSecret)
	form.Set("scope", defaultTokenScope)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return "", time.Time{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
