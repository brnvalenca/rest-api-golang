package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"

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
	var dog entities.Dog
	_ = json.NewDecoder(r.Body).Decode(&dog)

	err := dogService.Validate(&dog)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		dogService.Create(&dog)
		json.NewEncoder(w).Encode(dog)
	}
}

func (*dogController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dogs, err := dogService.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(dogs)
}

func (*dogController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	dog, err := dogService.FindById(id)

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(dog)
}

func (*dogController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	dog, err := dogService.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(dog)
}

func (*dogController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	var dog entities.Dog
	_ = json.NewDecoder(r.Body).Decode(&dog)

	check := dogService.Check(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := dogService.Update(&dog, id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = json.NewEncoder(w).Encode(&dog)
		}
	}
}
