package controllers

import (
	"encoding/json"
	"log"
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
	if err != nil {
		log.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(breeds)
}

func (b *breedController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	breed, err := b.breedService.FindBreedByID(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(breed)
}

func (b *breedController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var breedDTO dtos.BreedDTO
	err := json.NewDecoder(r.Body).Decode(&breedDTO)
	if err != nil {
		log.Fatal(err.Error())
	}
	breedValidation := b.breedService.ValidateBreed(&breedDTO)
	if breedValidation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(breedValidation.Error()))
	}
	err = b.breedService.CreateBreed(&breedDTO)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(breedDTO)
}

func (b *breedController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("only allowed to admin"))
}

func (b *breedController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var breedDTO dtos.BreedDTO

	err := json.NewDecoder(r.Body).Decode(&breedDTO)
	if err != nil {
		log.Fatal(err.Error())
	}
	breedValidation := b.breedService.ValidateBreed(&breedDTO)
	if breedValidation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(breedValidation.Error()))
	}
	err = b.breedService.UpdateBreed(&breedDTO)
	if err != nil {
		log.Fatal(err.Error())
	}
}
