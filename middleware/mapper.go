package middleware

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
)

func PartitionUserDTO(u *dtos.UserDTOSignUp) (*entities.UserDogPreferences, *entities.User) {
	uprefs := entities.NewUserDogPrefsBuilder()
	uprefs.Has().
		UserID(u.ID).
		GoodWithKidsAndDogs(
			u.UserPrefs.GoodWithKids,
			u.UserPrefs.GoodWithDogs).
		SheddGroomAndEnergy(
			u.UserPrefs.Shedding,
			u.UserPrefs.Grooming,
			u.UserPrefs.Energy)

	userprefs := uprefs.BuildUserPref()

	ub := entities.NewUserBuilder()
	ub.Has().
		Name(u.Name).
		Email(u.Email).
		Password(string(u.Password)).
		Uprefs(*userprefs)

	user := ub.BuildUser()

	return userprefs, user
}

func PartitionKennelDTO(k *dtos.KennelDTO) (*entities.Address, *entities.Kennel) {
	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(k.ID).
		Numero(k.Numero).
		Rua(k.Rua).
		Bairro(k.Bairro).
		CEP(k.CEP).
		Cidade(k.Cidade)
	addr := ad.BuildAddr()

	kn := entities.NewKennelBuilder()
	kn.Has().
		ID(k.ID).
		ContactNumber(k.ContactNumber).
		Name(k.Name).
		Address(*addr)
	kennel := kn.BuildKennel()

	return addr, kennel
}

func PartitionDogDTO(dto dtos.DogDTO) (*entities.Dog, *entities.DogBreed) {
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
