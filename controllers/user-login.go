package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/golang/exercise/authentication"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"rest-api/golang/exercise/services/middleware"
)

type login struct{}

func NewLoginController(service services.IUserService) LoginInterface {
	userService = service
	return &login{}
}

func (*login) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error(), "error reading request body")
		return
	}

	DBUser, check := userService.CheckEmailServ(&user)
	if !check {
		log.Fatal(err.Error(), "user not registered")
		return
	}
	checkPasswordHash := middleware.CheckPassword(user.Password, DBUser.Password)
	if !checkPasswordHash {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("password authentication failed"))
		return
	} else {
		token, err := authentication.GenerateJWT(DBUser.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to generate JWT token"))
			return
		}
		json.NewEncoder(w).Encode(token)
	}
}
