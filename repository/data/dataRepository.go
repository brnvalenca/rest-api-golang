package repository

import (
	"rest-api/golang/exercise/domain/entities"
)

var Users []entities.User
var Dogs []entities.Dog

func MakeUsers(u []entities.User) []entities.User {

	u = append(u, entities.BuildUser("1", "Bruno", "brnvalenca@gmail.com", "123"))
	u = append(u, entities.BuildUser("2", "Luis", "lrcv@gmail.com", "453"))
	u = append(u, entities.BuildUser("3", "Miranda", "mrnd@gmail.com", "897"))
	u = append(u, entities.BuildUser("4", "Mahmed", "mahmed@gmail.com", "568"))

	return u

}

func MakeDogs(d []entities.Dog) []entities.Dog {

	b := entities.BuildBreed("Yorkshire", "Small", "3", "5")
	d = append(d, entities.BuildDog(b, 1, true, "Max", "Black", "1"))

	b = entities.BuildBreed("Doberman", "Big", "7", "6")
	d = append(d, entities.BuildDog(b, 11, false, "Bella", "Grey", "2"))

	b = entities.BuildBreed("Pug", "Small", "2", "3")
	d = append(d, entities.BuildDog(b, 7, false, "Menina", "Brown", "3"))

	b = entities.BuildBreed("Pastor Alemão", "Big", "2", "3")
	d = append(d, entities.BuildDog(b, 5, true, "Pintado", "Brown", "4"))

	return d
}
