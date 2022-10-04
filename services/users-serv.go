package services

import (
	"errors"
	"fmt"
	"net/mail"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/services/middleware"
	"strconv"
)

type userv struct{}

var (
	userRepo  repository.IUserRepository
	prefsRepo repository.IPrefsRepository
)

func NewUserService(user repository.IUserRepository, pref repository.IPrefsRepository) IUserService {
	userRepo = user
	prefsRepo = pref
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
	idStr := strconv.Itoa(u.ID)
	if idStr == "" {
		err := errors.New("the user id is empty")
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

func (*userv) UpdateUser(u *entities.User, id string) error {
	idStr := id
	err := userRepo.Update(u, idStr)
	if err != nil {
		return fmt.Errorf(err.Error(), "error with userRepo.Update call in service")
	}
	return nil
}

func (*userv) Create(u *entities.User) (int, error) {
	userData, err := userRepo.Save(u)
	if err != nil {
		fmt.Println(err.Error(), "Erro no userRepo.Save()")
	}

	userPrefs := middleware.PartitionData(u, userData)
	err = prefsRepo.SavePrefs(userPrefs)
	if err != nil {
		fmt.Println(err.Error(), "error on the prefsRepo.Save() method")
	}

	return userData, nil
}

func (*userv) Check(id string) bool {
	return userRepo.CheckIfExists(id)
}
