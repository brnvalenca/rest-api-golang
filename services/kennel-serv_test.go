package services

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
	"strconv"
	"testing"
	"rest-api/golang/exercise/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type kennelMock struct {
	mock.Mock
}

type addrMock struct {
	mock.Mock
}

func (addr *addrMock) SaveAddress(ad *entities.Address, kennelid int) error {
	args := addr.Called(ad, kennelid)
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

func (k *kennelMock) UpdateRepo(u *entities.Kennel, addr *entities.Address, id string) error {
	args := k.Called(u, addr, id)
	return args.Error(0)
}

func (k *kennelMock) CheckIfExistsRepo(id string) bool {
	args := k.Called(id)
	return args.Bool(0)
}

func MakeKennel() *entities.Kennel {

	db := entities.NewDogBreedBuilder()
	db.Has().
		ID(1).
		Name("x").
		Img("1").
		GoodWithKidsAndDogs(3, 4).
		SheddGroomAndEnergy(5, 6, 7)
	breed := db.BuildBreed()

	d := entities.NewDogBuilder()
	d.Has().
		KennelID(2).
		DogID(1).
		NameAndSex("B", "M").
		Breed(*breed)
	dogs := d.BuildDog()

	ad := entities.NewAddressBuilder()
	ad.Has().
		IDKennel(1).
		Numero("2").
		Rua("3").
		Bairro("4").
		CEP("4").
		Cidade("R")
	addr := ad.BuildAddr()

	kennel := entities.Kennel{ID: 1, ContactNumber: "1", Dogs: []entities.Dog{*dogs}, Address: *addr}

	return &kennel
}

func TestFindAllKennels(t *testing.T) {
	mock := new(kennelMock)
	kennel := MakeKennel()

	//mock.On("FindAll").Return([]entities.Kennel{*kennel}, nil)
	mock.On("FindAllRepo").Return([]entities.Kennel{*kennel}, nil)

	testService := NewKennelService(mock, nil)
	result, err := testService.FindAllKennels()

	mock.AssertExpectations(t)

	assert.Equal(t, 1, result[0].Kennel.ID)
	assert.Equal(t, "1", result[0].Kennel.ContactNumber)
	assert.Equal(t, kennel.Dogs, result[0].Kennel.Dogs)
	assert.Equal(t, kennel.Address, result[0].Kennel.Address)

	assert.Nil(t, err)
}

func TestFindKennelById(t *testing.T) {
	mock := new(kennelMock)
	kennel := MakeKennel()
	idStr := strconv.Itoa(kennel.ID)
	mock.On("FindByIdRepo", idStr).Return(kennel, nil)

	testService := NewKennelService(mock, nil)
	result, err := testService.FindKennelByIdServ(idStr)
	mock.AssertExpectations(t)
	assert.Equal(t, 1, result.Kennel.ID)
	assert.Equal(t, "1", result.Kennel.ContactNumber)
	assert.Equal(t, kennel.Dogs, result.Kennel.Dogs)
	assert.Equal(t, kennel.Address, result.Kennel.Address)

	assert.Nil(t, err)
}

func TestSaveKennel(t *testing.T) {
	kennelMock := new(kennelMock)
	addrMock := new(addrMock)
	kennel := MakeKennel()
	//idStr := strconv.Itoa(kennel.ID)
	kennelDTO := dtos.KennelDTO{Kennel: *kennel}
	addrMock.On("SaveAddress", &kennel.Address, kennel.ID).Return(nil)
	kennelMock.On("SaveRepo", kennel).Return(kennel.ID, nil)

	middleware.PartitionKennelDTO(&kennelDTO)
	ValidateKennel(kennel)
	testService := NewKennelService(kennelMock, addrMock)
	result, err := testService.SaveKennel(&kennelDTO)

	addrMock.AssertExpectations(t)
	kennelMock.AssertExpectations(t)

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

	mock.On("DeleteRepo", idStr).Return(kennel, nil)

	testService := NewKennelService(mock, nil)
	result, err := testService.DeleteKennelServ(idStr)
	mock.AssertExpectations(t)

	assert.Equal(t, 1, result.Kennel.ID)
	assert.Equal(t, "1", result.Kennel.ContactNumber)
	assert.Equal(t, kennel.Dogs, result.Kennel.Dogs)
	assert.Equal(t, kennel.Address, result.Kennel.Address)

	assert.Nil(t, err)
}

func TestUpdateKennel(t *testing.T) {
	mock := new(kennelMock)
	kennel := MakeKennel()
	idStr := strconv.Itoa(kennel.ID)
	kennelDTO := dtos.KennelDTO{Kennel: *kennel}
	mock.On("UpdateRepo", kennel, &kennel.Address, idStr).Return(nil)

	testService := NewKennelService(mock, nil)
	err := testService.UpdateKennelServ(&kennelDTO, idStr)

	mock.AssertExpectations(t)

	assert.Nil(t, err)

}
