package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"testing"

	"github.com/stretchr/testify/mock"
)

// mocar o servico de breed e o repositorio de breed

type breedRepoMock struct {
	mock.Mock
}

type breedServMock struct {
	mock.Mock
}

func (br *breedRepoMock) Save(d *entities.DogBreed) (int, error) {
	args := br.Called(d)
	return args.Int(0), args.Error(1)
}

func (br *breedRepoMock) FindById(id string) (*entities.DogBreed, error) {
	args := br.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}

func (br *breedRepoMock) FindAll() ([]entities.DogBreed, error) {
	args := br.Called()
	return args.Get(0).([]entities.DogBreed), args.Error(1)
}

func (br *breedRepoMock) Update(d *entities.DogBreed, id string) error {
	args := br.Called(d, id)
	return args.Error(0)
}

func (br *breedRepoMock) Delete(id string) (*entities.DogBreed, error) {
	args := br.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}

func (br *breedRepoMock) CheckIfExists(id string) bool {
	args := br.Called(id)
	return args.Bool(0)
}

// service implementation

func (bs *breedServMock) CreateBreed(d *entities.DogBreed) error {
	args := bs.Called(d)
	return args.Error(0)
}

func (bs *breedServMock) UpdateBreed(d *entities.DogBreed, id string) error {
	args := bs.Called(d, id)
	return args.Error(0)
}

func (bs *breedServMock) FindBreedByID(id string) (*entities.DogBreed, error) {
	args := bs.Called(id)
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}

func (bs *breedServMock) FindBreeds() ([]entities.DogBreed, error) {
	args := bs.Called()
	return args.Get(0).([]entities.DogBreed), args.Error(1)
}

func TestCreateBreed(t *testing.T) {

	breedServMock := new(breedServMock)
	breedRepoMock := new(breedRepoMock)

	dogBreed := entities.BuildDogBreed("1", "x", 1, 1, 1, 1, 1, 1, 1)

	breedServMock.On("CreateBreed", dogBreed).Return(nil)
	breedRepoMock.On("Save", dogBreed).Return(dogBreed.ID, nil)

	jsonBody, err := json.Marshal(dogBreed)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling json body")
	}
	requestURL := "/breed/create/"
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error on post request creation")
	}

	testService := services.NewBreedService(breedRepoMock)
	testController := NewBreedController(breedServMock)

}
