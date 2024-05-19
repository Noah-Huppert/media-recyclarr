package jelly

import (
	"context"
	"fmt"
	"time"

	"github.com/Noah-Huppert/media-recyclarr/restclient"
)

// JellyClient wraps Jellyseerr API calls
type JellyClient struct {
	// apiClient is used to make API calls
	apiClient *restclient.APIClient
}

// NewJellyClientOpts are options for NewJellyseerrManager
type NewJellyClientOpts struct {
	// JellyseerrURL is the URL to the Jellyseerr server
	JellyseerrURL string

	// JellyseerrAPIKey is the API key used to authenticate with Jellyseerr
	JellyseerrAPIKey string
}

// NewJellyClient creates a new JellyseerrManager
func NewJellyClient(opts NewJellyClientOpts) (*JellyClient, error) {
	// Parse URL
	apiClient, err := restclient.NewAPIClient(restclient.NewAPIClientOpts{
		BaseURL: opts.JellyseerrURL,
		Headers: map[string]string{
			"X-Api-Key": opts.JellyseerrAPIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make underlying API client: %s", err)
	}

	return &JellyClient{
		apiClient: apiClient,
	}, nil
}

// PageInfo holds info about what page of results has been returned
type PageInfo struct {
	Pages    int `json:"pages"`
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
}

// PaginatedResp holds paginated results
type PaginatedResp[Item interface{}] struct {
	PageInfo PageInfo `json:"pageInfo" validate:"required"`
	Results  []Item   `json:"results" validate:"required"`
}

// makePaginatedReq is a function which makes a request for a specific page of results
type makePaginatedReq[Item interface{}] func(ctx context.Context, take int, skip int) (*PaginatedResp[Item], error)

// GET_ALL_PAGES_PAGE_SIZE is the page size used in getAllPages()
const GET_ALL_PAGES_PAGE_SIZE = 20

// getAllPages of a paginated endpoint
func getAllPages[Item interface{}](ctx context.Context, makeReq makePaginatedReq[Item]) ([]Item, error) {
	items := []Item{}

	pageNum := 1
	totalPages := -1
	for totalPages < 0 || pageNum < totalPages {
		resp, err := makeReq(
			ctx,
			GET_ALL_PAGES_PAGE_SIZE,
			(pageNum-1)*pageNum,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to make request for page %d: %s", pageNum, err)
		}

		items = append(items, resp.Results...)

		pageNum = resp.PageInfo.Page
		totalPages = resp.PageInfo.Pages
	}

	return items, nil
}

const (

)

// MediaRequest is a request for media in Jellyseerr
type MediaRequest struct {
	ID    uint   `json:"id" validate:"required"`
	Type  string `json:"type" validate:"required,oneof=movie tv"`
	Media struct {
		ID              uint      `json:"id" validate:"required"`
		JellyfinMediaID string    `json:"jellyfindMediaId" validate:"required"`
		MediaAddedAt    time.Time `json:"mediaAddedAt" validate:"required"`
	} `json:"media" validate:"required"`
	RequestedBy struct {
		ID             uint   `json:"id" validate:"required"`
		JellyfinUserID string `json:"jellyfinUserId" validate:"required"`
	} `json:"requestedBy" validate:"required"`
}

// GetAvailableMediaRequestsPage gets a page of media requests
func (client *JellyClient) GetAvailableMediaRequestsPage(ctx context.Context, take, skip int) (*PaginatedResp[MediaRequest], error) {
	var resp PaginatedResp[MediaRequest]

	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:    ctx,
		Method: "GET",
		Path:   "/api/v1/request",
		QueryParams: map[string]string{
			"take":   fmt.Sprint(take),
			"skip":   fmt.Sprint(skip),
			"filter": "available",
		},
		Resp: &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	return &resp, nil
}

// GetAvailableMediaRequests gets all media requests
func (client *JellyClient) GetAvailableMediaRequests(ctx context.Context) ([]MediaRequest, error) {
	// Get requests from Jellyseerr
	mediaReqs, err := getAllPages(
		ctx,
		client.GetAvailableMediaRequestsPage,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get all pages: %s", err)
	}

	return mediaReqs, nil
}
