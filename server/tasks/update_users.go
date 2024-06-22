package tasks

import (
	"context"
	"fmt"
	"github.com/Noah-Huppert/media-recyclarr/emby"
	"github.com/Noah-Huppert/media-recyclarr/models"
	"github.com/hashicorp/go-set/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UpdateUsersTask struct {
	// logger is used to output runtime information
	logger *zap.Logger

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
	if res := t.db.Model(&models.User{}).Find(&existingUsers); res.Error != nil {
		return fmt.Errorf("failed to retrieve existing user from DB: %s", res.Error)
	}

	inDBUserLibIDs := set.FromFunc(existingUsers, func(u models.User) string {
		return u.LibraryExternalID
	})
	existingUsersByLibID := map[string]models.User{}
	for _, user := range existingUsers {
		existingUsersByLibID[user.LibraryExternalID] = user
	}

	// Find difference of users
	toCreateLibIDs := embyUserIDs.Difference(inDBUserLibIDs)
	toDeleteLibIDs := inDBUserLibIDs.Difference(embyUserIDs)

	// Delete users that don't exist anymore
	if !toDeleteLibIDs.Empty() {
		t.logger.Info("deleting users", zap.Any("user_library_ids", toDeleteLibIDs))
		if res := t.db.Model(&models.User{}).Where("library_external_id IN ?", toDeleteLibIDs); res.Error != nil {
			return fmt.Errorf("failed to delete users which don't exist in library anymore: %s", res.Error)
		}
	}

	// Create users
	if !toCreateLibIDs.Empty() {
		t.logger.Info("creating users", zap.Any("user_library_ids", toCreateLibIDs))
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
	}

	// Ensure users are up to date
	usersToUpdate := []models.User{}
	inDBUserLibIDs.ForEach(func(libID string) bool {
		embyU := embyUsersByID[libID]
		dbU := existingUsersByLibID[libID]

		needsUpdate := false

		if embyU.Name != dbU.Name {
			dbU.Name = embyU.Name
			needsUpdate = true
		}

		if embyU.Policy.IsAdministrator != dbU.IsAdmin {
			dbU.IsAdmin = embyU.Policy.IsAdministrator
			needsUpdate = true
		}

		if needsUpdate {
			usersToUpdate = append(usersToUpdate, dbU)
		}

		return true
	})

	for _, user := range usersToUpdate {
		t.logger.Info("updating user", zap.Uint("id", user.ID), zap.String("library_external_id", user.LibraryExternalID))
		if res := t.db.Save(&user); res.Error != nil {
			return fmt.Errorf("failed to update user (ID: %d): %s", user.ID, res.Error)
		}
	}

	return nil
}
