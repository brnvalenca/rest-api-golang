package services

import (
	"errors"
	"rest-api/golang/exercise/domain/entities"
)

type dserv struct{}

func NewDogService() Service {
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
