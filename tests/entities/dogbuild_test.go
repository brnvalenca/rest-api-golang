package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"testing"
)

func TestBuildDog(t *testing.T) {
	breed := entities.BuildBreed("Cavalier", "Medium", "4", "7")
	got := entities.BuildDog(breed, 5, true, "pitomba", "black and white", "1")

	want := entities.Dog{
		Name:   "pitomba",
		Age:    5,
		IsMale: true,
		Color:  "black and white",
		Breed:  breed,
		ID: "1",
	}

	if got != want {
		t.Errorf("BuildDog error. Got: %v\n Want: %v\n", got, want)
	}
}
