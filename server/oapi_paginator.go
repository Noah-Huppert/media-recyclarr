package main

import (
	"fmt"
	"net/http"
)

// OpenAPIResponsePageInfo contains info about what page of results the response contains
type OpenAPIResponsePageInfo struct {
	// Page number response represents, starts at 1
	Page *float32

	// Pages is the total number of pages, starts at 1
	Pages *float32

	// Total number of results
	Results *float32
}

// OpenAPIResponse is an Open API endpoint response
type OpenAPIResponse[Item interface{}] struct {
	// PageInfo is information about which page of results
	PageInfo *OpenAPIResponsePageInfo

	// Results is the page of data
	Results []Item
}

// ToOpenAPIResponse converts an arbitrary type to an OpenAPIResponse
type ToOpenAPIResponse[From interface{}, Item interface{}] func(from From) (*OpenAPIResponse[Item], error)

// OpenAPIRequester defines methods provided by OpenAPI code generated requests
type OpenAPIRequester[Chain interface{}, Resp interface{}] interface {
	// Take specifies the number of results to return
	Take(take float32) Chain

	// Skip specifies the number of results to skip
	Skip(skip float32) Chain

	// Execute sends the request
	Execute() (*Resp, *http.Response, error)
}

// AllPagesDefaultPageSize is the number of items per page for AllPages() if not specified in AllPagesOpts
const AllPagesDefaultPageSize int = 20

// AllPagesOpts are options for AllPages()
type AllPagesOpts[
	Chain interface{},
	Resp interface{},
	Item interface{},
] struct {
	// Req is the request builder
	Req OpenAPIRequester[Chain, Resp]

	// ToOAPIResp converts the endpoint's response to an OpenAPIResponse
	ToOAPIResp ToOpenAPIResponse[Resp, Item]

	// PageSize optionally indicates how many items should be retrieved per page
	PageSize *int
}

// AllPages keeps calling an OpenAPIRequester until all pages have been consumed.
// Type args:
//
// - Chain is the type used to build request options
// - Resp is the endpoint's response type
// - Item is a single item on a page returned by the endpoint
func AllPages[
	Chain interface{},
	Resp interface{},
	Item interface{},
](opts AllPagesOpts[Chain, Resp, Item]) ([]Item, error) {
	pageSize := AllPagesDefaultPageSize
	if opts.PageSize != nil {
		pageSize = *opts.PageSize
	}

	items := []Item{}

	opts.Req.Take(float32(pageSize))

	pageNum := 0
	maxPage := 0
	for maxPage == 0 || pageNum >= maxPage {
		// Setup pagination
		opts.Req.Skip(float32(pageNum * pageSize))

		// Make request
		aResp, _, err := opts.Req.Execute()
		if err != nil {
			return nil, fmt.Errorf("failed to make request for page %d: %s", pageNum, &err)
		}

		// Convert to OpenAPIResponse
		oapiResp, err := opts.ToOAPIResp(*aResp)
		if err != nil {
			return nil, fmt.Errorf("failed to convert response for page %d into OpenAPIResponse: %s", pageNum, err)
		}

		items = append(items, oapiResp.Results...)

		maxPage = int(*oapiResp.PageInfo.Pages - 1)
		pageNum += 1
	}

	return items, nil
}
