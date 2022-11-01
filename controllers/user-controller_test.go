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
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func (m *MockUserService) FindAll() ([]dtos.UserDTO, error) {
	args := m.Called()
	return args.Get(0).([]dtos.UserDTO), args.Error(1)
}

func (m *MockUserService) FindById(id string) (*dtos.UserDTO, error) {
	args := m.Called()
	return args.Get(0).(*dtos.UserDTO), args.Error(1)
}

func (m *MockUserService) Delete(id string) (*dtos.UserDTO, error) {
	args := m.Called()
	return args.Get(0).(*dtos.UserDTO), args.Error(1)
}

func (m *MockUserService) UpdateUser(u *dtos.UserDTO) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserService) Create(u *dtos.UserDTO) (int, error) {
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
		GoodWithKidsAndDogs(2, 3).
		SheddGroomAndEnergy(4, 5, 6)
	userPrefs := upref.BuildUserPref()
	u := entities.NewUserBuilder()
	u.Has().
		ID(0).
		Name("bruno").
		Email("b@gmail.com").
		Password("123").
		Uprefs(*userPrefs)
	user := u.BuildUser()

	//user.Password, _ = security.GeneratePasswordHash(user.Password)
	userDTO := dtos.UserDTO{User: *user}

	mockUserServ.On("CheckEmailServ", userDTO.User.Email).Return(false, user)
	//mockUserRepo.On("CheckEmail", user.Email).Return(false, user)
	mockPassword.On("GeneratePasswordHash", userDTO.User.Password).Return(userDTO.User.Password, nil)
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
	testPassword := security.NewMyHashPassword()
	testController := NewUserController(mockUserServ, mockPassword)

	//Assign HTTP Handler function (controller, Create function)

	handler := http.HandlerFunc(testController.Create)
	testPassword.GeneratePasswordHash(userDTO.User.Password)
	testService.Create(&userDTO)

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
	var user dtos.UserDTO
	_, userInfo := middleware.PartitionUserDTO(&user)
	errReturned := errors.New("the user name is empty")
	fmt.Println(userInfo)
	mockUserServ.On("CheckEmailServ", user.User.Email).Return(false, userInfo)
	mockPassword.On("GeneratePasswordHash", user.User.Password).Return(user.User.Password, nil)
	mockUserServ.On("Create", &user).Return(user.User.ID, nil)
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
		Password("").
		Uprefs(*userPrefs)
	user := u.BuildUser()
	userDTO := dtos.UserDTO{User: *user}
	// Describe my expectations on each call of the mocked objects

	mockUserServ.On("FindAll").Return([]dtos.UserDTO{userDTO}, nil)
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

	var userResp []dtos.UserDTO
	err = json.NewDecoder(io.Reader(resp.Body)).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, userDTO.User.ID, userResp[0].User.ID)
	assert.Equal(t, userDTO.User.Name, userResp[0].User.Name)
	assert.Equal(t, userDTO.User.Email, userResp[0].User.Email)
	assert.Equal(t, userDTO.User.Password, userResp[0].User.Password)
	assert.Equal(t, userDTO.User.UserPreferences.UserID, userResp[0].User.UserPreferences.UserID)
	assert.Equal(t, userDTO.User.UserPreferences.GoodWithKids, userResp[0].User.UserPreferences.GoodWithKids)
	assert.Equal(t, userDTO.User.UserPreferences.GoodWithDogs, userResp[0].User.UserPreferences.GoodWithDogs)
	assert.Equal(t, userDTO.User.UserPreferences.Shedding, userResp[0].User.UserPreferences.Shedding)
	assert.Equal(t, userDTO.User.UserPreferences.Grooming, userResp[0].User.UserPreferences.Grooming)
	assert.Equal(t, userDTO.User.UserPreferences.Energy, userResp[0].User.UserPreferences.Energy)
}

func TestGetById(t *testing.T) {

	//Mock repositories and service
	mockPassword := new(MockPasswordHash)
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

	// Create an instance of user

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
	idStr := strconv.Itoa(user.ID)
	userDTO := dtos.UserDTO{User: *user}
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

	var userResp dtos.UserDTO
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, userDTO.User.ID, userResp.User.ID)
	assert.Equal(t, userDTO.User.Name, userResp.User.Name)
	assert.Equal(t, userDTO.User.Email, userResp.User.Email)
	assert.Equal(t, userDTO.User.Password, userResp.User.Password)
	assert.Equal(t, userDTO.User.UserPreferences.UserID, userResp.User.UserPreferences.UserID)
	assert.Equal(t, userDTO.User.UserPreferences.GoodWithKids, userResp.User.UserPreferences.GoodWithKids)
	assert.Equal(t, userDTO.User.UserPreferences.GoodWithDogs, userResp.User.UserPreferences.GoodWithDogs)
	assert.Equal(t, userDTO.User.UserPreferences.Shedding, userResp.User.UserPreferences.Shedding)
	assert.Equal(t, userDTO.User.UserPreferences.Grooming, userResp.User.UserPreferences.Grooming)
	assert.Equal(t, userDTO.User.UserPreferences.Energy, userResp.User.UserPreferences.Energy)

}

func TestGetByIdIfDontExist(t *testing.T) {

	//Mock repositories and service
	mockPassword := new(MockPasswordHash)
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

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
	//idStr := strconv.Itoa(user.ID)
	userDTO := dtos.UserDTO{User: *user}

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

	upref := entities.NewUserDogPrefsBuilder()
	upref.Has().
		UserID(1).
		GoodWithKidsAndDogs(2, 3).
		SheddGroomAndEnergy(4, 5, 6)

	userPrefs := upref.BuildUserPref()
	u := entities.NewUserBuilder()
	u.Has().
		ID(1).
		Name("bruno").
		Email("b@gmail.com").
		Password("123").
		Uprefs(*userPrefs)
	user := u.BuildUser()
	idStr := strconv.Itoa(user.ID)
	userDTO := dtos.UserDTO{User: *user}
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

	var userResp dtos.UserDTO
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, user.ID, userResp.User.ID)
	assert.Equal(t, user.Name, userResp.User.Name)
	assert.Equal(t, user.Email, userResp.User.Email)
	assert.Equal(t, user.Password, userResp.User.Password)
	assert.Equal(t, user.UserPreferences.UserID, userResp.User.UserPreferences.UserID)
	assert.Equal(t, user.UserPreferences.GoodWithKids, userResp.User.UserPreferences.GoodWithKids)
	assert.Equal(t, user.UserPreferences.GoodWithDogs, userResp.User.UserPreferences.GoodWithDogs)
	assert.Equal(t, user.UserPreferences.Shedding, userResp.User.UserPreferences.Shedding)
	assert.Equal(t, user.UserPreferences.Grooming, userResp.User.UserPreferences.Grooming)
	assert.Equal(t, user.UserPreferences.Energy, userResp.User.UserPreferences.Energy)

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

	upref := entities.NewUserDogPrefsBuilder()
	upref.Has().
		UserID(0).
		GoodWithKidsAndDogs(2, 3).
		SheddGroomAndEnergy(4, 5, 6)
	userPrefs := upref.BuildUserPref()

	u := entities.NewUserBuilder()
	u.Has().
		ID(0).
		Name("bruno").
		Email("b@gmail.com").
		Password("123").
		Uprefs(*userPrefs)
	user := u.BuildUser()
	userDTO := dtos.UserDTO{User: *user}
	// Assert the functions expect to create a new user
	mockUserServ.On("CheckEmailServ", userDTO.User.Email).Return(false, user)
	mockPassword.On("GeneratePasswordHash", userDTO.User.Password).Return(userDTO.User.Password, nil)
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
	testService.Create(&userDTO)

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
		Password("321").
		Uprefs(*userPrefs)
	newUser := u.BuildUser()
	idStr := strconv.Itoa(newUser.ID)
	newUserDTO := dtos.UserDTO{User: *newUser}
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
	userDTO = dtos.UserDTO{User: *newUser}
	// Call the functions
	services.Validate(newUser)
	testService.Check(idStr)
	testService.UpdateUser(&newUserDTO)

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

	var userResp dtos.UserDTO
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error(), "error during body decoding")
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, user.ID, userResp.User.ID)
	assert.Equal(t, newUserDTO.User.Name, userResp.User.Name)
	assert.Equal(t, newUserDTO.User.Email, userResp.User.Email)
	assert.Equal(t, newUserDTO.User.UserPreferences.UserID, userResp.User.UserPreferences.UserID)
	assert.Equal(t, newUserDTO.User.UserPreferences.GoodWithKids, userResp.User.UserPreferences.GoodWithKids)
	assert.Equal(t, newUserDTO.User.UserPreferences.GoodWithDogs, userResp.User.UserPreferences.GoodWithDogs)
	assert.Equal(t, newUserDTO.User.UserPreferences.Shedding, userResp.User.UserPreferences.Shedding)
	assert.Equal(t, newUserDTO.User.UserPreferences.Grooming, userResp.User.UserPreferences.Grooming)
	assert.Equal(t, newUserDTO.User.UserPreferences.Energy, userResp.User.UserPreferences.Energy)

}
