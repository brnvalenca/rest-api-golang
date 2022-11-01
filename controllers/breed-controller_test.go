package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
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
	args := br.Called()
	return args.Get(0).(*entities.DogBreed), args.Error(1)
}

func (br *breedRepoMock) FindAll() ([]entities.DogBreed, error) {
	args := br.Called()
	return args.Get(0).([]entities.DogBreed), args.Error(1)
}

func (br *breedRepoMock) Update(d *entities.DogBreed) error {
	args := br.Called(d)
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

func (bs *breedServMock) CreateBreed(d *dtos.BreedDTO) error {
	args := bs.Called(d)
	return args.Error(0)
}

func (bs *breedServMock) UpdateBreed(d *dtos.BreedDTO) error {
	args := bs.Called(d)
	return args.Error(0)
}

func (bs *breedServMock) FindBreedByID(id string) (*dtos.BreedDTO, error) {
	args := bs.Called()
	return args.Get(0).(*dtos.BreedDTO), args.Error(1)
}

func (bs *breedServMock) FindBreeds() ([]dtos.BreedDTO, error) {
	args := bs.Called()
	return args.Get(0).([]dtos.BreedDTO), args.Error(1)
}

func (bs *breedServMock) ValidateBreed(d *dtos.BreedDTO) error {
	args := bs.Called(d)
	return args.Error(0)
}

func TestCreateBreed(t *testing.T) {

	breedServMock := new(breedServMock)
	breedRepoMock := new(breedRepoMock)

	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(0).
		Name("Yorkshire").
		Img("imgurl").
		GoodWithKidsAndDogs(1, 2).
		SheddGroomAndEnergy(1, 2, 3)

	dogBreed := db.BuildBreed()

	bdto := dtos.NewBreedBuilderDTO()
	bdto.Has().
		ID(dogBreed.ID).
		Name(dogBreed.Name).
		Img(dogBreed.BreedImg).
		GoodWithKidsAndDogs(dogBreed.GoodWithKids, dogBreed.GoodWithDogs).
		SheddGroomAndEnergy(dogBreed.Shedding, dogBreed.Grooming, dogBreed.Energy)

	breedDTO := bdto.BuildBreedDTO()

	breedServMock.On("ValidateBreed", breedDTO).Return(nil)
	breedServMock.On("CreateBreed", breedDTO).Return(nil)
	breedRepoMock.On("Save", dogBreed).Return(dogBreed.ID, nil)

	jsonBody, err := json.Marshal(breedDTO)
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

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(testController.Create)
	testService.ValidateBreed(breedDTO)
	testService.CreateBreed(breedDTO)
	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting 200 code got: %v", status)
	}

	var respBody dtos.BreedDTO
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf(err.Error(), "error on body response decoding")
	}

	breedServMock.AssertExpectations(t)
	breedRepoMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, respBody.ID, dogBreed.ID)
	assert.Equal(t, respBody.BreedImg, dogBreed.BreedImg)
	assert.Equal(t, respBody.Energy, dogBreed.Energy)
	assert.Equal(t, respBody.GoodWithDogs, dogBreed.GoodWithDogs)
	assert.Equal(t, respBody.GoodWithKids, dogBreed.GoodWithKids)
	assert.Equal(t, respBody.Grooming, dogBreed.Grooming)
	assert.Equal(t, respBody.Shedding, dogBreed.Shedding)

}

/*
	Invalid memory address or nik pointer dereference
*/
func TestGetAllBreeds(t *testing.T) {

	breedRepoMock := new(breedRepoMock)
	servico := services.NewBreedService(breedRepoMock)

	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("Yorkshire").
		Img("imgurl").
		GoodWithKidsAndDogs(1, 2).
		SheddGroomAndEnergy(1, 2, 3)

	dogBreed := db.BuildBreed()
	//breedDTO := dtos.BreedDTO{Breed: *dogBreed}

	breedRepoMock.On("FindAll").Return([]entities.DogBreed{*dogBreed}, nil)

	requestURL := "/breeds/"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Errorf(err.Error(), "error during get request creation")
	}
	resp := httptest.NewRecorder()

	testController := NewBreedController(servico)
	handler := http.HandlerFunc(testController.GetAll)

	handler.ServeHTTP(resp, req)

	var breeds []dtos.BreedDTO
	err = json.NewDecoder(resp.Body).Decode(&breeds)
	if err != nil {
		t.Errorf(err.Error(), "error during resp body decoding")
	}

	breedRepoMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)
}

func TestGetBreedById(t *testing.T) {

	breedRepoMock := new(breedRepoMock)
	breedService := services.NewBreedService(breedRepoMock)

	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("Yorkshire").
		Img("imgurl").
		GoodWithKidsAndDogs(1, 2).
		SheddGroomAndEnergy(1, 2, 3)

	dogBreed := db.BuildBreed()

	breedRepoMock.On("FindById").Return(dogBreed, nil)

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

	testController := NewBreedController(breedService)
	handler := http.HandlerFunc(testController.GetById)

	handler.ServeHTTP(resp, req)

	var respBody dtos.BreedDTO
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf(err.Error(), "error during resp body decoding")
	}

	breedRepoMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, respBody.ID, dogBreed.ID)
	assert.Equal(t, respBody.BreedImg, dogBreed.BreedImg)
	assert.Equal(t, respBody.Energy, dogBreed.Energy)
	assert.Equal(t, respBody.GoodWithDogs, dogBreed.GoodWithDogs)
	assert.Equal(t, respBody.GoodWithKids, dogBreed.GoodWithKids)
	assert.Equal(t, respBody.Grooming, dogBreed.Grooming)
	assert.Equal(t, respBody.Shedding, dogBreed.Shedding)
}

func TestDeleteBreed(t *testing.T) {
	breedRepo := new(breedRepoMock)
	breedService := services.NewBreedService(breedRepo)

	jsonBody, _ := json.Marshal("10")
	requestURL := "/breed/delete/{id}/"
	req, err := http.NewRequest("DELETE", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error())
	}
	resp := httptest.NewRecorder()

	testController := NewBreedController(breedService)
	handler := http.HandlerFunc(testController.Delete)
	handler.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 405)
}

func TestUpdateBreed(t *testing.T) {

	breedRepoMock := new(breedRepoMock)
	breedService := services.NewBreedService(breedRepoMock)
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("Yorkshire").
		Img("imgurl").
		GoodWithKidsAndDogs(1, 2).
		SheddGroomAndEnergy(1, 2, 3)

	dogBreed := db.BuildBreed()

	bdto := dtos.NewBreedBuilderDTO()
	bdto.Has().
		ID(dogBreed.ID).
		Name(dogBreed.Name).
		Img(dogBreed.BreedImg).
		GoodWithKidsAndDogs(dogBreed.GoodWithKids, dogBreed.GoodWithDogs).
		SheddGroomAndEnergy(dogBreed.Shedding, dogBreed.Grooming, dogBreed.Energy)

	breedDTO := bdto.BuildBreedDTO()

	breedRepoMock.On("Save", dogBreed).Return(dogBreed.ID, nil)

	jsonBody, err := json.Marshal(breedDTO)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling json body")
	}
	requestURL := "/breed/create/"
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error on post request creation")
	}

	testController := NewBreedController(breedService)

	handler := http.HandlerFunc(testController.Create)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Error(err.Error(), "got wrong status code :%v", resp.Code)
	}

	db = entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("Maltes").
		Img("imgurl").
		GoodWithKidsAndDogs(1, 2).
		SheddGroomAndEnergy(1, 2, 3)

	dogBreed = db.BuildBreed()

	bdto = dtos.NewBreedBuilderDTO()
	bdto.Has().
		ID(dogBreed.ID).
		Name(dogBreed.Name).
		Img(dogBreed.BreedImg).
		GoodWithKidsAndDogs(dogBreed.GoodWithKids, dogBreed.GoodWithDogs).
		SheddGroomAndEnergy(dogBreed.Shedding, dogBreed.Grooming, dogBreed.Energy)

	breedDTO = bdto.BuildBreedDTO()

	breedRepoMock.On("Update", dogBreed).Return(nil)

	jsonBody, err = json.Marshal(breedDTO)

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

	if resp.Code != 200 {
		t.Error(err.Error(), "got wrong status code :%v", resp.Code)
	}

	fmt.Println(resp.Body)
	breedRepoMock.AssertExpectations(t)

	assert.Equal(t, resp.Code, 200)

}
