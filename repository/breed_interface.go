package repository

import "rest-api/golang/exercise/domain/entities"

type IBreedRepository interface {
	Save(b *entities.DogBreed) (int, error)
	FindAll() ([]entities.DogBreed, error)
	FindById(id string) (*entities.DogBreed, error)
	Delete(id string) (*entities.DogBreed, error)
	Update(b *entities.DogBreed, id string) error
	CheckIfExists(id string) bool
}
