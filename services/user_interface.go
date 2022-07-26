package services

import "rest-api/golang/exercise/domain/entities"

type IUserService interface {
	Validate(u *entities.User) error
	FindAll() ([]entities.User, error)
	FindById(id string) (*entities.User, error)
	Delete(id string) (*entities.User, error)
	UpdateUser(u *entities.User, id string) error
	Create(u *entities.User) (int, error)
	Check(id string) bool
}
