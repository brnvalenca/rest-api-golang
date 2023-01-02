package services

import (
	"errors"
	"fmt"
	"net/mail"
	"rest-api/golang/exercise/domain/dtos"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/utils"
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

type UserService struct {
	userRepo  repository.IUserRepository
	prefsRepo repository.IPrefsRepository
}

func NewUserService(userRepo repository.IUserRepository, prefsRepo repository.IPrefsRepository) *UserService {
	return &UserService{userRepo: userRepo, prefsRepo: prefsRepo}
}

func Validate(u *entities.User, isUpdate bool) error {

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
	if u.Password == "" && !isUpdate {
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

func (userService *UserService) FindAll() ([]dtos.UserDTOSignUp, error) {
	users, err := userService.userRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
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

func (userService *UserService) FindById(id string) (*dtos.UserDTOSignUp, error) {
	user, err := userService.userRepo.FindById(id)
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

func (userService *UserService) Delete(id string) (*dtos.UserDTOSignUp, error) {
	user, err := userService.userRepo.Delete(id)
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

func (userService *UserService) UpdateUser(u *dtos.UserDTOSignUp) error {

	userPrefs, userInfo := utils.PartitionUserDTO(u)
	isUpdate := true
	err := Validate(userInfo, isUpdate)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = userService.userRepo.Update(userInfo, userPrefs)
	if err != nil {
		return fmt.Errorf(err.Error(), "error with userRepo.Update call in service")
	}
	return nil
}

func (userService *UserService) Create(u *dtos.UserDTOSignUp) (int, error) {

	userPrefs, userInfo := utils.PartitionUserDTO(u)
	isUpdate := false
	err := Validate(userInfo, isUpdate)
	if err != nil {
		return 0, fmt.Errorf("error during user creation: %w", err)
	}
	userData, err := userService.userRepo.Save(userInfo)
	if err != nil {
		fmt.Println(err.Error(), "error no userRepo.Save()")
	}

	err = userService.prefsRepo.SavePrefs(userPrefs, userData)
	if err != nil {
		fmt.Println(err.Error(), "error on the prefsRepo.Save() method")
	}

	return userData, nil
}

func (userService *UserService) Check(id string) bool {
	check := userService.userRepo.CheckIfExists(id)
	return check
}

func (userService *UserService) CheckEmailServ(email string) (bool, *dtos.UserCheckDTO) {
	flagUser, userDB := userService.userRepo.CheckEmail(email)
	if !flagUser {
		return false, nil
	} else {
		passwordByteDTO := dtos.UserCheckDTO{PasswordDTO: userDB.Password}
		return true, &passwordByteDTO
	}
}
