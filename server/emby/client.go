package emby

import (
	"fmt"

	"github.com/Noah-Huppert/media-recyclarr/restclient"
)

// EmbyClient interacts with the Emby API client
type EmbyClient struct {
	// apiClient is used to make rest HTTP calls
	apiClient *restclient.APIClient
}

// NewEmbyClientOpts are options for NewEmbyManager
type NewEmbyClientOpts struct {
	// EmbyURL is the Emby API server URL
	EmbyURL string

	// EmbyAPIKey is the Emby API key
	EmbyAPIKey string
}

// NewEmbyClient creates a new EmbyManager
func NewEmbyClient(opts NewEmbyClientOpts) (*EmbyClient, error) {
	apiClient, err := restclient.NewAPIClient(restclient.NewAPIClientOpts{
		BaseURL: opts.EmbyURL,
		Headers: map[string]string{
			"X-Emby-Token": opts.EmbyAPIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create rest client: %s", err)
	}

	return &EmbyClient{
		apiClient: apiClient,
	}, nil
}
