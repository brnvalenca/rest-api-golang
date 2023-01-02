package repository

import (
	"rest-api/golang/exercise/domain/entities"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (ur *MockUserRepo) Save(u *entities.User) (int, error) {
	args := ur.Called(u)
	return args.Int(0), args.Error(1)
}

func (ur *MockUserRepo) FindAll() ([]entities.User, error) {
	args := ur.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

func (ur *MockUserRepo) FindById(id string) (*entities.User, error) {
	args := ur.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (ur *MockUserRepo) Delete(id string) (*entities.User, error) {
	args := ur.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (ur *MockUserRepo) Update(u *entities.User, id string) error {
	args := ur.Called(u, id)
	return args.Error(0)
}

func (ur *MockUserRepo) CheckIfExists(id string) bool {
	args := ur.Called(id)
	return args.Bool(0)
}

func TestSaveUser(t *testing.T) {
	mockUser := new(MockUserRepo)
	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()

	mockUser.On("Save", user).Return(user.ID, nil)

	got, err := mockUser.Save(user)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, got)
}

func TestFindAllUsers(t *testing.T) {
	mockUser := new(MockUserRepo)
	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()

	mockUser.On("FindAll").Return([]entities.User{*user}, nil)

	got, err := mockUser.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, got[0].ID, user.ID)
	assert.Equal(t, got[0].Email, user.Email)
	assert.Equal(t, got[0].Name, user.Name)
	assert.Equal(t, got[0].Password, user.Password)
	assert.Equal(t, got[0].UserPreferences, user.UserPreferences)
}

func TestFindUserById(t *testing.T) {
	mockUser := new(MockUserRepo)
	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	idStr := strconv.Itoa(user.ID)

	mockUser.On("FindById", idStr).Return(user, nil)

	got, err := mockUser.FindById(idStr)

	assert.Nil(t, err)
	assert.Equal(t, got.ID, user.ID)
	assert.Equal(t, got.Email, user.Email)
	assert.Equal(t, got.Name, user.Name)
	assert.Equal(t, got.Password, user.Password)
	assert.Equal(t, got.UserPreferences, user.UserPreferences)
}

func TestDeleteUser(t *testing.T) {
	mockUser := new(MockUserRepo)
	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	idStr := strconv.Itoa(user.ID)

	mockUser.On("Delete", idStr).Return(user, nil)

	got, err := mockUser.Delete(idStr)

	assert.Nil(t, err)
	assert.Equal(t, got.ID, user.ID)
	assert.Equal(t, got.Email, user.Email)
	assert.Equal(t, got.Name, user.Name)
	assert.Equal(t, got.Password, user.Password)
	assert.Equal(t, got.UserPreferences, user.UserPreferences)
}

func TestUpdateUser(t *testing.T) {
	mockUser := new(MockUserRepo)
	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()

	idStr := strconv.Itoa(user.ID)

	mockUser.On("Save", user).Return(user.ID, nil)
	_, _ = mockUser.Save(user)

	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("lucas").
		Email("l@gmail.com").
		Password("321")
	user = u.BuildUser()
	mockUser.On("Update", user, idStr).Return(nil)

	err := mockUser.Update(user, idStr)

	assert.Nil(t, err)

}

func TestCheckIfUserExists(t *testing.T) {
	mockUser := new(MockUserRepo)
	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	u := entities.NewUserBuilder()
	u.Has().
		Uprefs(prefs).
		ID(1).
		Name("bruno").
		Email("b@gmail.com").
		Password("123")
	user := u.BuildUser()
	idStr := strconv.Itoa(user.ID)
	mockUser.On("Save", user).Return(user.ID, nil)
	mockUser.On("CheckIfExists", idStr).Return(true)
	mockUser.Save(user)
	got := mockUser.CheckIfExists(idStr)

	assert.Equal(t, true, got)

}
