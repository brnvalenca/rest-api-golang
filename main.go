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

func handleRequestsDogs() {
	r.HandleFunc("/kennel", controllers.GetDogs).Methods("GET")
	r.HandleFunc("/dog/{id}", controllers.GetDogsByID).Methods("GET")
	r.HandleFunc("/dog/create", controllers.CreateDog).Methods("POST")
	r.HandleFunc("/dog/delete/{id}", controllers.DeleteDog).Methods("DELETE")
	r.HandleFunc("/dog/update/{id}", controllers.UpdateDog).Methods("PUT")
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
	repository.DogKennels = repository.MakeDogKennels()
	fmt.Println(repository.DogKennels)
	HandleRequest()

}
