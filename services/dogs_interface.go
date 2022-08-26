package services

import "rest-api/golang/exercise/domain/entities"

type IDogService interface {
	Validate(u *entities.Dog) error
	FindAll() ([]entities.Dog, error)
	FindById(id string) (*entities.Dog, error)
	Delete(id string) (*entities.Dog, error)
	Update(u *entities.Dog, id string) error
	Create(u *entities.Dog) (*entities.Dog, error)
	Check(id string) bool
}
