package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

var (
	userService services.IUserService // Instance of the UserService interface. That i'll use inside my controller
) // This interface will allow my controllers to 'talk' with my services, and perform actions before
// the calls to my database.

type userController struct{}

func NewUserController(service services.IUserService) IController {
	userService = service
	return &userController{}
}

func (*userController) Create(w http.ResponseWriter, r *http.Request) {
	/*
		Gateway: criar endpoints tanto em HTTP como em GRPC.
		Escrever os endpoints num protobuf e sair referenciando a partir dele.
	*/
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = userService.Validate(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	} else {
		row, err := userService.Create(&user)
		if err != nil {
			log.Fatal(err.Error(), "userService.Create() error")
		}
		json.NewEncoder(w).Encode(row) // Codifico a resposta guardada em w para JSON e mostro na tela.
	}
}

func (*userController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := userService.FindAll()
	if err != nil {
		fmt.Printf("Error with ListUsers: %v", err)
	}
	json.NewEncoder(w).Encode(users)
}

func (*userController) GetById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	user, err := userService.FindById(id)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		json.NewEncoder(w).Encode(user)
	}

}

func (*userController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	check := userService.Check(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		user, err := userService.Delete(id)
		if err != nil {
			fmt.Println(err.Error())
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (*userController) Update(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = userService.Validate(&user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	check := userService.Check(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := userService.UpdateUser(&user, id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = json.NewEncoder(w).Encode(&user)
		}
	}
}
