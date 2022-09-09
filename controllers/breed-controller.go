package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

type breedController struct{}

var breedService services.IBreedService

func NewBreedController(service services.IBreedService) IController {
	breedService = service
	return &breedController{}
}

func (*breedController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	breeds, err := breedService.FindBreeds()
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(breeds)

}

func (*breedController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	breed, err := breedService.FindBreedByID(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(breed)
}

/*
	The above function creates a new Breed. This action must be conditioned to a check
	where if the breed exists or not. In case of the breed already exists in the database
	a error message must be displayed and the proccess has to terminate.
*/

func (*breedController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var breed entities.DogBreed
	err := json.NewDecoder(r.Body).Decode(&breed)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = breedService.CreateBreed(&breed)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(breed)
}

func (*breedController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("only allowed to admin"))
}

func (*breedController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var breed entities.DogBreed

	err := json.NewDecoder(r.Body).Decode(&breed)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = breedService.UpdateBreed(&breed, id)
	if err != nil {
		log.Fatal(err.Error())
	}
}
