package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

type kennelController struct {
	kennelService services.IKennelService
}

func NewKennelController(kennelServ services.IKennelService) IController {
	return &kennelController{kennelService: kennelServ}
}

func (kc *kennelController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	kennels, err := kc.kennelService.FindAllKennels()
	if err != nil {
		fmt.Printf("error with get all kennels: %v", err)
	}
	json.NewEncoder(w).Encode(kennels)
}

func (kc *kennelController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id := params["id"]
	kennel, err := kc.kennelService.FindKennelByIdServ(id)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		json.NewEncoder(w).Encode(kennel)
	}
}

func (kc *kennelController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var kennelDTO dtos.KennelDTO
	err := json.NewDecoder(r.Body).Decode(&kennelDTO)
	if err != nil {
		log.Fatal(err.Error(), "error during request body decoding")
	}
	row, err := kc.kennelService.SaveKennel(&kennelDTO)
	if err != nil {
		log.Fatal(err.Error(), "kennelService.Create() error")
	}
	json.NewEncoder(w).Encode(row)
}

func (kc *kennelController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	check := kc.kennelService.CheckIfExists(id)
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		kennel, err := kc.kennelService.DeleteKennelServ(id)
		if err != nil {
			log.Fatal(err.Error(), "error in kennelService.Delete() func")
		}
		json.NewEncoder(w).Encode(kennel)
	}
}

func (kc *kennelController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var kennelDTO dtos.KennelDTO

	err := json.NewDecoder(r.Body).Decode(&kennelDTO)
	if err != nil {
		log.Fatal(err.Error(), "error decoding request body")
	}
	check := kc.kennelService.CheckIfExists(id)

	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	} else {
		err := kc.kennelService.UpdateKennelServ(&kennelDTO, id)
		if err != nil {
			log.Fatal(err.Error(), "error during kennelService.Update() func")
		} else {
			json.NewEncoder(w).Encode(&kennelDTO)
		}
	}

}
