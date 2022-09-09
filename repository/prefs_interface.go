package repository

import "rest-api/golang/exercise/domain/entities"

type IPrefsRepository interface {
	SavePrefs(u *entities.UserDogPreferences) error
	DeletePrefs(id string) error
	UpdatePrefs(u *entities.UserDogPreferences, id string) error
}
