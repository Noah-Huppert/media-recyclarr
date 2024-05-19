package emby

import (
	"fmt"
	"golang.org/x/net/context"
	"strings"

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

// PaginatableResponse is the format in which the Emby API will return a page of results
type PaginatableResponse[Item interface{}] struct {
	// Items returned for page
	Items []Item `validate:"required"`

	// TotalRecordCount is the total number of items available
	TotalRecordCount int `validate:"required"`
}

// MediaItem contains details about a piece of media in the server
type MediaItem struct {
	// ID is the unique ID of the media item in the server
	ID string `json:"Id" validate:"required"`

	// Name is the name of the media item
	Name string `validate:"required"`

	// IndexNumber is the episode number (starting at 1) if the media item is part of a series
	IndexNumber *int

	// ParentIndexNumber is the season number (starting at 1) if the media item is part of a series
	ParentIndexNumber *int

	// IsFolder indicates if the media item is actually made up multiple child media items (ex., a TV series or season)
	IsFolder bool `validate:"required"`

	// Type of media
	Type string `validate:"required oneof=Episode,Season,Movie"`
}

// UserMediaItem is a media item with details specific to a user
type UserMediaItem struct {
	MediaItem

	// UserData is user specific information about the media item
	UserData struct {
		// Played indicates if the user has watched the media item
		Played bool `validate:"required"`
	} `validate:"required"`
}

func (client *EmbyClient) GetUserMediaItems(ctx context.Context, userID string, ids []string) ([]UserMediaItem, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("ids cannot be empty")
	}

	var resp PaginatableResponse[UserMediaItem]
	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:    ctx,
		Method: "GET",
		Path:   fmt.Sprintf("/emby/Users/%s/items", userID),
		QueryParams: map[string]string{
			"ids": strings.Join(ids, ","),
		},
		Resp: &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	if resp.TotalRecordCount != len(resp.Items) {
		return nil, fmt.Errorf("no pagination configured but results do not contain all items")
	}

	return resp.Items, nil
}
