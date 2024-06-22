package tasks

import (
	"context"
	"fmt"
	"github.com/Noah-Huppert/media-recyclarr/emby"
	"github.com/Noah-Huppert/media-recyclarr/models"
	"github.com/hashicorp/go-set/v2"
	"gorm.io/gorm"
)

type UpdateUsersTask struct {
	// db is a database client
	db *gorm.DB

	// embyClient is the Emby client
	embyClient *emby.EmbyClient
}

// Run updates users in the database
func (t UpdateUsersTask) Run(ctx context.Context) error {
	// Get users from Emby
	embyUsers, err := t.embyClient.ListUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve users from Emby: %s", err)
	}
	embyUserIDs := set.FromFunc(embyUsers, func(u emby.User) string {
		return u.ID
	})
	embyUsersByID := map[string]emby.User{}
	for _, user := range embyUsers {
		embyUsersByID[user.ID] = user
	}

	// Try to find users in db
	var existingUsers []models.User
	if res := t.db.Model(&models.User{}).Select("library_external_id").Find(&existingUsers); res.Error != nil {
		return fmt.Errorf("failed to retrieve existing user from DB: %s", res.Error)
	}

	inDBUserLibIDs := set.FromFunc(existingUsers, func(u models.User) string {
		return u.LibraryExternalID
	})

	// Find difference of users
	toCreateLibIDs := embyUserIDs.Difference(inDBUserLibIDs)
	toDeleteLibIDs := inDBUserLibIDs.Difference(embyUserIDs)

	// Delete users that don't exist anymore
	if res := t.db.Model(&models.User{}).Where("library_external_id IN ?", toDeleteLibIDs); res.Error != nil {
		return fmt.Errorf("failed to delete users which don't exist in library anymore: %s", res.Error)
	}

	// Create users
	toCreateUsers := []models.User{}
	toCreateLibIDs.ForEach(func(libID string) bool {
		oneEmbyUser := embyUsersByID[libID]

		toCreateUsers = append(toCreateUsers, models.User{
			LibraryExternalID: libID,
			Name:              oneEmbyUser.Name,
			IsAdmin:           oneEmbyUser.Policy.IsAdministrator,
		})
		return true
	})

	if res := t.db.Create(&toCreateUsers); res.Error != nil {
		return fmt.Errorf("failed to insert new users: %s", res.Error)
	}

	return nil
}
