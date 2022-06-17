package habitify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const DefaultEndpoint = "https://api.habitify.me"

const (
	urlHabits = "/habits"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	endpoint   string
}

type Option func(*Client)

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

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

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

func (resp *httpResponse) Close() {
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
}

func (c *Client) do(ctx context.Context, method, path string, body io.Reader) (*httpResponse, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.endpoint+path, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "application/json; charset=utf8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return nil, errors.New("client has issues an invalid request")
	case http.StatusUnauthorized:
		return nil, errors.New("authorization for the API is required but the request has not been authenticated")
	case http.StatusForbidden:
		return nil, errors.New("the request has been authenticated but does not have permission or the resource is not found")
	case http.StatusNotAcceptable:
		return nil, errors.New("the client has requestd a MIM typ via the Accept header for a value not supported by the server")
	case http.StatusUnsupportedMediaType:
		return nil, errors.New("the client has defined a Content-Type header that is not supported by the server")
	case http.StatusUnprocessableEntity:
		return nil, errors.New("the client has made a valid request but the server cannot process it")
	case http.StatusTooManyRequests:
		return nil, errors.New("the client has exceeded the number of requests allowed for a givn time window")
	case http.StatusInternalServerError:
		return nil, errors.New("an unexpected error on the server has occurred")
	}

	return &httpResponse{Response: resp}, nil
}

func (c *Client) get(ctx context.Context, path string) (*httpResponse, error) {
	return c.do(ctx, http.MethodGet, path, nil)
}

func (c *Client) post(ctx context.Context, path string, body interface{}) (*httpResponse, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, err
	}

	return c.do(ctx, http.MethodPost, path, &buf)
}
