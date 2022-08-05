package infra

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
)

func PostUser(u entities.User, db *sql.DB) {
	insert, err := db.Query("INSERT INTO `rampup`.`users` (`nome`,`email`,`passwd`) VALUES (?, ?, ?)", u.Name, u.Email, u.Password)
	if err != nil {
		fmt.Println("Insert Query failed")
		log.Fatal(err)
	}
	defer insert.Close()
}
