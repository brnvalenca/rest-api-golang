package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

var (
	userService services.UserServiceI // Instance of the UserService interface. That i'll use inside my controller
) // This interface will allow my controllers to 'talk' with my services, and perform actions before
// the calls to my database.

type userController struct{}

func NewUserController(service services.UserServiceI) Controller {
	userService = service
	return &userController{}
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

func (*userController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	_ = json.NewDecoder(r.Body).Decode(&user) // Aqui eu decodifico o body da requisicao, que estará em JSON, contendo os dados do user
	err := userService.Validate(&user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		userService.Create(&user)
		json.NewEncoder(w).Encode(user) // Codifico a resposta guardada em w para JSON e mostro na tela.
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
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := userService.Validate(&user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	check := userService.Check(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := userService.Update(&user, id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = json.NewEncoder(w).Encode(&user)
		}
	}

}
