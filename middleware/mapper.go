package middleware

import "rest-api/golang/exercise/domain/entities"

func PartitionData(u *entities.User, userID int) *entities.UserDogPreferences {
	prefs :=
		entities.BuildUserDogPreferences(
			userID,
			u.UserPreferences.GoodWithKids,
			u.UserPreferences.GoodWithDogs,
			u.UserPreferences.Shedding,
			u.UserPreferences.Grooming,
			u.UserPreferences.Energy,
		)
	return &prefs
}
