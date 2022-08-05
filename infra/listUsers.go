package infra

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ListUsers(db *sql.DB, w http.ResponseWriter) {
	get, err := db.Query("SELECT * FROM `rampup`.`users`")
	if err != nil {
		fmt.Println("Insert Query failed")
		log.Fatal(err)
	}
	defer get.Close()
	json.NewEncoder(w).Encode(get)

}
