package services

import (
	"fmt"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
)

type breedService struct{}

var breedRepository repository.IBreedRepository

func NewBreedService(repo repository.IBreedRepository) IBreedService {
	breedRepository = repo
	return &breedService{}
}

/*
	The create and delete breed will only be able with admin users
*/

func (*breedService) CreateBreed(d *entities.DogBreed) error {
	_, err := breedRepository.Save(d)
	if err != nil {
		return fmt.Errorf(err.Error(), "eror during CreateBreed function")
	}
	return nil
}

func (*breedService) UpdateBreed(d *entities.DogBreed, id string) error {
	err := breedRepository.Update(d, id)
	if err != nil {
		return fmt.Errorf(err.Error(), "eror during UpdateBreed function")
	}
	return nil
}

func (*breedService) FindBreedByID(id string) (*entities.DogBreed, error) {
	return breedRepository.FindById(id)
}

func (*breedService) FindBreeds() ([]entities.DogBreed, error) {
	return breedRepository.FindAll()
}
