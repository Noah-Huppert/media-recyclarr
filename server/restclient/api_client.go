package restclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
)

// APIClient is a generic client for HTTP APIs
type APIClient struct {
	baseURL   url.URL
	headers   map[string]string
	validator validator.Validate
}

type NewAPIClientOpts struct {
	BaseURL   string
	Headers   map[string]string
	Validator *validator.Validate
}

func NewAPIClient(opts NewAPIClientOpts) (*APIClient, error) {
	baseURL, err := url.Parse(opts.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL as URL: %s", err)
	}

	if opts.Validator == nil {
		opts.Validator = validator.New(validator.WithRequiredStructEnabled())
	}

	return &APIClient{
		baseURL:   *baseURL,
		headers:   opts.Headers,
		validator: *opts.Validator,
	}, nil
}

type MakeRequestOpts struct {
	Ctx         context.Context
	Method      string
	Path        string
	QueryParams map[string]string
	Body        interface{}
	Resp        interface{}
}

func (c *APIClient) MakeRequest(opts MakeRequestOpts) error {
	// Setup request
	var bodyReader io.Reader = nil
	if opts.Body != nil {
		jsonBody, err := json.Marshal(opts.Body)
		if err != nil {
			return fmt.Errorf("failed to marshal body into JSON: %s", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	reqURL := c.baseURL.JoinPath(opts.Path)
	for k, v := range opts.QueryParams {
		reqURL.Query().Set(k, v)
	}

	req, err := http.NewRequestWithContext(opts.Ctx, opts.Method, reqURL.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to setup request: %s", err)
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %s", err)
	}

	// Check response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("response had non-OK status code %d %s: %s", resp.StatusCode, resp.Status, respBody)
	}

	// Get response into provided struct
	if opts.Resp != nil {
		// Unmarshall
		err = json.Unmarshal(respBody, opts.Resp)
		if err != nil {
			return fmt.Errorf("failed to unmarshall response as JSON: %s", err)
		}

		// Validate
		err = c.validator.StructCtx(opts.Ctx, opts.Resp)
		return fmt.Errorf("failed to validate response: %s", err)
	}

	return nil
}
