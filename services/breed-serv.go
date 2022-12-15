package services

import (
	"errors"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/repository"
	"strconv"
)

type IBreedService interface {
	ValidateBreed(d *dtos.BreedDTO) error
	FindBreeds() ([]dtos.BreedDTO, error)
	FindBreedByID(id string) (*dtos.BreedDTO, error)
	UpdateBreed(d *dtos.BreedDTO) error
	CreateBreed(d *dtos.BreedDTO) error
	DeleteBreed(id string) (*dtos.BreedDTO, error)
	CheckIfBreedExist(id string) bool
}

type breedService struct {
	breedRepository repository.IBreedRepository
}

func NewBreedService(repo repository.IBreedRepository) IBreedService {
	return &breedService{breedRepository: repo}
}

func (bs *breedService) CreateBreed(d *dtos.BreedDTO) error {
	b := entities.NewDogBreedBuilder()
	b.Has().
		ID(d.ID).
		Name(d.Name).
		Img(d.BreedImg).
		GoodWithKidsAndDogs(d.GoodWithKids, d.GoodWithDogs).
		SheddGroomAndEnergy(d.Shedding, d.Grooming, d.Energy)

	breed := b.BuildBreed()
	_, err := bs.breedRepository.Save(breed)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (bs *breedService) UpdateBreed(d *dtos.BreedDTO) error {
	b := entities.NewDogBreedBuilder()
	b.Has().
		ID(d.ID).
		Name(d.Name).
		Img(d.BreedImg).
		GoodWithKidsAndDogs(d.GoodWithKids, d.GoodWithDogs).
		SheddGroomAndEnergy(d.Shedding, d.Grooming, d.Energy)

	breed := b.BuildBreed()
	err := bs.breedRepository.Update(breed)
	if err != nil {
		return fmt.Errorf(err.Error(), "error during UpdateBreed function")
	}
	return nil
}

func (bs *breedService) FindBreedByID(id string) (*dtos.BreedDTO, error) {

	breed, err := bs.breedRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("breed not found %w", err)
	}

	bdto := dtos.NewBreedBuilderDTO()
	bdto.Has().
		ID(breed.ID).
		Name(breed.Name).
		Img(breed.BreedImg).
		GoodWithKidsAndDogs(breed.GoodWithKids, breed.GoodWithDogs).
		SheddGroomAndEnergy(breed.Shedding, breed.Grooming, breed.Energy)

	breedDTO := bdto.BuildBreedDTO()

	return breedDTO, nil

}

func (bs *breedService) FindBreeds() ([]dtos.BreedDTO, error) {
	breeds, err := bs.breedRepository.FindAll()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	var breedsDTO []dtos.BreedDTO
	breedDtoBuilder := dtos.NewBreedBuilderDTO()

	for i := 0; i < len(breeds); i++ {
		breedDtoBuilder.Has().
			ID(breeds[i].ID).
			Name(breeds[i].Name).
			Img(breeds[i].BreedImg).
			GoodWithKidsAndDogs(breeds[i].GoodWithKids, breeds[i].GoodWithDogs).
			SheddGroomAndEnergy(breeds[i].Shedding, breeds[i].Grooming, breeds[i].Energy)
		breedDTO := breedDtoBuilder.BuildBreedDTO()
		breedsDTO = append(breedsDTO, *breedDTO)
	}

	return breedsDTO, nil
}

func (bs *breedService) DeleteBreed(id string) (*dtos.BreedDTO, error) {
	breed, err := bs.breedRepository.Delete(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	} else {
		breedBuilder := dtos.NewBreedBuilderDTO()
		breedBuilder.Has().
			ID(breed.ID).
			Name(breed.Name).
			Img(breed.BreedImg).
			GoodWithKidsAndDogs(breed.GoodWithKids, breed.GoodWithDogs).
			SheddGroomAndEnergy(breed.Shedding, breed.Grooming, breed.Energy)
		breedDto := breedBuilder.BuildBreedDTO()
		return breedDto, nil
	}
}

func (bs *breedService) CheckIfBreedExist(id string) bool {
	return bs.breedRepository.CheckIfExists(id)

}

func (*breedService) ValidateBreed(d *dtos.BreedDTO) error {
	idStr := strconv.Itoa(d.ID)
	if idStr == "" {
		err := errors.New("breed must have an valid ID")
		return err
	}
	if d.BreedImg == "" {
		err := errors.New("breed image is empty")
		return err
	}
	if d.Energy < 0 {
		err := errors.New("breed energy cannot be negative")
		return err
	}
	if d.GoodWithDogs < 0 {
		err := errors.New("good with dogs field cannot be negative")
		return err
	}
	if d.GoodWithKids < 0 {
		err := errors.New("good with kids field cannot be negative")
		return err
	}
	if d.Grooming < 0 {
		err := errors.New("grooming field cannot be negative")
		return err
	}
	if d.Name == "" {
		err := errors.New("breed must have a name")
		return err
	}
	if d.Shedding < 0 {
		err := errors.New("shedding field cannot be negative")
		return err
	}
	if d == nil {
		err := errors.New("breed must not be nil")
		return err
	}
	return nil
}
