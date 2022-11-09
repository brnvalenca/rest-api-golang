package controllers

import (
	"encoding/json"
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
	var appErr dtos.AppErrorDTO
	kennels, err := kc.kennelService.FindAllKennels()
	if err != nil {
		appErr.Code = http.StatusInternalServerError
		appErr.Message = "Failed to return kennels"
		json.NewEncoder(w).Encode(appErr)
		return
	}
	json.NewEncoder(w).Encode(kennels)
}

func (kc *kennelController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id := params["id"]
	kennel, err := kc.kennelService.FindKennelByIdServ(id)
	var appErr dtos.AppErrorDTO
	if err != nil {
		appErr.Code = http.StatusNotFound
		appErr.Message = "Kennel not found"
		json.NewEncoder(w).Encode(appErr)
		return
	} else {
		json.NewEncoder(w).Encode(kennel)
	}
}

func (kc *kennelController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var kennelDTO dtos.KennelDTO
	var appErr dtos.AppErrorDTO
	err := json.NewDecoder(r.Body).Decode(&kennelDTO)
	if err != nil {
		appErr.Code = http.StatusBadRequest
		appErr.Message = "Could not parse the request body"
		json.NewEncoder(w).Encode(appErr)
		return
	}
	row, err := kc.kennelService.SaveKennel(&kennelDTO)
	if err != nil {
		appErr.Code = http.StatusInternalServerError
		appErr.Message = "Failed creating new kennel"
	}
	json.NewEncoder(w).Encode(row)
}

func (kc *kennelController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appErr dtos.AppErrorDTO
	params := mux.Vars(r)
	id := params["id"]

	check := kc.kennelService.CheckIfExists(id)
	if !check {
		appErr.Code = http.StatusNotFound
		appErr.Message = "Kennel not found"
		json.NewEncoder(w).Encode(appErr)
		return
	} else {
		kennel, err := kc.kennelService.DeleteKennelServ(id)
		if err != nil {
			appErr.Code = http.StatusInternalServerError
			appErr.Message = "Failed deleting kennel"
		}
		json.NewEncoder(w).Encode(kennel)
	}
}

func (kc *kennelController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var appError dtos.AppErrorDTO
	var kennelDTO dtos.KennelDTO

	err := json.NewDecoder(r.Body).Decode(&kennelDTO)
	if err != nil {
		appError.Code = http.StatusBadRequest
		appError.Message = "Could not parse the request body"
		json.NewEncoder(w).Encode(appError)
		return
	}
	check := kc.kennelService.CheckIfExists(id)

	if !check {
		appError.Code = http.StatusNotFound
		appError.Message = "Kennel not found"
		json.NewEncoder(w).Encode(appError)
		return
	} else {
		err := kc.kennelService.UpdateKennelServ(&kennelDTO, id)
		if err != nil {
			appError.Code = http.StatusInternalServerError
			appError.Message = "Failed updating kennel"
			json.NewEncoder(w).Encode(appError)
			return
		} else {
			json.NewEncoder(w).Encode(&kennelDTO)
		}
	}
}
