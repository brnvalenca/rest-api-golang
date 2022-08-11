package services

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/infra"
	"rest-api/golang/exercise/utils"
)

/*
	This function should not call the infra.PostUser directly. The service function should deal with all
	the logic business and validation of data requests that are made by the controllers. The functions
	that make queries to the database should be defined in a repository layer.
*/

func StoreUser(u entities.User) {
	user := u

	infra.PostUser(user, utils.DB) // repository

}
