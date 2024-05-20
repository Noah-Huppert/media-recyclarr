package emby

import (
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"math"
	"strings"
	"time"

	"github.com/Noah-Huppert/media-recyclarr/restclient"
)

// EmbyClient interacts with the Emby API client
type EmbyClient struct {
	// logger is used to output runtime information
	logger *zap.Logger

	// apiClient is used to make rest HTTP calls
	apiClient *restclient.APIClient
}

// NewEmbyClientOpts are options for NewEmbyManager
type NewEmbyClientOpts struct {
	// Logger is used to output runtime information
	Logger *zap.Logger

	// EmbyURL is the Emby API server URL
	EmbyURL string

	// EmbyAPIKey is the Emby API key
	EmbyAPIKey string
}

// NewEmbyClient creates a new EmbyManager
func NewEmbyClient(opts NewEmbyClientOpts) (*EmbyClient, error) {
	apiClient, err := restclient.NewAPIClient(restclient.NewAPIClientOpts{
		Logger:  opts.Logger.Named("api-client"),
		BaseURL: opts.EmbyURL,
		Headers: map[string]string{
			"X-Emby-Token": opts.EmbyAPIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create rest client: %s", err)
	}

	return &EmbyClient{
		logger:    opts.Logger,
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
		queryParams["Recursive"] = fmt.Sprint(*opts.Recursive)
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

// UserMediaItemArray is a list of UserMediaItem
type UserMediaItemArray []UserMediaItem

// CollectAllIDs returns the IDs of all the UserMediaItem objects
func (arr UserMediaItemArray) CollectAllIDs() []string {
	ids := []string{}

	for _, item := range arr {
		ids = append(ids, item.ID)
	}

	return ids
}

// ListUserMediaItemsFilterOpts are filter options for ListUserMediaItems()
type ListUserMediaItemsFilterOpts struct {
	// IDs filter by IDs
	IDs []string

	// IsPlayed filters by media items which have been played by the user
	IsPlayed *bool
}

// ListUserMediaItems retrieves user details about media items
func (client *EmbyClient) ListUserMediaItems(ctx context.Context, userID string, filter ListUserMediaItemsFilterOpts) (UserMediaItemArray, error) {
	queryParams := map[string]string{}

	if filter.IDs != nil {
		queryParams["Ids"] = strings.Join(filter.IDs, ",")
	}
	if filter.IsPlayed != nil {
		queryParams["IsPlayed"] = fmt.Sprint(*filter.IsPlayed)
	}

	var resp PaginatableResponse[UserMediaItem]
	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:         ctx,
		Method:      "GET",
		Path:        fmt.Sprintf("/emby/Users/%s/items", userID),
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
		Path:   fmt.Sprintf("/emby/Shows/%s/Seasons", showID),
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
		Path:   fmt.Sprintf("/emby/Shows/%s/Episodes", showID),
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
	Children MediaItemNodeArray
}

// NewMediaItemNode makes a new MediaItemNode from a MediaItem
func NewMediaItemNode(item MediaItem) MediaItemNode {
	return MediaItemNode{
		MediaItem: item,
		Children:  MediaItemNodeArray{},
	}
}

// MediaItemNodeArray is a list of media item nodes
type MediaItemNodeArray []MediaItemNode

// NewMediaItemNodeArray makes a new array of MediaItemNode from an array of MediaItem
func NewMediaItemNodeArray(items []MediaItem) MediaItemNodeArray {
	nodes := MediaItemNodeArray{}

	for _, one_item := range items {
		nodes = append(nodes, NewMediaItemNode(one_item))
	}

	return nodes
}

// CollectIDs returns the IDs of all media node items.
// If filterType is not nil then will only include IDs of media items of those types. Just because a parent doesn't match a filter type doesn't mean its children's IDs won't be collected
func (arr MediaItemNodeArray) CollectIDs(filterType []string) []string {
	// Build hashmap so filter check is not a linear search of an array
	var filterTypeMap map[string]interface{} = nil
	if filterType != nil {
		filterTypeMap = map[string]interface{}{}
		for _, key := range filterType {
			filterTypeMap[key] = nil
		}
	}

	// Collect IDs
	ids := []string{}
	for _, item := range arr {
		// Collect children Ids before filtering
		ids = append(ids, item.Children.CollectIDs(filterType)...)

		// Filter item based on type
		if filterTypeMap != nil {
			if _, ok := filterTypeMap[item.Type]; !ok {
				continue
			}
		}

		ids = append(ids, item.ID)
	}

	return ids
}

// MediaItemNodeIDMap organizes a tree of MediaItemNode objects by ID
type MediaItemNodeIDMap map[string]*MediaItemNode

// MediaItemNodeIDMap creates a MediaItemNodeIDMap from a MediaItemNodeArray
func (arr MediaItemNodeArray) MediaItemNodeIDMap() map[string]*MediaItemNode {
	idMap := map[string]*MediaItemNode{}

	for _, item := range arr {
		idMap[item.ID] = &item

		for childID, childPtr := range item.Children.MediaItemNodeIDMap() {
			idMap[childID] = childPtr
		}
	}

	return idMap
}

// GetMediaTree recursively retrieves all the media items and their children
// Optionally parent can be specified to start building the tree with the children of the specified media item
func (client *EmbyClient) GetMediaTree(ctx context.Context, parent *MediaItemNode) (MediaItemNodeArray, error) {
	children := MediaItemNodeArray{}

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

			episodes, err := client.ListShowEpisodes(ctx, *parent.SeriesID, parent.ID)
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

// UserPlayActivity is a record of a piece of media being watched
type UserPlayActivity struct {
	// Date on which activity occurred
	Date time.Time `json:"date" validate:"required"`

	// Time at which activity occurred
	Time time.Time `json:"time" validate:"required"`

	// UserID is the ID of the user who watched
	UserID string `json:"user_id" validate:"required"`

	// ItemID is the ID of the item which was watched
	ItemID int `json:"item_id" validate:"required"`

	// ItemType is the type of the item which was watched
	ItemType string `json:"item_type" validate:"required oneof=Series,Season,Episode,Movie"`

	// Duration is the number of seconds the media item was watched
	Duration string `json:"duration" validate:"required"`
}

// WatchedAt combines the date and time in the UserPlayActivity
func (a UserPlayActivity) WatchedAt() time.Time {
	return time.Date(a.Date.Year(), a.Date.Month(), a.Date.Day(), a.Time.Hour(), a.Time.Minute(), a.Time.Second(), a.Time.Nanosecond(), a.Time.Location())
}

// ListUserPlayActivity retrieves a list of all the play activity items for a user
// If lookBackDays is nil then calculates the number of days from today back to EPOCH
func (client *EmbyClient) ListUserPlayActivity(ctx context.Context, userID string, lookBackDays *int) ([]UserPlayActivity, error) {
	if lookBackDays == nil {
		epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, nil)
		daysFromEpoch := int(math.Ceil(time.Now().Sub(epoch).Hours() / 24))
		lookBackDays = &daysFromEpoch
	}

	var resp []UserPlayActivity
	err := client.apiClient.MakeRequest(restclient.MakeRequestOpts{
		Ctx:    ctx,
		Method: "GET",
		Path:   "/emby/user_usage_stats/UserPlaylist",
		QueryParams: map[string]string{
			"user_id": userID,
			"days":    fmt.Sprint(*lookBackDays),
		},
		Resp: &resp,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s", err)
	}

	return resp, nil
}
