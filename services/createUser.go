package services

import (
	"database/sql"
	"fmt"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/infra"
)

func StoreUser(u entities.User) {

	db, err := sql.Open("mysql", "root:*P*ndor*2018*@tcp(localhost:3306)/rampup")
	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}
	user := u
	defer db.Close()

	infra.PostUser(user, db)

}
