package controllers

import (
	"encoding/json"
	"net/http"
	repository "rest-api/golang/exercise/repository/data"

	"github.com/gorilla/mux"
)

func GetDogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repository.Dogs)
}

func GetDogsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, dog := range repository.Dogs {
		if dog.ID == params["id"] {
			json.NewEncoder(w).Encode(dog)
		}
	}
}
