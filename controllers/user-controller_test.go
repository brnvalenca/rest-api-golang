package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/middleware"
	"rest-api/golang/exercise/services"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO: Rodar os testes!

type MockPrefsRepository struct {
	mock.Mock
}

type MockUserRepository struct {
	mock.Mock
}

type MockUserService struct {
	mock.Mock
}

type MockPasswordHash struct {
	mock.Mock
}

// Security Password Hash Mock

func (hash *MockPasswordHash) GeneratePasswordHash(password string) (string, error) {
	args := hash.Called(password)
	return args.String(0), args.Error(1)
}

func (hash *MockPasswordHash) CheckPassword(hashedPassword, passwordString string) bool {
	args := hash.Called(hashedPassword, passwordString)
	return args.Bool(0)
}

// Prefs Repository Mock

func (m *MockPrefsRepository) SavePrefs(u *entities.UserDogPreferences, userID int) error {
	args := m.Called(u, userID)
	return args.Error(0)
}

func (m *MockPrefsRepository) DeletePrefs(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPrefsRepository) UpdatePrefs(u *entities.UserDogPreferences, id string) error {
	args := m.Called(u, id)
	return args.Error(0)
}

// User Repository Mock

func (mr *MockUserRepository) FindAll() ([]entities.User, error) {
	args := mr.Called()
	result := args.Get(0)
	return result.([]entities.User), args.Error(1)
}

func (mr *MockUserRepository) FindById(id string) (*entities.User, error) {
	args := mr.Called(id)
	result := args.Get(0)
	return result.(*entities.User), args.Error(1)
}

func (mr *MockUserRepository) Delete(id string) (*entities.User, error) {
	args := mr.Called()
	result := args.Get(0)
	return result.(*entities.User), args.Error(1)
}

func (mr *MockUserRepository) Update(u *entities.User, uprefs *entities.UserDogPreferences) error {
	args := mr.Called(u, uprefs)
	return args.Error(0)
}

func (mr *MockUserRepository) Save(u *entities.User) (int, error) {
	args := mr.Called(u)
	return args.Int(0), args.Error(1)
}

func (mr *MockUserRepository) CheckIfExists(id string) bool {
	args := mr.Called(id)
	return args.Bool(0)
}

func (mr *MockUserRepository) CheckEmail(email string) (bool, *entities.User) {
	args := mr.Called(email)
	return args.Bool(0), args.Get(1).(*entities.User)
}

// User Service Mock

func (m *MockUserService) Validate(u *entities.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserService) FindAll() ([]dtos.UserDTOSignUp, error) {
	args := m.Called()
	return args.Get(0).([]dtos.UserDTOSignUp), args.Error(1)
}

func (m *MockUserService) FindById(id string) (*dtos.UserDTOSignUp, error) {
	args := m.Called()
	return args.Get(0).(*dtos.UserDTOSignUp), args.Error(1)
}

func (m *MockUserService) Delete(id string) (*dtos.UserDTOSignUp, error) {
	args := m.Called()
	return args.Get(0).(*dtos.UserDTOSignUp), args.Error(1)
}

func (m *MockUserService) UpdateUser(u *dtos.UserDTOSignUp) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserService) Create(u *dtos.UserDTOSignUp) (int, error) {
	args := m.Called(u)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) Check(id string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockUserService) CheckEmailServ(email string) (bool, *entities.User) {
	args := m.Called(email)
	return args.Bool(0), args.Get(1).(*entities.User)
}

const (
	ID        int    = 1
	Name      string = "b"
	Email     string = "b@gmail.com"
	Password  string = "123"
	uID       int    = 1
	GwithKids int    = 2
	GwithDogs int    = 3
	Shed      int    = 4
	Groom     int    = 5
	Energy    int    = 6
)

func TestCreateUser(t *testing.T) {
	//Mock repositories and service
	mockPassword := new(MockPasswordHash)
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

	upref := entities.NewUserDogPrefsBuilder()
	upref.Has().
		UserID(0).
		GoodWithKidsAndDogs(0, 0).
		SheddGroomAndEnergy(0, 0, 0)
	userPrefs := upref.BuildUserPref()

	u := entities.NewUserBuilder()
	u.Has().
		ID(0).
		Name("bruno").
		Email("b@gmail.com").
		Password("1")
	user := u.BuildUser()

	//user.Password, _ = security.GeneratePasswordHash(user.Password)

	dtoBuilder := dtos.NewUserDTOBuilder()
	dtoBuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password)
	userDTO := dtoBuilder.BuildUser()

	mockUserServ.On("CheckEmailServ", userDTO.Email).Return(false, user)
	mockPassword.On("GeneratePasswordHash", userDTO.Password).Return(userDTO.Password, nil)
	mockUserServ.On("Create", &userDTO).Return(user.ID, nil)
	mockUserRepo.On("Save", user).Return(user.ID, nil)
	mockPrefsRepo.On("SavePrefs", userPrefs, user.ID).Return(nil)

	//Create an HTTP Post Request

	jsonUser, err := json.Marshal(userDTO)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonUser))

	testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	testController := NewUserController(mockUserServ, mockPassword)

	//Assign HTTP Handler function (controller, Create function)

	handler := http.HandlerFunc(testController.Create)
	testService.Create(userDTO)

	//Record the HTTP Response(httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request

	handler.ServeHTTP(response, req)
	//Add the assertions on the HTTP Status code and the response
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got: %v", status)
	}

	// Decode the HTTP response

	var userID int
	err = json.NewDecoder(io.Reader(response.Body)).Decode(&userID)
	if err != nil {
		t.Errorf(err.Error())
	}
	mockPassword.AssertExpectations(t)
	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockPrefsRepo.AssertExpectations(t)

	assert.NotNil(t, userID)
	assert.Equal(t, 200, status)
	assert.Equal(t, 0, userID)
}

func TestCreateEmptyUser(t *testing.T) {
	//Mock repositories and service
	//mockUserRepo := new(MockUserRepository)
	//mockPrefsRepo := new(MockPrefsRepository)
	mockPassword := new(MockPasswordHash)
	mockUserServ := new(MockUserService)
	var user dtos.UserDTOSignUp
	_, userInfo := middleware.PartitionUserDTO(&user)
	errReturned := errors.New("the user name is empty")
	fmt.Println(userInfo)
	mockUserServ.On("CheckEmailServ", user.Email).Return(false, userInfo)
	mockPassword.On("GeneratePasswordHash", user.Password).Return(user.Password, nil)
	mockUserServ.On("Create", &user).Return(user.ID, nil)
	//mockUserServ.On("Validate", userInfo).Return(errReturned)

	//Create an HTTP Post Request

	jsonUser, err := json.Marshal(nil)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonUser))
	//testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	statusReturned := services.Validate(userInfo)
	fmt.Println(statusReturned)
	testController := NewUserController(mockUserServ, mockPassword)

	//Assign HTTP Handler function (controller, Create function)

	handler := http.HandlerFunc(testController.Create)

	//Record the HTTP Response(httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request

	handler.ServeHTTP(response, req)

	// Decode the HTTP response

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	mockUserServ.AssertExpectations(t)
	mockPassword.AssertExpectations(t)

	assert.EqualError(t, statusReturned, errReturned.Error())
}

func TestGetAllUsers(t *testing.T) {
	// Mock user service and repository to be used
	mockPassword := new(MockPasswordHash)
	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)

	upref := dtos.NewUserPrefsDTOBuilder()
	upref.Has().
		UserID(0).
		GoodWithKids(2).
		GoodWithDogs(3).
		SheddAndGroom(4, 5).
		Energy(6)
	userPrefs := upref.BuildUserPrefsDTO()

	u := entities.NewUserBuilder()
	u.Has().
		Name("bruno").
		Email("b@gmail.com").
		Password("")
	user := u.BuildUser()

	dtoBuilder := dtos.NewUserDTOBuilder()
	dtoBuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password).
		UserPrefs(*userPrefs)
	userDTO := dtoBuilder.BuildUser()
	// Describe my expectations on each call of the mocked objects

	mockUserServ.On("FindAll").Return([]dtos.UserDTOSignUp{*userDTO}, nil)
	mockUserRepo.On("FindAll").Return([]entities.User{*user}, nil)

	// Create a new HTTP GET Request

	jsonData, err := json.Marshal(user) // Marshalling the user to json format
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err := http.NewRequest("GET", "/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	// Create a service and controller instance

	testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	testController := NewUserController(mockUserServ, mockPassword)
	testService.FindAll() // Call the FindAll service function

	// Create a handler func
	handler := http.HandlerFunc(testController.GetAll)

	// Create a response type
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)
	status := resp.Code
	if status != 200 {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)

	var userResp []dtos.UserDTOSignUp
	err = json.NewDecoder(io.Reader(resp.Body)).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, userDTO.ID, userResp[0].ID)
	assert.Equal(t, userDTO.Name, userResp[0].Name)
	assert.Equal(t, userDTO.Email, userResp[0].Email)
	assert.Equal(t, userDTO.Password, userResp[0].Password)
	assert.Equal(t, userDTO.UserPrefs.UserID, userResp[0].UserPrefs.UserID)
	assert.Equal(t, userDTO.UserPrefs.GoodWithKids, userResp[0].UserPrefs.GoodWithKids)
	assert.Equal(t, userDTO.UserPrefs.GoodWithDogs, userResp[0].UserPrefs.GoodWithDogs)
	assert.Equal(t, userDTO.UserPrefs.Shedding, userResp[0].UserPrefs.Shedding)
	assert.Equal(t, userDTO.UserPrefs.Grooming, userResp[0].UserPrefs.Grooming)
	assert.Equal(t, userDTO.UserPrefs.Energy, userResp[0].UserPrefs.Energy)
}

func TestGetById(t *testing.T) {

	//Mock repositories and service
	mockPassword := new(MockPasswordHash)
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

	upref := dtos.NewUserPrefsDTOBuilder()
	upref.Has().
		UserID(0).
		GoodWithKids(2).
		GoodWithDogs(3).
		SheddAndGroom(4, 5).
		Energy(6)
	userPrefs := upref.BuildUserPrefsDTO()

	u := entities.NewUserBuilder()
	u.Has().
		Name("bruno").
		Email("b@gmail.com").
		Password("")
	user := u.BuildUser()

	dtoBuilder := dtos.NewUserDTOBuilder()
	dtoBuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password).
		UserPrefs(*userPrefs)
	userDTO := dtoBuilder.BuildUser()

	idStr := strconv.Itoa(user.ID)
	// Make the expectations for the mocked functions

	mockUserServ.On("FindById").Return(&userDTO, nil)
	mockUserRepo.On("FindById", idStr).Return(user, nil)

	// Create a HTTP GET request

	jsonData, err := json.Marshal(userDTO)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err := http.NewRequest("GET", "/users/{id}", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	// Call the mock functions that will be nedded

	testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	testController := NewUserController(mockUserServ, mockPassword)
	handler := http.HandlerFunc(testController.GetById) // Assign a HTTP handler func calling the controller function
	testService.FindById(idStr)

	// Assign a response to the request

	resp := httptest.NewRecorder()

	// Create a HTTP server that will call the request made and store the response on the NewRecorder response assigned

	handler.ServeHTTP(resp, req)

	// Now we see if the status code of the resp is ok

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)

	var userResp dtos.UserDTOSignUp
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, userDTO.ID, userResp.ID)
	assert.Equal(t, userDTO.Name, userResp.Name)
	assert.Equal(t, userDTO.Email, userResp.Email)
	assert.Equal(t, userDTO.Password, userResp.Password)
	assert.Equal(t, userDTO.UserPrefs.UserID, userResp.UserPrefs.UserID)
	assert.Equal(t, userDTO.UserPrefs.GoodWithKids, userResp.UserPrefs.GoodWithKids)
	assert.Equal(t, userDTO.UserPrefs.GoodWithDogs, userResp.UserPrefs.GoodWithDogs)
	assert.Equal(t, userDTO.UserPrefs.Shedding, userResp.UserPrefs.Shedding)
	assert.Equal(t, userDTO.UserPrefs.Grooming, userResp.UserPrefs.Grooming)
	assert.Equal(t, userDTO.UserPrefs.Energy, userResp.UserPrefs.Energy)

}

func TestGetByIdIfDontExist(t *testing.T) {

	//Mock repositories and service
	mockPassword := new(MockPasswordHash)
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

	upref := dtos.NewUserPrefsDTOBuilder()
	upref.Has().
		UserID(0).
		GoodWithKids(2).
		GoodWithDogs(3).
		SheddAndGroom(4, 5).
		Energy(6)
	userPrefs := upref.BuildUserPrefsDTO()

	u := entities.NewUserBuilder()
	u.Has().
		Name("bruno").
		Email("b@gmail.com").
		Password("")
	user := u.BuildUser()

	dtoBuilder := dtos.NewUserDTOBuilder()
	dtoBuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password).
		UserPrefs(*userPrefs)
	userDTO := dtoBuilder.BuildUser()

	errReturned := errors.New("user by ID 5: no such user")
	mockUserServ.On("FindById").Return(&userDTO, errReturned)
	mockUserRepo.On("FindById", "5").Return(user, errReturned)

	jsonData, err := json.Marshal("5")
	if err != nil {
		t.Errorf(err.Error(), "error marshalling id to json")
	}
	req, err := http.NewRequest("GET", "/users/{id}", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	testController := NewUserController(mockUserServ, mockPassword)

	_, errService := testService.FindById("5")
	handler := http.HandlerFunc(testController.GetById)

	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusNotFound {
		t.Errorf(err.Error(), "expecting 404 status code but got: %v", status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	serverResponse := errors.New(string(b))

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)

	assert.EqualError(t, serverResponse, "404 Not Found")
	assert.EqualError(t, errService, "user by ID 5: no such user")

}

func TestDelete(t *testing.T) {

	// Mock the repository and service that will be needed
	mockPassword := new(MockPasswordHash)
	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockUserPref := new(MockPrefsRepository)

	upref := dtos.NewUserPrefsDTOBuilder()
	upref.Has().
		UserID(0).
		GoodWithKids(2).
		GoodWithDogs(3).
		SheddAndGroom(4, 5).
		Energy(6)
	userPrefs := upref.BuildUserPrefsDTO()

	u := entities.NewUserBuilder()
	u.Has().
		Name("bruno").
		Email("b@gmail.com").
		Password("")
	user := u.BuildUser()

	dtoBuilder := dtos.NewUserDTOBuilder()
	dtoBuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password).
		UserPrefs(*userPrefs)
	userDTO := dtoBuilder.BuildUser()
	idStr := strconv.Itoa(user.ID)
	// Assert the functions to delete the user created

	mockUserServ.On("Check").Return(true)
	mockUserRepo.On("CheckIfExists", "1").Return(true)
	mockUserServ.On("Delete").Return(&userDTO, nil)
	mockUserRepo.On("Delete").Return(user, nil)

	// Create a HTTP Delete request

	jsonData, err := json.Marshal(idStr)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err := http.NewRequest("DELETE", "/users/delete/{id}", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}
	testService := services.NewUserService(mockUserRepo, mockUserPref)
	testController := NewUserController(mockUserServ, mockPassword)
	handler := http.HandlerFunc(testController.Delete)

	testService.Check("1")
	testService.Delete("1")

	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockUserPref.AssertExpectations(t)

	var userResp dtos.UserDTOSignUp
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, user.ID, userResp.ID)
	assert.Equal(t, user.Name, userResp.Name)
	assert.Equal(t, user.Email, userResp.Email)
	assert.Equal(t, user.Password, userResp.Password)
	assert.Equal(t, user.UserPreferences.UserID, userResp.UserPrefs.UserID)
	assert.Equal(t, user.UserPreferences.GoodWithKids, userResp.UserPrefs.GoodWithKids)
	assert.Equal(t, user.UserPreferences.GoodWithDogs, userResp.UserPrefs.GoodWithDogs)
	assert.Equal(t, user.UserPreferences.Shedding, userResp.UserPrefs.Shedding)
	assert.Equal(t, user.UserPreferences.Grooming, userResp.UserPrefs.Grooming)
	assert.Equal(t, user.UserPreferences.Energy, userResp.UserPrefs.Energy)

}

func TestDeleteIfDontExists(t *testing.T) {
	// Mock the repository and service that will be needed
	mockPassword := new(MockPasswordHash)
	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockUserPref := new(MockPrefsRepository)

	upref := entities.NewUserDogPrefsBuilder()
	upref.Has().
		UserID(1).
		GoodWithKidsAndDogs(2, 3).
		SheddGroomAndEnergy(4, 5, 6)

	userPrefs := upref.BuildUserPref()
	u := entities.NewUserBuilder()
	u.Has().
		Name("bruno").
		Email("b@gmail.com").
		Password("123").
		Uprefs(*userPrefs)
	user := u.BuildUser()
	idStr := "6"

	mockUserServ.On("Check").Return(false)
	mockUserRepo.On("CheckIfExists", idStr).Return(false)
	mockUserRepo.On("Delete").Return(user, nil)

	jsonData, err := json.Marshal(idStr)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err := http.NewRequest("DELETE", "/users/delete/{id}", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating delete request")
	}

	testService := services.NewUserService(mockUserRepo, mockUserPref)
	testController := NewUserController(mockUserServ, mockPassword)
	handler := http.HandlerFunc(testController.Delete)

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	testService.Check(idStr)
	testService.Delete(idStr)

	resp = httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusNotFound {
		t.Errorf(err.Error(), "expecting status code of 404, got: %v", status)
	}

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	serverResponse := errors.New(string(responseBody))

	assert.EqualError(t, serverResponse, "404 Not Found")

}

func TestUpdate(t *testing.T) {
	// Create a NEW USER

	// Mock the repository and service that will be needed
	mockPassword := new(MockPasswordHash)
	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockUserPref := new(MockPrefsRepository)

	upref := dtos.NewUserPrefsDTOBuilder()
	upref.Has().
		UserID(0).
		GoodWithKids(2).
		GoodWithDogs(3).
		SheddAndGroom(4, 5).
		Energy(6)
	userPrefs := upref.BuildUserPrefsDTO()

	u := entities.NewUserBuilder()
	u.Has().
		Name("bruno").
		Email("b@gmail.com").
		Password("")
	user := u.BuildUser()

	dtoBuilder := dtos.NewUserDTOBuilder()
	dtoBuilder.Has().
		ID(user.ID).
		Name(user.Name).
		Email(user.Email).
		Password(user.Password).
		UserPrefs(*userPrefs)
	userDTO := dtoBuilder.BuildUser()

	// Assert the functions expect to create a new user
	mockUserServ.On("CheckEmailServ", userDTO.Email).Return(false, user)
	mockPassword.On("GeneratePasswordHash", userDTO.Password).Return(userDTO.Password, nil)
	mockUserServ.On("Create", &userDTO).Return(user.ID, nil)
	mockUserRepo.On("Save", user).Return(user.ID, nil)
	mockUserPref.On("SavePrefs", userPrefs, user.ID).Return(nil)

	// Create a HTTP POST request

	jsonUser, err := json.Marshal(userDTO)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonUser))

	testService := services.NewUserService(mockUserRepo, mockUserPref)
	testController := NewUserController(mockUserServ, mockPassword)
	handler := http.HandlerFunc(testController.Create)

	services.Validate(user)
	testService.Create(userDTO)

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	// Ending user creation
	/* --------------------- */
	// Update A USER

	u = entities.NewUserBuilder()
	u.Has().
		ID(0).
		Name("c").
		Email("c@gmail.com").
		Password("321")
	newUser := u.BuildUser()

	dtoBuilder.Has().
		ID(user.ID).
		Name(newUser.Name).
		Email(newUser.Email).
		Password(newUser.Password).
		UserPrefs(*userPrefs)
	newUserDTO := dtoBuilder.BuildUser()

	idStr := strconv.Itoa(user.ID)
	mockUserServ.On("Check").Return(true)
	mockUserRepo.On("CheckIfExists", idStr).Return(true)
	mockUserServ.On("UpdateUser", &newUserDTO).Return(nil)
	mockUserRepo.On("Update", newUser, userPrefs).Return(nil)

	// Update Request

	urlString := "/users/update/" + idStr

	requestBody, err := json.Marshal(newUserDTO)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ = http.NewRequest("UPDATE", urlString, bytes.NewBuffer(requestBody))

	handler = http.HandlerFunc(testController.Update)
	resp = httptest.NewRecorder()
	// Call the functions
	services.Validate(newUser)
	testService.Check(idStr)
	testService.UpdateUser(newUserDTO)

	// Servers UP

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockUserPref.AssertExpectations(t)

	//Decode de body response

	var userResp dtos.UserDTOSignUp
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error(), "error during body decoding")
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, user.ID, userResp.ID)
	assert.Equal(t, newUserDTO.Name, userResp.Name)
	assert.Equal(t, newUserDTO.Email, userResp.Email)
	assert.Equal(t, newUserDTO.UserPrefs.UserID, userResp.UserPrefs.UserID)
	assert.Equal(t, newUserDTO.UserPrefs.GoodWithKids, userResp.UserPrefs.GoodWithKids)
	assert.Equal(t, newUserDTO.UserPrefs.GoodWithDogs, userResp.UserPrefs.GoodWithDogs)
	assert.Equal(t, newUserDTO.UserPrefs.Shedding, userResp.UserPrefs.Shedding)
	assert.Equal(t, newUserDTO.UserPrefs.Grooming, userResp.UserPrefs.Grooming)
	assert.Equal(t, newUserDTO.UserPrefs.Energy, userResp.UserPrefs.Energy)

}
