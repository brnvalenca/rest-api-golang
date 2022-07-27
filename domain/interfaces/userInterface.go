package interfaces

import (
	"rest-api/golang/exercise/domain/entities"
)

type IManipulateUsers interface {
	GetUsers() *[]entities.User
	GetUsersById(id string) *entities.User
	CreateUser() *[]entities.User
	DeleteUser(id string) *[]entities.User
	UpdateUser(id string) *entities.User
}
