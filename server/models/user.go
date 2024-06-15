package models

import "gorm.io/gorm"

// User is an external user
type User struct {
	gorm.Model

	// LibraryExternalID is the ID of the user in the external media library
	LibraryExternalID string

	// Name is the user friendly name of the user
	Name string

	// FinishedBy are users who have finished the media
	FinishedBy []UserMediaFinish
}
