package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/controllers"
	repository "rest-api/golang/exercise/repository/data"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func getDogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repository.Dogs)
}

func getDogsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, dog := range repository.Dogs {
		if dog.ID == params["id"] {
			json.NewEncoder(w).Encode(dog)
		}

	}
}
func handleRequestsDogs() {
	r.HandleFunc("/dogs", getDogs).Methods("GET")
	r.HandleFunc("/dogs/{id}", getDogsByID).Methods("GET")
}

func HandleRequest() {

	fmt.Println("Starting server at port: 8080")
	handleRequestsDogs()
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUsersById).Methods("GET")
	r.HandleFunc("/users/create", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/delete/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/update/{id}", controllers.UpdateUser).Methods("PUT")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {

	repository.Users = repository.MakeUsers(repository.Users)
	repository.Dogs = repository.MakeDogs(repository.Dogs)

	HandleRequest()

}
