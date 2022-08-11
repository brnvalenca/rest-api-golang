package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var DB *sql.DB

func DBConn() *sql.DB {
	db, err := sql.Open("mysql", "root:*P*ndor*2018*@tcp(localhost:3306)/rampup")
	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}

	fmt.Println("Connection Succeded")
	return db
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
