package utils

import (
	"database/sql"
	"fmt"
	"rest-api/golang/exercise/config"
)

var DB *sql.DB

func DBConn() *sql.DB {
	DBconfig := config.DBConfInit()
	db, err := sql.Open("mysql", config.DSN(*DBconfig))
	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}

	fmt.Println("Connection Succeded")
	return db
}
