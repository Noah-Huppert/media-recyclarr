package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Noah-Huppert/media-recyclarr/emby"
	"github.com/Noah-Huppert/media-recyclarr/jelly"
)

// Trasher is responsible for finding media which should be elimated
type Trasher struct {
	// embyMgr is the Emby manager to use
	embyClient *emby.EmbyClient

	// jellyClient is the Jellyseerr manager to use
	jellyClient *jelly.JellyClient

	// expireAfter is the amount of time after media is available and the most recent time a user watched it when the media will be considered expired
	expireAfter time.Duration
}

// NewTrasherOpts are opts for NewTrasher
type NewTrasherOpts struct {
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
	PlayedBy []User

	// AvailableAt is when the media became available to watch
	AvailableAt time.Time

	// Children are pieces of media which are encompassed by this piece of media (ex., special features or episodes)
	Children []RequestedMedia
}

// GetRequestedMedia returns RequestedMedia
func (trasher *Trasher) GetRequestedMedia(ctx context.Context) ([]RequestedMedia, error) {
	// Find media requests in Jellyseerr
	mediaRequests, err := trasher.jellyClient.GetAvailableMediaRequests(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get media requests: %s", err)
	}

	media := []RequestedMedia{}
	for _, req := range mediaRequests {
		media = append(media, RequestedMedia{
			JellyseerrID: req.ID,
			EmbyID:       req.Media.JellyfinMediaID,
			RequestedBy: User{
				EmbyID: req.RequestedBy.JellyfinUserID,
			},
			AvailableAt: req.Media.MediaAddedAt,
		})
	}

	// Get watch status from Emby
	return nil, nil
}

// GetExpiredMedia returns all RequestedMedia which has passed the expired deadline
func (trasher *Trasher) GetExpiredMedia(ctx context.Context) ([]RequestedMedia, error) {
	return nil, nil
}
