package services

import (
	"rest-api/golang/exercise/domain/dtos"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type kennelMockRepo struct {
	mock.Mock
}

type addrMockRepo struct {
	mock.Mock
}

func (addr *addrMockRepo) SaveAddress(ad *entities.Address, kennelid int) error {
	args := addr.Called(ad, kennelid)
	return args.Error(0)
}

func (k *kennelMockRepo) FindAllKennelRepo() ([]entities.Kennel, error) {
	args := k.Called()
	return args.Get(0).([]entities.Kennel), args.Error(1)
}

func (k *kennelMockRepo) SaveKennelRepo(u *entities.Kennel) (int, error) {
	args := k.Called(u)
	return args.Int(0), args.Error(1)
}

func (k *kennelMockRepo) FindByIdKennelRepo(id string) (*entities.Kennel, error) {
	args := k.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (k *kennelMockRepo) DeleteKennelRepo(id string) (*entities.Kennel, error) {
	args := k.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (k *kennelMockRepo) UpdateKennelRepo(u *entities.Kennel, addr *entities.Address, id string) error {
	args := k.Called(u, addr, id)
	return args.Error(0)
}

func (k *kennelMockRepo) CheckIfKennelExistsRepo(id string) bool {
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

func MakeKennelDTO() *dtos.KennelDTO {
	dogs := MakeDogDTO()
	kennelBuilder := dtos.NewKennelBuilderDTO()
	kennelBuilder.Has().
		Bairro("4").
		ID(1).
		Numero("2").
		Rua("3").
		CEP("4").
		Cidade("R").
		ContactNumber("5").
		Dogs([]dtos.DogDTO{dogs})

	kennel := kennelBuilder.BuildKennel()
	return kennel
}

func TestFindAllKennels(t *testing.T) {
	kennelMockRepo := new(kennelMockRepo)
	kennel := MakeKennel()

	//mock.On("FindAll").Return([]entities.Kennel{*kennel}, nil)
	kennelMockRepo.On("FindAllRepo").Return([]entities.Kennel{*kennel}, nil)

	testService := NewKennelService(kennelMockRepo, nil)
	result, err := testService.FindAllKennels()

	kennelMockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "1", result[0].ContactNumber)
	assert.Equal(t, kennel.Dogs, result[0].Dogs)

	assert.Nil(t, err)
}

func TestFindKennelById(t *testing.T) {
	kennelMockRepo := new(kennelMockRepo)
	kennel := MakeKennel()
	idStr := strconv.Itoa(kennel.ID)
	kennelMockRepo.On("FindByIdRepo", idStr).Return(kennel, nil)

	testService := NewKennelService(kennelMockRepo, nil)
	result, err := testService.FindKennelByIdServ(idStr)
	kennelMockRepo.AssertExpectations(t)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "1", result.ContactNumber)
	assert.Equal(t, kennel.Dogs, result.Dogs)

	assert.Nil(t, err)
}

func TestSaveKennel(t *testing.T) {
	kennelMockRepo := new(kennelMockRepo)
	addrMockRepo := new(addrMockRepo)
	kennelDTO := MakeKennelDTO()
	kennelAddr, kennel := utils.PartitionKennelDTO(kennelDTO)
	//idStr := strconv.Itoa(kennel.ID)
	addrMockRepo.On("SaveAddress", kennelAddr, kennel.ID).Return(nil)
	kennelMockRepo.On("SaveRepo", kennel).Return(kennel.ID, nil)

	testService := NewKennelService(kennelMockRepo, addrMockRepo)
	testService.ValidateKennel(kennel)
	result, err := testService.SaveKennel(kennelDTO)

	addrMockRepo.AssertExpectations(t)
	kennelMockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result)
	assert.Nil(t, err)

}

func TestDeleteKennel(t *testing.T) {
	kennelMockRepo := new(kennelMockRepo)
	kennel := MakeKennel()
	idStr := strconv.Itoa(kennel.ID)

	kennelMockRepo.On("DeleteRepo", idStr).Return(kennel, nil)

	testService := NewKennelService(kennelMockRepo, nil)
	result, err := testService.DeleteKennelServ(idStr)
	kennelMockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "1", result.ContactNumber)
	assert.Equal(t, kennel.Dogs, result.Dogs)

	assert.Nil(t, err)
}

func TestUpdateKennel(t *testing.T) {
	kennelMockRepo := new(kennelMockRepo)
	kennelDTO := MakeKennelDTO()
	kennelAddr, kennel := utils.PartitionKennelDTO(kennelDTO)
	idStr := strconv.Itoa(kennelDTO.ID)
	kennelMockRepo.On("UpdateRepo", kennel, kennelAddr, idStr).Return(nil)

	testService := NewKennelService(kennelMockRepo, nil)
	err := testService.UpdateKennelServ(kennelDTO, idStr)

	kennelMockRepo.AssertExpectations(t)

	assert.Nil(t, err)

}
