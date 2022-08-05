package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"testing"
)

func TestBuildUser(t *testing.T) {
	//	a := entities.BuildAddress("Rua 1", "Graças", "500", "Recife")
	//	udog := entities.BuildUserDogPreferences(8, 6, "Large")
	got := entities.BuildUser(123, "Bruno", "brn@gmail.com", "321")

	want := entities.User{
		ID:       123,
		Name:     "Bruno",
		Email:    "brn@gmail.com",
		Password: "321",
	}
	if got != want {
		t.Errorf("BuildUser error. Got: %v\n Want: %v\n", got, want)
	}

}
