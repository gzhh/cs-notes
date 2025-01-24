package functionaloptions

import (
	"net/http"
	"time"
)

// UserClient implements a client for the users' service.
type (
	UserClient struct {
		httpClient *http.Client

		maxRetries   uint
		timeout      time.Duration
		baseURL      string
		authUsername string
		authPassword string
	}

	// UserClientOption allows customising the underlying UserClient.
	UserClientOption func(*UserClient)
)

// WithMaxRetries allows customising the max retries.
func WithMaxRetries(retries uint) UserClientOption {
	return func(c *UserClient) {
		c.maxRetries = retries
	}
}

// WithTimeout allows customising the timeout.
func WithTimeout(timeout time.Duration) UserClientOption {
	return func(c *UserClient) {
		c.timeout = timeout
		// Also update the HTTP client's timeout
		c.httpClient.Timeout = timeout
	}
}

// WithBaseURL allows customising the base url.
func WithBaseURL(url string) UserClientOption {
	return func(c *UserClient) {
		c.baseURL = url
	}
}

// WithAuth allows customising the aith.
func WithAuth(username, password string) UserClientOption {
	return func(c *UserClient) {
		c.authUsername = username
		c.authPassword = password
	}
}

// WithHTTPClient allows customising the http client.
func WithHTTPClient(client *http.Client) UserClientOption {
	return func(c *UserClient) {
		c.httpClient = client
	}
}

// NewUserClient returns a new pointer to a UserClient.
func NewUserClient(opts ...UserClientOption) *UserClient {
	// Set up defaults
	client := &UserClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second, // Default timeout
		},
		maxRetries:   3, // Default to 3 retries
		timeout:      30 * time.Second,
		baseURL:      "http://localhost:8080", // Default base URL
		authUsername: "",                      // No auth by default
		authPassword: "",
	}

	// Apply all options
	for _, opt := range opts {
		opt(client)
	}

	return client
}
