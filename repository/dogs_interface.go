package repository

import "rest-api/golang/exercise/domain/entities"

type IDogRepository interface {
	Save(u *entities.Dog) (*entities.Dog, error)
	FindAll() ([]entities.Dog, error)
	FindById(id string) (*entities.Dog, error)
	Delete(id string) (*entities.Dog, error)
	Update(u *entities.Dog, id string) error
	CheckIfExists(id string) bool
}
