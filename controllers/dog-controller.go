package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"

	"github.com/gorilla/mux"
)

/*
	Dog controller will need an instance of Dog Service to work
	Router - Controller - Service - Repo - Database
*/

type dogController struct {}

func NewDogController() Controller {
	return &dogController{}
}

var (

)

func (*dogController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dog entities.Dog
	_ = json.NewDecoder(r.Body).Decode(&dog)

	d, err := repository.Save(&dog)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(d)
}

func (*dogController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dogs, err := repository.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(dogs)
}

func (*dogController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	dog, err := repository.FindById(id)

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(dog)
}

func (*dogController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["id"]
	dog, err := repository.Delete(id)
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

	check := repository.CheckIfExists(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := repository.Update(&dog, id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = json.NewEncoder(w).Encode(&dog)
		}
	}
}
