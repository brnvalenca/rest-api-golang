package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/infra"
	"rest-api/golang/exercise/services"
	"rest-api/golang/exercise/utils"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := infra.ListUsers(utils.DB)
	if err != nil {
		fmt.Printf("Error with ListUsers: %v", err)
	}
	json.NewEncoder(w).Encode(users)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	user, err := infra.ListUserById(utils.DB, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	_ = json.NewDecoder(r.Body).Decode(&user) // Aqui eu decodifico o body da requisicao, que estará em JSON, contendo os dados do user
	services.StoreUser(user)
	json.NewEncoder(w).Encode(user) // Codifico a resposta guardada em w para JSON e mostro na tela.
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	user, err := infra.DeleteUser(utils.DB, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entities.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	rows, err := infra.UpdateUser(utils.DB, user)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.NewEncoder(w).Encode(rows)

}
