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

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error(), "error reading request body")
		return
	}

	check, userDB := logserv.userService.CheckEmailServ(user.Email)
	if !check {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("user not registered"))
		return
	}

	checkPasswordHash := logserv.passwordService.CheckPassword(user.Password, userDB.Password)
	if !checkPasswordHash {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("password authentication failed"))
		return
	} else {
		token, err := authentication.GenerateJWT(userDB.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to generate JWT token"))
			return
		}
		json.NewEncoder(w).Encode(token)
	}
}
