package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-api/golang/exercise/infra"
)

func ListUsers(ch chan http.ResponseWriter) {
	db, err := sql.Open("mysql", "root:*P*ndor*2018*@tcp(localhost:3306)/rampup")
	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}
	w := <-ch
	defer db.Close()

	infra.ListUsers(db, w)

}
