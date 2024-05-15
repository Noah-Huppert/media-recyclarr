package main

import (
	"context"
	"fmt"
)

// OpenAPIResponsePageInfo contains info about what page of results the response contains
type OpenAPIResponsePageInfo interface {
	// Page number response represents, starts at 1
	GetPage() float32

	// Pages is the total number of pages, starts at 1
	GetPages() float32

	// Total number of results
	GetResults() float32
}

// OpenAPIRequesterOpts are options provided to an OpenAPIRequester
type OpenAPIRequesterOpts[PageInfo OpenAPIResponsePageInfo] struct {
	// Ctx is the context used
	Ctx context.Context

	// PageSize is the number of items to retrieve per page
	PageSize int

	// PageNum is the page to retrieve, starts at 0
	PageNum int
}

// OpenAPIRequesterResult is the result of retrieving one page of an OpenAPI endpoint
type OpenAPIRequesterResult[PageInfo OpenAPIResponsePageInfo, Item interface{}] struct {
	PageInfo PageInfo
	Items    []Item
}

// OpenAPIRequester retrieves one page from an OpenAPI endpoint
type OpenAPIRequester[PageInfo OpenAPIResponsePageInfo, Item interface{}] func(opts OpenAPIRequesterOpts[PageInfo]) (*OpenAPIRequesterResult[PageInfo, Item], error)

// AllPagesDefaultPageSize is the number of items per page for AllPages() if not specified in AllPagesOpts
const AllPagesDefaultPageSize int = 20

// AllPages keeps calling the OpenAPIRequester untl all pages have been consumed
type AllPages[PageInfo OpenAPIResponsePageInfo, Item interface{}] struct {
	// Req is the request builder
	Req OpenAPIRequester[PageInfo, Item]

	// PageSize optionally indicates how many items should be retrieved per page
	PageSize *int
}

// Execute performs the pagination
func (p AllPages[PageInfo, Item]) Execute(ctx context.Context) ([]Item, error) {
	// Setup request
	pageSize := AllPagesDefaultPageSize
	if p.PageSize != nil {
		pageSize = *p.PageSize
	}

	items := []Item{}

	pageNum := 0
	maxPage := 0
	for maxPage == 0 || pageNum >= maxPage {
		// Make request
		res, err := p.Req(OpenAPIRequesterOpts[PageInfo]{
			PageSize: pageSize,
			PageNum:  pageNum,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to make request for page %d: %s", pageNum, &err)
		}

		// Convert to OpenAPIResponse
		items = append(items, res.Items...)

		maxPage = int(res.PageInfo.GetPages() - 1)
		pageNum += 1
	}

	return items, nil
}
