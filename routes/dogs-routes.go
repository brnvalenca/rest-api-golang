package routes

import "rest-api/golang/exercise/controllers"

func HandleRequestsDogs() {
	r.HandleFunc("/dogs", controllers.GetDogs).Methods("GET")
	r.HandleFunc("/dogs/{id}", controllers.GetDogsByID).Methods("GET")
	r.HandleFunc("/dogs/create", controllers.CreateDog).Methods("POST")
	r.HandleFunc("/dogs/delete/{id}", controllers.DeleteDog).Methods("DELETE")
	r.HandleFunc("/dogs/update/{id}", controllers.UpdateDog).Methods("PUT")
}
