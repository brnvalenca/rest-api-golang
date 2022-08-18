package services

import (
	"errors"
	"net/mail"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
)

/*
	This function should not call the infra.PostUser directly. The service function should deal with all
	the logic business and validation of data requests that are made by the controllers. The functions
	that make queries to the database should be defined in a repository layer.
*/

type userv struct{}

var (
	userRepo repository.Repository
)

func NewUserService(repo repository.Repository) Service {
	userRepo = repo
	return &userv{}
}

func (*userv) Validate(u *entities.User) error {
	if u == nil {
		err := errors.New("the user is empty")
		return err
	}
	if u.Name == "" {
		err := errors.New("the user name is empty")
		return err
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		err := errors.New("user email not valid")
		return err
	}
	if u.Password == "" {
		err := errors.New("the user password is empty")
		return err
	}
	return nil
}

func (*userv) FindAll() ([]entities.User, error) {
	return userRepo.FindAll()
}

func (*userv) FindById(id string) (*entities.User, error) {
	return userRepo.FindById(id)
}

func (*userv) Delete(id string) (*entities.User, error) {
	return userRepo.Delete(id)
}

func (*userv) Update(u *entities.User, id string) error {
	return userRepo.Update(u, id)
}

func (*userv) Create(u *entities.User) (*entities.User, error) {
	return userRepo.Save(u)
}

func (*userv) Check(id string) bool {
	return userRepo.CheckIfExists(id)
}
