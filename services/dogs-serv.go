package services

import (
	"errors"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/repository"
	"strconv"
)

type IDogService interface {
	ValidateDog(d *entities.Dog) error
	FindDogs() ([]dtos.DogDTO, error)
	FindDogByID(id string) (*dtos.DogDTO, error)
	DeleteDog(id string) (*dtos.DogDTO, error)
	UpdateDog(d *dtos.DogDTO, id string) error
	CreateDog(d *entities.Dog, b *entities.DogBreed) error
	CheckIfDogExist(id string) bool
	CheckIfKennelExist(d *dtos.DogDTO) bool
	CheckIfBreedExist(d *dtos.DogDTO) bool
}

type dserv struct {
	breedRepo  repository.IBreedRepository
	dogRepo    repository.IDogRepository
	kennelRepo repository.IKennelRepository
}

func NewDogService(dogRepo repository.IDogRepository, breedRepo repository.IBreedRepository, kennelBreed repository.IKennelRepository) IDogService {
	return &dserv{breedRepo: breedRepo, dogRepo: dogRepo, kennelRepo: kennelBreed}
}

func (*dserv) ValidateDog(d *entities.Dog) error {
	if d == nil {
		err := errors.New("dog is empty")
		return err
	}
	if d.DogName == "" {
		err := errors.New("dog name is empty")
		return err
	}
	if d.Sex == "" {
		err := errors.New("dog sex is empty")
		return err
	}
	return nil

}

func (dogserv *dserv) FindDogs() ([]dtos.DogDTO, error) {
	dogs, err := dogserv.dogRepo.FindAll()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	var dogsDTO []dtos.DogDTO
	dogDtoBuilder := dtos.NewDogDTOBuilder()

	for i := 0; i < len(dogs); i++ {
		dogDtoBuilder.Has().
			KennelID(dogs[i].KennelID).
			BreedID(dogs[i].BreedID).
			DogID(dogs[i].DogID).
			NameAndSex(dogs[i].DogName, dogs[i].Sex).
			DogDTOBreedName(dogs[i].Breed.Name).
			DogDTOBreedImg(dogs[i].Breed.BreedImg).
			DogDTOGoodWithKidsAndDogs(dogs[i].Breed.GoodWithKids, dogs[i].Breed.GoodWithKids).
			DogDTOSheddGroomAndEnergy(dogs[i].Breed.Shedding, dogs[i].Breed.Grooming, dogs[i].Breed.Energy)
		dogDto := dogDtoBuilder.BuildDogDTO()
		dogsDTO = append(dogsDTO, *dogDto)
	}

	return dogsDTO, nil

}

func (dogserv *dserv) FindDogByID(id string) (*dtos.DogDTO, error) {
	dog, err := dogserv.dogRepo.FindById(id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	dogDtoBuilder := dtos.NewDogDTOBuilder()
	dogDtoBuilder.Has().
		KennelID(dog.KennelID).
		BreedID(dog.BreedID).
		DogID(dog.DogID).
		NameAndSex(dog.DogName, dog.Sex).
		DogDTOBreedName(dog.Breed.Name).
		DogDTOBreedImg(dog.Breed.BreedImg).
		DogDTOGoodWithKidsAndDogs(dog.Breed.GoodWithKids, dog.Breed.GoodWithKids).
		DogDTOSheddGroomAndEnergy(dog.Breed.Shedding, dog.Breed.Grooming, dog.Breed.Energy)

	dogDto := dogDtoBuilder.BuildDogDTO()

	return dogDto, nil

}

func (dogserv *dserv) DeleteDog(id string) (*dtos.DogDTO, error) {
	dog, err := dogserv.dogRepo.Delete(id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	dogDtoBuilder := dtos.NewDogDTOBuilder()
	dogDtoBuilder.Has().
		KennelID(dog.KennelID).
		BreedID(dog.BreedID).
		DogID(dog.DogID).
		NameAndSex(dog.DogName, dog.Sex).
		DogDTOBreedName(dog.Breed.Name).
		DogDTOBreedImg(dog.Breed.BreedImg).
		DogDTOGoodWithKidsAndDogs(dog.Breed.GoodWithKids, dog.Breed.GoodWithKids).
		DogDTOSheddGroomAndEnergy(dog.Breed.Shedding, dog.Breed.Grooming, dog.Breed.Energy)
	dogDto := dogDtoBuilder.BuildDogDTO()
	return dogDto, nil

}

func (dogserv *dserv) UpdateDog(u *dtos.DogDTO, id string) error {
	dogBuilder := entities.NewDogBuilder()
	dogBuilder.Has().
		KennelID(u.KennelID).
		BreedID(u.BreedID).
		DogID(u.DogID).
		NameAndSex(u.DogName, u.Sex)
	dog := dogBuilder.BuildDog()
	err := dogserv.dogRepo.Update(dog, id)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func (dogserv *dserv) CreateDog(d *entities.Dog, b *entities.DogBreed) error {
	if b != nil {
		err := dogserv.dogRepo.Save(d, b.ID)
		if err != nil {
			log.Fatal(err.Error(), "\n service error during dog creation")
		}
		return nil
	} else {
		_, err := dogserv.breedRepo.Save(b)
		if err != nil {
			log.Fatal(err.Error(), "\n service error during breed creation")
		}
		dogserv.dogRepo.Save(d, b.ID)
	}
	return nil
}

func (dogserv *dserv) CheckIfDogExist(id string) bool {
	return dogserv.dogRepo.CheckIfExists(id)
}

func (dogserv *dserv) CheckIfKennelExist(d *dtos.DogDTO) bool {
	id := strconv.Itoa(d.KennelID)
	return dogserv.kennelRepo.CheckIfExistsRepo(id)
}

func (dogserv *dserv) CheckIfBreedExist(d *dtos.DogDTO) bool {
	id := strconv.Itoa(d.BreedID)
	return dogserv.breedRepo.CheckIfExists(id)
}
