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
	json.NewEncoder(w).Encode(data.DogKennels)
}

func GetDogsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, kennel := range data.DogKennels {
		if kennel.ID == params["id"] {
			json.NewEncoder(w).Encode(kennel)
		}
	}
}

func CreateDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kennel entities.DogKennel

	_ = json.NewDecoder(r.Body).Decode(&kennel)
	kennel.ID = strconv.Itoa(len(data.DogKennels) + 1)
	data.DogKennels = append(data.DogKennels, kennel)
	json.NewEncoder(w).Encode(kennel)
}

func DeleteDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, kennel := range data.DogKennels {
		if kennel.ID == params["id"] {
			data.DogKennels = append(data.DogKennels[:index], data.DogKennels[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.DogKennels)
}

func UpdateDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, kennel := range data.DogKennels {
		if kennel.ID == params["id"] {
			data.DogKennels = append(data.DogKennels[:index], data.DogKennels[index+1:]...)
			var kennel entities.DogKennel
			kennel.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&kennel)
			data.DogKennels = append(data.DogKennels, kennel)
			json.NewEncoder(w).Encode(data.DogKennels)
		}
	}
}
