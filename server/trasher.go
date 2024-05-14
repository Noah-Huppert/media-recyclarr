package main

import "time"

// Trasher is responsible for finding media which should be elimated
type Trasher struct {
	// embyMgr is the Emby manager to use
	embyMgr EmbyManager

	// jellyMgr is the Jellyseerr manager to use
	jellyMgr JellyseerrManager
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
