package middleware

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dto"

	"golang.org/x/crypto/bcrypt"
)

func PartitionData(u *entities.User, userID int) *entities.UserDogPreferences {

	upref := entities.NewUserDogPrefsBuilder()
	upref.Has().
		UserID(userID).
		GoodWithKidsAndDogs(u.UserPreferences.GoodWithKids, u.UserPreferences.GoodWithDogs).
		SheddGroomAndEnergy(u.UserPreferences.Shedding, u.UserPreferences.Grooming, u.UserPreferences.Energy)

	prefs := upref.BuildUserPref()

	return prefs
}

func PartitionKennelAddress(k *entities.Kennel, kennelID int) *entities.Address {
	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(kennelID).
		Numero(k.Address.Numero).
		Rua(k.Address.Rua).
		Bairro(k.Address.Bairro).
		CEP(k.Address.Bairro).
		Cidade(k.Address.Cidade)

	addr := ad.BuildAddr()
	return addr
}

func PartitionDogDTO(dto dto.DogDTO) (*entities.Dog, *entities.DogBreed) {
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(dto.BreedID).
		Name(dto.DogName).
		Img(dto.BreedImg).
		GoodWithKidsAndDogs(dto.GoodWithKids, dto.GoodWithDogs).
		SheddGroomAndEnergy(dto.Shedding, dto.Grooming, dto.Energy)
	dogbreed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(dto.KennelID).
		DogID(dto.DogID).
		NameAndSex(dto.DogName, dto.Sex).
		Breed(*dogbreed)
	dog := d.BuildDog()

	return dog, dogbreed
}

func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(hashedPassword, passwordString string) bool {
	var passwordOK bool
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordString))
	if passwordCheck != nil {
		passwordOK = true
	}
	return passwordOK
}
