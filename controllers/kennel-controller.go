package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/golang/exercise/data"
	"rest-api/golang/exercise/domain/entities"
	"strconv"

	"github.com/gorilla/mux"
)

func GetDogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.DogShelters)
}

func GetDogsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, kennel := range data.DogShelters {
		if kennel.ID == params["id"] {
			json.NewEncoder(w).Encode(kennel)
		}
	}
}

func CreateDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kennel entities.DogShelter

	_ = json.NewDecoder(r.Body).Decode(&kennel)
	kennel.ID = strconv.Itoa(len(data.DogShelters) + 1)
	data.DogShelters = append(data.DogShelters, kennel)
	json.NewEncoder(w).Encode(kennel)
}

func DeleteDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, kennel := range data.DogShelters {
		if kennel.ID == params["id"] {
			data.DogShelters = append(data.DogShelters[:index], data.DogShelters[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.DogShelters)
}

func UpdateDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, kennel := range data.DogShelters {
		if kennel.ID == params["id"] {
			data.DogShelters = append(data.DogShelters[:index], data.DogShelters[index+1:]...)
			var kennel entities.DogShelter
			kennel.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&kennel)
			data.DogShelters = append(data.DogShelters, kennel)
			json.NewEncoder(w).Encode(data.DogShelters)
		}
	}
}
