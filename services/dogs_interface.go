package services

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dto"
)

type IDogService interface {
	ValidateDog(d *entities.Dog) error
	FindDogs() ([]entities.Dog, error)
	FindDogByID(id string) (*entities.Dog, error)
	DeleteDog(id string) (*entities.Dog, error)
	UpdateDog(d *entities.Dog, id string) error
	CreateDog(d *entities.Dog, b *entities.DogBreed) error
	CheckIfDogExist(id string) bool
	CheckIfKennelExist(d *dto.DogDTO) bool
	CheckIfBreedExist(d *dto.DogDTO) bool
}
