package repository

import (
	"rest-api/golang/exercise/domain/entities"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockKennelRepo struct {
	mock.Mock
}

func (kr *MockKennelRepo) FindAllRepo() ([]entities.Kennel, error) {
	args := kr.Called()
	return args.Get(0).([]entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) SaveRepo(u *entities.Kennel) (int, error) {
	args := kr.Called(u)
	return args.Int(0), args.Error(1)
}

func (kr *MockKennelRepo) FindByIdRepo(id string) (*entities.Kennel, error) {
	args := kr.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) DeleteRepo(id string) (*entities.Kennel, error) {
	args := kr.Called(id)
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

func TestFindAllKennel(t *testing.T) {
	kennelMock := new(MockKennelRepo)

	breed := entities.BuildDogBreed("x", "z", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "a", "b")
	address := entities.BuildAddress(1, "a", "b", "c", "d", "e")

	kennel := entities.BuildKennel(1, []entities.Dog{*dog}, *address, "a", "b")

	kennelMock.On("FindAllRepo").Return([]entities.Kennel{*kennel}, nil)

	got, err := kennelMock.FindAllRepo()

	assert.Nil(t, err)
	assert.Equal(t, got[0].ID, kennel.ID)
	assert.Equal(t, got[0].ContactNumber, kennel.ContactNumber)
	assert.Equal(t, got[0].Name, kennel.Name)
	assert.Equal(t, got[0].Dogs, kennel.Dogs)
	assert.Equal(t, got[0].Address, kennel.Address)
}

func TestSaveRepoKennel(t *testing.T) {
	kennelMock := new(MockKennelRepo)

	breed := entities.BuildDogBreed("x", "z", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "a", "b")
	address := entities.BuildAddress(1, "a", "b", "c", "d", "e")

	kennel := entities.BuildKennel(1, []entities.Dog{*dog}, *address, "a", "b")

	kennelMock.On("SaveRepo", kennel).Return(kennel.ID, nil)

	got, err := kennelMock.SaveRepo(kennel)

	assert.Nil(t, err)
	assert.Equal(t, got, kennel.ID)
}

func TestFindByIdRepo(t *testing.T) {
	kennelMock := new(MockKennelRepo)

	breed := entities.BuildDogBreed("x", "z", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "a", "b")
	address := entities.BuildAddress(1, "a", "b", "c", "d", "e")

	kennel := entities.BuildKennel(1, []entities.Dog{*dog}, *address, "a", "b")
	idStr := strconv.Itoa(kennel.ID)

	kennelMock.On("SaveRepo", kennel).Return(kennel.ID, nil)
	kennelMock.On("FindByIdRepo", idStr).Return(kennel, nil)

	_, _ = kennelMock.SaveRepo(kennel)

	got, err := kennelMock.FindByIdRepo(idStr)

	assert.Nil(t, err)
	assert.Equal(t, got, kennel)
}

func TestDeleteKennelRepo(t *testing.T) {
	kennelMock := new(MockKennelRepo)

	breed := entities.BuildDogBreed("x", "z", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "a", "b")
	address := entities.BuildAddress(1, "a", "b", "c", "d", "e")

	kennel := entities.BuildKennel(1, []entities.Dog{*dog}, *address, "a", "b")
	idStr := strconv.Itoa(kennel.ID)

	kennelMock.On("SaveRepo", kennel).Return(kennel.ID, nil)
	kennelMock.On("DeleteRepo", idStr).Return(kennel, nil)

	_, _ = kennelMock.SaveRepo(kennel)

	got, err := kennelMock.DeleteRepo(idStr)

	assert.Nil(t, err)
	assert.Equal(t, got, kennel)

}

func TestUpdateKennelRepo(t *testing.T) {
	kennelMock := new(MockKennelRepo)

	breed := entities.BuildDogBreed("x", "z", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "a", "b")
	address := entities.BuildAddress(1, "a", "b", "c", "d", "e")

	kennel := entities.BuildKennel(1, []entities.Dog{*dog}, *address, "a", "b")
	newKennel := entities.BuildKennel(1, []entities.Dog{*dog}, *address, "c", "d")

	idStr := strconv.Itoa(kennel.ID)

	kennelMock.On("SaveRepo", kennel).Return(kennel.ID, nil)
	kennelMock.On("UpdateRepo", newKennel, idStr).Return(nil)

	_, _ = kennelMock.SaveRepo(kennel)

	err := kennelMock.UpdateRepo(newKennel, idStr)

	assert.Nil(t, err)
}

func TestCheckIfKennelExists(t *testing.T) {
	kennelMock := new(MockKennelRepo)

	breed := entities.BuildDogBreed("x", "z", 1, 1, 1, 1, 1, 1, 1)
	dog := entities.BuildDog(*breed, 1, 1, "a", "b")
	address := entities.BuildAddress(1, "a", "b", "c", "d", "e")

	kennel := entities.BuildKennel(1, []entities.Dog{*dog}, *address, "a", "b")
	idStr := strconv.Itoa(kennel.ID)

	kennelMock.On("SaveRepo", kennel).Return(kennel.ID, nil)
	kennelMock.On("CheckIfExistsRepo", idStr).Return(true)

	_, _ = kennelMock.SaveRepo(kennel)
	got := kennelMock.CheckIfExistsRepo(idStr)

	assert.Equal(t, true, got)
}
