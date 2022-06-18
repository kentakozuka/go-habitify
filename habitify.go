package habitify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// DefaultEndpoint is the default endpoint of Habitify API.
const DefaultEndpoint = "https://api.habitify.me"

const (
	urlHabits = "/habits"
	urlNotes  = "/notes"
)

type apiResponse struct {
	Message string          `json:"message"`
	Version string          `json:"version"`
	Status  bool            `json:"status"`
	Data    json.RawMessage `json:"data"`
}

// Client is a client for Habitify API.
type Client struct {
	httpClient *http.Client
	apiKey     string
	endpoint   string
}

// Option is an option for Client.
type Option func(*Client)

// New returns a new client associated with given api key.
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		endpoint:   DefaultEndpoint,
		apiKey:     apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithHTTPClient allows you to pass your http client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithEndpoint allows you to set an endpoint.
func WithEndpoint(endpoint string) Option {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

type httpResponse struct {
	*http.Response
}

func (resp *httpResponse) DecodeJSON(data interface{}) error {
	if err := json.NewDecoder(resp.Response.Body).Decode(data); err != nil {
		return fmt.Errorf("decoding JSON data: %w", err)
	}

	return nil
}

// Close closes the client.
func (resp *httpResponse) Close() {
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
}

func (c *Client) do(ctx context.Context, method, path string, body io.Reader, data interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, c.endpoint+path, body)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "application/json; charset=utf8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	httpResp := &httpResponse{Response: resp}
	var apiResp apiResponse
	if err := httpResp.DecodeJSON(&apiResp); err != nil {
		return err
	}
	if !apiResp.Status {
		return newError(resp.StatusCode, apiResp.Message)
	}

	if err := json.Unmarshal(apiResp.Data, data); err != nil {
		return fmt.Errorf("decoding JSON data: %w", err)
	}

	return nil
}

func (c *Client) get(ctx context.Context, path string, data interface{}) error {
	return c.do(ctx, http.MethodGet, path, nil, data)
}

func (c *Client) post(ctx context.Context, path string, body interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return err
	}

	return c.do(ctx, http.MethodPost, path, &buf, nil)
}
