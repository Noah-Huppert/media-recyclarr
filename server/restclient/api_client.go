package restclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
)

// APIClient is a generic client for HTTP APIs
type APIClient struct {
	// logger outputs runtime information
	logger *zap.Logger

	// baseURL of API server
	baseURL url.URL

	// headers to send in every request
	headers map[string]string

	// timeout is how long a request is allowed to run before it is cancelled
	timeout time.Duration

	// validator to use to validate responses
	validator validator.Validate

	// logRequests indicates if details about each request should be logged
	logRequests bool
}

const (
	// DefaultTimeout is the timeout to use if one is not provided to NewAPIClient()
	DefaultTimeout = time.Second * 30
)

// NewAPIClientOpts are options for NewAPIClient()
type NewAPIClientOpts struct {
	// Logger is used to output runtime information
	Logger *zap.Logger

	// BaseURL of API server
	BaseURL string

	// Headers to include in every request
	Headers map[string]string

	// Timeout is the length of time a request is allowed to last before being cancelled, defaults to DefaultTimeout
	Timeout *time.Duration

	// Validator is used to validate responses, reads validate tags in the Resp struct, if not provided a default one is made
	Validator *validator.Validate

	// LogRequests indicates if requests should be logged, defaults to false
	LogRequests *bool
}

func NewAPIClient(opts NewAPIClientOpts) (*APIClient, error) {
	baseURL, err := url.Parse(opts.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL as URL: %s", err)
	}

	timeout := DefaultTimeout
	if opts.Timeout != nil {
		timeout = *opts.Timeout
	}

	validatorInst := opts.Validator
	if opts.Validator == nil {
		validatorInst = validator.New(validator.WithRequiredStructEnabled())
	}

	logRequests := false
	if opts.LogRequests != nil {
		logRequests = *opts.LogRequests
	}

	return &APIClient{
		logger:      opts.Logger,
		baseURL:     *baseURL,
		headers:     opts.Headers,
		timeout:     timeout,
		validator:   *validatorInst,
		logRequests: logRequests,
	}, nil
}

// MakeRequestOpts are options for making a REST request
type MakeRequestOpts struct {
	// Ctx is the context to use for the request
	Ctx context.Context

	// Method is the HTTP method
	Method string

	// Path is the request path
	Path string

	// QueryParams are optional query params
	QueryParams map[string]string

	// Body is an optional request body which will be serialized as JSON
	Body interface{}

	// Resp is an optional variable to put the JSON response into
	Resp interface{}

	// Validate indicates if the response should be validated, defaults to True
	Validate *bool
}

// CheckCanValidate determines if the input object can be validated
func (c *APIClient) CheckCanValidate(obj interface{}) bool {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	return val.Kind() == reflect.Struct
}

// Validate a struct
func (c *APIClient) Validate(ctx context.Context, obj interface{}) error {
	if !c.CheckCanValidate(obj) {
		return fmt.Errorf("cannot validate a non-struct")
	}

	return c.validator.StructCtx(ctx, obj)
}

// MakeRequest executes an HTTP request
// Note: Cannot validate opts.Resp if it isn't a struct, be sure to set opts.Validate = false if this is the case and
// then call APIClient.Validate() with a type of struct
func (c *APIClient) MakeRequest(opts MakeRequestOpts) error {
	// Check pre-conditions
	shouldValidate := true
	if opts.Validate != nil {
		shouldValidate = *opts.Validate
	}

	if opts.Resp != nil && !c.CheckCanValidate(opts.Resp) && shouldValidate {
		return fmt.Errorf("cannot validate a non-struct response object")
	}

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
	if opts.QueryParams != nil {
		query := reqURL.Query()

		for k, v := range opts.QueryParams {
			query.Set(k, v)
		}

		reqURL.RawQuery = query.Encode()
	}

	if c.logRequests {
		c.logger.Debug("request URL", zap.String("URL", reqURL.String()), zap.String("Path", opts.Path))
	}

	ctx, _ := context.WithDeadline(opts.Ctx, time.Now().Add(c.timeout))

	req, err := http.NewRequestWithContext(ctx, opts.Method, reqURL.String(), bodyReader)
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
		return fmt.Errorf("response had non-OK status code %s: %s", resp.Status, respBody)
	}

	// Get response into provided struct
	if opts.Resp != nil && shouldValidate {
		// Unmarshall
		err = json.Unmarshal(respBody, opts.Resp)
		if err != nil {
			return fmt.Errorf("failed to unmarshall response as JSON: %s", err)
		}

		// Validate
		err = c.validator.StructCtx(opts.Ctx, opts.Resp)
		if err != nil {
			return fmt.Errorf("failed to validate response: %s", err)
		}
	}

	return nil
}
