package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"testing"
)

func TestBuildUser(t *testing.T) {
	a := entities.BuildAddress("Rua 1", "Graças", "500", "Recife")
	udog := entities.BuildUserDogPreferences(8, 6, "Large")
	got := entities.BuildUser("123", "Bruno", "brn@gmail.com", "321", a, udog)

	want := entities.User{
		ID:       "123",
		Name:     "Bruno",
		Email:    "brn@gmail.com",
		Password: "321",
		Address: entities.Address{
			Street:     "Rua 1",
			District:   "Graças",
			PostalCode: "500",
			City:       "Recife",
		},
		DogPreferences: entities.UserDogPreferences{
			DogLoudness: 8,
			DogEnergy:   6,
			DogAVGSize:  "Large",
		},
	}
	if got != want {
		t.Errorf("BuildUser error. Got: %v\n Want: %v\n", got, want)
	}

}
