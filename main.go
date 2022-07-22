package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
	json.NewEncoder(w).Encode(entities.UsersData)
}

func printHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. Second Page")

}

func handleRequest() {
	http.HandleFunc("/users", homePage)
	http.HandleFunc("/second", printHello)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {

	name := "Bruno Valença"
	email := "brunovalenca10@gmail.com"
	password := "123"

	users := entities.BuildUser(name, email, password)

	entities.UsersData = append(entities.UsersData, users)

	fmt.Println(entities.UsersData)
	handleRequest()
}
