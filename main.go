package main

import (
	"rest-api/golang/exercise/data"
	"rest-api/golang/exercise/routes"
	"rest-api/golang/exercise/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	utils.DB = utils.DBConn()
	data.Users = data.MakeUsers(data.Users)
	data.Dogs = data.MakeDogs(data.Dogs)
	routes.HandleRequest()
	defer utils.DB.Close()

}
