package controllers

import (
	"encoding/json"
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
	var appError dtos.AppErrorDTO
	err := json.NewDecoder(r.Body).Decode(&dogDto)
	if err != nil {
		appError.Code = http.StatusBadRequest
		appError.Message = "Could not read request body"
		json.NewEncoder(w).Encode(appError)
		return
	}
	breedCheck := dc.dogService.CheckIfBreedExist(&dogDto)
	if !breedCheck {
		appError.Code = http.StatusNotFound
		appError.Message = "Breed doesnt exist"
		json.NewEncoder(w).Encode(appError)
		return
	}
	checkKennel := dc.dogService.CheckIfKennelExist(&dogDto)
	if !checkKennel {
		appError.Code = http.StatusNotFound
		appError.Message = "Kennel doesnt exist"
		json.NewEncoder(w).Encode(appError)
		return
	} else {
		dog, breed := middleware.PartitionDogDTO(dogDto)
		dc.dogService.CreateDog(dog, breed)
		json.NewEncoder(w).Encode(1)
	}
}

func (dc *dogController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dogs, err := dc.dogService.FindDogs()
	var appError dtos.AppErrorDTO
	if err != nil {
		appError.Code = http.StatusInternalServerError
		appError.Message = "Failed to return dogs"
		json.NewEncoder(w).Encode(appError)
		return
	}
	json.NewEncoder(w).Encode(dogs)
}

func (dc *dogController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var appError dtos.AppErrorDTO
	id := params["id"]
	check := dc.dogService.CheckIfDogExist(id)
	if !check {
		appError.Code = http.StatusNotFound
		appError.Message = "Dog doesnt exist"
		json.NewEncoder(w).Encode(appError)
		return
	} else {
		dog, err := dc.dogService.FindDogByID(id)
		if err != nil {
			appError.Code = http.StatusInternalServerError
			appError.Message = "Failed to return dog"
			json.NewEncoder(w).Encode(appError)
			return
		}
		json.NewEncoder(w).Encode(dog)
	}
}

func (dc *dogController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var appError dtos.AppErrorDTO
	id := params["id"]
	check := dc.dogService.CheckIfDogExist(id)
	if !check {
		appError.Code = http.StatusNotFound
		appError.Message = "Dog doesnt exist"
		json.NewEncoder(w).Encode(appError)
		return
	} else {
		dog, err := dc.dogService.DeleteDog(id)
		if err != nil {
			appError.Code = http.StatusInternalServerError
			appError.Message = "Failed to return dog"
			json.NewEncoder(w).Encode(appError)
			return
		}
		json.NewEncoder(w).Encode(dog)
	}
}

func (dc *dogController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var appError dtos.AppErrorDTO
	id := params["id"]
	var dog dtos.DogDTO
	err := json.NewDecoder(r.Body).Decode(&dog)
	if err != nil {
		appError.Code = http.StatusBadRequest
		appError.Message = "Could not read request body"
		json.NewEncoder(w).Encode(appError)
		return
	}
	check := dc.dogService.CheckIfDogExist(id)
	if !check {
		appError.Code = http.StatusNotFound
		appError.Message = "Dog doesnt exist"
		json.NewEncoder(w).Encode(appError)
		return
	} else {
		err := dc.dogService.UpdateDog(&dog, id)
		if err != nil {
			appError.Code = http.StatusInternalServerError
			appError.Message = "Failed to update dog"
			json.NewEncoder(w).Encode(appError)
			return
		} else {
			json.NewEncoder(w).Encode(&dog)
		}
	}
}
