package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"rest-api/golang/exercise/data"
	"rest-api/golang/exercise/domain/entities"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Users)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, user := range data.Users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entities.User

	_ = json.NewDecoder(r.Body).Decode(&user) // Aqui eu decodifico o body da requisicao, que estará em JSON, contendo os dados do user
	user.ID = strconv.Itoa(rand.Intn(1000000))
	data.Users = append(data.Users, user)
	json.NewEncoder(w).Encode(user) // Codifico a resposta guardada em w para JSON e mostro na tela.

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.Users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range data.Users {

		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			var user entities.User
			user.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&user)
			data.Users = append(data.Users, user)
			json.NewEncoder(w).Encode(data.Users)
		}
	}
}
