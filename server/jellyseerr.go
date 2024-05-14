package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Noah-Huppert/media-recyclarr/jellyseerrclient"
)

// JellyseerrManager wraps Jellyseerr API calls
type JellyseerrManager struct {
	// apiClient is used to make API calls
	apiClient jellyseerrclient.APIClient
}

// NewJellyseerrManagerOpts are options for NewJellyseerrManager
type NewJellyseerrManagerOpts struct {
	// JellyseerrURL is the URL to the Jellyseerr server
	JellyseerrURL string

	// JellyseerrAPIKey is the API key used to authenticate with Jellyseerr
	JellyseerrAPIKey string
}

// NewJellyseerrManager creates a new JellyseerrManager
func NewJellyseerrManager(opts NewJellyseerrManagerOpts) (*JellyseerrManager, error) {
	// Parse URL
	jellyURL, err := url.Parse(opts.JellyseerrURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Jellyseerr URL: %s", err)
	}

	apiClient := jellyseerrclient.NewAPIClient(&jellyseerrclient.Configuration{
		Scheme: jellyURL.Scheme,
		Host:   jellyURL.Host,
		DefaultHeader: map[string]string{
			"X-Api-Key": opts.JellyseerrAPIKey,
		},
	})

	return &JellyseerrManager{
		apiClient: *apiClient,
	}, nil
}

// GetRequestMedia retrieves all pieces of media which are available
func (jellyMgr *JellyseerrManager) GetRequestMedia(ctx context.Context) ([]RequestedMedia, error) {
	items, err := AllPages(AllPagesOpts[
		jellyseerrclient.RequestAPIRequestGetRequest,
		jellyseerrclient.UserUserIdRequestsGet200Response,
		jellyseerrclient.MediaRequest,
	]{
		Req: jellyMgr.apiClient.RequestAPI.RequestGet(ctx),
		ToOAPIResp: func(from jellyseerrclient.UserUserIdRequestsGet200Response) (*OpenAPIResponse[jellyseerrclient.MediaRequest], error) {
			return &OpenAPIResponse[jellyseerrclient.MediaRequest]{
				PageInfo: &OpenAPIResponsePageInfo{
					Page:    from.PageInfo.Page,
					Pages:   from.PageInfo.Pages,
					Results: from.PageInfo.Results,
				},
				Results: from.Results,
			}, nil
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get all pages: %s", err)
	}

	return nil, nil
}
