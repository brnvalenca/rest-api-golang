package services

import (
	"errors"
	"fmt"
	"log"
	"net/mail"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/middleware"
	"rest-api/golang/exercise/repository"
	"strconv"
)

type IUserService interface {
	FindAll() ([]dtos.UserDTOSignUp, error)
	FindById(id string) (*dtos.UserDTOSignUp, error)
	Delete(id string) (*dtos.UserDTOSignUp, error)
	UpdateUser(u *dtos.UserDTOSignUp) error
	Create(u *dtos.UserDTOSignUp) (int, error)
	Check(id string) bool
	CheckEmailServ(email string) (bool, *dtos.UserCheckDTO)
}

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

func Validate(u *entities.User) error {

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
	if u.Email == "" {
		err := errors.New("the user email is empty")
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

func (*userv) FindAll() ([]dtos.UserDTOSignUp, error) {
	users, err := userRepo.FindAll()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	ubuilder := dtos.NewUserDTOBuilder()
	var usersDTO []dtos.UserDTOSignUp
	var uDTO dtos.UserDTOSignUp
	for i := 0; i < len(users); i++ {
		ubuilder.Has().
			ID(users[i].ID).
			Name(users[i].Name).
			Email(users[i].Email).
			Password(users[i].Password).
			UserPrefs(dtos.UserPrefsDTO(users[i].UserPreferences))
		uDTO = *ubuilder.BuildUser()
		usersDTO = append(usersDTO, uDTO)
	}

	return usersDTO, nil
}

func (*userv) FindById(id string) (*dtos.UserDTOSignUp, error) {
	user, err := userRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	ubuilder := dtos.NewUserDTOBuilder()
	ubuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password).
		UserPrefs(dtos.UserPrefsDTO(user.UserPreferences))
	uDTO := *ubuilder.BuildUser()
	return &uDTO, nil
}

func (*userv) Delete(id string) (*dtos.UserDTOSignUp, error) {
	user, err := userRepo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	ubuilder := dtos.NewUserDTOBuilder()
	ubuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password).
		UserPrefs(dtos.UserPrefsDTO(user.UserPreferences))
	uDTO := *ubuilder.BuildUser()
	return &uDTO, nil

}

func (*userv) UpdateUser(u *dtos.UserDTOSignUp) error {

	userPrefs, userInfo := middleware.PartitionUserDTO(u)

	err := Validate(userInfo)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	err = userRepo.Update(userInfo, userPrefs)
	if err != nil {
		return fmt.Errorf(err.Error(), "error with userRepo.Update call in service")
	}
	return nil
}

func (*userv) Create(u *dtos.UserDTOSignUp) (int, error) {

	userPrefs, userInfo := middleware.PartitionUserDTO(u)

	err := Validate(userInfo)
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	userData, err := userRepo.Save(userInfo)
	if err != nil {
		fmt.Println(err.Error(), "error no userRepo.Save()")
	}

	err = prefsRepo.SavePrefs(userPrefs, userData)
	if err != nil {
		fmt.Println(err.Error(), "error on the prefsRepo.Save() method")
	}

	return userData, nil
}

func (*userv) Check(id string) bool {
	return userRepo.CheckIfExists(id)
}

func (*userv) CheckEmailServ(email string) (bool, *dtos.UserCheckDTO) {
	flagUser, userDB := userRepo.CheckEmail(email)
	if !flagUser {
		return false, nil
	} else {
		passwordByteDTO := dtos.UserCheckDTO{PasswordDTO: userDB.Password}
		return true, &passwordByteDTO
	}
}
