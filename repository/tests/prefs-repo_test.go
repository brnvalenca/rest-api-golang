package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPrefsRepo struct {
	mock.Mock
}

func (pr *MockPrefsRepo) SavePrefs(u *entities.UserDogPreferences) error {
	args := pr.Called(u)
	return args.Error(0)
}

func (pr *MockPrefsRepo) DeletePrefs(id string) error {
	args := pr.Called(id)
	return args.Error(0)
}

func (pr *MockPrefsRepo) UpdatePrefs(u *entities.UserDogPreferences, id string) error {
	args := pr.Called(u, id)
	return args.Error(0)
}

func TestSavePrefs(t *testing.T) {
	prefsMock := new(MockPrefsRepo)

	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)

	prefsMock.On("SavePrefs", &prefs).Return(nil)

	err := prefsMock.SavePrefs(&prefs)
	assert.Nil(t, err)

}

func TestDeletePrefs(t *testing.T) {
	prefsMock := new(MockPrefsRepo)

	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	idStr := strconv.Itoa(prefs.UserID)

	prefsMock.On("SavePrefs", &prefs).Return(nil)
	prefsMock.On("DeletePrefs", idStr).Return(nil)

	_ = prefsMock.SavePrefs(&prefs)
	err := prefsMock.DeletePrefs(idStr)
	assert.Nil(t, err)

}

func TestUpdatePrefs(t *testing.T) {
	prefsMock := new(MockPrefsRepo)

	prefs := entities.BuildUserDogPreferences(1, 1, 1, 1, 1, 1)
	idStr := strconv.Itoa(prefs.UserID)
	newPrefs := entities.BuildUserDogPreferences(1, 2, 2, 2, 2, 2)

	prefsMock.On("SavePrefs", &prefs).Return(nil)
	prefsMock.On("UpdatePrefs", &newPrefs, idStr).Return(nil)
	_ = prefsMock.SavePrefs(&prefs)
	err := prefsMock.UpdatePrefs(&newPrefs, idStr)

	assert.Nil(t, err)
}
