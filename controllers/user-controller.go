package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

var (
// Instance of the UserService interface. That i'll use inside my controller
) // This interface will allow my controllers to 'talk' with my services, and perform actions before
// the calls to my database.

type userController struct {
	passwordHash security.IPasswordHash
	userService  services.IUserService
}

func NewUserController(service services.IUserService, password security.IPasswordHash) IController {

	return &userController{passwordHash: password, userService: service}
}

func (u *userController) Create(w http.ResponseWriter, r *http.Request) {
	/*
		Gateway: criar endpoints tanto em HTTP como em GRPC.
		Escrever os endpoints num protobuf e sair referenciando a partir dele.
	*/
	w.Header().Set("Content-Type", "application/json")
	var userDTO dtos.UserDTOSignUp
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		log.Fatal(err.Error())
	}
	check, _ := u.userService.CheckEmailServ(userDTO.Email)
	if check {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("email already registered in our system"))
		return
	}
	userDTO.Password, err = u.passwordHash.GeneratePasswordHash(userDTO.Password)
	if err != nil {
		log.Fatal(err.Error(), "error hashing user password")
	}

	user, err := u.userService.Create(&userDTO)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(user)

}

func (u *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := u.userService.FindAll()
	if err != nil {
		fmt.Printf("Error with ListUsers: %v", err)
	}

	json.NewEncoder(w).Encode(users)
}

func (u *userController) GetById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	user, err := u.userService.FindById(id)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	}
	json.NewEncoder(w).Encode(user)
}

func (u *userController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	check := u.userService.Check(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	} else {
		user, err := u.userService.Delete(id)
		if err != nil {
			fmt.Println(err.Error())
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (u *userController) Update(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var userDTO dtos.UserDTOSignUp
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		log.Fatal(err.Error())
	}

	params := mux.Vars(r)
	id := params["id"]

	check := u.userService.Check(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := u.userService.UpdateUser(&userDTO)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(&userDTO)
	}
}
