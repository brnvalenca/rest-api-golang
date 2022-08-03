package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/controllers"
	"rest-api/golang/exercise/data"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func HandleRequestsKennels() {
	r.HandleFunc("/kennel", controllers.GetDogs).Methods("GET")
	r.HandleFunc("/kennel/{id}", controllers.GetDogsByID).Methods("GET")
	r.HandleFunc("/kennel/create", controllers.CreateDog).Methods("POST")
	r.HandleFunc("/kennel/delete/{id}", controllers.DeleteDog).Methods("DELETE")
	r.HandleFunc("/kennel/update/{id}", controllers.UpdateDog).Methods("PUT")
}

func HandleRequest() {

	fmt.Println("Starting server at port: 8080")
	HandleRequestsKennels()
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUsersById).Methods("GET")
	r.HandleFunc("/users/create", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/delete/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/update/{id}", controllers.UpdateUser).Methods("PUT")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {

	data.Users = data.MakeUsers(data.Users)
	data.DogKennels = data.MakeDogKennels()
	HandleRequest()

}
