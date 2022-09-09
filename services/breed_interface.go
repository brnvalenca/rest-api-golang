package services

import "rest-api/golang/exercise/domain/entities"

type IBreedService interface {
	//ValidateBreed(d *entities.DogBreed) error
	FindBreeds() ([]entities.DogBreed, error)
	FindBreedByID(id string) (*entities.DogBreed, error)
	UpdateBreed(d *entities.DogBreed, id string) error
	CreateBreed(d *entities.DogBreed) error
}
