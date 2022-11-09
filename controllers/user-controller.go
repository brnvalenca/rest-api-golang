package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"

	"github.com/gorilla/mux"
)

type userController struct {
	passwordHash security.IPasswordHash
	userService  services.IUserService
}

func NewUserController(service services.IUserService, password security.IPasswordHash) IController {

	return &userController{passwordHash: password, userService: service}
}

func (u *userController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appErr dtos.AppErrorDTO
	var userDTO dtos.UserDTOSignUp
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		appErr.Code = http.StatusBadRequest
		appErr.Message = "Could not read request body"
		json.NewEncoder(w).Encode(appErr)
	}
	check, _ := u.userService.CheckEmailServ(userDTO.Email)
	if check {
		appErr.Code = http.StatusBadRequest
		appErr.Message = "User Already Registered"
		json.NewEncoder(w).Encode(appErr)
	}
	userDTO.Password, err = u.passwordHash.GeneratePasswordHash(userDTO.Password)
	if err != nil {
		appErr.Code = http.StatusInternalServerError
		appErr.Message = "Failed Generating Hash Password"
		json.NewEncoder(w).Encode(appErr)
	}
	user, err := u.userService.Create(&userDTO)
	// TODO : retornar um objeto com o ID dentro, por exemplo um DTO.
	if err != nil {
		appErr.Code = http.StatusBadRequest
		appErr.Message = "Status Bad Request"
		json.NewEncoder(w).Encode(appErr)
	}
	json.NewEncoder(w).Encode(user)
}

func (u *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appErr dtos.AppErrorDTO
	users, err := u.userService.FindAll()
	if err != nil {
		appErr.Code = http.StatusBadRequest
		appErr.Message = "Status Bad Request"
		json.NewEncoder(w).Encode(appErr)
	}
	json.NewEncoder(w).Encode(users)
}

func (u *userController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	user, err := u.userService.FindById(id)
	if err != nil {
		var appErr dtos.AppErrorDTO
		appErr.Code = http.StatusNotFound
		appErr.Message = "Status Not Found"
		json.NewEncoder(w).Encode(appErr)
	} else {
		json.NewEncoder(w).Encode(user)
	}

}

func (u *userController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appErr dtos.AppErrorDTO
	params := mux.Vars(r)
	id := params["id"]
	check := u.userService.Check(id)
	if !check {
		appErr.Code = http.StatusNotFound
		appErr.Message = "Status Not Found"
		json.NewEncoder(w).Encode(appErr)
	} else {
		user, err := u.userService.Delete(id)
		if err != nil {
			appErr.Code = 400
			appErr.Message = "Status Bad Request"
			json.NewEncoder(w).Encode(appErr)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (u *userController) Update(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var userDTO dtos.UserDTOSignUp
	var appErr dtos.AppErrorDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		appErr.Code = http.StatusInternalServerError
		appErr.Message = "Failed Generating Hash Password"
		json.NewEncoder(w).Encode(appErr)
	}
	params := mux.Vars(r)
	id := params["id"]
	check := u.userService.Check(id)
	if !check {
		var appErr dtos.AppErrorDTO
		appErr.Code = http.StatusNotFound
		appErr.Message = "Status Not Found"
		json.NewEncoder(w).Encode(appErr)
	} else {
		err := u.userService.UpdateUser(&userDTO)
		if err != nil {
			appErr.Code = http.StatusBadRequest
			appErr.Message = "Status Bad Request"
			json.NewEncoder(w).Encode(appErr)
		}
		json.NewEncoder(w).Encode(&userDTO)
	}
}
