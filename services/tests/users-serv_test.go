package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPrefsRepository struct {
	mock.Mock
}

type MockRepository struct {
	mock.Mock
}

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

func (mock *MockRepository) FindAll() ([]entities.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.User), args.Error(1)
}

func (mock *MockRepository) FindById(id string) (*entities.User, error) {
	args := mock.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (mock *MockRepository) Delete(id string) (*entities.User, error) {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(*entities.User), args.Error(1)
}

func (mock *MockRepository) Update(u *entities.User, id string) error {
	args := mock.Called(u, id)
	return args.Error(0)
}

func (mock *MockRepository) Save(u *entities.User) (int, error) {
	args := mock.Called(u)
	return args.Int(0), args.Error(1)
}

func (mock *MockRepository) CheckIfExists(id string) bool {
	args := mock.Called()
	result := args.Bool(0)
	return result
}
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository) // Struct that will hold the repository methods

	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockRepo.On("FindAll").Return([]entities.User{user}, nil)
	/*
		Defining the expectations. Im basically saying that when a call de FindAll method, it is expected to return the
		elements insed the Return function.
	*/

	testService := services.NewUserService(mockRepo, nil) // Instance testService that will implement the mockRepo interface

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "B", result[0].Name)
	assert.Equal(t, "b@gmail.com", result[0].Email)
	assert.Equal(t, "123", result[0].Password)
}

func TestFindById(t *testing.T) {
	mockRepo := new(MockRepository)

	// Setup the expectations
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockRepo.On("FindById", "1").Return(&user, nil) // mockRepo when calls FindUserById is expected to return &user, nil

	testService := services.NewUserService(mockRepo, nil)

	result, _ := testService.FindById("1")

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "B", result.Name)
	assert.Equal(t, "b@gmail.com", result.Email)
	assert.Equal(t, "123", result.Password)

}

func TestDelete(t *testing.T) {
	mockRepo := new(MockRepository)

	// Setup the expectations
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockRepo.On("Delete", "1").Return(&user, nil) // mockRepo when calls FindUserById is expected to return &user, nil

	testService := services.NewUserService(mockRepo, nil)

	result, _ := testService.Delete("1")

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "B", result.Name)
	assert.Equal(t, "b@gmail.com", result.Email)
	assert.Equal(t, "123", result.Password)

}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockRepository)
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	// Setup the expectations
	mockRepo.On("Update", &user, "1").Return(nil)

	testService := services.NewUserService(mockRepo, nil)
	result := testService.Update(&user, "1")

	mockRepo.AssertExpectations(t)

	assert.Equal(t, nil, result)

}

func TestSave(t *testing.T) {

	mockRepo := new(MockRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	prefs := entities.UserDogPreferences{UserID: 1, GoodWithKids: 2, GoodWithDogs: 3, Shedding: 4, Grooming: 5, Energy: 5}
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123", UserPreferences: prefs}

	mockPrefsRepo.On("SavePrefs", &prefs).Return(nil)
	mockRepo.On("Save", &user).Return(user.ID, nil)

	testService := services.NewUserService(mockRepo, mockPrefsRepo)
	result, err := testService.Create(&user)

	mockRepo.AssertExpectations(t)
	mockPrefsRepo.AssertExpectations(t)

	assert.Equal(t, 1, result)
	assert.Nil(t, nil, err)

}

func TestCheckIfExists(t *testing.T) {
	mockRepo := new(MockRepository)
	_ = entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockRepo.On("CheckIfExists").Return(true)

	testService := services.NewUserService(mockRepo, nil)
	result := testService.Check("1")
	mockRepo.AssertExpectations(t)

	assert.Equal(t, true, result)
}

func TestValidateEmptyUser(t *testing.T) {
	testService := services.NewUserService(nil, nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "the user is empty", err.Error())
}

func TestValidateEmptyName(t *testing.T) {
	user := entities.User{ID: 1, Name: "", Email: "b", Password: "1"}
	testService := services.NewUserService(nil, nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "the user name is empty", err.Error())
}

func TestValidateEmptyEmail(t *testing.T) {
	user := entities.User{ID: 1, Name: "b", Email: "", Password: "1"}
	testService := services.NewUserService(nil, nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "user email not valid", err.Error())
}

func TestValidateNonValidEmail(t *testing.T) {
	user := entities.User{ID: 1, Name: "b", Email: "b", Password: "1"}
	testService := services.NewUserService(nil, nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "user email not valid", err.Error())
}

func TestValidateEmptyPassword(t *testing.T) {
	user := entities.User{ID: 1, Name: "b", Email: "b@gmail.com", Password: ""}
	testService := services.NewUserService(nil, nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "the user password is empty", err.Error())
}
