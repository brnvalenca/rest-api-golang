package services

import (
	"rest-api/golang/exercise/domain/dtos"
	"rest-api/golang/exercise/domain/entities"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPrefsRepository struct {
	mock.Mock
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockPrefsRepository) SavePrefs(u *entities.UserDogPreferences, userid int) error {
	args := m.Called(u, userid)
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

func (mock *MockUserRepository) FindAll() ([]entities.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.User), args.Error(1)
}

func (mock *MockUserRepository) FindById(id string) (*entities.User, error) {
	args := mock.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (mock *MockUserRepository) Delete(id string) (*entities.User, error) {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(*entities.User), args.Error(1)
}

func (mock *MockUserRepository) Update(u *entities.User, uprefs *entities.UserDogPreferences) error {
	args := mock.Called(u, uprefs)
	return args.Error(0)
}

func (mock *MockUserRepository) Save(u *entities.User) (int, error) {
	args := mock.Called(u)
	return args.Int(0), args.Error(1)
}

func (mock *MockUserRepository) CheckIfExists(id string) bool {
	args := mock.Called()
	result := args.Bool(0)
	return result
}

func (mock *MockUserRepository) CheckEmail(email string) (bool, *entities.User) {
	args := mock.Called(email)
	return args.Bool(0), args.Get(1).(*entities.User)
}

func MakeUserDTO() *dtos.UserDTOSignUp {
	userDtoBuilder := dtos.NewUserDTOBuilder()
	userDtoBuilder.Has().
		ID(0).
		Name("Bruno").
		Email("b@gmail.com").
		Password("123")

	userDTO := userDtoBuilder.BuildUser()

	return userDTO
}

func TestFindAll(t *testing.T) {
	mockUserRepo := new(MockUserRepository) // Struct that will hold the repository methods

	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockUserRepo.On("FindAll").Return([]entities.User{user}, nil)

	testService := NewUserService(mockUserRepo, nil) // Instance testService that will implement the mockRepo interface

	result, _ := testService.FindAll()

	mockUserRepo.AssertExpectations(t)

	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "B", result[0].Name)
	assert.Equal(t, "b@gmail.com", result[0].Email)
	assert.Equal(t, "123", result[0].Password)
}

func TestFindById(t *testing.T) {
	mockUserRepo := new(MockUserRepository)

	// Setup the expectations
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockUserRepo.On("FindById", "1").Return(&user, nil) // mockRepo when calls FindUserById is expected to return &user, nil

	testService := NewUserService(mockUserRepo, nil)

	result, _ := testService.FindById("1")

	// Mock Assertion: Behavioral
	mockUserRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "B", result.Name)
	assert.Equal(t, "b@gmail.com", result.Email)
	assert.Equal(t, "123", result.Password)

}

func TestDelete(t *testing.T) {
	mockUserRepo := new(MockUserRepository)

	// Setup the expectations
	user := entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockUserRepo.On("Delete", "1").Return(&user, nil) // mockRepo when calls FindUserById is expected to return &user, nil

	testService := NewUserService(mockUserRepo, nil)

	result, _ := testService.Delete("1")

	// Mock Assertion: Behavioral
	mockUserRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "B", result.Name)
	assert.Equal(t, "b@gmail.com", result.Email)
	assert.Equal(t, "123", result.Password)

}

func TestUpdate(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	prefs := entities.UserDogPreferences{UserID: 0, GoodWithKids: 0, GoodWithDogs: 0, Shedding: 0, Grooming: 0, Energy: 0}
	u := entities.NewUserBuilder()
	u.Has().
		ID(0).
		Name("Bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	userDTO := MakeUserDTO()
	// Setup the expectations
	mockUserRepo.On("Update", user, &prefs).Return(nil)

	testService := NewUserService(mockUserRepo, nil)
	result := testService.UpdateUser(userDTO)

	mockUserRepo.AssertExpectations(t)

	assert.Equal(t, nil, result)

}

func TestSave(t *testing.T) {

	mockUserRepo := new(MockUserRepository)
	mockPrefsRepo := new(MockPrefsRepository)
	prefs := entities.UserDogPreferences{UserID: 0, GoodWithKids: 2, GoodWithDogs: 3, Shedding: 4, Grooming: 5, Energy: 5}
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(0).
		Name("Bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	userDTO := MakeUserDTO()

	mockPrefsRepo.On("SavePrefs", &prefs, user.ID).Return(nil)
	mockUserRepo.On("Save", user).Return(user.ID, nil)

	testService := NewUserService(mockUserRepo, mockPrefsRepo)
	result, err := testService.Create(userDTO)

	mockUserRepo.AssertExpectations(t)
	mockPrefsRepo.AssertExpectations(t)

	assert.Equal(t, 0, result)
	assert.Nil(t, nil, err)

}

func TestCheckIfExists(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	_ = entities.User{ID: 1, Name: "B", Email: "b@gmail.com", Password: "123"}
	mockUserRepo.On("CheckIfExists").Return(true)

	testService := NewUserService(mockUserRepo, nil)
	result := testService.Check("1")
	mockUserRepo.AssertExpectations(t)

	assert.Equal(t, true, result)
}

func TestValidateEmptyUser(t *testing.T) {
	user := entities.User{}
	var flag bool
	err := Validate(&user, flag)

	assert.NotNil(t, err)

	assert.Equal(t, "the user is empty", err.Error())
}

func TestValidateEmptyName(t *testing.T) {
	prefs := entities.UserDogPreferences{UserID: 1, GoodWithKids: 2, GoodWithDogs: 3, Shedding: 4, Grooming: 5, Energy: 5}
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("Bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	var flag bool
	err := Validate(user, flag)

	assert.NotNil(t, err)

	assert.Equal(t, "the user name is empty", err.Error())
}

func TestValidateEmptyEmail(t *testing.T) {
	prefs := entities.UserDogPreferences{UserID: 1, GoodWithKids: 2, GoodWithDogs: 3, Shedding: 4, Grooming: 5, Energy: 5}
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("Bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	var flag bool
	err := Validate(user, flag)

	assert.NotNil(t, err)

	assert.Equal(t, "user email not valid", err.Error())
}

func TestValidateNonValidEmail(t *testing.T) {
	prefs := entities.UserDogPreferences{UserID: 1, GoodWithKids: 2, GoodWithDogs: 3, Shedding: 4, Grooming: 5, Energy: 5}
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("Bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()

	var flag bool
	err := Validate(user, flag)

	assert.NotNil(t, err)

	assert.Equal(t, "user email not valid", err.Error())
}

func TestValidateEmptyPassword(t *testing.T) {
	prefs := entities.UserDogPreferences{UserID: 1, GoodWithKids: 2, GoodWithDogs: 3, Shedding: 4, Grooming: 5, Energy: 5}
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("Bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	var flag bool
	err := Validate(user, flag)

	assert.NotNil(t, err)

	assert.Equal(t, "the user password is empty", err.Error())
}
