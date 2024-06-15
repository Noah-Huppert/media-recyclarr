package models

import (
	"gorm.io/gorm"
	"time"
)

// MediaRequest is an intent to download a piece of media on behalf of a user
type MediaRequest struct {
	gorm.Model

	// RequesterExternalID is the ID of the request in the external media request management service
	RequesterExternalID string

	// RequestedByUserID is the ID of the user who created the request
	RequestedByUserID int

	// RequestedByUser is an ORM filled field based on RequestedByUserID
	RequestedByUser User

	// AvailableAt is when the media became available to watch
	AvailableAt time.Time
}
