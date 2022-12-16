package utils

import (
	"database/sql"
	"fmt"
	"rest-api/golang/exercise/config"
)

var DB *sql.DB

func DBConn(cfg config.AplicationConfig) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.DBConfig.User, cfg.DBConfig.Passwd, cfg.DBConfig.ConnType, cfg.DBConfig.HostName, cfg.DBConfig.DBName))
	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}

	fmt.Println("DB Connection Succeded")
	return db
}
