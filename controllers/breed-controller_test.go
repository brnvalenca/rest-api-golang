package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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
	args := bs.Called(d)
	return args.Error(0)
}

func (bs *breedServMock) FindBreedByID(id string) (*entities.DogBreed, error) {
	args := bs.Called()
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
	testService.CreateBreed(dogBreed)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(testController.Create)
	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting 200 code got: %v", status)
	}

	var respBody entities.DogBreed
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf(err.Error(), "error on body response decoding")
	}

	breedServMock.AssertExpectations(t)
	breedRepoMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, respBody.Grooming, dogBreed.ID)
	assert.Equal(t, respBody.BreedImg, dogBreed.BreedImg)
	assert.Equal(t, respBody.Energy, dogBreed.Energy)
	assert.Equal(t, respBody.GoodWithDogs, dogBreed.GoodWithDogs)
	assert.Equal(t, respBody.GoodWithKids, dogBreed.GoodWithKids)
	assert.Equal(t, respBody.Grooming, dogBreed.Grooming)
	assert.Equal(t, respBody.Shedding, dogBreed.Shedding)

}

func TestGetAllBreeds(t *testing.T) {

	breedServMock := new(breedServMock)
	breedRepoMock := new(breedRepoMock)

	dogBreed := entities.BuildDogBreed("1", "x", 1, 1, 1, 1, 1, 1, 1)

	breedServMock.On("FindBreeds").Return([]entities.DogBreed{*dogBreed}, nil)
	breedRepoMock.On("FindAll").Return([]entities.DogBreed{*dogBreed}, nil)

	requestURL := "/breeds/"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Errorf(err.Error(), "error during get request creation")
	}
	resp := httptest.NewRecorder()

	testService := services.NewBreedService(breedRepoMock)
	testController := NewBreedController(breedServMock)
	handler := http.HandlerFunc(testController.GetAll)

	handler.ServeHTTP(resp, req)

	testService.FindBreeds()

	var breeds []entities.DogBreed
	err = json.NewDecoder(resp.Body).Decode(&breeds)
	if err != nil {
		t.Errorf(err.Error(), "error during resp body decoding")
	}

	breedRepoMock.AssertExpectations(t)
	breedServMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)
}

func TestGetBreedById(t *testing.T) {

	breedServMock := new(breedServMock)
	breedRepoMock := new(breedRepoMock)

	dogBreed := entities.BuildDogBreed("1", "x", 1, 1, 1, 1, 1, 1, 1)

	idStr := strconv.Itoa(dogBreed.ID)

	breedServMock.On("FindBreedByID").Return(dogBreed, nil)
	breedRepoMock.On("FindById", idStr).Return(dogBreed, nil)

	requestURL := "/breed/{id}/"
	jsonBody, err := json.Marshal(dogBreed.ID)
	if err != nil {
		t.Errorf(err.Error())
	}
	req, err := http.NewRequest("GET", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error during get request creation")
	}
	resp := httptest.NewRecorder()

	testService := services.NewBreedService(breedRepoMock)
	testController := NewBreedController(breedServMock)
	handler := http.HandlerFunc(testController.GetById)

	handler.ServeHTTP(resp, req)

	testService.FindBreedByID(idStr)

	var respBody entities.DogBreed
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf(err.Error(), "error during resp body decoding")
	}

	breedRepoMock.AssertExpectations(t)
	breedServMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, respBody.Grooming, dogBreed.ID)
	assert.Equal(t, respBody.BreedImg, dogBreed.BreedImg)
	assert.Equal(t, respBody.Energy, dogBreed.Energy)
	assert.Equal(t, respBody.GoodWithDogs, dogBreed.GoodWithDogs)
	assert.Equal(t, respBody.GoodWithKids, dogBreed.GoodWithKids)
	assert.Equal(t, respBody.Grooming, dogBreed.Grooming)
	assert.Equal(t, respBody.Shedding, dogBreed.Shedding)
}

func TestDeleteBreed(t *testing.T) {

	breedServMock := new(breedServMock)

	jsonBody, _ := json.Marshal("10")
	requestURL := "/breed/delete/{id}/"
	req, err := http.NewRequest("DELETE", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error())
	}
	resp := httptest.NewRecorder()

	testController := NewBreedController(breedServMock)
	handler := http.HandlerFunc(testController.Delete)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 405)
}

func TestUpdateBreed(t *testing.T) {

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
	testService.CreateBreed(dogBreed)

	handler := http.HandlerFunc(testController.Create)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Error(err.Error(), "got wrong status code :%v", resp.Code)
	}

	dogBreed = entities.BuildDogBreed("1", "a", 2, 2, 2, 2, 2, 2, 2)
	idStr := strconv.Itoa(dogBreed.ID)
	breedServMock.On("UpdateBreed", dogBreed).Return(nil)
	breedRepoMock.On("Update", dogBreed, idStr).Return(nil)

	jsonBody, err = json.Marshal(dogBreed)

	if err != nil {
		t.Errorf(err.Error(), "error marshalling json body")
	}
	requestURL = "/breed/update/" + strconv.Itoa(dogBreed.ID)
	req, err = http.NewRequest("UPDATE", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error on update request creation")
	}

	resp = httptest.NewRecorder()
	handler = http.HandlerFunc(testController.Update)

	handler.ServeHTTP(resp, req)
	testService.UpdateBreed(dogBreed, idStr)

	if resp.Code != 200 {
		t.Error(err.Error(), "got wrong status code :%v", resp.Code)
	}

	fmt.Println(resp.Body)
	breedRepoMock.AssertExpectations(t)
	breedServMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)

}
