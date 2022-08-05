package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/data"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ch := make(chan http.ResponseWriter)

	db, err := sql.Open("mysql", "root:*P*ndor*2018*@tcp(localhost:3306)/rampup")
	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM `rampup`.`users`")
	if err != nil {
		fmt.Println("Insert Query failed")
		log.Fatal(err)
	}
	defer rows.Close()

	
	json.NewEncoder(w).Encode(get)

	//services.ListUsers(ch)
	fmt.Println(w)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var idString string
	for _, user := range data.Users {
		idString = strconv.Itoa(user.ID)
		if idString == params["id"] {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	_ = json.NewDecoder(r.Body).Decode(&user) // Aqui eu decodifico o body da requisicao, que estará em JSON, contendo os dados do user
	services.StoreUser(user)
	json.NewEncoder(w).Encode(user) // Codifico a resposta guardada em w para JSON e mostro na tela.
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var idString string
	for index, user := range data.Users {
		idString = strconv.Itoa(user.ID)
		if idString == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.Users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var idString string
	for index, user := range data.Users {
		idString = strconv.Itoa(user.ID)
		if idString == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			var user entities.User
			userIDInt, _ := strconv.Atoi(params["id"])
			user.ID = userIDInt
			_ = json.NewDecoder(r.Body).Decode(&user)
			data.Users = append(data.Users, user)
			json.NewEncoder(w).Encode(data.Users)
		}
	}
}
