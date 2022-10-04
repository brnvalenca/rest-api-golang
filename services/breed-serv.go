package services

import (
	"errors"
	"fmt"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"strconv"
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

func (*breedService) UpdateBreed(d *entities.DogBreed) error {
	err := breedRepository.Update(d)
	if err != nil {
		return fmt.Errorf(err.Error(), "error during UpdateBreed function")
	}
	return nil
}

func (*breedService) FindBreedByID(id string) (*entities.DogBreed, error) {
	return breedRepository.FindById(id)
}

func (*breedService) FindBreeds() ([]entities.DogBreed, error) {
	return breedRepository.FindAll()
}

func (*breedService) ValidateBreed(d *entities.DogBreed) error {
	idStr := strconv.Itoa(d.ID)
	if idStr == "" {
		err := errors.New("breed must have an valid ID")
		return err
	}
	if d.BreedImg == "" {
		err := errors.New("breed image is empty")
		return err
	}
	if d.Energy < 0 {
		err := errors.New("breed energy cannot be negative")
		return err
	}
	if d.GoodWithDogs < 0 {
		err := errors.New("good with dogs field cannot be negative")
		return err
	}
	if d.GoodWithKids < 0 {
		err := errors.New("good with kids field cannot be negative")
		return err
	}
	if d.Grooming < 0 {
		err := errors.New("grooming field cannot be negative")
		return err
	}
	if d.Name == "" {
		err := errors.New("breed must have a name")
		return err
	}
	if d.Shedding < 0 {
		err := errors.New("shedding field cannot be negative")
		return err
	}
	if d == nil {
		err := errors.New("breed must not be nil")
		return err
	}
	return nil
}
