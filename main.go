package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"strconv"

	"github.com/gorilla/mux"
)

var users []entities.User
var r = mux.NewRouter()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, user := range users {
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
	users = append(users, user)
	json.NewEncoder(w).Encode(user) // Codifico a resposta guardada em w para JSON e mostro na tela.

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range users {

		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user entities.User
			user.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&user)
			users = append(users, user)
			json.NewEncoder(w).Encode(users)
		}
	}
}

func HandleRequest() {

	fmt.Println("Starting server at port: 8000")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUsersById).Methods("GET")
	r.HandleFunc("/users/create", CreateUser).Methods("POST")
	r.HandleFunc("/users/delete/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/update/{id}", UpdateUser).Methods("PUT")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func main() {

	users = repository.MakeUsers(users)
	HandleRequest()

}
