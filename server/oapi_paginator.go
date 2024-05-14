package main

import "net/http"

// OpenAPIRequest defines methods provided by OpenAPI code generated requests
type OpenAPIRequest[Resp interface{}] interface {
	// Take specifies the number of results to return
	Take(take float32) OpenAPIRequest[Resp]

	// Skip specifies the number of results to skip
	Skip(skip float32) OpenAPIRequest[Resp]

	// Execute sends the request
	Execute() (*Resp, *http.Response, error)
}
