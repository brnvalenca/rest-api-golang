package tests

import (
	"errors"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockar os repositorios

type dogMock struct {
	mock.Mock
}

type breedMock struct {
	mock.Mock
}

//Breed Mock

func (breed *breedMock) Save(b *entities.DogBreed) (int, error) {
	args := breed.Called(b)
	return args.Int(0), args.Error(1)
}
func (breed *breedMock) FindAll() ([]entities.DogBreed, error) {
	args := breed.Called()
	return args.Get(0).([]entities.DogBreed), args.Error(1)
}
func (breed *breedMock) FindById(id string) (*entities.DogBreed, error) {
	args := breed.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}
func (breed *breedMock) Delete(id string) (*entities.DogBreed, error) {
	args := breed.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}
func (breed *breedMock) Update(b *entities.DogBreed) error {
	args := breed.Called(b)
	return args.Error(0)
}
func (breed *breedMock) CheckIfExists(id string) bool {
	args := breed.Called(id)
	return args.Bool(0)
}

//Dog Mock Methods

func (dog *dogMock) Save(d *entities.Dog, id interface{}) error {
	args := dog.Called(d, id)
	return args.Error(0)
}
func (dog *dogMock) FindAll() ([]entities.Dog, error) {
	args := dog.Called()
	return args.Get(0).([]entities.Dog), args.Error(1)
}
func (dog *dogMock) FindById(id string) (*entities.Dog, error) {
	args := dog.Called(id)
	return args.Get(0).(*entities.Dog), args.Error(1)
}
func (dog *dogMock) Delete(id string) (*entities.Dog, error) {
	args := dog.Called(id)
	return args.Get(0).(*entities.Dog), args.Error(1)
}
func (dog *dogMock) Update(d *entities.Dog, id string) error {
	args := dog.Called(d, id)
	return args.Error(0)
}
func (dog *dogMock) CheckIfExists(id string) bool {
	args := dog.Called(id)
	return args.Bool(0)
}

// Make Breed and Dog Functions

func MakeDog() (*entities.Dog, *entities.DogBreed) {
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(3, 4).
		SheddGroomAndEnergy(5, 6, 7)
	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(2).
		DogID(1).
		NameAndSex("M", "B").
		Breed(*breed)
	dog := d.BuildDog()
	return dog, breed
}

// Test Dog Functions

func TestSaveDog(t *testing.T) {
	dogMock := new(dogMock)     // mokando dog
	breedMock := new(breedMock) // mokando breed
	dog, breed := MakeDog()     // criando um dog

	dogMock.On("Save", dog, 1).Return(nil)
	//breedMock.On("Save", breed).Return(1, nil)

	// Eu desfiz a chamada do Save breed pois essa é uma funcionalidade que é acessivel apenas a usuarios admins

	testService := services.NewDogService(dogMock, breedMock)
	err := testService.CreateDog(dog, breed)
	//breedMock.AssertExpectations(t)
	dogMock.AssertExpectations(t)

	assert.Nil(t, err)

}

func TestFindAllDogs(t *testing.T) {
	dogMock := new(dogMock)
	breedMock := new(breedMock) // mokando breed
	dog, _ := MakeDog()

	dogMock.On("FindAll").Return([]entities.Dog{*dog}, nil)

	testService := services.NewDogService(dogMock, breedMock)
	dogs, err := testService.FindDogs()

	dogMock.AssertExpectations(t)

	assert.Equal(t, 1, dogs[0].DogID)
	assert.Equal(t, 2, dogs[0].KennelID)
	assert.Equal(t, "B", dogs[0].Sex)
	assert.Equal(t, "M", dogs[0].DogName)
	assert.Equal(t, dog.Breed, dogs[0].Breed)
	assert.Nil(t, err)
}

func TestFindDogById(t *testing.T) {

	dogMock := new(dogMock)
	breedMock := new(breedMock)
	dog, _ := MakeDog()
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("FindById", idStr).Return(dog, nil)

	testService := services.NewDogService(dogMock, breedMock)
	dog, err := testService.FindDogByID(idStr)

	dogMock.AssertExpectations(t)
	assert.Equal(t, 1, dog.DogID)
	assert.Equal(t, 2, dog.KennelID)
	assert.Equal(t, "B", dog.Sex)
	assert.Equal(t, "M", dog.DogName)
	assert.Equal(t, dog.Breed, dog.Breed)
	assert.Nil(t, err)
}

func TestDeleteDog(t *testing.T) {
	dogMock := new(dogMock)
	breedMock := new(breedMock)
	dog, _ := MakeDog()

	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("Delete", idStr).Return(dog, nil)

	testService := services.NewDogService(dogMock, breedMock)
	result, err := testService.DeleteDog(idStr)

	dogMock.AssertExpectations(t)

	assert.Equal(t, 1, result.DogID)
	assert.Equal(t, 2, result.KennelID)
	assert.Equal(t, "B", result.Sex)
	assert.Equal(t, "M", result.DogName)
	assert.Equal(t, result.Breed, result.Breed)
	assert.Nil(t, err)

}

func TestUpdateDog(t *testing.T) {
	dogMock := new(dogMock)
	breedMock := new(breedMock)

	dog, _ := MakeDog()
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("Update", dog, idStr).Return(nil)

	dog.BreedID = 7
	dog.DogName = "Z"

	testService := services.NewDogService(dogMock, breedMock)
	err := testService.UpdateDog(dog, idStr)

	dogMock.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestUpdateDogDontExists(t *testing.T) {
	dogMock := new(dogMock)
	breedMock := new(breedMock)

	dog, _ := MakeDog()
	//idStr := strconv.Itoa(dog.DogID)

	dogMock.On("Update", dog, "34").Return(errors.New("update dog failed"))

	dog.BreedID = 7
	dog.DogName = "Z"

	testService := services.NewDogService(dogMock, breedMock)
	err := testService.UpdateDog(dog, "34")

	dogMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Error(t, err, "update dog failed")

}

func TestCheckIfDontExistsDog(t *testing.T) {
	dogMock := new(dogMock)
	breedMock := new(breedMock)

	//dog, _ := MakeDog()
	//idStr := strconv.Itoa(dog.DogID)

	dogMock.On("CheckIfExists", "31").Return(false)

	testService := services.NewDogService(dogMock, breedMock)
	result := testService.CheckIfDogExist("31")

	dogMock.AssertExpectations(t)

	assert.Equal(t, false, result)

}

func TestCheckIfExistsDogs(t *testing.T) {
	dogMock := new(dogMock)
	breedMock := new(breedMock)

	dog, _ := MakeDog()
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("CheckIfExists", idStr).Return(true)

	testService := services.NewDogService(dogMock, breedMock)
	result := testService.CheckIfDogExist(idStr)

	dogMock.AssertExpectations(t)

	assert.Equal(t, true, result)
}
