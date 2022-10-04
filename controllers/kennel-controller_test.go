package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAddressRepo struct {
	mock.Mock
}

type MockKennelRepo struct {
	mock.Mock
}

type MockKennelServ struct {
	mock.Mock
}

// Address Repository Mock

func (addr *MockAddressRepo) SaveAddress(address *entities.Address) error {
	args := addr.Called(address)
	return args.Error(0)
}

// Kennel Repository Mock

func (kr *MockKennelRepo) FindAllRepo() ([]entities.Kennel, error) {
	args := kr.Called()
	return args.Get(0).([]entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) SaveRepo(k *entities.Kennel) (int, error) {
	args := kr.Called(k)
	return args.Int(0), args.Error(1)
}

func (kr *MockKennelRepo) FindByIdRepo(id string) (*entities.Kennel, error) {
	args := kr.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) DeleteRepo(id string) (*entities.Kennel, error) {
	args := kr.Called()
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) UpdateRepo(u *entities.Kennel, id string) error {
	args := kr.Called(u, id)
	return args.Error(0)
}

func (kr *MockKennelRepo) CheckIfExistsRepo(id string) bool {
	args := kr.Called(id)
	return args.Bool(0)
}

// Kennel Service Mock

func (ks *MockKennelServ) FindAllKennels() ([]entities.Kennel, error) {
	args := ks.Called()
	return args.Get(0).([]entities.Kennel), args.Error(1)
}

func (ks *MockKennelServ) Save(u *entities.Kennel) (int, error) {
	args := ks.Called(u)
	return args.Int(0), args.Error(1)
}

func (ks *MockKennelServ) FindKennelByIdServ(id string) (*entities.Kennel, error) {
	args := ks.Called()
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (ks *MockKennelServ) DeleteKennelServ(id string) (*entities.Kennel, error) {
	args := ks.Called()
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (ks *MockKennelServ) UpdateKennelServ(u *entities.Kennel, id string) error {
	args := ks.Called(u)
	return args.Error(0)
}

func (ks *MockKennelServ) CheckIfExists(id string) bool {
	args := ks.Called()
	return args.Bool(0)
}

const (
	serverPort int = 8080
)

func TestGetAllKennels(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(1, 1).
		SheddGroomAndEnergy(1, 1, 1)

	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(1).
		BreedID(1).
		DogID(1).
		NameAndSex("m", "b").
		Breed(*breed)
	dog := d.BuildDog()
	// Create an array of dogs with just one dog
	var dogs []entities.Dog
	dogs = append(dogs, *dog)

	// Create a new Kennel

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("432").
		Rua("um").
		Bairro("dois").
		CEP("tres").
		Cidade("quatro")

	addr := ad.BuildAddr()
	kennel := entities.BuildKennel(1, dogs, *addr, "1", "x")

	// Fazer uma composicao da função para criar um Kennel

	kennelMockServ.On("FindAllKennels").Return([]entities.Kennel{*kennel}, nil)
	kennelMockRepo.On("FindAllRepo").Return([]entities.Kennel{*kennel}, nil)

	// Create a new GET Request

	requestURL := fmt.Sprintf("http://localhost:%d/kennels/", serverPort)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	// Create service and controller instance
	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)
	testService.FindAllKennels()

	handler := http.HandlerFunc(testController.GetAll)

	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)
	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	kennelMockServ.AssertExpectations(t)
	kennelMockRepo.AssertExpectations(t)

	var serverResp []entities.Kennel
	err = json.NewDecoder(io.Reader(resp.Body)).Decode(&serverResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, kennel.ID, serverResp[0].ID)
	assert.Equal(t, kennel.Name, serverResp[0].Name)
	assert.Equal(t, kennel.ContactNumber, serverResp[0].ContactNumber)
	assert.Equal(t, kennel.Address.ID_Kennel, serverResp[0].Address.ID_Kennel)
	assert.Equal(t, kennel.Address.Numero, serverResp[0].Address.Numero)
	assert.Equal(t, kennel.Address.Rua, serverResp[0].Address.Rua)
	assert.Equal(t, kennel.Address.Bairro, serverResp[0].Address.Bairro)
	assert.Equal(t, kennel.Address.Cidade, serverResp[0].Address.Cidade)
	assert.Equal(t, kennel.Address.CEP, serverResp[0].Address.CEP)
	assert.Equal(t, kennel.Dogs[0], serverResp[0].Dogs[0])
}

func TestGetKennelById(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	// Create an array of dogs with just one dog
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(1, 1).
		SheddGroomAndEnergy(1, 1, 1)

	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(1).
		BreedID(1).
		DogID(1).
		NameAndSex("m", "b").
		Breed(*breed)
	dog := d.BuildDog()

	var dogs []entities.Dog
	dogs = append(dogs, *dog)

	// Create a new Kennel

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("432").
		Rua("um").
		Bairro("dois").
		CEP("tres").
		Cidade("quatro")

	addr := ad.BuildAddr()

	kennel := entities.BuildKennel(1, dogs, *addr, "1", "x")
	idStr := strconv.Itoa(kennel.ID)

	kennelMockServ.On("FindKennelByIdServ").Return(kennel, nil)
	kennelMockRepo.On("FindByIdRepo", idStr).Return(kennel, nil)

	// Create a new GET Request

	requestURL := fmt.Sprintf("http://localhost:%d/kennels/%s/", serverPort, idStr)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)
	testService.FindKennelByIdServ(idStr)

	handler := http.HandlerFunc(testController.GetById)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)
	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	kennelMockServ.AssertExpectations(t)
	kennelMockRepo.AssertExpectations(t)

	var serverResp entities.Kennel
	err = json.NewDecoder(io.Reader(resp.Body)).Decode(&serverResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
}

func TestGetKennelByIdIfDoesntExist(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	// Create an array of dogs with just one dog
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(1, 1).
		SheddGroomAndEnergy(1, 1, 1)

	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(1).
		BreedID(1).
		DogID(1).
		NameAndSex("m", "b").
		Breed(*breed)
	dog := d.BuildDog()

	var dogs []entities.Dog
	dogs = append(dogs, *dog)
	// Create a new Kennel

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("432").
		Rua("um").
		Bairro("dois").
		CEP("tres").
		Cidade("quatro")

	addr := ad.BuildAddr()

	kennel := entities.BuildKennel(1, dogs, *addr, "1", "x")
	idStr := "25"

	errReturned := errors.New("kennel by ID 25: no such kennel")
	kennelMockServ.On("FindKennelByIdServ").Return(kennel, errReturned)
	kennelMockRepo.On("FindByIdRepo", idStr).Return(kennel, errReturned)

	// Create a new GET Request

	requestURL := fmt.Sprintf("http://localhost:%d/kennels/%s/", serverPort, idStr)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)
	_, errService := testService.FindKennelByIdServ(idStr)
	fmt.Println(errService)
	handler := http.HandlerFunc(testController.GetById)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusNotFound {
		t.Errorf(err.Error(), "expecting status code of 404, got: %v", status)
	}

	kennelMockServ.AssertExpectations(t)
	kennelMockRepo.AssertExpectations(t)

	b, err := io.ReadAll(resp.Body)
	fmt.Println(b)
	if err != nil {
		t.Errorf(err.Error())
	}
	serverResponse := errors.New(string(b))

	assert.EqualError(t, serverResponse, "404 Not Found")
	assert.EqualError(t, errService, "kennel by ID 25: no such kennel")
}

func TestCreateKennel(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	// Create an array of dogs with just one dog
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(1, 1).
		SheddGroomAndEnergy(1, 1, 1)

	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(1).
		BreedID(1).
		DogID(1).
		NameAndSex("m", "b").
		Breed(*breed)
	dog := d.BuildDog()

	var dogs []entities.Dog
	dogs = append(dogs, *dog)

	// Create a new Kennel

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("432").
		Rua("um").
		Bairro("dois").
		CEP("tres").
		Cidade("quatro")

	addr := ad.BuildAddr()

	kennel := entities.BuildKennel(1, dogs, *addr, "1", "x")

	kennelMockServ.On("Save", kennel).Return(kennel.ID, nil)
	kennelMockRepo.On("SaveRepo", kennel).Return(kennel.ID, nil)
	addressMockRepo.On("SaveAddress", &kennel.Address).Return(nil)

	// Create a new POST Request

	jsonBody, err := json.Marshal(kennel)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling json body")
	}

	requestURL := fmt.Sprintf("http://localhost:%d/kennels/create/", serverPort)

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error creating post request")
	}

	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)
	kennelID, err := testService.Save(kennel)

	if err != nil {
		t.Errorf(err.Error())
	}

	handler := http.HandlerFunc(testController.Create)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got %v", status)
	}

	var respBody int
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf(err.Error())
	}

	addressMockRepo.AssertExpectations(t)
	kennelMockRepo.AssertExpectations(t)
	kennelMockServ.AssertExpectations(t)

	assert.NotNil(t, kennelID)
	assert.Equal(t, 200, status)
	assert.Equal(t, kennel.ID, kennelID)

}

func TestDeleteKennel(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	// Create an array of dogs with just one dog
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(1, 1).
		SheddGroomAndEnergy(1, 1, 1)

	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(1).
		BreedID(1).
		DogID(1).
		NameAndSex("m", "b").
		Breed(*breed)
	dog := d.BuildDog()

	var dogs []entities.Dog
	dogs = append(dogs, *dog)

	// Create a new Kennel

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("432").
		Rua("um").
		Bairro("dois").
		CEP("tres").
		Cidade("quatro")

	addr := ad.BuildAddr()
	kennel := entities.BuildKennel(1, dogs, *addr, "1", "x")

	kennelMockServ.On("Save", kennel).Return(kennel.ID, nil)
	kennelMockRepo.On("SaveRepo", kennel).Return(kennel.ID, nil)
	addressMockRepo.On("SaveAddress", &kennel.Address).Return(nil)

	// Create a new POST Request

	jsonBody, err := json.Marshal(kennel)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling json body")
	}

	requestURL := fmt.Sprintf("http://localhost:%d/kennels/create/", serverPort)

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error creating post request")
	}
	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)

	kennelID, err := testService.Save(kennel)
	idStr := strconv.Itoa(kennelID)
	if err != nil {
		t.Errorf(err.Error())
	}

	handler := http.HandlerFunc(testController.Create)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got %v", status)
	}

	// Start the DELETE part:

	kennelMockServ.On("CheckIfExists").Return(true)
	kennelMockRepo.On("CheckIfExistsRepo", idStr).Return(true)
	kennelMockServ.On("DeleteKennelServ").Return(kennel, nil)
	kennelMockRepo.On("DeleteRepo").Return(kennel, nil)

	requestURL = fmt.Sprintf("http://localhost:%d/kennels/delete/%d/", serverPort, kennelID)

	req, err = http.NewRequest("DELETE", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error creating delete request")
	}

	testService.CheckIfExists(idStr)
	testService.DeleteKennelServ(idStr)
	handler = http.HandlerFunc(testController.Delete)
	resp = httptest.NewRecorder()

	handler.ServeHTTP(resp, req)
	status = resp.Code
	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got %v", status)
	}

	var respBody entities.Kennel
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf(err.Error())
	}

	addressMockRepo.AssertExpectations(t)
	kennelMockRepo.AssertExpectations(t)
	kennelMockServ.AssertExpectations(t)

}

func TestDeleteKennelIfDontExists(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	idStr := "2"
	// Start the DELETE part:

	kennelMockServ.On("CheckIfExists").Return(false)
	kennelMockRepo.On("CheckIfExistsRepo", idStr).Return(false)

	jsonData, err := json.Marshal(idStr)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err := http.NewRequest("DELETE", "/kennels/delete/{id}/", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating delete request")
	}

	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)
	handler := http.HandlerFunc(testController.Delete)

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	testService.CheckIfExists(idStr)

	status := resp.Code
	if status != http.StatusNotFound {
		t.Errorf(err.Error(), "expecting status code of 404, got: %v", status)
	}

	kennelMockServ.AssertExpectations(t)
	kennelMockRepo.AssertExpectations(t)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	serverResponse := errors.New(string(responseBody))

	assert.EqualError(t, serverResponse, "404 Not Found")
}

func TestUpdateKennel(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	// Create an array of dogs with just one dog
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(1, 1).
		SheddGroomAndEnergy(1, 1, 1)

	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(1).
		BreedID(1).
		DogID(1).
		NameAndSex("m", "b").
		Breed(*breed)
	dog := d.BuildDog()

	var dogs []entities.Dog
	dogs = append(dogs, *dog)

	// Create a new Kennel

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("432").
		Rua("um").
		Bairro("dois").
		CEP("tres").
		Cidade("quatro")

	addr := ad.BuildAddr()
	kennel := entities.BuildKennel(1, dogs, *addr, "1", "x")

	kennelMockServ.On("Save", kennel).Return(kennel.ID, nil)
	kennelMockRepo.On("SaveRepo", kennel).Return(kennel.ID, nil)
	addressMockRepo.On("SaveAddress", &kennel.Address).Return(nil)

	// Create a new POST Request
	jsonBody, err := json.Marshal(kennel)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling json body")
	}

	requestURL := fmt.Sprintf("http://localhost:%d/kennels/create/", serverPort)

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error(), "error creating post request")
	}

	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)

	kennelID, err := testService.Save(kennel)
	if err != nil {
		t.Errorf(err.Error())
	}

	handler := http.HandlerFunc(testController.Create)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	// Ending Kennel Creation
	/*------------------------*/
	// Update A KENNEL

	newKennel := entities.BuildKennel(1, dogs, *addr, "1", "new")
	idStr := strconv.Itoa(kennelID)

	kennelMockServ.On("CheckIfExists").Return(true)
	kennelMockRepo.On("CheckIfExistsRepo", idStr).Return(true)
	kennelMockServ.On("UpdateKennelServ", newKennel).Return(nil)
	kennelMockRepo.On("UpdateRepo", newKennel, idStr).Return(nil)

	// Update Request

	urlString := "/kennels/update/{id}/" + idStr

	requestBody, err := json.Marshal(newKennel)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ = http.NewRequest("UPDATE", urlString, bytes.NewBuffer(requestBody))

	handler = http.HandlerFunc(testController.Update)
	resp = httptest.NewRecorder()

	testService.CheckIfExists(idStr)
	testService.UpdateKennelServ(newKennel, idStr)
	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	kennelMockRepo.AssertExpectations(t)
	kennelMockServ.AssertExpectations(t)
	addressMockRepo.AssertExpectations(t)

	var kennelResp entities.Kennel
	err = json.NewDecoder(resp.Body).Decode(&kennelResp)
	if err != nil {
		t.Errorf(err.Error(), "error during response body decoding")
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, kennelID, kennelResp.ID)
	assert.Equal(t, newKennel.Name, kennelResp.Name)
	assert.Equal(t, newKennel.ContactNumber, kennelResp.ContactNumber)
	assert.Equal(t, newKennel.Address.Bairro, kennelResp.Address.Bairro)
	assert.Equal(t, newKennel.Address.CEP, kennelResp.Address.CEP)
	assert.Equal(t, newKennel.Address.Cidade, kennelResp.Address.Cidade)
	assert.Equal(t, newKennel.Address.Numero, kennelResp.Address.Numero)
	assert.Equal(t, newKennel.Address.Rua, kennelResp.Address.Rua)
	assert.Equal(t, newKennel.Dogs, kennelResp.Dogs)

}

func TestUpdateKennelIfDontExists(t *testing.T) {
	// Instance the mock objects

	kennelMockServ := new(MockKennelServ)
	kennelMockRepo := new(MockKennelRepo)
	addressMockRepo := new(MockAddressRepo)

	// Create an array of dogs with just one dog
	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(1, 1).
		SheddGroomAndEnergy(1, 1, 1)

	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(1).
		BreedID(1).
		DogID(1).
		NameAndSex("m", "b").
		Breed(*breed)
	dog := d.BuildDog()

	var dogs []entities.Dog
	dogs = append(dogs, *dog)

	// Create a new Kennel

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("432").
		Rua("um").
		Bairro("dois").
		CEP("tres").
		Cidade("quatro")

	addr := ad.BuildAddr()
	kennel := entities.BuildKennel(1, dogs, *addr, "1", "x")

	idStr := strconv.Itoa(kennel.ID)

	kennelMockServ.On("CheckIfExists").Return(false)
	kennelMockRepo.On("CheckIfExistsRepo", idStr).Return(false)

	testService := services.NewKennelService(kennelMockRepo, addressMockRepo)
	testController := NewKennelController(kennelMockServ)
	// Update Request

	urlString := "/kennels/update/{id}/" + idStr

	requestBody, err := json.Marshal(kennel)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("UPDATE", urlString, bytes.NewBuffer(requestBody))

	handler := http.HandlerFunc(testController.Update)
	resp := httptest.NewRecorder()

	testService.CheckIfExists(idStr)
	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusNotFound {
		t.Errorf(err.Error(), "expecting status code of 404, got: %v", status)
	}

	kennelMockRepo.AssertExpectations(t)
	kennelMockServ.AssertExpectations(t)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	serverResponse := errors.New(string(responseBody))

	assert.EqualError(t, serverResponse, "404 Not Found")
}
