package repository

import (
	"rest-api/golang/exercise/domain/entities"
)

func MakeUsers(u []entities.User) []entities.User {

	u = append(u, entities.BuildUser("1", "Bruno", "brnvalenca@gmail.com", "123"))
	u = append(u, entities.BuildUser("2", "Luis", "lrcv@gmail.com", "453"))
	u = append(u, entities.BuildUser("3", "Miranda", "mrnd@gmail.com", "897"))
	u = append(u, entities.BuildUser("4", "Mahmed", "mahmed@gmail.com", "568"))

	return u
	
}

/* 
	This layer is has the responsability of populate the slice of users with data so that can be 
*/

