package models

import "gorm.io/gorm"

// User is an external user
type User struct {
	gorm.Model

	// LibraryExternalID is the ID of the user in the external media library
	LibraryExternalID string

	// Name is the user friendly name of the user
	Name string

	// IsAdmin records if the user is an admin
	IsAdmin bool

	// FinishedBy are users who have finished the media
	FinishedBy []UserMediaFinish
}
