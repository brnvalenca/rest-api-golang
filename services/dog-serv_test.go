package services

import (
	"errors"
	"rest-api/golang/exercise/domain/dtos"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO: Mudar o nome do mock

// mockar os repositorios

type dogRepoMock struct {
	mock.Mock
}

type breedRepoMock struct {
	mock.Mock
}

//Breed Mock

func (breed *breedRepoMock) Save(b *entities.DogBreed) (int, error) {
	args := breed.Called(b)
	return args.Int(0), args.Error(1)
}
func (breed *breedRepoMock) FindAll() ([]entities.DogBreed, error) {
	args := breed.Called()
	return args.Get(0).([]entities.DogBreed), args.Error(1)
}
func (breed *breedRepoMock) FindById(id string) (*entities.DogBreed, error) {
	args := breed.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}
func (breed *breedRepoMock) Delete(id string) (*entities.DogBreed, error) {
	args := breed.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}
func (breed *breedRepoMock) Update(b *entities.DogBreed) error {
	args := breed.Called(b)
	return args.Error(0)
}
func (breed *breedRepoMock) CheckIfExists(id string) bool {
	args := breed.Called(id)
	return args.Bool(0)
}

//Dog Mock Methods

func (dog *dogRepoMock) Save(d *entities.Dog, id interface{}) error {
	args := dog.Called(d, id)
	return args.Error(0)
}
func (dog *dogRepoMock) FindAll() ([]entities.Dog, error) {
	args := dog.Called()
	return args.Get(0).([]entities.Dog), args.Error(1)
}
func (dog *dogRepoMock) FindById(id string) (*entities.Dog, error) {
	args := dog.Called(id)
	return args.Get(0).(*entities.Dog), args.Error(1)
}
func (dog *dogRepoMock) Delete(id string) (*entities.Dog, error) {
	args := dog.Called(id)
	return args.Get(0).(*entities.Dog), args.Error(1)
}
func (dog *dogRepoMock) Update(d *entities.Dog, id string) error {
	args := dog.Called(d, id)
	return args.Error(0)
}
func (dog *dogRepoMock) CheckIfExists(id string) bool {
	args := dog.Called(id)
	return args.Bool(0)
}

// Make Breed and Dog Functions

func MakeDogDTO() dtos.DogDTO {

	d := dtos.NewDogDTOBuilder()
	d.Has().
		KennelID(2).
		DogID(1).
		BreedID(1).
		NameAndSex("M", "B").
		DogDTOBreedName("Chucky").
		DogDTOBreedImg("1").
		DogDTOGoodWithKidsAndDogs(3, 4).
		DogDTOSheddGroomAndEnergy(5, 6, 7)
	dog := d.BuildDogDTO()
	return *dog
}

// Test Dog Functions

func TestSaveDog(t *testing.T) {
	dogRepoMock := new(dogRepoMock) // mokando dog
	breedMock := new(breedRepoMock) // mokando breed
	kennelMock := new(kennelMockRepo)
	dogDto := MakeDogDTO() // criando um dog

	dogRepoMock.On("Save", dogDto, 1).Return(nil)

	// Eu desfiz a chamada do Save breed pois essa é uma funcionalidade que é acessivel apenas a usuarios admins
	testService := NewDogService(dogRepoMock, breedMock, kennelMock)
	err := testService.CreateDog(&dogDto)
	dogRepoMock.AssertExpectations(t)

	assert.Nil(t, err)

}

func TestFindAllDogs(t *testing.T) {
	dogMock := new(dogRepoMock)
	breedMock := new(breedRepoMock) // mokando breed
	kennelMock := new(kennelMockRepo)
	dogDto := MakeDogDTO()
	dog, _ := utils.PartitionDogDTO(dogDto)

	dogMock.On("FindAll").Return([]entities.Dog{*dog}, nil)

	testService := NewDogService(dogMock, breedMock, kennelMock)
	dogs, err := testService.FindDogs()

	dogMock.AssertExpectations(t)

	assert.Equal(t, 1, dogs[0].DogID)
	assert.Equal(t, 2, dogs[0].KennelID)
	assert.Equal(t, "B", dogs[0].Sex)
	assert.Equal(t, "M", dogs[0].DogName)
	assert.Equal(t, dog.BreedID, dogs[0].BreedID)
	assert.Nil(t, err)
}

func TestFindDogById(t *testing.T) {

	dogMock := new(dogRepoMock)
	breedMock := new(breedRepoMock)
	kennelMock := new(kennelMockRepo)
	dog := MakeDogDTO()
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("FindById", idStr).Return(dog, nil)

	testService := NewDogService(dogMock, breedMock, kennelMock)
	dogDto, err := testService.FindDogByID(idStr)

	dogMock.AssertExpectations(t)
	assert.Equal(t, 1, dogDto.DogID)
	assert.Equal(t, 2, dogDto.KennelID)
	assert.Equal(t, "B", dogDto.Sex)
	assert.Equal(t, "M", dogDto.DogName)
	assert.Equal(t, dog.BreedID, dogDto.BreedID)
	assert.Nil(t, err)
}

func TestDeleteDog(t *testing.T) {
	dogMock := new(dogRepoMock)
	breedMock := new(breedRepoMock)
	kennelMock := new(kennelMockRepo)
	dog := MakeDogDTO()

	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("Delete", idStr).Return(dog, nil)

	testService := NewDogService(dogMock, breedMock, kennelMock)
	result, err := testService.DeleteDog(idStr)

	dogMock.AssertExpectations(t)

	assert.Equal(t, 1, result.DogID)
	assert.Equal(t, 2, result.KennelID)
	assert.Equal(t, "B", result.Sex)
	assert.Equal(t, "M", result.DogName)
	assert.Equal(t, result.BreedID, result.BreedID)
	assert.Nil(t, err)

}

func TestUpdateDog(t *testing.T) {
	dogMock := new(dogRepoMock)
	breedMock := new(breedRepoMock)
	kennelMock := new(kennelMockRepo)

	dog := MakeDogDTO()
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("Update", dog, idStr).Return(nil)

	dog.BreedID = 7
	dog.DogName = "Z"

	testService := NewDogService(dogMock, breedMock, kennelMock)
	err := testService.UpdateDog(&dog, idStr)

	dogMock.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestUpdateDogDontExists(t *testing.T) {
	dogMock := new(dogRepoMock)
	breedMock := new(breedRepoMock)
	kennelMock := new(kennelMockRepo)
	dog := MakeDogDTO()

	dogMock.On("Update", dog, "34").Return(errors.New("update dog failed"))

	dog.BreedID = 7
	dog.DogName = "Z"
	dogDto := dtos.DogDTO{}
	testService := NewDogService(dogMock, breedMock, kennelMock)
	err := testService.UpdateDog(&dogDto, "34")

	dogMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Error(t, err, "update dog failed")

}

func TestCheckIfDontExistsDog(t *testing.T) {
	dogMock := new(dogRepoMock)
	breedMock := new(breedRepoMock)
	kennelMock := new(kennelMockRepo)

	dogMock.On("CheckIfExists", "31").Return(false)

	testService := NewDogService(dogMock, breedMock, kennelMock)
	result := testService.CheckIfDogExistServ("31")

	dogMock.AssertExpectations(t)

	assert.Equal(t, false, result)

}

func TestCheckIfExistsDogs(t *testing.T) {
	dogMock := new(dogRepoMock)
	breedMock := new(breedRepoMock)
	kennelMock := new(kennelMockRepo)
	dog := MakeDogDTO()
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("CheckIfExists", idStr).Return(true)

	testService := NewDogService(dogMock, breedMock, kennelMock)
	result := testService.CheckIfDogExistServ(idStr)

	dogMock.AssertExpectations(t)

	assert.Equal(t, true, result)
}
