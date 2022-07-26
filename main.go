package main

import (
	"rest-api/golang/exercise/routes"
	"rest-api/golang/exercise/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.DB = utils.DBConn()
	routes.HandleAllReq()
	defer utils.DB.Close()
}
