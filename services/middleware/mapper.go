package middleware

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dto"
)

func PartitionData(u *entities.User, userID int) *entities.UserDogPreferences {
	prefs :=
		entities.BuildUserDogPreferences(
			userID,
			u.UserPreferences.GoodWithKids,
			u.UserPreferences.GoodWithDogs,
			u.UserPreferences.Shedding,
			u.UserPreferences.Grooming,
			u.UserPreferences.Energy,
		)
	return &prefs
}

func PartitionKennelAddress(k *entities.Kennel, kennelID int) *entities.Address {
	addr :=
		entities.BuildAddress(
			kennelID,
			k.Address.Numero,
			k.Address.Rua,
			k.Address.Bairro,
			k.Address.CEP,
			k.Address.Cidade,
		)
	return addr
}

func PartitionDogDTO(dto dto.DogDTO) (*entities.Dog, *entities.DogBreed) {
	breed := entities.BuildDogBreed(
		dto.BreedImg,
		dto.DogName,
		dto.BreedID,
		dto.DogID,
		dto.GoodWithKids,
		dto.GoodWithDogs,
		dto.Shedding,
		dto.Grooming,
		dto.Energy,
	)
	dog := entities.BuildDog(
		*breed,
		dto.DogID,
		dto.KennelID,
		dto.Sex,
		dto.DogName,
	)
	return dog, breed
}
