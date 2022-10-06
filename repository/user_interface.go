package repository

import "rest-api/golang/exercise/domain/entities"

type IUserRepository interface {
	Save(u *entities.User) (int, error)
	FindAll() ([]entities.User, error)
	FindById(id string) (*entities.User, error)
	Delete(id string) (*entities.User, error)
	Update(u *entities.User, id string) error
	CheckIfExists(id string) bool
	CheckEmail(email string) (*entities.User, bool)
}
