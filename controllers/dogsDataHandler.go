package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	repository "rest-api/golang/exercise/repository/data"
	"strconv"

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

func CreateDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dog entities.Dog

	_ = json.NewDecoder(r.Body).Decode(&dog)
	dog.ID = strconv.Itoa(len(repository.Dogs) + 1)
	repository.Dogs = append(repository.Dogs, dog)
	json.NewEncoder(w).Encode(dog)
}

func DeleteDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, dog := range repository.Dogs {
		if dog.ID == params["id"] {
			repository.Dogs = append(repository.Dogs[:index], repository.Dogs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(repository.Dogs)
}

func UpdateDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, dog := range repository.Dogs {
		if dog.ID == params["id"] {
			repository.Dogs = append(repository.Dogs[:index], repository.Dogs[index+1:]...)
			var dog entities.Dog
			dog.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&dog)
			repository.Dogs = append(repository.Dogs, dog)
			json.NewEncoder(w).Encode(repository.Dogs)
		}
	}
}
