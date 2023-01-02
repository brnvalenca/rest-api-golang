package services

import (
	"errors"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/dtos"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/utils"
	"strconv"
)

type IDogService interface {
	ValidateDog(d *entities.Dog) error
	FindDogs() ([]dtos.DogDTO, error)
	FindDogByID(id string) (*dtos.DogDTO, error)
	DeleteDog(id string) (*dtos.DogDTO, error)
	UpdateDog(d *dtos.DogDTO, id string) error
	CreateDog(dogDto *dtos.DogDTO) error
	CheckIfDogExistServ(id string) bool
	CheckIfKennelExistServ(dogDto *dtos.DogDTO) bool
	CheckIfBreedExistServ(dogDto *dtos.DogDTO) bool
}

type dogService struct {
	breedRepo  repository.IBreedRepository
	dogRepo    repository.IDogRepository
	kennelRepo repository.IKennelRepository
}

func NewDogService(dogRepo repository.IDogRepository, breedRepo repository.IBreedRepository, kennelBreed repository.IKennelRepository) *dogService {
	return &dogService{breedRepo: breedRepo, dogRepo: dogRepo, kennelRepo: kennelBreed}
}

func (dogserv *dogService) ValidateDog(d *entities.Dog) error {
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

func (dogserv *dogService) FindDogs() ([]dtos.DogDTO, error) {
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

func (dogserv *dogService) FindDogByID(id string) (*dtos.DogDTO, error) {
	dog, err := dogserv.dogRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
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

func (dogserv *dogService) DeleteDog(id string) (*dtos.DogDTO, error) {
	dog, err := dogserv.dogRepo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
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

func (dogserv *dogService) UpdateDog(dogDto *dtos.DogDTO, id string) error {

	dogBuilder := entities.NewDogBuilder()
	dogBuilder.Has().
		KennelID(dogDto.KennelID).
		BreedID(dogDto.BreedID).
		DogID(dogDto.DogID).
		NameAndSex(dogDto.DogName, dogDto.Sex)
	dog := dogBuilder.BuildDog()

	check := dogserv.CheckIfDogExistServ(id)
	if !check {
		return fmt.Errorf("dog not found")
	}
	// TODO: This check should recieve only the kennelID as argument, and be called in the beginning of the function to spare unecessary processing spent
	kennelCheck := dogserv.CheckIfKennelExistServ(dogDto)
	if !kennelCheck {
		return fmt.Errorf("kennel not found")
	}

	err := dogserv.dogRepo.Update(dog, id)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func (dogserv *dogService) CreateDog(dogDto *dtos.DogDTO) error {
	dog, breed := utils.PartitionDogDTO(*dogDto)
	breedCheck := dogserv.CheckIfBreedExistServ(dogDto)
	if !breedCheck {
		return fmt.Errorf("breed not found")
	}
	kennelCheck := dogserv.CheckIfKennelExistServ(dogDto)
	if !kennelCheck {
		return fmt.Errorf("kennel not found")
	}

	if breed != nil {
		err := dogserv.dogRepo.Save(dog, breed.ID)
		if err != nil {
			log.Fatal(err.Error(), "\n service error during dog creation")
		}
		return nil
	} else {
		_, err := dogserv.breedRepo.Save(breed)
		if err != nil {
			log.Fatal(err.Error(), "\n service error during breed creation")
		}
		dogserv.dogRepo.Save(dog, breed.ID)
	}
	return nil
}

func (dogserv *dogService) CheckIfDogExistServ(id string) bool {
	return dogserv.dogRepo.CheckIfExists(id)
}

func (dogserv *dogService) CheckIfKennelExistServ(dogDto *dtos.DogDTO) bool {
	id := strconv.Itoa(dogDto.KennelID)
	return dogserv.kennelRepo.CheckIfKennelExistsRepo(id)
}

func (dogserv *dogService) CheckIfBreedExistServ(dogDto *dtos.DogDTO) bool {
	id := strconv.Itoa(dogDto.BreedID)
	return dogserv.breedRepo.CheckIfExists(id)
}
