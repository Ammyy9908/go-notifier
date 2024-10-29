package api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// APIClient defines the interface for making HTTP requests
type IAPIClient interface {
	Request(method, path string, body interface{}) (*http.Response, error)
	Get(path string) ([]byte, error)
	Post(path string, body interface{}) (*http.Response, error)
	Put(path string, body interface{}) (*http.Response, error)
	Delete(path string) (*http.Response, error)
	SetBaseURL(url string)
}

// Client represents an HTTP client for making API calls
type Client struct {
	baseURL    string
	httpClient *http.Client
	headers    map[string]string
}

// ClientOption is a function type to configure the Client
type ClientOption func(*Client)

// WithTimeout sets the client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHeader adds a header to all requests
func WithHeader(key, value string) ClientOption {
	return func(c *Client) {
		c.headers[key] = value
	}
}

// NewClient creates a new API client
func NewClient(baseURL string, options ...ClientOption) *Client {
	client := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: make(map[string]string),
	}

	// Apply options
	for _, option := range options {
		option(client)
	}

	return client
}

// SetBaseURL sets the base URL for the client
func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
}

// Request makes an HTTP request and returns the response
func (c *Client) Request(method, path string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	req.Header.Set("Content-Type", "application/json")

	// Set custom headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return resp, nil
}

// Get makes a GET request
func (c *Client) Get(path string) ([]byte, error) {
	response, err := c.Request(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read and return the response body as a byte slice
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

// Post makes a POST request
func (c *Client) Post(path string, body interface{}) (*http.Response, error) {
	return c.Request(http.MethodPost, path, body)
}

// Put makes a PUT request
func (c *Client) Put(path string, body interface{}) (*http.Response, error) {
	return c.Request(http.MethodPut, path, body)
}

// Delete makes a DELETE request
func (c *Client) Delete(path string) (*http.Response, error) {
	return c.Request(http.MethodDelete, path, nil)
}
