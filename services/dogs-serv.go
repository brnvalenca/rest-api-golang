package services

import (
	"errors"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dto"
	"rest-api/golang/exercise/repository"
	"strconv"
)

type dserv struct{}

var (
	breedRepo repository.IBreedRepository
	dogRepo   repository.IDogRepository
)

func NewDogService(drepo repository.IDogRepository, brepo repository.IBreedRepository) IDogService {
	dogRepo = drepo
	breedRepo = brepo
	return &dserv{}
}

func (*dserv) ValidateDog(d *entities.Dog) error {
	if d == nil {
		err := errors.New("dog is empty")
		return err
	}
	if d.DogName == "" {
		err := errors.New("dog name is empty")
		return err
	}
	if d.Sex == "" {
		err := errors.New("dog sex is empty")
		return err
	}
	return nil

}

func (*dserv) FindDogs() ([]entities.Dog, error) {
	return dogRepo.FindAll()
}

func (*dserv) FindDogByID(id string) (*entities.Dog, error) {
	return dogRepo.FindById(id)
}

func (*dserv) DeleteDog(id string) (*entities.Dog, error) {
	return dogRepo.Delete(id)
}

func (*dserv) UpdateDog(u *entities.Dog, id string) error {
	return dogRepo.Update(u, id)
}

func (*dserv) CreateDog(d *entities.Dog, b *entities.DogBreed) error {
	if b != nil {
		err := dogRepo.Save(d, b.ID)
		if err != nil {
			log.Fatal(err.Error(), "\n service error during dog creation")
		}
		return nil
	} else {
		_, err := breedRepo.Save(b)
		if err != nil {
			log.Fatal(err.Error(), "\n service error during breed creation")
		}
		dogRepo.Save(d, b.ID)
	}
	return nil
}

func (*dserv) CheckIfDogExist(id string) bool {
	return dogRepo.CheckIfExists(id)
}

func (*dserv) CheckIfKennelExist(d *dto.DogDTO) bool {
	id := strconv.Itoa(d.KennelID)
	return kennelRepo.CheckIfExistsRepo(id)
}

func (*dserv) CheckIfBreedExist(d *dto.DogDTO) bool {
	id := strconv.Itoa(d.BreedID)
	return breedRepo.CheckIfExists(id)
}
