package creatorsapi

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"
)

const (
	defaultTokenScope = "creatorsapi/default"
	tokenExpirySkew   = time.Minute
	tokenRetryBackoff = 5 * time.Second
)

var (
	ErrEmptyCredentialID      = errors.New("empty credential id")
	ErrEmptyCredentialSecret  = errors.New("empty credential secret")
	ErrEmptyCredentialVersion = errors.New("empty credential version")
	ErrEmptyPartnerTag        = errors.New("empty partner tag")
	ErrEmptyMarketplace       = errors.New("empty marketplace")
	ErrNilHTTPClient          = errors.New("nil http client")
)

type Client struct {
	credentialID      string
	credentialSecret  string
	credentialVersion string
	partnerTag        string
	marketplace       string
	httpClient        *http.Client

	mu             sync.Mutex
	accessToken    string
	tokenExpiry    time.Time
	inflight       *tokenFetchState
	lastTokenError error
	retryTokenAt   time.Time
}

type tokenFetchState struct {
	done chan struct{}
	err  error
}

type Config struct {
	CredentialID      string
	CredentialSecret  string
	CredentialVersion string
	PartnerTag        string
	Marketplace       string
	HTTPClient        *http.Client
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.CredentialID == "" {
		return nil, ErrEmptyCredentialID
	}
	if cfg.CredentialSecret == "" {
		return nil, ErrEmptyCredentialSecret
	}
	if cfg.CredentialVersion == "" {
		return nil, ErrEmptyCredentialVersion
	}
	if cfg.PartnerTag == "" {
		return nil, ErrEmptyPartnerTag
	}
	if cfg.Marketplace == "" {
		return nil, ErrEmptyMarketplace
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{
		credentialID:      cfg.CredentialID,
		credentialSecret:  cfg.CredentialSecret,
		credentialVersion: cfg.CredentialVersion,
		partnerTag:        cfg.PartnerTag,
		marketplace:       cfg.Marketplace,
		httpClient:        httpClient,
	}, nil
}

func (c *Client) SetHTTPClient(httpClient *http.Client) error {
	if httpClient == nil {
		return ErrNilHTTPClient
	}

	c.mu.Lock()
	c.httpClient = httpClient
	c.mu.Unlock()

	return nil
}

func (c *Client) token(ctx context.Context) (string, error) {
	for {
		c.mu.Lock()
		now := time.Now()

		if c.accessToken != "" && now.Before(c.tokenExpiry) {
			token := c.accessToken
			c.mu.Unlock()
			return token, nil
		}

		if c.inflight != nil {
			inflight := c.inflight
			c.mu.Unlock()

			select {
			case <-inflight.done:
			case <-ctx.Done():
				return "", ctx.Err()
			}
			continue
		}

		if c.lastTokenError != nil && now.Before(c.retryTokenAt) {
			err := c.lastTokenError
			c.mu.Unlock()
			return "", err
		}

		inflight := &tokenFetchState{done: make(chan struct{})}
		c.inflight = inflight
		c.mu.Unlock()

		token, expiry, err := c.fetchAccessToken(ctx)

		c.mu.Lock()
		if c.inflight == inflight {
			c.inflight = nil
		}
		if err == nil {
			c.accessToken = token
			c.tokenExpiry = expiry
			c.lastTokenError = nil
			c.retryTokenAt = time.Time{}
		} else if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			c.lastTokenError = nil
			c.retryTokenAt = time.Time{}
		} else {
			c.lastTokenError = err
			c.retryTokenAt = time.Now().Add(tokenRetryBackoff)
		}
		inflight.err = err
		close(inflight.done)
		c.mu.Unlock()

		if err != nil {
			return "", err
		}

		return token, nil
	}
}

func (c *Client) httpClientSnapshot() *http.Client {
	c.mu.Lock()
	httpClient := c.httpClient
	c.mu.Unlock()

	return httpClient
}
