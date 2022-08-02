package repository

import (
	"rest-api/golang/exercise/domain/entities"
)

var Users []entities.User
var Dogs []entities.Dog

func MakeUsers(u []entities.User) []entities.User {
	udog := entities.BuildUserDogPreferences(4, 6, "Medium")
	a := entities.BuildAddress("Rua Salomao", "Graças", "52050-002", "Recife")
	u = append(u, entities.BuildUser("1", "Bruno", "brnvalenca@gmail.com", "123", a, udog))

	udog = entities.BuildUserDogPreferences(1, 3, "Small" )
	a = entities.BuildAddress("Rua 11", "Ibura", "51528-052", "Recife")
	u = append(u, entities.BuildUser("2", "Luis", "lrcv@gmail.com", "453", a, udog))

	udog = entities.BuildUserDogPreferences(8, 6, "Dont mind")
	a = entities.BuildAddress("Rua do Futuro", "Jaqueira", "52050-002", "Recife")
	u = append(u, entities.BuildUser("3", "Miranda", "mrnd@gmail.com", "897", a, udog))

	udog = entities.BuildUserDogPreferences(10, 8, "Large")
	a = entities.BuildAddress("Rua do Forte", "Brum", "52050-002", "Recife")
	u = append(u, entities.BuildUser("4", "Mahmed", "mahmed@gmail.com", "568", a, udog))

	return u

}

func MakeDogs(d []entities.Dog) []entities.Dog {

	b := entities.BuildBreed("Yorkshire", "Small", "7", "5")
	d = append(d, entities.BuildDog(b, 1, "Male", "Max", "1"))

	b = entities.BuildBreed("Doberman", "Big", "3", "6")
	d = append(d, entities.BuildDog(b, 2, "Female", "Bella", "2"))

	b = entities.BuildBreed("Pug", "Small", "2", "3")
	d = append(d, entities.BuildDog(b, 3, "Female", "Menina", "3"))

	b = entities.BuildBreed("Pastor Alemão", "Big", "2", "3")
	d = append(d, entities.BuildDog(b, 4, "Female", "Pintado", "4"))

	return d
}
