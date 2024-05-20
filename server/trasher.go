package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/Noah-Huppert/media-recyclarr/emby"
	"github.com/Noah-Huppert/media-recyclarr/jelly"
)

// Trasher is responsible for finding media which should be elimated
type Trasher struct {
	// logger is used to output information
	logger *zap.Logger

	// embyMgr is the Emby manager to use
	embyClient *emby.EmbyClient

	// jellyClient is the Jellyseerr manager to use
	jellyClient *jelly.JellyClient

	// expireAfter is the amount of time after media is available and the most recent time a user watched it when the media will be considered expired
	expireAfter time.Duration
}

// NewTrasherOpts are opts for NewTrasher
type NewTrasherOpts struct {
	// Logger is used to output information
	Logger *zap.Logger

	// EmbyClient is the Emby API client
	EmbyClient *emby.EmbyClient

	// JellyClient is the Jellyseerr API client
	JellyClient *jelly.JellyClient

	// ExpireAfter is amount of time after media is available and the most recent time a user watched it when the media will be considered expired
	ExpireAfter time.Duration
}

// NewTrasher creates a new Trasher
func NewTrasher(opts NewTrasherOpts) *Trasher {
	return &Trasher{
		logger:      opts.Logger,
		embyClient:  opts.EmbyClient,
		jellyClient: opts.JellyClient,
		expireAfter: opts.ExpireAfter,
	}
}

// User represents a user who requested media
type User struct {
	// Name is the username
	Name string

	// EmbyID is the ID of the user in Emby
	EmbyID string
}

// UserWatch represents a user who has watched a piece of media
type UserWatch struct {
	User

	// WatchedAt is when the piece of media was watched
	WatchedAt time.Time
}

// UserWatchArray is a list of UserWatch objects
type UserWatchArray []UserWatch

// MaxWatchedAt gets the newest watched at date from the array
// Returns nil if the array is empty
func (arr UserWatchArray) Newest() *UserWatch {
	if len(arr) == 0 {
		return nil
	} else if len(arr) == 1 {
		return &arr[0]
	}

	newest := arr[0]

	for _, item := range arr[1:] {
		if item.WatchedAt.After(newest.WatchedAt) {
			newest = item
		}
	}

	return &newest
}

// RequestedMediaType is the kind of media
type RequestedMediaType string

const (
	Movie   RequestedMediaType = "movie"
	Series  RequestedMediaType = "series"
	Season  RequestedMediaType = "season"
	Episode RequestedMediaType = "episode"
)

// RequestedMediaTypeFromEmbyType converts an emby media type to a RequestedMediaType
func RequestedMediaTypeFromEmbyType(embyType string) (*RequestedMediaType, error) {
	val := Movie
	switch embyType {
	case emby.MediaItemTypeMovie:
		val = Movie
	case emby.MediaItemTypeSeries:
		val = Series
	case emby.MediaItemTypeSeason:
		val = Season
	case emby.MediaItemTypeEpisode:
		val = Episode
	default:
		return nil, fmt.Errorf("unknown emby media type '%s' cannot map to RequestedMediaType", embyType)
	}

	return &val, nil
}

// RequestedMedia represents a piece of media which exists because it was requested and retrieved
type RequestedMedia struct {
	// JellyseerrID is the ID of the request in Jellyseerr
	JellyseerrID uint

	// Name is the user friendly name of the media
	Name string

	// EmbyID is the ID of the media in Emby
	EmbyID string

	// RequestedBy is the user who requested the media
	RequestedBy User

	// PlayedBy are users who have played this piece of media
	PlayedBy UserWatchArray

	// AvailableAt is when the media became available to watch
	AvailableAt time.Time

	// Type is the kind of media
	Type RequestedMediaType

	// Children are pieces of media which are encompassed by this piece of media (ex., special features or episodes)
	Children RequestedMediaArray
}

// NewestWatched gets the most recent UserWatch out of this node and all its children
func (media RequestedMedia) NewestWatched() *UserWatch {
	newest := media.PlayedBy.Newest()

	for _, child := range media.Children {
		childNewest := child.NewestWatched()
		if childNewest == nil {
			continue
		}

		if newest == nil || childNewest.WatchedAt.After(newest.WatchedAt) {
			newest = childNewest
		}
	}

	return newest
}

// RequestedMediaArray is a list of RequestedMedia
type RequestedMediaArray []RequestedMedia

func (arr RequestedMediaArray) FormatTree(depth int) []string {
	out := []string{}

	for _, node := range arr {
		watchedStrs := []string{}
		for _, watched := range node.PlayedBy {
			watchedStrs = append(watchedStrs, fmt.Sprintf("%s %s", watched.Name, watched.WatchedAt))
		}

		out = append(out, fmt.Sprintf("%s%s (%s)", strings.Repeat("  ", depth), node.Name, strings.Join(watchedStrs, ", ")))
		out = append(out, node.Children.FormatTree(depth+1)...)
	}

	return out
}

// EmbyIDMap returns a map where keys are emby IDs and values are references to requested media
func (arr RequestedMediaArray) EmbyIDMap() map[string]*RequestedMedia {
	idMap := map[string]*RequestedMedia{}

	for _, media := range arr {
		idMap[media.EmbyID] = &media

		childMap := media.Children.EmbyIDMap()
		for id, childMedia := range childMap {
			idMap[id] = childMedia
		}
	}

	return idMap
}

// requestedMediaBuilder builds a tree of RequestedMedia
// It holds information from jellyseerr and emby APIs
type requestedMediaBuilder struct {
	// mediaRequests are Jellyseerr media requests
	mediaRequests []jelly.MediaRequest

	// mediaRequestsByID is a map of jellyseerr media requests references organized by ID
	mediaRequestsByID map[uint]*jelly.MediaRequest

	// embyUsersByID is a map of emby users where keys are user IDs
	embyUsersByID map[string]emby.User

	// mediaTree is a tree of emby media items
	mediaTree emby.MediaItemNodeArray

	// mediaTreeIDMap is a tree of references to nodes in mediaTree, where keys are media item IDs
	mediaTreeIDMap emby.MediaItemNodeIDMap

	// mediaWatchedBy indicates which Emby users watched which pieces of media
	// First level keys are emby media IDs
	// Second level keys are emby user IDs
	mediaWatchedBy map[string]map[string]UserWatch
}

// buildNode creates a RequestedMedia object for the specified emby media
func (builder *requestedMediaBuilder) buildNode(embyMediaID string, jellyRequestID uint) (*RequestedMedia, error) {
	// Ensure media is found in our intermediary structures
	mediaReq, ok := builder.mediaRequestsByID[jellyRequestID]
	if !ok {
		return nil, fmt.Errorf("no Jellyseer media request found with id '%s'", jellyRequestID)
	}

	media, ok := builder.mediaTreeIDMap[embyMediaID]
	if !ok {
		return nil, fmt.Errorf("no Emby media item found with id '%s'", embyMediaID)
	}

	// Assemble
	playedBy := []UserWatch{}
	for _, user := range builder.mediaWatchedBy[embyMediaID] {
		playedBy = append(playedBy, user)
	}

	mediaType, err := RequestedMediaTypeFromEmbyType(media.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to get media type: %s", err)
	}

	requested := RequestedMedia{
		JellyseerrID: mediaReq.ID,
		Name:         media.Name,
		EmbyID:       embyMediaID,
		RequestedBy: User{
			Name:   builder.embyUsersByID[mediaReq.RequestedBy.JellyfinUserID].Name,
			EmbyID: mediaReq.RequestedBy.JellyfinUserID,
		},
		PlayedBy:    playedBy,
		Type:        *mediaType,
		AvailableAt: *mediaReq.Media.MediaAddedAt,
	}
	requestedChildren := []RequestedMedia{}

	for _, mediaChild := range media.Children {
		child, err := builder.buildNode(mediaChild.ID, jellyRequestID)
		if err != nil {
			return nil, fmt.Errorf("failed to build requested media item emby media item '%s' (Child of emby media item '%s'): %s", mediaChild.ID, media.ID, err)
		}

		requestedChildren = append(requestedChildren, *child)
	}

	requested.Children = requestedChildren

	return &requested, nil
}

// buildTree creates a RequestedMedia object for each known piece of media
func (builder *requestedMediaBuilder) buildTree() (RequestedMediaArray, error) {
	children := []RequestedMedia{}

	for _, jellyReq := range builder.mediaRequests {
		child, err := builder.buildNode(jellyReq.Media.JellyfinMediaID, jellyReq.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to create RequestedMedia object for Emby media '%s': %s", jellyReq.Media.JellyfinMediaID, err)
		}

		children = append(children, *child)
	}

	return children, nil
}

// populateRequestedMediaBuilder makes all necessary API calls and creates all required data structures to create a valid requestedMediaBuilder
func (trasher *Trasher) populateRequestedMediaBuilder(ctx context.Context) (*requestedMediaBuilder, error) {
	// Find media requests in Jellyseerr
	trasher.logger.Debug("getting media requests")
	mediaRequests, err := trasher.jellyClient.GetAvailableMediaRequests(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get media requests: %s", err)
	}

	mediaRequestsByID := map[uint]*jelly.MediaRequest{}

	for _, req := range mediaRequests {
		mediaRequestsByID[req.ID] = &req
	}

	// Get emby users
	embyUsers, err := trasher.embyClient.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of emby users: %s", err)
	}

	embyUsersByID := map[string]emby.User{}
	for _, user := range embyUsers {
		embyUsersByID[user.ID] = user
	}

	// Get tree of all media
	trasher.logger.Debug("getting media tree")
	children, err := trasher.embyClient.GetMediaTree(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve media tree: %s", err)
	}
	leafMediaIDs := children.CollectIDs([]string{
		emby.MediaItemTypeEpisode,
		emby.MediaItemTypeMovie,
	})

	// Get watch status for Emby
	mediaWatchedBy := map[string]map[string]UserWatch{}
	for _, id := range leafMediaIDs {
		mediaWatchedBy[id] = map[string]UserWatch{}
	}

	for _, user := range embyUsers {
		userActivity, err := trasher.embyClient.ListUserPlayActivity(ctx, user.ID, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to Emby get user play activity for user '%s' (%s): %s", user.Name, user.ID, err)
		}

		for _, activity := range userActivity {
			itemID := fmt.Sprint(activity.ItemID)

			if _, ok := mediaWatchedBy[itemID]; !ok {
				continue
			}

			mediaWatchedBy[itemID][activity.UserID] = UserWatch{
				User: User{
					Name:   embyUsersByID[activity.UserID].Name,
					EmbyID: activity.UserID,
				},
				WatchedAt: activity.WatchedAt(),
			}
		}
	}

	return &requestedMediaBuilder{
		mediaRequests:     mediaRequests,
		mediaRequestsByID: mediaRequestsByID,
		embyUsersByID:     embyUsersByID,
		mediaTree:         children,
		mediaTreeIDMap:    children.MediaItemNodeIDMap(),
		mediaWatchedBy:    mediaWatchedBy,
	}, nil
}

// GetRequestedMedia returns RequestedMedia
func (trasher *Trasher) GetRequestedMedia(ctx context.Context) (RequestedMediaArray, error) {
	builder, err := trasher.populateRequestedMediaBuilder(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to gather information about media: %s", err)
	}

	requested, err := builder.buildTree()
	if err != nil {
		return nil, fmt.Errorf("failed to build tree of requested media: %s", err)
	}

	return requested, nil
}

// GetExpiredMedia returns all RequestedMedia which has passed the expired deadline
func (trasher *Trasher) GetExpiredMedia(ctx context.Context, requestedMedia RequestedMediaArray) (RequestedMediaArray, error) {
	requestedMediaByEmbyID := requestedMedia.EmbyIDMap()

	// Get newest watch for each parent piece of media
	// Also get newest watch for any seasons
	newestWatchForEmbyID := map[string]*UserWatch{}
	for _, mediaReq := range requestedMedia {
		newestWatchForEmbyID[mediaReq.EmbyID] = mediaReq.NewestWatched()

		// Get newest watch for any seasons too
		if mediaReq.Type != Series {
			continue
		}

		for _, childMediaReq := range mediaReq.Children {
			if childMediaReq.Type != Season {
				continue
			}

			newestWatchForEmbyID[childMediaReq.EmbyID] = childMediaReq.NewestWatched()
		}
	}

	// Find media which hasn't been watched in too long
	now := time.Now()
	expiredMediaReqs := []RequestedMedia{}

	for embyID, newestWatch := range newestWatchForEmbyID {
		mediaReq := requestedMediaByEmbyID[embyID]

		// If hasn't been watched and available too long
		if newestWatch == nil {
			if now.Sub(mediaReq.AvailableAt) > trasher.expireAfter {
				// Never watched
				expiredMediaReqs = append(expiredMediaReqs, *mediaReq)
			}
		} else if now.Sub(newestWatch.WatchedAt) > trasher.expireAfter {
			// Watched too long ago
			expiredMediaReqs = append(expiredMediaReqs, *mediaReq)
		}
	}

	return expiredMediaReqs, nil
}
