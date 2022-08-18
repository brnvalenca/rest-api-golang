package services

import (
	"rest-api/golang/exercise/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) FindAll() ([]entities.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.User), args.Error(1)
}

func (mock *MockRepository) FindById(id string) (*entities.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.User), args.Error(1)
}

func (mock *MockRepository) Delete(id string) (*entities.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.User), args.Error(1)
}

func (mock *MockRepository) Update(u *entities.User, id string) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) Save(u *entities.User) (*entities.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.User), args.Error(1)
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

	testService := NewUserService(mockRepo) // Instance testService that will implement the mockRepo interface

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
	mockRepo.On("FindUserById").Return(&user, nil) // mockRepo when calls FindUserById is expected to return &user, nil

	testService := NewUserService(mockRepo)

	result, _ := testService.FindById("1")

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "B", result.Name)
	assert.Equal(t, "b@gmail.com", result.Email)
	assert.Equal(t, "123", result.Password)

}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockRepository)

	// Setup the expectations
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockRepo.On("DeleteUser").Return(&user, nil) // mockRepo when calls FindUserById is expected to return &user, nil

	testService := NewUserService(mockRepo)

	result, _ := testService.Delete("1")

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "B", result.Name)
	assert.Equal(t, "b@gmail.com", result.Email)
	assert.Equal(t, "123", result.Password)

}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockRepository)

	// Setup the expectations
	mockRepo.On("UpdateUser").Return(nil)

	testService := NewUserService(mockRepo)
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	result := testService.Update(&user, "1")

	mockRepo.AssertExpectations(t)

	assert.Equal(t, nil, result)

}

func TestSaveUser(t *testing.T) {
	mockRepo := new(MockRepository)
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockRepo.On("SaveUser").Return(&user, nil)

	testService := NewUserService(mockRepo)
	result, err := testService.Create(&user)
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "B", result.Name)
	assert.Equal(t, "b@gmail.com", result.Email)
	assert.Equal(t, "123", result.Password)
	assert.Nil(t, err)
}

func TestCheckIfExists(t *testing.T) {
	mockRepo := new(MockRepository)
	_ = entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockRepo.On("CheckIfExists").Return(true)

	testService := NewUserService(mockRepo)
	result := testService.Check("1")
	mockRepo.AssertExpectations(t)

	assert.Equal(t, true, result)
}

func TestValidateEmptyUser(t *testing.T) {
	testService := NewUserService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "the user is empty", err.Error())
}

func TestValidateEmptyName(t *testing.T) {
	user := entities.User{ID: 1, Name: "", Email: "b", Password: "1"}
	testService := NewUserService(nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "the user name is empty", err.Error())
}

func TestValidateEmptyEmail(t *testing.T) {
	user := entities.User{ID: 1, Name: "b", Email: "", Password: "1"}
	testService := NewUserService(nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "user email not valid", err.Error())
}

func TestValidateNonValidEmail(t *testing.T) {
	user := entities.User{ID: 1, Name: "b", Email: "b", Password: "1"}
	testService := NewUserService(nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "user email not valid", err.Error())
}

func TestValidateEmptyPassword(t *testing.T) {
	user := entities.User{ID: 1, Name: "b", Email: "b@gmail.com", Password: ""}
	testService := NewUserService(nil)
	err := testService.Validate(&user)

	assert.NotNil(t, err)

	assert.Equal(t, "the user password is empty", err.Error())
}
