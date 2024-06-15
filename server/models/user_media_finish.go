package models

import "gorm.io/gorm"

// UserMediaFinish records a user finishing a piece of media
type UserMediaFinish struct {
	gorm.Model

	// UserID is the ID of the user who requested the media
	UserID int

	// User is an ORM populated by UserID
	User User

	// MediaRequestID is the ID of the media request
	MediaRequestID int

	// MediaRequest is an ORM populated field by MediaRequestID
	MediaRequest MediaRequest
}
