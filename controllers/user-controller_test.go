package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
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

type MockPrefsRepository struct {
	mock.Mock
}

type MockUserRepository struct {
	mock.Mock
}

type MockUserService struct {
	mock.Mock
}

// Prefs Repository Mock

func (m *MockPrefsRepository) SavePrefs(u *entities.UserDogPreferences) error {
	args := m.Called(u)
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

func (mr *MockUserRepository) Update(u *entities.User, id string) error {
	args := mr.Called(u, id)
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

// User Service Mock

func (m *MockUserService) Validate(u *entities.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserService) FindAll() ([]entities.User, error) {
	args := m.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *MockUserService) FindById(id string) (*entities.User, error) {
	args := m.Called()
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) Delete(id string) (*entities.User, error) {
	args := m.Called()
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(u *entities.User, id string) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserService) Create(u *entities.User) (int, error) {
	args := m.Called(u)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) Check(id string) bool {
	args := m.Called()
	return args.Bool(0)
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
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

	userPrefs := entities.BuildUserDogPreferences(1, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 1, "b", "b@gmail.com", "123")

	mockUserServ.On("Create", user).Return(user.ID, nil)
	mockUserRepo.On("Save", user).Return(user.ID, nil)
	mockPrefsRepo.On("SavePrefs", &userPrefs).Return(nil)

	//Create an HTTP Post Request

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonUser))

	testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	testService.Create(user)
	testController := NewUserController(mockUserServ)

	//Assign HTTP Handler function (controller, Create function)

	handler := http.HandlerFunc(testController.Create)

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

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockPrefsRepo.AssertExpectations(t)

	assert.NotNil(t, userID)
	assert.Equal(t, 200, status)
	assert.Equal(t, ID, userID)
}

func TestCreateEmptyUser(t *testing.T) {
	//Mock repositories and service
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)
	var user entities.User
	var userPrefs entities.UserDogPreferences

	mockUserServ.On("Create", &user).Return(user.ID, nil)
	mockUserRepo.On("Save", &user).Return(user.ID, nil)
	mockPrefsRepo.On("SavePrefs", &userPrefs).Return(nil)

	//Create an HTTP Post Request

	jsonUser, err := json.Marshal(nil)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonUser))
	testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	testService.Create(&user)
	testController := NewUserController(mockUserServ)

	//Assign HTTP Handler function (controller, Create function)

	handler := http.HandlerFunc(testController.Create)

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

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockPrefsRepo.AssertExpectations(t)

	assert.NotEqual(t, ID, userID)
	assert.Equal(t, 0, user.ID)
	assert.Equal(t, "", user.Name)
	assert.Equal(t, "", user.Email)
	assert.Equal(t, "", user.Password)
	assert.Equal(t, 0, user.UserPreferences.UserID)
	assert.Equal(t, 0, user.UserPreferences.GoodWithKids)
	assert.Equal(t, 0, user.UserPreferences.GoodWithDogs)
	assert.Equal(t, 0, user.UserPreferences.Shedding)
	assert.Equal(t, 0, user.UserPreferences.Grooming)
	assert.Equal(t, 0, user.UserPreferences.Energy)
}

func TestGetAllUsers(t *testing.T) {
	// Mock user service and repository to be used
	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)

	// Instance a user
	userPrefs := entities.BuildUserDogPreferences(1, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 1, "b", "b@gmail.com", "123")

	// Describe my expectations on each call of the mocked objects

	mockUserServ.On("FindAll").Return([]entities.User{*user}, nil)
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
	testController := NewUserController(mockUserServ)
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

	var userResp []entities.User
	err = json.NewDecoder(io.Reader(resp.Body)).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, ID, userResp[0].ID)
	assert.Equal(t, Name, userResp[0].Name)
	assert.Equal(t, Email, userResp[0].Email)
	assert.Equal(t, Password, userResp[0].Password)
	assert.Equal(t, uID, userResp[0].UserPreferences.UserID)
	assert.Equal(t, GwithKids, userResp[0].UserPreferences.GoodWithKids)
	assert.Equal(t, GwithDogs, userResp[0].UserPreferences.GoodWithDogs)
	assert.Equal(t, Shed, userResp[0].UserPreferences.Shedding)
	assert.Equal(t, Groom, userResp[0].UserPreferences.Grooming)
	assert.Equal(t, Energy, userResp[0].UserPreferences.Energy)
}

func TestGetById(t *testing.T) {

	//Mock repositories and service

	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

	// Create an instance of user

	userPrefs := entities.BuildUserDogPreferences(1, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 1, "b", "b@gmail.com", "123")
	idStr := strconv.Itoa(user.ID)

	// Make the expectations for the mocked functions

	mockUserServ.On("FindById").Return(user, nil)
	mockUserRepo.On("FindById", idStr).Return(user, nil)

	// Create a HTTP GET request

	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err := http.NewRequest("GET", "/users/{id}", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	// Call the mock functions that will be nedded

	testService := services.NewUserService(mockUserRepo, mockPrefsRepo)
	testController := NewUserController(mockUserServ)
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

	var userResp entities.User
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, ID, userResp.ID)
	assert.Equal(t, Name, userResp.Name)
	assert.Equal(t, Email, userResp.Email)
	assert.Equal(t, Password, userResp.Password)
	assert.Equal(t, uID, userResp.UserPreferences.UserID)
	assert.Equal(t, GwithKids, userResp.UserPreferences.GoodWithKids)
	assert.Equal(t, GwithDogs, userResp.UserPreferences.GoodWithDogs)
	assert.Equal(t, Shed, userResp.UserPreferences.Shedding)
	assert.Equal(t, Groom, userResp.UserPreferences.Grooming)
	assert.Equal(t, Energy, userResp.UserPreferences.Energy)

}

func TestGetByIdIfDontExist(t *testing.T) {

	//Mock repositories and service

	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	mockUserServ := new(MockUserService)

	userPrefs := entities.BuildUserDogPreferences(1, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 1, "b", "b@gmail.com", "123")
	//idStr := strconv.Itoa(user.ID)

	errReturned := errors.New("user by ID 5: no such user")
	mockUserServ.On("FindById").Return(user, errReturned)
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
	testController := NewUserController(mockUserServ)

	_, errService := testService.FindById("5")
	handler := http.HandlerFunc(testController.GetById)

	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusNotFound {
		t.Errorf(err.Error(), "expecting 404 status code but got: %v", status)
	}

	b, err := io.ReadAll(resp.Body)
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

	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockUserPref := new(MockPrefsRepository)

	userPrefs := entities.BuildUserDogPreferences(1, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 1, "b", "b@gmail.com", "123")
	idStr := strconv.Itoa(user.ID)

	// Assert the functions expect to create a new user

	mockUserServ.On("Create", user).Return(user.ID, nil)
	mockUserRepo.On("Save", user).Return(user.ID, nil)
	mockUserPref.On("SavePrefs", &userPrefs).Return(nil)

	// Assert the functions to delete the user created

	mockUserServ.On("Check").Return(true)
	mockUserRepo.On("CheckIfExists", "1").Return(true)
	mockUserServ.On("Delete").Return(user, nil)
	mockUserRepo.On("Delete").Return(user, nil)

	// Create a HTTP POST request

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonUser))

	testService := services.NewUserService(mockUserRepo, mockUserPref)
	testController := NewUserController(mockUserServ)
	handler := http.HandlerFunc(testController.Create)

	testService.Create(user)

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	// Create a HTTP Delete request

	jsonData, err := json.Marshal(idStr)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err = http.NewRequest("DELETE", "/users/delete/{id}", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	handler = http.HandlerFunc(testController.Delete)

	testService.Check("1")
	testService.Delete("1")

	resp = httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf(err.Error(), "expecting status code of 200, got: %v", status)
	}

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockUserPref.AssertExpectations(t)

	var userResp entities.User
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, ID, userResp.ID)
	assert.Equal(t, Name, userResp.Name)
	assert.Equal(t, Email, userResp.Email)
	assert.Equal(t, Password, userResp.Password)
	assert.Equal(t, uID, userResp.UserPreferences.UserID)
	assert.Equal(t, GwithKids, userResp.UserPreferences.GoodWithKids)
	assert.Equal(t, GwithDogs, userResp.UserPreferences.GoodWithDogs)
	assert.Equal(t, Shed, userResp.UserPreferences.Shedding)
	assert.Equal(t, Groom, userResp.UserPreferences.Grooming)
	assert.Equal(t, Energy, userResp.UserPreferences.Energy)

}

func TestDeleteIfDontExists(t *testing.T) {
	// Mock the repository and service that will be needed

	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockUserPref := new(MockPrefsRepository)

	userPrefs := entities.BuildUserDogPreferences(5, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 5, "b", "b@gmail.com", "123")
	idStr := "5"

	mockUserServ.On("Check").Return(false)
	mockUserRepo.On("CheckIfExists", idStr).Return(false)
	mockUserRepo.On("Delete").Return(user, nil)

	jsonData, err := json.Marshal(idStr)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, err := http.NewRequest("DELETE", "/users/delete/{id}", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf(err.Error(), "error creating get request")
	}

	testService := services.NewUserService(mockUserRepo, mockUserPref)
	testController := NewUserController(mockUserServ)
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

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	serverResponse := errors.New(string(responseBody))

	assert.EqualError(t, serverResponse, "404 Not Found")

}

func TestUpdate(t *testing.T) {
	// Create a NEW USER

	// Mock the repository and service that will be needed

	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockUserPref := new(MockPrefsRepository)

	userPrefs := entities.BuildUserDogPreferences(1, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 1, "b", "b@gmail.com", "123")

	// Assert the functions expect to create a new user

	mockUserServ.On("Create", user).Return(user.ID, nil)
	mockUserRepo.On("Save", user).Return(user.ID, nil)
	mockUserPref.On("SavePrefs", &userPrefs).Return(nil)

	// Create a HTTP POST request

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonUser))

	testService := services.NewUserService(mockUserRepo, mockUserPref)
	testController := NewUserController(mockUserServ)
	handler := http.HandlerFunc(testController.Create)
	testService.Create(user)

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	// Ending user creation
	/* --------------------- */
	// Update A USER

	newUser := entities.BuildUser(userPrefs, 1, "c", "c@gmail.com", "321")
	idStr := strconv.Itoa(newUser.ID)

	mockUserServ.On("Check").Return(true)
	mockUserRepo.On("CheckIfExists", idStr).Return(true)
	mockUserServ.On("UpdateUser", newUser).Return(nil)
	mockUserRepo.On("Update", newUser, idStr).Return(nil)

	// Update Request

	urlString := "/users/update/" + idStr

	requestBody, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ = http.NewRequest("UPDATE", urlString, bytes.NewBuffer(requestBody))

	handler = http.HandlerFunc(testController.Update)
	resp = httptest.NewRecorder()

	// Call the functions

	testService.Check(idStr)
	testService.UpdateUser(newUser, idStr)

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

	var userResp entities.User
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		t.Errorf(err.Error(), "error during body decoding")
	}

	assert.Equal(t, status, 200)
	assert.Equal(t, user.ID, userResp.ID)
	assert.Equal(t, newUser.Name, userResp.Name)
	assert.Equal(t, newUser.Email, userResp.Email)
	assert.Equal(t, newUser.Password, userResp.Password)
	assert.Equal(t, newUser.UserPreferences.UserID, userResp.UserPreferences.UserID)
	assert.Equal(t, newUser.UserPreferences.GoodWithKids, userResp.UserPreferences.GoodWithKids)
	assert.Equal(t, newUser.UserPreferences.GoodWithDogs, userResp.UserPreferences.GoodWithDogs)
	assert.Equal(t, newUser.UserPreferences.Shedding, userResp.UserPreferences.Shedding)
	assert.Equal(t, newUser.UserPreferences.Grooming, userResp.UserPreferences.Grooming)
	assert.Equal(t, newUser.UserPreferences.Energy, userResp.UserPreferences.Energy)

}

func TestUpdateIfDontExist(t *testing.T) {
	// Mock the repository and service that will be needed

	mockUserServ := new(MockUserService)
	mockUserRepo := new(MockUserRepository)
	mockUserPref := new(MockPrefsRepository)

	userPrefs := entities.BuildUserDogPreferences(4, 2, 3, 4, 5, 6)
	user := entities.BuildUser(userPrefs, 4, "b", "b@gmail.com", "123")

	idStr := "5"

	mockUserServ.On("Check").Return(false)
	mockUserRepo.On("CheckIfExists", idStr).Return(false)

	// Update Request

	urlString := "/users/update/" + idStr

	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Errorf(err.Error(), "error marshalling user to json")
	}
	req, _ := http.NewRequest("UPDATE", urlString, bytes.NewBuffer(requestBody))

	resp := httptest.NewRecorder()

	// Call the functions

	testService := services.NewUserService(mockUserRepo, mockUserPref)
	testController := NewUserController(mockUserServ)
	handler := http.HandlerFunc(testController.Update)

	testService.Check(idStr)

	// Servers UP
	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusNotFound {
		t.Errorf(err.Error(), "expecting status code of 404, got: %v", status)
	}

	mockUserServ.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	serverResponse := errors.New(string(responseBody))

	assert.EqualError(t, serverResponse, "404 Not Found")
}
