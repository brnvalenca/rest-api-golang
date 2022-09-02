package services

import (
	"errors"
	"fmt"
	"net/mail"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services/middleware"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/repository/repos"
)

type userv struct{}

var (
	userRepo  repository.IUserRepository
	prefsRepo repository.IPrefsRepository = repos.NewPrefsRepo()
)

func NewUserService(repo repository.IUserRepository) IUserService {
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

func (*userv) Create(u *entities.User) (int, error) {
	userData, err := userRepo.Save(u)
	if err != nil {
		fmt.Println(err.Error(), "Erro no userRepo.Save()")
	}
	_, err = fmt.Println(userData)
	if err != nil {
		fmt.Println(err.Error(), "Erro no Println(*userData)")
	}
	userPrefs := middleware.PartitionData(u, userData)
	err = prefsRepo.Save(userPrefs)
	if err != nil {
		fmt.Println(err.Error(), "error on the prefsRepo.Save() method")
	}

	return userData, nil
}

func (*userv) Check(id string) bool {
	return userRepo.CheckIfExists(id)
}
