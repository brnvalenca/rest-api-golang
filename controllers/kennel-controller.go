package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

var (
	kennelService services.IKennelService
)

type kennelController struct{}

func NewKennelController(service services.IKennelService) IController {
	kennelService = service
	return &kennelController{}
}

func (*kennelController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	kennels, err := kennelService.FindAll()
	if err != nil {
		fmt.Printf("Error with ListUsers: %v", err)
	}
	json.NewEncoder(w).Encode(kennels)
}

func (*kennelController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)                     // take the parameters of the request
	id := params["id"]                        // take the id from the parameters
	kennel, err := kennelService.FindById(id) // call the service function
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		json.NewEncoder(w).Encode(kennel)
	}
}

func (*kennelController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var kennel entities.Kennel                     // declares a kennel obj that will store the data
	err := json.NewDecoder(r.Body).Decode(&kennel) // decode the request body to the kennel obj
	if err != nil {
		log.Fatal(err.Error(), "error during request body decoding")
	}
	row, err := kennelService.Save(&kennel)
	if err != nil {
		log.Fatal(err.Error(), "kennelService.Create() error")
	}
	json.NewEncoder(w).Encode(row) // encode the created element to the w response
}

func (*kennelController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // getting the parameters of the request
	id := params["id"]    // getting the id from the parameters

	check := kennelService.CheckIfExists(id) // checking if there is and id
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		kennel, err := kennelService.Delete(id)
		if err != nil {
			log.Fatal(err.Error(), "error in kennelService.Delete() func")
		}
		json.NewEncoder(w).Encode(kennel)
	}
}

func (*kennelController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var kennel entities.Kennel
	err := json.NewDecoder(r.Body).Decode(&kennel)
	if err != nil {
		log.Fatal(err.Error(), "error decoding request body")
	}
	check := kennelService.CheckIfExists(id)

	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := kennelService.Update(&kennel, id)
		if err != nil {
			log.Fatal(err.Error(), "error during kennelService.Update() func")
		} else {
			json.NewEncoder(w).Encode(&kennel)
		}
	}

}
