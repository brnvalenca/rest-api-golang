package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
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
func (breed *breedMock) Update(b *entities.DogBreed, id string) error {
	args := breed.Called(b, id)
	return args.Error(0)
}
func (breed *breedMock) CheckIfExists(id string) bool {
	args := breed.Called(id)
	return args.Bool(0)
}

//Dog Mocks

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
	breed := entities.BuildDogBreed("1", 1, 2, 3, 4, 5, 6, 7)
	dogs := entities.BuildDog(*breed, 1, 2, "M", "B")
	return dogs, breed
}

// Test Dog Functions

func TestSaveDog(t *testing.T) {
	dogMock := new(dogMock)
	breedMock := new(breedMock)
	dog, breed := MakeDog()

	dogMock.On("Save", dog, 1).Return(nil)
	breedMock.On("Save", breed).Return("1", nil)

	// Ele reclama dizendo que o breed precisa fazer mais outra chamada, desconfio que é pq o breed tem um service, ver isso dps, to cansadaaao

	testService := services.NewDogService(dogMock, breedMock)
	err := testService.CreateDog(dog, breed)

	dogMock.AssertExpectations(t)
	breedMock.AssertExpectations(t)

	assert.Nil(t, err)

}
