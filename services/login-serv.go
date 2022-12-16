package services

import (
	"fmt"
	"rest-api/golang/exercise/authentication"
	"rest-api/golang/exercise/security"
)

type ILoginService interface {
	AuthenticateUser(userEmail, userPassword string) (string, error)
}

type loginServ struct {
	passwordServ security.IPasswordHash
	userServ     IUserService
}

func NewLoginService(passwordServ security.IPasswordHash, userServ IUserService) ILoginService {
	return &loginServ{passwordServ: passwordServ, userServ: userServ}
}

func (loginServ *loginServ) AuthenticateUser(userEmail, userPassword string) (string, error) {
	check, userDB := loginServ.userServ.CheckEmailServ(userEmail)
	if !check {
		return "", fmt.Errorf("email not registered")
	}

	checkPassword := loginServ.passwordServ.CheckPassword(userPassword, userDB.PasswordDTO)
	if !checkPassword {
		return "", fmt.Errorf("email not registered")
	}

	token, err := authentication.GenerateJWT(userDB.ID)
	if err != nil {
		return "", fmt.Errorf("internal error during token generation: %w", err)
	}

	return token, nil
}
