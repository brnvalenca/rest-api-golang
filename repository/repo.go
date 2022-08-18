package repository

import (
	"rest-api/golang/exercise/domain/entities"
)

type Repository interface {
	Save(u *entities.User) (*entities.User, error)
	FindAll() ([]entities.User, error)
	FindById(id string) (*entities.User, error)
	Delete(id string) (*entities.User, error)
	Update(u *entities.User, id string) error
	CheckIfExists(id string) bool
}

/*

	This repository layer defines a UserRepository with CRUD methods that may be implemented
	by any database configuration that needs to implement it. Doing this, the application core
	logic will not be dependent on DB implementation. In other words, we can define a implementation
	with mySQL, or mongo, or firebase etc, they will implement this interface and the changes in my
	core logic will be the less possible.

*/
