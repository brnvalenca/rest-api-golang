package repository

import (
	"rest-api/golang/exercise/domain/entities"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBreedRepo struct {
	mock.Mock
}

func (br *MockBreedRepo) Save(b *entities.DogBreed) (int, error) {
	args := br.Called(b)
	return args.Int(0), args.Error(1)

}

func (br *MockBreedRepo) FindById(id string) (*entities.DogBreed, error) {
	args := br.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}

func (br *MockBreedRepo) Delete(id string) (*entities.DogBreed, error) {
	args := br.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}

func (br *MockBreedRepo) Update(b *entities.DogBreed, id string) error {
	args := br.Called(b, id)
	return args.Error(0)
}

func (br *MockBreedRepo) CheckIfExists(id string) bool {
	args := br.Called(id)
	return args.Bool(0)
}

func TestSaveBreed(t *testing.T) {
	breedMock := new(MockBreedRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)

	breedMock.On("Save", breed).Return(breed.ID, nil)

	id, err := breedMock.Save(breed)

	assert.Nil(t, err)
	assert.Equal(t, id, breed.ID)
}

func TestFindByIdBreed(t *testing.T) {
	breedMock := new(MockBreedRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	idStr := strconv.Itoa(breed.ID)
	breedMock.On("FindById", idStr).Return(breed, nil)

	breedReturn, err := breedMock.FindById(idStr)
	assert.Nil(t, err)
	assert.Equal(t, breedReturn.ID, breed.ID)
}

func TestDeleteBreed(t *testing.T) {
	breedMock := new(MockBreedRepo)

	breed := entities.BuildDogBreed("1", "kxzs", 1, 1, 1, 1, 1, 1, 1)
	idStr := strconv.Itoa(breed.ID)
	breedMock.On("Delete", idStr).Return(breed, nil)

	breedReturn, err := breedMock.Delete(idStr)
	assert.Nil(t, err)
	assert.Equal(t, breedReturn.ID, breed.ID)
}

func TestUpdateBreed(t *testing.T) {
	breedMock := new(MockBreedRepo)

	breed := entities.BuildDogBreed("1", "kxzs", 1, 1, 1, 1, 1, 1, 1)
	newBreed := entities.BuildDogBreed("1", "abcd", 1, 1, 1, 1, 1, 1, 1)
	idStr := strconv.Itoa(breed.ID)
	breedMock.On("Save", breed).Return(breed.ID, nil)
	breedMock.On("Update", newBreed, idStr).Return(nil)

	breedMock.Save(breed)

	err := breedMock.Update(newBreed, idStr)

	assert.Nil(t, err)
}

func TestCheckIfExistsBreed(t *testing.T) {
	breedMock := new(MockBreedRepo)

	breed := entities.BuildDogBreed("1", "1", 1, 1, 1, 1, 1, 1, 1)
	idStr := strconv.Itoa(breed.ID)
	breedMock.On("CheckIfExists", idStr).Return(true)

	breedReturn := breedMock.CheckIfExists(idStr)
	assert.Equal(t, true, breedReturn)
}
