package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dto"
	"rest-api/golang/exercise/services"
	"rest-api/golang/exercise/services/middleware"

	"github.com/gorilla/mux"
)

/*
	Dog controller will need an instance of Dog Service to work
	Router - Controller - Service - Repo - Database
*/

type dogController struct{}

var (
	dogService services.IDogService
)

func NewDogController(service services.IDogService) IController {
	dogService = service
	return &dogController{}
}

func (*dogController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dogDto dto.DogDTO
	err := json.NewDecoder(r.Body).Decode(&dogDto)

	if err != nil {
		log.Fatal(err.Error(), "error during body request decoding")
	}

	breedCheck := dogService.CheckIfBreedExist(&dogDto)
	if !breedCheck {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("breed doesnt exist. please send request with breed fields"))
		return
	}

	checkKennel := dogService.CheckIfKennelExist(&dogDto)
	if !checkKennel {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kennel doesnt exist. please send request with kennel fields"))
	} else {
		dog, breed := middleware.PartitionDogDTO(dogDto)
		dogService.CreateDog(dog, breed)
		json.NewEncoder(w).Encode(1)
	}
}

func (*dogController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dogs, err := dogService.FindDogs()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(dogs)
}

func (*dogController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	check := dogService.CheckIfDogExist(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		dog, err := dogService.FindDogByID(id)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(dog)
	}
}

func (*dogController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	check := dogService.CheckIfDogExist(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		dog, err := dogService.DeleteDog(id)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(dog)
	}
}

func (*dogController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	var dog entities.Dog
	_ = json.NewDecoder(r.Body).Decode(&dog)

	check := dogService.CheckIfDogExist(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := dogService.UpdateDog(&dog, id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = json.NewEncoder(w).Encode(&dog)
		}
	}
}
