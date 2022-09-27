package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDogRepo struct {
	mock.Mock
}

func (dr *MockDogRepo) Save(d *entities.Dog, id interface{}) error {
	args := dr.Called(d, id)
	return args.Error(0)
}

func (dr *MockDogRepo) FindAll() ([]entities.Dog, error) {
	args := dr.Called()
	return args.Get(0).([]entities.Dog), args.Error(1)
}

func (dr *MockDogRepo) FindById(id string) (*entities.Dog, error) {
	args := dr.Called(id)
	return args.Get(0).(*entities.Dog), args.Error(1)
}

func (dr *MockDogRepo) Delete(id string) (*entities.Dog, error) {
	args := dr.Called(id)
	return args.Get(0).(*entities.Dog), args.Error(1)
}

func (dr *MockDogRepo) Update(d *entities.Dog, id string) error {
	args := dr.Called(d, id)
	return args.Error(0)
}

func (dr *MockDogRepo) CheckIfExists(id string) bool {
	args := dr.Called(id)
	return args.Bool(0)
}

func TestSaveDog(t *testing.T) {
	dogMock := new(MockDogRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "x", "x")

	dogMock.On("Save", dog, breed.ID).Return(nil)

	err := dogMock.Save(dog, breed.ID)
	assert.Nil(t, err)
}

func TestFindAllDogs(t *testing.T) {
	dogMock := new(MockDogRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "x", "x")

	dogMock.On("FindAll").Return([]entities.Dog{*dog}, nil)

	arrDog, err := dogMock.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, arrDog[0].Breed, dog.Breed)
	assert.Equal(t, arrDog[0].BreedID, dog.BreedID)
	assert.Equal(t, arrDog[0].DogID, dog.DogID)
	assert.Equal(t, arrDog[0].DogName, dog.DogName)
	assert.Equal(t, arrDog[0].KennelID, dog.KennelID)
	assert.Equal(t, arrDog[0].Sex, dog.Sex)
	assert.Equal(t, arrDog[0].Breed.GoodWithDogs, dog.Breed.GoodWithDogs)
	assert.Equal(t, arrDog[0].Breed.GoodWithKids, dog.Breed.GoodWithKids)
	assert.Equal(t, arrDog[0].Breed.BreedImg, dog.Breed.BreedImg)
	assert.Equal(t, arrDog[0].Breed.Energy, dog.Breed.Energy)
	assert.Equal(t, arrDog[0].Breed.Grooming, dog.Breed.Grooming)
	assert.Equal(t, arrDog[0].Breed.Shedding, dog.Breed.Shedding)
	assert.Equal(t, arrDog[0].Breed.ID, dog.Breed.ID)
	assert.Equal(t, arrDog[0].Breed.Name, dog.Breed.Name)

}

func TestFindById(t *testing.T) {
	dogMock := new(MockDogRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "x", "x")
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("Save", dog, breed.ID).Return(nil)
	dogMock.On("FindById", idStr).Return(dog, nil)

	_ = dogMock.Save(dog, breed.ID)

	dogReturn, err := dogMock.FindById(idStr)

	assert.Nil(t, err)
	assert.Equal(t, dogReturn, dog)
}

func TestDeleteDog(t *testing.T) {
	dogMock := new(MockDogRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "x", "x")
	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("Save", dog, breed.ID).Return(nil)
	dogMock.On("Delete", idStr).Return(dog, nil)

	_ = dogMock.Save(dog, breed.ID)
	dogReturn, err := dogMock.Delete(idStr)

	assert.Nil(t, err)
	assert.Equal(t, dogReturn, dog)
}

func TestUpdateDog(t *testing.T) {
	dogMock := new(MockDogRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "x", "x")

	idStr := strconv.Itoa(dog.DogID)
	newDog := entities.BuildDog(*breed, 1, 1, "m", "b")

	dogMock.On("Save", dog, breed.ID).Return(nil)
	dogMock.On("Update", newDog, idStr).Return(nil)

	_ = dogMock.Save(dog, breed.ID)
	err := dogMock.Update(newDog, idStr)

	assert.Nil(t, err)

}

func TestCheckIfDogExists(t *testing.T) {
	dogMock := new(MockDogRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "x", "x")

	idStr := strconv.Itoa(dog.DogID)

	dogMock.On("CheckIfExists", idStr).Return(true)

	got := dogMock.CheckIfExists(idStr)

	assert.Equal(t, true, got)
}
