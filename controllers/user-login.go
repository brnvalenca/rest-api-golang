package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/golang/exercise/authentication"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"
)

type login struct {
	passwordService security.IPasswordHash
	userService     services.IUserService
}

func NewLoginController(userServ services.IUserService, passwordServ security.IPasswordHash) LoginInterface {
	return &login{passwordService: passwordServ, userService: userServ}
}

func (logserv *login) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user dtos.UserDTOSignIn
	var appErr dtos.AppErrorDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error(), "error reading request body")
		return
	}

	check, userDB := logserv.userService.CheckEmailServ(user.Email)
	if !check {
		appErr.Code = http.StatusNotAcceptable
		appErr.Message = "User not registered"
		json.NewEncoder(w).Encode(appErr)
		return
	}

	checkPasswordHash := logserv.passwordService.CheckPassword(user.Password, userDB.Password)
	if !checkPasswordHash {
		appErr.Code = http.StatusUnauthorized
		appErr.Message = "Password incorrect"
		json.NewEncoder(w).Encode(appErr)
		return
	} else {
		token, err := authentication.GenerateJWT(userDB.ID)
		if err != nil {
			appErr.Code = http.StatusInternalServerError
			appErr.Message = "Failed to generate JWT Token"
			json.NewEncoder(w).Encode(appErr)
			return
		}
		json.NewEncoder(w).Encode(token)
	}
}
