package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/middleware"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

/*
	Dog controller will need an instance of Dog Service to work
	Router - Controller - Service - Repo - Database
*/

type dogController struct {
	dogService services.IDogService
}

func NewDogController(service services.IDogService) IController {
	return &dogController{dogService: service}
}

func (dc *dogController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dogDto dtos.DogDTO
	err := json.NewDecoder(r.Body).Decode(&dogDto)
	if err != nil {
		log.Fatal(err.Error(), "error during body request decoding")
	}

	breedCheck := dc.dogService.CheckIfBreedExist(&dogDto)
	if !breedCheck {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("breed doesnt exist. create breed"))
		return
	}

	checkKennel := dc.dogService.CheckIfKennelExist(&dogDto)
	if !checkKennel {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kennel doesnt exist. create kennel"))
	} else {
		dog, breed := middleware.PartitionDogDTO(dogDto)
		dc.dogService.CreateDog(dog, breed)
		json.NewEncoder(w).Encode(1)
	}
}

func (dc *dogController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dogs, err := dc.dogService.FindDogs()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(dogs)
}

func (dc *dogController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	check := dc.dogService.CheckIfDogExist(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		dog, err := dc.dogService.FindDogByID(id)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(dog)
	}
}

func (dc *dogController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	check := dc.dogService.CheckIfDogExist(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		dog, err := dc.dogService.DeleteDog(id)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(dog)
	}
}

func (dc *dogController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	var dog dtos.DogDTO
	err := json.NewDecoder(r.Body).Decode(&dog)
	if err != nil {
		log.Fatal(err.Error(), "error during body request decoding")
	}
	check := dc.dogService.CheckIfDogExist(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := dc.dogService.UpdateDog(&dog, id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = json.NewEncoder(w).Encode(&dog)
		}
	}
}
