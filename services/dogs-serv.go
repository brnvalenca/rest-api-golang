package services

import (
	"errors"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
)

type dserv struct{}

var (
	dogRepo repository.DogRepositoryI
)

func NewDogService(repo repository.DogRepositoryI) DogServiceI {
	dogRepo = repo
	return &dserv{}
}

func (*dserv) Validate(d *entities.Dog) error {
	if d == nil {
		err := errors.New("dog is empty")
		return err
	}
	if d.Name == "" {
		err := errors.New("dog name is empty")
		return err
	}
	if d.Sex == "" {
		err := errors.New("dog sex is empty")
		return err
	}
	if d.Breed == "" {
		err := errors.New("dog breed is empty")
		return err
	}
	if d.Age == 0 {
		err := errors.New("dog age is empty")
		return err
	}
	return nil

}

func (*dserv) FindAll() ([]entities.Dog, error) {
	return dogRepo.FindAll()
}

func (*dserv) FindById(id string) (*entities.Dog, error) {
	return dogRepo.FindById(id)
}

func (*dserv) Delete(id string) (*entities.Dog, error) {
	return dogRepo.Delete(id)
}

func (*dserv) Update(u *entities.Dog, id string) error {
	return dogRepo.Update(u, id)
}

func (*dserv) Create(d *entities.Dog) (*entities.Dog, error) {
	return dogRepo.Save(d)
}

func (*dserv) Check(id string) bool {
	return dogRepo.CheckIfExists(id)
}
