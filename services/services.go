package services

import "rest-api/golang/exercise/domain/entities"

type Service interface {
	Validate(u *entities.User) error
	FindAll() ([]entities.User, error)
	FindById(id string) (*entities.User, error)
	Delete(id string) (*entities.User, error)
	Update(u *entities.User, id string) error
	Create(u *entities.User) (*entities.User, error)
	Check(id string) bool
}
