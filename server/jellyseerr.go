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

// getMediaRequests calls the jellyseerr get media requests endpoint
func (jellyMgr *JellyseerrManager) getMediaRequests(opts OpenAPIRequesterOpts[*jellyseerrclient.PageInfo]) (*OpenAPIRequesterResult[*jellyseerrclient.PageInfo, jellyseerrclient.MediaRequest], error) {
	req := jellyMgr.apiClient.RequestAPI.RequestGet(opts.Ctx).Filter("available")

	res, _, err := req.Take(float32(opts.PageSize)).Skip(float32(opts.PageSize * opts.PageNum)).Execute()
	if err != nil {
		return nil, err
	}

	return &OpenAPIRequesterResult[*jellyseerrclient.PageInfo, jellyseerrclient.MediaRequest]{
		PageInfo: res.PageInfo,
		Items:    res.Results,
	}, nil
}

// GetRequestMedia retrieves all pieces of media which are available
func (jellyMgr *JellyseerrManager) GetRequestMedia(ctx context.Context) ([]RequestedMedia, error) {
	// Get requests from Jellyseerr
	jellyItems, err := AllPages[*jellyseerrclient.PageInfo, jellyseerrclient.MediaRequest]{
		Req: jellyMgr.getMediaRequests,
	}.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get requests from jellyseerr: %s", err)
	}

	// Tranform
	reqMedia := []RequestedMedia{}

	for _, item := range jellyItems {
		var f jellyseerrclient.MediaRequest
		var a jellyseerrclient.MediaInfo
		var b jellyseerrclient.User

		reqMedia = append(reqMedia, RequestedMedia{
			JellyseerrID: item.Id,
			EmbyID: item.Media.,
		})
	}

	return nil, nil
}
