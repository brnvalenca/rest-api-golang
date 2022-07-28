package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"testing"
)

func TestBuildBreed(t *testing.T) {
	got := entities.BuildBreed("Cavalier", "Medium", "4", "7")

	want := entities.Breed{
		BreedName:     "Cavalier",
		BreedAVGSize:  "Medium",
		BreedLoudness: "4",
		BreedEnergy:   "7",
	}

	if got != want {
		t.Errorf("BuildBreed error. Got: %v\n Want: %v\n", got, want)
	}

}
