package routes

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/controllers"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func HandleRequest() {

	fmt.Println("Starting server at port: 8080")
	HandleRequestsDogs()
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUsersById).Methods("GET")
	r.HandleFunc("/users/create", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/delete/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/update/{id}", controllers.UpdateUser).Methods("PUT")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
