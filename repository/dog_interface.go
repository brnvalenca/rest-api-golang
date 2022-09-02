package repository

import "rest-api/golang/exercise/domain/entities"

type IDogRepository interface {
	Save(d *entities.Dog, id interface{}) error
	FindAll() ([]entities.Dog, error)
	FindById(id string) (*entities.Dog, error)
	Delete(id string) (*entities.Dog, error)
	Update(d *entities.Dog, id string) error
	CheckIfExists(id string) bool
}
