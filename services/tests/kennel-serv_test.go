package tests

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/services"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type kennelMock struct {
	mock.Mock
}

type addrMock struct {
	mock.Mock
}

func (addr *addrMock) SaveAddress(ad *entities.Address) error {
	args := addr.Called(ad)
	return args.Error(0)
}

func (k *kennelMock) FindAllRepo() ([]entities.Kennel, error) {
	args := k.Called()
	return args.Get(0).([]entities.Kennel), args.Error(1)
}

func (k *kennelMock) SaveRepo(u *entities.Kennel) (int, error) {
	args := k.Called(u)
	return args.Int(0), args.Error(1)
}

func (k *kennelMock) FindByIdRepo(id string) (*entities.Kennel, error) {
	args := k.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (k *kennelMock) DeleteRepo(id string) (*entities.Kennel, error) {
	args := k.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (k *kennelMock) UpdateRepo(u *entities.Kennel, id string) error {
	args := k.Called(u, id)
	return args.Error(0)
}

func (k *kennelMock) CheckIfExistsRepo(id string) bool {
	args := k.Called(id)
	return args.Bool(0)
}

func MakeKennel() *entities.Kennel {
	breed := entities.BuildDogBreed("1", 1, 2, 3, 4, 5, 6, 7)
	dogs := entities.BuildDog(*breed, 1, 2, "M", "B")
	addr := entities.Address{ID_Kennel: 1, Numero: "2", Rua: "3", Bairro: "4", CEP: "5", Cidade: "R"}
	kennel := entities.Kennel{ID: 1, ContactNumber: "1", Dogs: []entities.Dog{*dogs}, Address: addr}

	return &kennel
}

func TestFindAllKennels(t *testing.T) {
	mock := new(kennelMock)
	kennel := MakeKennel()

	mock.On("FindAll").Return([]entities.Kennel{*kennel}, nil)

	testService := services.NewKennelService(mock, nil)
	result, err := testService.FindAllKennels()

	mock.AssertExpectations(t)

	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "1", result[0].ContactNumber)
	assert.Equal(t, kennel.Dogs, result[0].Dogs)
	assert.Equal(t, kennel.Address, result[0].Address)

	assert.Nil(t, err)
}

func TestFindKennelById(t *testing.T) {
	mock := new(kennelMock)
	kennel := MakeKennel()
	idStr := strconv.Itoa(kennel.ID)
	mock.On("FindById", idStr).Return(kennel, nil)

	testService := services.NewKennelService(mock, nil)
	result, err := testService.FindKennelByIdServ(idStr)
	mock.AssertExpectations(t)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "1", result.ContactNumber)
	assert.Equal(t, kennel.Dogs, result.Dogs)
	assert.Equal(t, kennel.Address, result.Address)

	assert.Nil(t, err)
}

func TestSaveKennel(t *testing.T) {
	mock := new(kennelMock)
	mockad := new(addrMock)
	kennel := MakeKennel()
	//idStr := strconv.Itoa(kennel.ID)
	addr := entities.Address{
		ID_Kennel: 1,
		Numero:    "2",
		Rua:       "3",
		Bairro:    "4",
		CEP:       "5",
		Cidade:    "R",
	}

	mockad.On("Save", &addr).Return(nil)
	mock.On("Save", kennel).Return(kennel.ID, nil)

	testService := services.NewKennelService(mock, mockad)
	result, err := testService.Save(kennel)

	mock.AssertExpectations(t)
	mockad.AssertExpectations(t)

	assert.Equal(t, 1, result)
	assert.Nil(t, err)

	/*
		Como o service de kennel chama o repo de Address, eu tenho que mockar tbm
		o repositorio de Address e fazer as chamadas certas dele no teste. BUT HOW???
	*/
}

func TestDeleteKennel(t *testing.T) {
	mock := new(kennelMock)
	kennel := MakeKennel()
	idStr := strconv.Itoa(kennel.ID)

	mock.On("Delete", idStr).Return(kennel, nil)

	testService := services.NewKennelService(mock, nil)
	result, err := testService.DeleteKennelServ(idStr)
	mock.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "1", result.ContactNumber)
	assert.Equal(t, kennel.Dogs, result.Dogs)
	assert.Equal(t, kennel.Address, result.Address)

	assert.Nil(t, err)
}

func TestUpdateKennel(t *testing.T) {
	mock := new(kennelMock)
	kennel := MakeKennel()
	idStr := strconv.Itoa(kennel.ID)

	mock.On("Update", kennel, idStr).Return(nil)

	testService := services.NewKennelService(mock, nil)
	err := testService.UpdateKennelServ(kennel, idStr)

	mock.AssertExpectations(t)

	assert.Nil(t, err)

}
