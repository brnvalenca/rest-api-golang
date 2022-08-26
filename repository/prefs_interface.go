package repository

import "rest-api/golang/exercise/domain/entities"

type IPrefsRepository interface {
	Save(u *entities.UserDogPreferences) error
	Delete(id string) error
	Update(u *entities.UserDogPreferences, id string) error
}
