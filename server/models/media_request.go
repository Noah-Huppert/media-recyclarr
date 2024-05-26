package models

import (
	"gorm.io/gorm"
	"time"
)

// RequestedMediaType is the kind of media
type RequestedMediaType string

const (
	Movie   RequestedMediaType = "movie"
	Series  RequestedMediaType = "series"
	Season  RequestedMediaType = "season"
	Episode RequestedMediaType = "episode"
)

// MediaRequest is an intent to download a piece of media on behalf of a user
type MediaRequest struct {
	gorm.Model

	// RequesterExternalID is the ID of the request in the external media request management service
	RequesterExternalID string

	// LibraryExternalID is the Id of the request in the external media library
	LibraryExternalID string

	// Name is a user friendly name of the piece of media
	Name string

	// RequestedByUserID is the ID of the user who created the request
	RequestedByUserID int

	// RequestedByUser is an ORM filled field based on RequestedByUserID
	RequestedByUser User

	// FinishedBy are users who have finished the media
	FinishedBy []UserMediaFinish

	// AvailableAt is when the media became available to watch
	AvailableAt time.Time

	// Type is the kind of media requested
	Type RequestedMediaType

	// NodePath is a Postgres LTree path organizing the media request in a tree
	// This is used to record child media like seasons and episodes of tv shows
	NodePath string
}
