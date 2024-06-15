package models

import "gorm.io/gorm"

// MediaType is the kind of media
type MediaType string

const (
	Movie   MediaType = "movie"
	Series  MediaType = "series"
	Season  MediaType = "season"
	Episode MediaType = "episode"
)

// Media item
type Media struct {
	gorm.Model

	// LibraryExternalID is the ID of the requested media in the external media library
	LibraryExternalID string

	// Name is a user friendly name of the piece of media
	Name string

	// Type is the kind of media
	Type MediaType

	// NodePath is a Postgres LTree path organizing the media in a tree
	// This is used to record child media like seasons and episodes of tv shows
	NodePath string
}
