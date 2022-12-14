package config

import "fmt"

/*
 db_config file has the objective to hold all the configs about the db connection that i need
 to establish a connection with my database. It holds a datastructure with the configs and
 two functions that will initialize the struct with the data that i need and a function that
 will return a DSN with all these configs to be used by the utils.go file inside the utils folder
 to perform a sql.Open function
*/

type DBConf struct {
	User     string
	Passwd   string
	ConnType string
	HostName string
	DBName   string
}

func DBConfInit() *DBConf {

	dbconf := DBConf{
		User:     "root",
		Passwd:   "*P*ndor*2018*",
		ConnType: "tcp",
		HostName: "localhost:3306",
		DBName:   "rampup",
	}

	return &dbconf
}

func DSN(db DBConf) string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", db.User, db.Passwd, db.ConnType, db.HostName, db.DBName)
}
