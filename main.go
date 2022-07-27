package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/controllers"
	repository "rest-api/golang/exercise/repository/data"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func HandleRequest() {

	fmt.Println("Starting server at port: 8000")
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
	HandleRequest()

}
