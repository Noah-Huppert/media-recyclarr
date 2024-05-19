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

// CheckItemsCount ensures that all the items where returned in one page. As right now we don't implement pagination.
func (resp PaginatableResponse[Item]) CheckItemsCount() error {
	if resp.TotalRecordCount != len(resp.Items) {
		return fmt.Errorf("no pagination configured but results do not contain all items")
	}

	return nil
}

const (
	MediaItemTypeSeries  = "Series"
	MediaItemTypeSeason  = "Season"
	MediaItemTypeEpisode = "Episode"
	MediaItemTypeMovie   = "Movie"
)

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

	// SeriesID is the ID of the parent series if the media item is part of a series
	SeriesID *string `json:"SeriesId"`

	// Type of media
	Type string `validate:"required oneof=Series,Season,Episode,Movie"`
}

// ListMediaItemsOpts are query options for ListMediaItems
type ListMediaItemsOpts struct {
	// Recursive indicates items should be recursively searched for
	Recursive *bool

	// IncludeItemTypes are media item types to include in results
	IncludeItemTypes []string

	// IDs of media items to return
	IDs []string
}

// ListMediaItems retrieves media items
func (client *EmbyClient) ListMediaItems(ctx context.Context, opts ListMediaItemsOpts) ([]MediaItem, error) {
	queryParams := map[string]string{}
	if opts.Recursive != nil {
		queryParams["Recursive"] = fmt.Sprint(opts.Recursive)
	}
	if opts.IncludeItemTypes != nil {
		queryParams["IncludeItemTypes"] = strings.Join(opts.IncludeItemTypes, ",")
	}
	if opts.IDs != nil {
		queryParams["Ids"] = strings.Join(opts.IDs, ",")
	}

	var resp PaginatableResponse[MediaItem]
	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:         ctx,
		Method:      "GET",
		Path:        "/emby/Items",
		QueryParams: queryParams,
		Resp:        &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	if err := resp.CheckItemsCount(); err != nil {
		return nil, err
	}

	return resp.Items, nil
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

// ListUserMediaItems retrieves user details about media items
func (client *EmbyClient) ListUserMediaItems(ctx context.Context, userID string, ids []string) ([]UserMediaItem, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("ids cannot be empty")
	}

	var resp PaginatableResponse[UserMediaItem]
	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:    ctx,
		Method: "GET",
		Path:   fmt.Sprintf("/emby/Users/%s/items", userID),
		QueryParams: map[string]string{
			"Ids": strings.Join(ids, ","),
		},
		Resp: &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	if err := resp.CheckItemsCount(); err != nil {
		return nil, err
	}

	return resp.Items, nil
}

// User is an Emby user
type User struct {
	// ID of user
	ID string `json:"Id" validate:"required"`

	// Name of user
	Name string `validate:"required"`

	// Policy contains permission details for user
	Policy struct {
		// IsAdministrator indicates if the user is a server admin
		IsAdministrator bool `validate:"required"`
	} `validate:"required"`
}

// ListUsers retrieves all users
func (client *EmbyClient) ListUsers(ctx context.Context) ([]User, error) {
	var resp PaginatableResponse[User]

	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:    ctx,
		Method: "GET",
		Path:   "/emby/Users/Query",
		Resp:   &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	if err := resp.CheckItemsCount(); err != nil {
		return nil, err
	}

	return resp.Items, nil
}

// ListShowSeasons retrieves seasons for a TV show
func (client *EmbyClient) ListShowSeasons(ctx context.Context, showID string) ([]MediaItem, error) {
	var resp PaginatableResponse[MediaItem]

	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:    ctx,
		Method: "GET",
		Path:   fmt.Sprintf("/emby/Shows/%d/Seasons", showID),
		Resp:   &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	if err := resp.CheckItemsCount(); err != nil {
		return nil, err
	}

	return resp.Items, nil
}

// ListShowEpisodes retrieves episodes for a TV show
func (client *EmbyClient) ListShowEpisodes(ctx context.Context, showID string, seasonID string) ([]MediaItem, error) {
	var resp PaginatableResponse[MediaItem]

	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:    ctx,
		Method: "GET",
		Path:   fmt.Sprintf("/emby/Shows/%d/Episodes", showID),
		QueryParams: map[string]string{
			"SeasonId": seasonID,
		},
		Resp: &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	if err := resp.CheckItemsCount(); err != nil {
		return nil, err
	}

	return resp.Items, nil
}

// MediaItemNode is a media item which can hold child media items
type MediaItemNode struct {
	MediaItem

	// Children are media items which belong to the parent media item, like episodes of a season
	Children []MediaItemNode
}

// NewMediaItemNode makes a new MediaItemNode from a MediaItem
func NewMediaItemNode(item MediaItem) MediaItemNode {
	return MediaItemNode{
		MediaItem: item,
		Children:  []MediaItemNode{},
	}
}

// NewMediaItemNodeArray makes a new array of MediaItemNode from an array of MediaItem
func NewMediaItemNodeArray(items []MediaItem) []MediaItemNode {
	nodes := []MediaItemNode{}

	for _, one_item := range items {
		nodes = append(nodes, NewMediaItemNode(one_item))
	}

	return nodes
}

// GetMediaTree recursively retrieves all the media items and their children
// Optionally parent can be specified to start building the tree with the children of the specified media item
func (client *EmbyClient) GetMediaTree(ctx context.Context, parent *MediaItemNode) ([]MediaItemNode, error) {
	children := []MediaItemNode{}

	// Either get all media items or start with parent
	if parent == nil {
		// Get all movies and series
		trueVal := true

		items, err := client.ListMediaItems(ctx, ListMediaItemsOpts{
			Recursive: &trueVal,
			IncludeItemTypes: []string{
				MediaItemTypeMovie,
				MediaItemTypeSeries,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to list movies and series: %s", err)
		}

		children = NewMediaItemNodeArray(items)
	} else {
		// Get children of specific type of parent
		if parent.Type == MediaItemTypeSeries {
			// Get seasons of parent series
			seasons, err := client.ListShowSeasons(ctx, parent.ID)
			if err != nil {
				return nil, fmt.Errorf("failed to list seasons for series: %s", err)
			}

			children = NewMediaItemNodeArray(seasons)
		} else if parent.Type == MediaItemTypeSeason {
			// Get episodes of parent season
			if parent.SeriesID == nil {
				return nil, fmt.Errorf("cannot get episodes of season if SeriesID is nil")
			}

			episodes, err := client.ListShowEpisodes(ctx, parent.ID, *parent.SeriesID)
			if err != nil {
				return nil, fmt.Errorf("failed to list episodes for season: %s", err)
			}

			children = NewMediaItemNodeArray(episodes)
		}

		// No other media type has children
	}

	// Recursively get children of each child
	for childIdx, child := range children {
		subChildren, err := client.GetMediaTree(ctx, &child)
		if err != nil {
			return nil, fmt.Errorf("failed to get children of media item ID '%s': %s", child.ID, err)
		}

		children[childIdx].Children = subChildren
	}

	return children, nil
}
