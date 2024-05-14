package main

import (
	"fmt"
	"net/url"

	"github.com/Noah-Huppert/media-recyclarr/embyclient"
)

// EmbyManager interacts with the Emby API client
type EmbyManager struct {
	// apiClient is the Emby API client
	apiClient embyclient.APIClient
}

// NewEmbyManagerOpts are options for NewEmbyManager
type NewEmbyManagerOpts struct {
	// EmbyURL is the Emby API server URL
	EmbyURL string

	// EmbyAPIKey is the Emby API key
	EmbyAPIKey string
}

// NewEmbyManager creates a new EmbyManager
func NewEmbyManager(opts NewEmbyManagerOpts) (*EmbyManager, error) {
	// Parse URL
	embyURL, err := url.Parse(opts.EmbyURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Emby URL: %s", err)
	}

	apiClient := embyclient.NewAPIClient(&embyclient.Configuration{
		Scheme: embyURL.Scheme,
		Host:   embyURL.Host,
		DefaultHeader: map[string]string{
			"X-Emby-Token": opts.EmbyAPIKey,
		},
	})

	return &EmbyManager{
		apiClient: *apiClient,
	}, nil
}
