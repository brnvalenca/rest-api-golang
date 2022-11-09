package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

type breedController struct {
	breedService services.IBreedService
}

func NewBreedController(service services.IBreedService) IController {

	return &breedController{breedService: service}
}

// TODO: Essa funcao de retorno de Breed por id chama uma funcao do service que retorna um
// entities.DogBreed e não um DTO, é correto isso?

func (b *breedController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	breeds, err := b.breedService.FindBreeds()
	var appError dtos.AppErrorDTO
	if err != nil {
		appError.Code = http.StatusInternalServerError
		appError.Message = "Failed to get breeds"
		json.NewEncoder(w).Encode(appError)

		return
	}
	json.NewEncoder(w).Encode(breeds)
}

func (b *breedController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var appError dtos.AppErrorDTO
	breed, err := b.breedService.FindBreedByID(id)
	if err != nil {
		appError.Code = http.StatusNotFound
		appError.Message = "Breed not found"
		json.NewEncoder(w).Encode(appError)
		return
	}
	json.NewEncoder(w).Encode(breed)
}

func (b *breedController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var breedDTO dtos.BreedDTO
	var appError dtos.AppErrorDTO
	err := json.NewDecoder(r.Body).Decode(&breedDTO)
	if err != nil {
		appError.Code = http.StatusNotFound
		appError.Message = "Breed not found"
		json.NewEncoder(w).Encode(appError)
		return
	}
	breedValidation := b.breedService.ValidateBreed(&breedDTO)
	if breedValidation != nil {
		appError.Code = http.StatusBadRequest
		appError.Message = "Breed not valid"
		json.NewEncoder(w).Encode(appError)
		return
	}
	err = b.breedService.CreateBreed(&breedDTO)
	if err != nil {
		appError.Code = http.StatusInternalServerError
		appError.Message = "Failed to create breed"
		json.NewEncoder(w).Encode(appError)
		return
	}
	json.NewEncoder(w).Encode(breedDTO)
}

func (b *breedController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appError dtos.AppErrorDTO
	appError.Code = http.StatusMethodNotAllowed
	appError.Message = "Method only allowed for admin"
	json.NewEncoder(w).Encode(appError)
}

func (b *breedController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appError dtos.AppErrorDTO
	var breedDTO dtos.BreedDTO

	err := json.NewDecoder(r.Body).Decode(&breedDTO)
	if err != nil {
		appError.Code = http.StatusBadRequest
		appError.Message = "Could not read request body"
		json.NewEncoder(w).Encode(appError)
		return
	}
	breedValidation := b.breedService.ValidateBreed(&breedDTO)
	if breedValidation != nil {
		appError.Code = http.StatusBadRequest
		appError.Message = "Breed not valid"
		json.NewEncoder(w).Encode(appError)
		return
	}
	err = b.breedService.UpdateBreed(&breedDTO)
	if err != nil {
		appError.Code = http.StatusInternalServerError
		appError.Message = "Failed to update breed"
		json.NewEncoder(w).Encode(appError)
		return
	}
}
